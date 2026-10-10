/*
copyright 2020 The Vitess Authors.

licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at

    http://www.apache.org/licenses/license-2.0

unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyHashedMysqlNativePassword(t *testing.T) {
	salt := []byte{86, 53, 93, 237, 10, 40, 56, 94, 44, 13, 193, 32, 21, 69, 5, 55, 43, 9, 42, 94}
	password := "secret"

	// Double SHA1 of "secret"
	passwordHash := []byte{0x38, 0x81, 0x21, 0x9d, 0x08, 0x7d, 0xd9, 0xc6, 0x34, 0x37, 0x3f, 0xd3,
		0x3d, 0xfa, 0x33, 0xa2, 0xcb, 0x6b, 0xfc, 0x6c, 0x52, 0x0b, 0x64, 0xb8,
		0xbb, 0x60, 0xef, 0x2c, 0xeb, 0x53, 0x4a, 0xe7}

	reply := ScrambleCachingSha2Password(salt, []byte(password))

	assert.True(t, VerifyHashedCachingSha2Password(reply, salt, passwordHash), "password hash mismatch")

	passwordHash[0] = 0x00
	assert.False(t, VerifyHashedCachingSha2Password(reply, salt, passwordHash), "password hash match")
}

func TestVerifyHashedCachingSha2Password(t *testing.T) {
	salt := []byte{13, 20, 78, 112, 18, 98, 66, 72, 73, 81, 247, 76, 19, 82, 5, 27, 58, 1, 99, 61}
	password := "secret"

	// Double SHA1 of "secret"
	passwordHash := []byte{0x14, 0xe6, 0x55, 0x67, 0xab, 0xdb, 0x51, 0x35, 0xd0, 0xcf, 0xd9, 0xa7,
		0x0b, 0x30, 0x32, 0xc1, 0x79, 0xa4, 0x9e, 0xe7}

	reply := ScrambleMysqlNativePassword(salt, []byte(password))
	assert.True(t, VerifyHashedMysqlNativePassword(reply, salt, passwordHash), "password hash mismatch")

	passwordHash[0] = 0x00
	assert.False(t, VerifyHashedMysqlNativePassword(reply, salt, passwordHash), "password hash match")
}
