// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2018 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"bytes"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

var testPubKey = []byte("-----BEGIN PUBLIC KEY-----\n" +
	"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAol0Z8G8U+25Btxk/g/fm\n" +
	"UAW/wEKjQCTjkibDE4B+qkuWeiumg6miIRhtilU6m9BFmLQSy1ltYQuu4k17A4tQ\n" +
	"rIPpOQYZges/qsDFkZh3wyK5jL5WEFVdOasf6wsfszExnPmcZS4axxoYJfiuilrN\n" +
	"hnwinBAqfi3S0sw5MpSI4Zl1AbOrHG4zDI62Gti2PKiMGyYDZTS9xPrBLbN95Kby\n" +
	"FFclQLEzA9RJcS1nHFsWtRgHjGPhhjCQxEm9NQ1nePFhCfBfApyfH1VM2VCOQum6\n" +
	"Ci9bMuHWjTjckC84mzF99kOxOWVU7mwS6gnJqBzpuz8t3zq8/iQ2y7QrmZV+jTJP\n" +
	"WQIDAQAB\n" +
	"-----END PUBLIC KEY-----\n")

var testPubKeyRSA *rsa.PublicKey

func init() {
	block, _ := pem.Decode(testPubKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	testPubKeyRSA = pub.(*rsa.PublicKey)
}

func TestScrambleOldPass(t *testing.T) {
	scramble := []byte{9, 8, 7, 6, 5, 4, 3, 2}
	vectors := []struct {
		pass string
		out  string
	}{
		{" pass", "47575c5a435b4251"},
		{"pass ", "47575c5a435b4251"},
		{"123\t456", "575c47505b5b5559"},
		{"C0mpl!ca ted#PASS123", "5d5d554849584a45"},
	}
	for _, tuple := range vectors {
		ours := scrambleOldPassword(scramble, tuple.pass)
		if tuple.out != fmt.Sprintf("%x", ours) {
			t.Errorf("Failed old password %q", tuple.pass)
		}
	}
}

func TestScrambleSHA256Pass(t *testing.T) {
	scramble := []byte{10, 47, 74, 111, 75, 73, 34, 48, 88, 76, 114, 74, 37, 13, 3, 80, 82, 2, 23, 21}
	vectors := []struct {
		pass string
		out  string
	}{
		{"secret", "f490e76f66d9d86665ce54d98c78d0acfe2fb0b08b423da807144873d30b312c"},
		{"secret2", "abc3934a012cf342e876071c8ee202de51785b430258a7a0138bc79c4d800bc6"},
	}
	for _, tuple := range vectors {
		ours := scrambleSHA256Password(scramble, tuple.pass)
		if tuple.out != fmt.Sprintf("%x", ours) {
			t.Errorf("Failed SHA256 password %q", tuple.pass)
		}
	}
}

func TestAuthFastCachingSHA256PasswordCached(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"

	authData := []byte{18, 170, 22, 142, 90, 59, 56, 77, 8, 65, 225, 157, 53,
		22, 41, 84, 32, 123, 43, 118}
	plugin := "caching_sha2_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{102, 32, 5, 35, 143, 161, 140, 241, 171, 232, 56,
		139, 43, 14, 107, 196, 249, 170, 147, 60, 220, 204, 120, 178, 214, 15,
		184, 150, 26, 61, 57, 235}
	if writtenAuthRespLen != 32 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response
	conn.data = []byte{
		2, 0, 0, 2, 1, 3, // Fast Auth Success
		7, 0, 0, 3, 0, 0, 0, 2, 0, 0, 0, // OK
	}
	conn.maxReads = 1

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}
}

func TestAuthFastCachingSHA256PasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = ""

	authData := []byte{72, 165, 80, 208, 22, 91, 63, 14, 3, 56, 253, 149, 40,
		22, 41, 84, 32, 123, 43, 118}
	plugin := "caching_sha2_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	if writtenAuthRespLen != 0 {
		t.Fatalf("unexpected written auth response (%d bytes): %v",
			writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response
	conn.data = []byte{
		7, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, // OK
	}
	conn.maxReads = 1

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}
}

func TestAuthFastCachingSHA256PasswordFullRSA(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"

	authData := []byte{7, 57, 96, 216, 18, 46, 77, 67, 30, 67, 5, 18, 138, 96,
		62, 94, 83, 80, 52, 85}
	plugin := "caching_sha2_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{171, 201, 138, 146, 89, 159, 11, 170, 0, 67, 165,
		49, 175, 94, 218, 68, 177, 109, 110, 86, 34, 33, 44, 190, 67, 240, 70,
		110, 40, 139, 124, 41}
	if writtenAuthRespLen != 32 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response
	conn.data = []byte{
		2, 0, 0, 2, 1, 4, // Perform Full Authentication
	}
	conn.queuedReplies = [][]byte{
		// pub key response
		append([]byte{byte(1 + len(testPubKey)), 1, 0, 4, 1}, testPubKey...),

		// OK
		{7, 0, 0, 6, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 3

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	if !bytes.HasPrefix(conn.written, []byte{1, 0, 0, 3, 2, 0, 1, 0, 5}) {
		t.Errorf("unexpected written data: %v", conn.written)
	}
}

func TestAuthFastCachingSHA256PasswordFullRSAWithKey(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"
	mc.cfg.pubKey = testPubKeyRSA

	authData := []byte{0, 52, 35, 155, 71, 16, 63, 45, 95, 62, 9, 64, 154, 41,
		62, 94, 83, 80, 52, 85}
	plugin := "caching_sha2_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{171, 201, 138, 146, 89, 159, 11, 170, 0, 67, 165,
		49, 175, 94, 218, 68, 177, 109, 110, 86, 34, 33, 44, 190, 67, 240, 70,
		110, 40, 139, 124, 41}
	if writtenAuthRespLen != 32 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response
	conn.data = []byte{
		2, 0, 0, 2, 1, 4, // Perform Full Authentication
	}
	conn.queuedReplies = [][]byte{
		// OK
		{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 2

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	if !bytes.HasPrefix(conn.written, []byte{0, 1, 0, 3}) {
		t.Errorf("unexpected written data: %v", conn.written)
	}
}

func TestAuthFastCachingSHA256PasswordFullSecure(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"

	authData := []byte{8, 79, 94, 254, 77, 56, 92, 96, 31, 95, 8, 95, 179, 27,
		62, 94, 83, 80, 52, 85}
	plugin := "caching_sha2_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// Hack to make the caching_sha2_password plugin believe that the connection
	// is secure
	mc.cfg.tls = &tls.Config{InsecureSkipVerify: true}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{171, 201, 138, 146, 89, 159, 11, 170, 0, 67, 165,
		49, 175, 94, 218, 68, 177, 109, 110, 86, 34, 33, 44, 190, 67, 240, 70,
		110, 40, 139, 124, 41}
	if writtenAuthRespLen != 32 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response
	conn.data = []byte{
		2, 0, 0, 2, 1, 4, // Perform Full Authentication
	}
	conn.queuedReplies = [][]byte{
		// OK
		{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 3

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	if !bytes.Equal(conn.written, []byte{7, 0, 0, 3, 115, 101, 99, 114, 101, 116, 0}) {
		t.Errorf("unexpected written data: %v", conn.written)
	}
}

func TestAuthFastCleartextPasswordNotAllowed(t *testing.T) {
	_, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"

	authData := []byte{29, 155, 94, 63, 4, 46, 15, 109, 91, 171, 29, 224, 251,
		103, 26, 95, 81, 17, 24, 21}
	plugin := "mysql_clear_password"

	// Send Client Authentication Packet
	_, err := mc.auth(authData, plugin)
	if err != ErrCleartextPassword {
		t.Errorf("expected ErrCleartextPassword, got %v", err)
	}
}

func TestAuthFastCleartextPassword(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"
	mc.cfg.AllowCleartextPasswords = true

	authData := []byte{96, 206, 38, 38, 3, 78, 71, 213, 82, 240, 31, 128, 195,
		103, 26, 95, 81, 17, 24, 21}
	plugin := "mysql_clear_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{115, 101, 99, 114, 101, 116, 0}
	if writtenAuthRespLen != 7 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response
	conn.data = []byte{
		7, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, // OK
	}
	conn.maxReads = 1

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}
}

func TestAuthFastCleartextPasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = ""
	mc.cfg.AllowCleartextPasswords = true

	authData := []byte{13, 124, 91, 57, 8, 95, 15, 201, 35, 215, 46, 189, 159,
		103, 26, 95, 81, 17, 24, 21}
	plugin := "mysql_clear_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{0}
	if writtenAuthRespLen != 1 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response
	conn.data = []byte{
		7, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, // OK
	}
	conn.maxReads = 1

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}
}

func TestAuthFastNativePasswordNotAllowed(t *testing.T) {
	_, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"
	mc.cfg.AllowNativePasswords = false

	authData := []byte{71, 212, 78, 88, 5, 67, 89, 165, 51, 109, 36, 172, 111,
		103, 26, 95, 81, 17, 24, 21}
	plugin := "mysql_native_password"

	// Send Client Authentication Packet
	_, err := mc.auth(authData, plugin)
	if err != ErrNativePassword {
		t.Errorf("expected ErrNativePassword, got %v", err)
	}
}

func TestAuthFastNativePassword(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"

	authData := []byte{82, 109, 36, 88, 1, 42, 68, 130, 55, 165, 44, 240, 156,
		103, 26, 95, 81, 17, 24, 21}
	plugin := "mysql_native_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{53, 177, 140, 159, 251, 189, 127, 53, 109, 252,
		172, 50, 211, 192, 240, 164, 26, 48, 207, 45}
	if writtenAuthRespLen != 20 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response
	conn.data = []byte{
		7, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, // OK
	}
	conn.maxReads = 1

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}
}

func TestAuthFastNativePasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = ""

	authData := []byte{26, 232, 70, 44, 3, 46, 94, 220, 28, 130, 74, 121, 140,
		103, 26, 95, 81, 17, 24, 21}
	plugin := "mysql_native_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	if writtenAuthRespLen != 0 {
		t.Fatalf("unexpected written auth response (%d bytes): %v",
			writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response
	conn.data = []byte{
		7, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, // OK
	}
	conn.maxReads = 1

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}
}

func TestAuthFastSHA256PasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = ""

	authData := []byte{6, 54, 34, 254, 66, 54, 31, 84, 69, 21, 7, 38, 217, 66,
		62, 94, 83, 80, 52, 85}
	plugin := "sha256_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{0}
	if writtenAuthRespLen != 1 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response (pub key response)
	conn.data = append([]byte{byte(1 + len(testPubKey)), 1, 0, 2, 1}, testPubKey...)
	conn.queuedReplies = [][]byte{
		// OK
		{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 2

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	if !bytes.HasPrefix(conn.written, []byte{0, 1, 0, 3}) {
		t.Errorf("unexpected written data: %v", conn.written)
	}
}

func TestAuthFastSHA256PasswordRSA(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"

	authData := []byte{4, 29, 77, 133, 52, 58, 42, 19, 34, 93, 0, 72, 218, 89,
		62, 94, 83, 80, 52, 85}
	plugin := "sha256_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{1}
	if writtenAuthRespLen != 1 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response (pub key response)
	conn.data = append([]byte{byte(1 + len(testPubKey)), 1, 0, 2, 1}, testPubKey...)
	conn.queuedReplies = [][]byte{
		// OK
		{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 2

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	if !bytes.HasPrefix(conn.written, []byte{0, 1, 0, 3}) {
		t.Errorf("unexpected written data: %v", conn.written)
	}
}

func TestAuthFastSHA256PasswordRSAWithKey(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"
	mc.cfg.pubKey = testPubKeyRSA

	authData := []byte{1, 18, 38, 124, 95, 97, 32, 22, 66, 77, 9, 25, 214, 45,
		62, 94, 83, 80, 52, 85}
	plugin := "sha256_password"

	// Send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}
	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// auth response (OK)
	conn.data = []byte{7, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0}
	conn.maxReads = 1

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}
}

func TestAuthFastSHA256PasswordSecure(t *testing.T) {
	conn, mc := newRWMockConn(1)
	mc.cfg.User = "root"
	mc.cfg.Passwd = "secret"

	// hack to make the caching_sha2_password plugin believe that the connection
	// is secure
	mc.cfg.tls = &tls.Config{InsecureSkipVerify: true}

	authData := []byte{7, 92, 62, 226, 94, 62, 65, 53, 86, 49, 4, 42, 213, 15,
		62, 94, 83, 80, 52, 85}
	plugin := "sha256_password"

	// send Client Authentication Packet
	authResp, err := mc.auth(authData, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// unset TLS config to prevent the actual establishment of a TLS wrapper
	mc.cfg.tls = nil

	err = mc.writeHandshakeResponsePacket(authResp, plugin)
	if err != nil {
		t.Fatal(err)
	}

	// check written auth response
	authRespStart := 4 + 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1
	authRespEnd := authRespStart + 1 + len(authResp)
	writtenAuthRespLen := conn.written[authRespStart]
	writtenAuthResp := conn.written[authRespStart+1 : authRespEnd]
	expectedAuthResp := []byte{115, 101, 99, 114, 101, 116, 0}
	if writtenAuthRespLen != 7 || !bytes.Equal(writtenAuthResp, expectedAuthResp) {
		t.Fatalf("unexpected written auth response (%d bytes): %v", writtenAuthRespLen, writtenAuthResp)
	}
	conn.written = nil

	// auth response (OK)
	conn.data = []byte{7, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0}
	conn.maxReads = 1

	// Handle response to auth packet
	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	if !bytes.Equal(conn.written, []byte{}) {
		t.Errorf("unexpected written data: %v", conn.written)
	}
}

func TestAuthSwitchCachingSHA256PasswordCached(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.Passwd = "secret"

	// auth switch request
	conn.data = []byte{44, 0, 0, 2, 254, 99, 97, 99, 104, 105, 110, 103, 95,
		115, 104, 97, 50, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 101,
		11, 26, 18, 94, 97, 22, 72, 2, 46, 70, 106, 29, 55, 45, 94, 76, 90, 84,
		50, 0}

	// auth response
	conn.queuedReplies = [][]byte{
		{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0}, // OK
	}
	conn.maxReads = 3

	authData := []byte{206, 77, 40, 36, 97, 66, 17, 146, 15, 201, 94, 28, 20,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{
		// 1. Packet: Hash
		32, 0, 0, 3, 129, 93, 132, 95, 114, 48, 79, 215, 128, 62, 193, 118, 128,
		54, 75, 208, 159, 252, 227, 215, 129, 15, 242, 97, 19, 159, 31, 20, 58,
		153, 9, 130,
	}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchCachingSHA256PasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.Passwd = ""

	// auth switch request
	conn.data = []byte{44, 0, 0, 2, 254, 99, 97, 99, 104, 105, 110, 103, 95,
		115, 104, 97, 50, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 101,
		11, 26, 18, 94, 97, 22, 72, 2, 46, 70, 106, 29, 55, 45, 94, 76, 90, 84,
		50, 0}

	// auth response
	conn.queuedReplies = [][]byte{{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0}}
	conn.maxReads = 2

	authData := []byte{252, 21, 59, 42, 31, 38, 43, 102, 96, 184, 91, 36, 70,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{0, 0, 0, 3}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchCachingSHA256PasswordFullRSA(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.Passwd = "secret"

	// auth switch request
	conn.data = []byte{44, 0, 0, 2, 254, 99, 97, 99, 104, 105, 110, 103, 95,
		115, 104, 97, 50, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 101,
		11, 26, 18, 94, 97, 22, 72, 2, 46, 70, 106, 29, 55, 45, 94, 76, 90, 84,
		50, 0}

	conn.queuedReplies = [][]byte{
		// Perform Full Authentication
		{2, 0, 0, 4, 1, 4},

		// Pub Key Response
		append([]byte{byte(1 + len(testPubKey)), 1, 0, 6, 1}, testPubKey...),

		// OK
		{7, 0, 0, 8, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 4

	authData := []byte{202, 68, 58, 80, 70, 32, 29, 250, 14, 250, 15, 40, 42,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReplyPrefix := []byte{
		// 1. Packet: Hash
		32, 0, 0, 3, 129, 93, 132, 95, 114, 48, 79, 215, 128, 62, 193, 118, 128,
		54, 75, 208, 159, 252, 227, 215, 129, 15, 242, 97, 19, 159, 31, 20, 58,
		153, 9, 130,

		// 2. Packet: Pub Key Request
		1, 0, 0, 5, 2,

		// 3. Packet: Encrypted Password
		0, 1, 0, 7, // [changing bytes]
	}
	if !bytes.HasPrefix(conn.written, expectedReplyPrefix) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchCachingSHA256PasswordFullRSAWithKey(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.Passwd = "secret"
	mc.cfg.pubKey = testPubKeyRSA

	// auth switch request
	conn.data = []byte{44, 0, 0, 2, 254, 99, 97, 99, 104, 105, 110, 103, 95,
		115, 104, 97, 50, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 101,
		11, 26, 18, 94, 97, 22, 72, 2, 46, 70, 106, 29, 55, 45, 94, 76, 90, 84,
		50, 0}

	conn.queuedReplies = [][]byte{
		// Perform Full Authentication
		{2, 0, 0, 4, 1, 4},

		// OK
		{7, 0, 0, 6, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 3

	authData := []byte{201, 94, 22, 37, 96, 21, 16, 106, 75, 242, 23, 68, 93,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReplyPrefix := []byte{
		// 1. Packet: Hash
		32, 0, 0, 3, 129, 93, 132, 95, 114, 48, 79, 215, 128, 62, 193, 118, 128,
		54, 75, 208, 159, 252, 227, 215, 129, 15, 242, 97, 19, 159, 31, 20, 58,
		153, 9, 130,

		// 2. Packet: Encrypted Password
		0, 1, 0, 5, // [changing bytes]
	}
	if !bytes.HasPrefix(conn.written, expectedReplyPrefix) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchCachingSHA256PasswordFullSecure(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.Passwd = "secret"

	// Hack to make the caching_sha2_password plugin believe that the connection
	// is secure
	mc.cfg.tls = &tls.Config{InsecureSkipVerify: true}

	// auth switch request
	conn.data = []byte{44, 0, 0, 2, 254, 99, 97, 99, 104, 105, 110, 103, 95,
		115, 104, 97, 50, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 101,
		11, 26, 18, 94, 97, 22, 72, 2, 46, 70, 106, 29, 55, 45, 94, 76, 90, 84,
		50, 0}

	// auth response
	conn.queuedReplies = [][]byte{
		{2, 0, 0, 4, 1, 4},                // Perform Full Authentication
		{7, 0, 0, 6, 0, 0, 0, 2, 0, 0, 0}, // OK
	}
	conn.maxReads = 3

	authData := []byte{197, 36, 59, 29, 19, 97, 60, 242, 98, 167, 64, 75, 70,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{
		// 1. Packet: Hash
		32, 0, 0, 3, 129, 93, 132, 95, 114, 48, 79, 215, 128, 62, 193, 118, 128,
		54, 75, 208, 159, 252, 227, 215, 129, 15, 242, 97, 19, 159, 31, 20, 58,
		153, 9, 130,

		// 2. Packet: Cleartext password
		7, 0, 0, 5, 115, 101, 99, 114, 101, 116, 0,
	}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchCleartextPasswordNotAllowed(t *testing.T) {
	conn, mc := newRWMockConn(2)

	conn.data = []byte{22, 0, 0, 2, 254, 109, 121, 115, 113, 108, 95, 99, 108,
		101, 97, 114, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0}
	conn.maxReads = 1
	authData := []byte{173, 48, 51, 80, 39, 97, 45, 177, 50, 140, 48, 37, 52,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"
	err := mc.handleAuthResult(authData, plugin)
	if err != ErrCleartextPassword {
		t.Errorf("expected ErrCleartextPassword, got %v", err)
	}
}

func TestAuthSwitchCleartextPassword(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.AllowCleartextPasswords = true
	mc.cfg.Passwd = "secret"

	// auth switch request
	conn.data = []byte{22, 0, 0, 2, 254, 109, 121, 115, 113, 108, 95, 99, 108,
		101, 97, 114, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0}

	// auth response
	conn.queuedReplies = [][]byte{{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0}}
	conn.maxReads = 2

	authData := []byte{231, 76, 55, 87, 96, 38, 46, 215, 31, 107, 43, 96, 81,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{7, 0, 0, 3, 115, 101, 99, 114, 101, 116, 0}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchCleartextPasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.AllowCleartextPasswords = true
	mc.cfg.Passwd = ""

	// auth switch request
	conn.data = []byte{22, 0, 0, 2, 254, 109, 121, 115, 113, 108, 95, 99, 108,
		101, 97, 114, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0}

	// auth response
	conn.queuedReplies = [][]byte{{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0}}
	conn.maxReads = 2

	authData := []byte{169, 80, 92, 76, 87, 77, 13, 135, 27, 187, 22, 90, 44,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{1, 0, 0, 3, 0}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchNativePasswordNotAllowed(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.AllowNativePasswords = false

	conn.data = []byte{44, 0, 0, 2, 254, 109, 121, 115, 113, 108, 95, 110, 97,
		116, 105, 118, 101, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 96,
		71, 63, 8, 1, 58, 75, 12, 69, 95, 66, 60, 117, 31, 48, 31, 89, 39, 55,
		31, 0}
	conn.maxReads = 1
	authData := []byte{57, 13, 83, 0, 4, 21, 70, 54, 37, 63, 22, 95, 131, 24,
		48, 31, 89, 39, 55, 31}
	plugin := "caching_sha2_password"
	err := mc.handleAuthResult(authData, plugin)
	if err != ErrNativePassword {
		t.Errorf("expected ErrNativePassword, got %v", err)
	}
}

func TestAuthSwitchNativePassword(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.AllowNativePasswords = true
	mc.cfg.Passwd = "secret"

	// auth switch request
	conn.data = []byte{44, 0, 0, 2, 254, 109, 121, 115, 113, 108, 95, 110, 97,
		116, 105, 118, 101, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 96,
		71, 63, 8, 1, 58, 75, 12, 69, 95, 66, 60, 117, 31, 48, 31, 89, 39, 55,
		31, 0}

	// auth response
	conn.queuedReplies = [][]byte{{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0}}
	conn.maxReads = 2

	authData := []byte{89, 59, 36, 8, 1, 55, 85, 43, 36, 53, 34, 58, 133, 41,
		48, 31, 89, 39, 55, 31}
	plugin := "caching_sha2_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{20, 0, 0, 3, 202, 41, 195, 164, 34, 226, 49, 103,
		21, 211, 167, 199, 227, 116, 8, 48, 57, 71, 149, 146}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchNativePasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.AllowNativePasswords = true
	mc.cfg.Passwd = ""

	// auth switch request
	conn.data = []byte{44, 0, 0, 2, 254, 109, 121, 115, 113, 108, 95, 110, 97,
		116, 105, 118, 101, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 96,
		71, 63, 8, 1, 58, 75, 12, 69, 95, 66, 60, 117, 31, 48, 31, 89, 39, 55,
		31, 0}

	// auth response
	conn.queuedReplies = [][]byte{{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0}}
	conn.maxReads = 2

	authData := []byte{22, 53, 50, 6, 5, 44, 62, 18, 81, 40, 32, 33, 105, 89,
		48, 31, 89, 39, 55, 31}
	plugin := "caching_sha2_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{0, 0, 0, 3}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchOldPasswordNotAllowed(t *testing.T) {
	conn, mc := newRWMockConn(2)

	conn.data = []byte{41, 0, 0, 2, 254, 109, 121, 115, 113, 108, 95, 111, 108,
		100, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 95, 84, 103, 43, 61,
		49, 123, 61, 91, 50, 40, 113, 35, 84, 96, 101, 92, 123, 121, 107, 0}
	conn.maxReads = 1
	authData := []byte{52, 47, 132, 45, 79, 27, 252, 36, 11, 96, 86, 208, 70,
		84, 96, 101, 92, 123, 121, 107}
	plugin := "mysql_native_password"
	err := mc.handleAuthResult(authData, plugin)
	if err != ErrOldPassword {
		t.Errorf("expected ErrOldPassword, got %v", err)
	}
}

// Same to TestAuthSwitchOldPasswordNotAllowed, but use OldAuthSwitch request.
func TestOldAuthSwitchNotAllowed(t *testing.T) {
	conn, mc := newRWMockConn(2)

	// OldAuthSwitch request
	conn.data = []byte{1, 0, 0, 2, 0xfe}
	conn.maxReads = 1
	authData := []byte{10, 22, 126, 27, 36, 65, 218, 24, 32, 41, 75, 224, 42,
		84, 96, 101, 92, 123, 121, 107}
	plugin := "mysql_native_password"
	err := mc.handleAuthResult(authData, plugin)
	if err != ErrOldPassword {
		t.Errorf("expected ErrOldPassword, got %v", err)
	}
}

func TestAuthSwitchOldPassword(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.AllowOldPasswords = true
	mc.cfg.Passwd = "secret"

	// auth switch request
	conn.data = []byte{41, 0, 0, 2, 254, 109, 121, 115, 113, 108, 95, 111, 108,
		100, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 95, 84, 103, 43, 61,
		49, 123, 61, 91, 50, 40, 113, 35, 84, 96, 101, 92, 123, 121, 107, 0}

	// auth response
	conn.queuedReplies = [][]byte{{8, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0}}
	conn.maxReads = 2

	authData := []byte{12, 68, 232, 23, 73, 33, 161, 38, 43, 87, 58, 146, 17,
		84, 96, 101, 92, 123, 121, 107}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{9, 0, 0, 3, 86, 83, 83, 79, 74, 78, 65, 66, 0}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

// Same to TestAuthSwitchOldPassword, but use OldAuthSwitch request.
func TestOldAuthSwitch(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.AllowOldPasswords = true
	mc.cfg.Passwd = "secret"

	// OldAuthSwitch request
	conn.data = []byte{1, 0, 0, 2, 0xfe}

	// auth response
	conn.queuedReplies = [][]byte{{8, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0}}
	conn.maxReads = 2

	authData := []byte{30, 29, 210, 45, 82, 40, 201, 84, 84, 49, 87, 201, 64,
		84, 96, 101, 92, 123, 121, 107}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{9, 0, 0, 3, 86, 83, 83, 79, 74, 78, 65, 66, 0}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}
func TestAuthSwitchOldPasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.AllowOldPasswords = true
	mc.cfg.Passwd = ""

	// auth switch request
	conn.data = []byte{41, 0, 0, 2, 254, 109, 121, 115, 113, 108, 95, 111, 108,
		100, 95, 112, 97, 115, 115, 119, 111, 114, 100, 0, 95, 84, 103, 43, 61,
		49, 123, 61, 91, 50, 40, 113, 35, 84, 96, 101, 92, 123, 121, 107, 0}

	// auth response
	conn.queuedReplies = [][]byte{{8, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0}}
	conn.maxReads = 2

	authData := []byte{56, 68, 216, 27, 36, 67, 214, 74, 62, 18, 61, 255, 21,
		84, 96, 101, 92, 123, 121, 107}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{1, 0, 0, 3, 0}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

// Same to TestAuthSwitchOldPasswordEmpty, but use OldAuthSwitch request.
func TestOldAuthSwitchPasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.AllowOldPasswords = true
	mc.cfg.Passwd = ""

	// OldAuthSwitch request.
	conn.data = []byte{1, 0, 0, 2, 0xfe}

	// auth response
	conn.queuedReplies = [][]byte{{8, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0}}
	conn.maxReads = 2

	authData := []byte{88, 99, 139, 98, 34, 97, 125, 75, 40, 93, 18, 204, 69,
		84, 96, 101, 92, 123, 121, 107}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReply := []byte{1, 0, 0, 3, 0}
	if !bytes.Equal(conn.written, expectedReply) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchSHA256PasswordEmpty(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.Passwd = ""

	// auth switch request
	conn.data = []byte{38, 0, 0, 2, 254, 115, 104, 97, 50, 53, 54, 95, 112, 97,
		115, 115, 119, 111, 114, 100, 0, 78, 82, 62, 40, 100, 1, 59, 31, 44, 69,
		33, 112, 8, 81, 51, 96, 65, 82, 16, 114, 0}

	conn.queuedReplies = [][]byte{
		// OK
		{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 3

	authData := []byte{219, 59, 76, 17, 39, 63, 76, 136, 64, 163, 23, 18, 71,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReplyPrefix := []byte{
		// 1. Packet: Empty Password
		1, 0, 0, 3, 0,
	}
	if !bytes.HasPrefix(conn.written, expectedReplyPrefix) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchSHA256PasswordRSA(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.Passwd = "secret"

	// auth switch request
	conn.data = []byte{38, 0, 0, 2, 254, 115, 104, 97, 50, 53, 54, 95, 112, 97,
		115, 115, 119, 111, 114, 100, 0, 78, 82, 62, 40, 100, 1, 59, 31, 44, 69,
		33, 112, 8, 81, 51, 96, 65, 82, 16, 114, 0}

	conn.queuedReplies = [][]byte{
		// Pub Key Response
		append([]byte{byte(1 + len(testPubKey)), 1, 0, 4, 1}, testPubKey...),

		// OK
		{7, 0, 0, 6, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 3

	authData := []byte{124, 23, 98, 41, 59, 44, 16, 229, 12, 149, 11, 41, 52,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReplyPrefix := []byte{
		// 1. Packet: Pub Key Request
		1, 0, 0, 3, 1,

		// 2. Packet: Encrypted Password
		0, 1, 0, 5, // [changing bytes]
	}
	if !bytes.HasPrefix(conn.written, expectedReplyPrefix) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchSHA256PasswordRSAWithKey(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.Passwd = "secret"
	mc.cfg.pubKey = testPubKeyRSA

	// auth switch request
	conn.data = []byte{38, 0, 0, 2, 254, 115, 104, 97, 50, 53, 54, 95, 112, 97,
		115, 115, 119, 111, 114, 100, 0, 78, 82, 62, 40, 100, 1, 59, 31, 44, 69,
		33, 112, 8, 81, 51, 96, 65, 82, 16, 114, 0}

	conn.queuedReplies = [][]byte{
		// OK
		{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 2

	authData := []byte{120, 95, 12, 90, 57, 13, 20, 178, 66, 149, 44, 18, 72,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReplyPrefix := []byte{
		// 1. Packet: Encrypted Password
		0, 1, 0, 3, // [changing bytes]
	}
	if !bytes.HasPrefix(conn.written, expectedReplyPrefix) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}

func TestAuthSwitchSHA256PasswordSecure(t *testing.T) {
	conn, mc := newRWMockConn(2)
	mc.cfg.Passwd = "secret"

	// Hack to make the caching_sha2_password plugin believe that the connection
	// is secure
	mc.cfg.tls = &tls.Config{InsecureSkipVerify: true}

	// auth switch request
	conn.data = []byte{38, 0, 0, 2, 254, 115, 104, 97, 50, 53, 54, 95, 112, 97,
		115, 115, 119, 111, 114, 100, 0, 78, 82, 62, 40, 100, 1, 59, 31, 44, 69,
		33, 112, 8, 81, 51, 96, 65, 82, 16, 114, 0}

	conn.queuedReplies = [][]byte{
		// OK
		{7, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0},
	}
	conn.maxReads = 2

	authData := []byte{227, 48, 41, 61, 84, 62, 25, 161, 78, 102, 84, 37, 21,
		47, 43, 9, 41, 112, 67, 110}
	plugin := "mysql_native_password"

	if err := mc.handleAuthResult(authData, plugin); err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedReplyPrefix := []byte{
		// 1. Packet: Cleartext Password
		7, 0, 0, 3, 115, 101, 99, 114, 101, 116, 0,
	}
	if !bytes.Equal(conn.written, expectedReplyPrefix) {
		t.Errorf("got unexpected data: %v", conn.written)
	}
}
