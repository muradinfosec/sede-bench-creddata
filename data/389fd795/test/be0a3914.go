// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/crypto/openpgp"
)

func TestCommit_Marshal(t *testing.T) {
	testJSONMarshal(t, &Commit{}, "{}")

	u := &Commit{
		SHA: String("s"),
		Author: &CommitAuthor{
			Date:  &referenceTime,
			Name:  String("n"),
			Email: String("e"),
			Login: String("u"),
		},
		Committer: &CommitAuthor{
			Date:  &referenceTime,
			Name:  String("n"),
			Email: String("e"),
			Login: String("u"),
		},
		Message: String("m"),
		Tree: &Tree{
			SHA: String("s"),
			Entries: []*TreeEntry{{
				SHA:     String("s"),
				Path:    String("p"),
				Mode:    String("m"),
				Type:    String("t"),
				Size:    Int(1),
				Content: String("c"),
				URL:     String("u"),
			}},
			Truncated: Bool(false),
		},
		Parents: nil,
		Stats: &CommitStats{
			Additions: Int(1),
			Deletions: Int(1),
			Total:     Int(1),
		},
		HTMLURL: String("h"),
		URL:     String("u"),
		Verification: &SignatureVerification{
			Verified:  Bool(false),
			Reason:    String("r"),
			Signature: String("s"),
			Payload:   String("p"),
		},
		NodeID:       String("n"),
		CommentCount: Int(1),
		SigningKey:   &openpgp.Entity{},
	}

	want := `{
		"sha": "s",
		"author": {
			"date": ` + referenceTimeStr + `,
			"name": "n",
			"email": "e",
			"username": "u"
		},
		"committer": {
			"date": ` + referenceTimeStr + `,
			"name": "n",
			"email": "e",
			"username": "u"
		},
		"message": "m",
		"tree": {
			"sha": "s",
			"tree": [
				{
					"sha": "s",
					"path": "p",
					"mode": "m",
					"type": "t",
					"size": 1,
					"content": "c",
					"url": "u"
				}
			],
			"truncated": false
		},
		"stats": {
			"additions": 1,
			"deletions": 1,
			"total": 1
		},
		"html_url": "h",
		"url": "u",
		"verification": {
			"verified": false,
			"reason": "r",
			"signature": "s",
			"payload": "p"
		},
		"node_id": "n",
		"comment_count": 1
	}`

	testJSONMarshal(t, u, want)
}

func TestGitService_GetCommit(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/git/commits/s", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"sha":"s","message":"Commit Message.","author":{"name":"n"}}`)
	})

	ctx := context.Background()
	commit, _, err := client.Git.GetCommit(ctx, "o", "r", "s")
	if err != nil {
		t.Errorf("Git.GetCommit returned error: %v", err)
	}

	want := &Commit{SHA: String("s"), Message: String("Commit Message."), Author: &CommitAuthor{Name: String("n")}}
	if !cmp.Equal(commit, want) {
		t.Errorf("Git.GetCommit returned %+v, want %+v", commit, want)
	}

	const methodName = "GetCommit"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Git.GetCommit(ctx, "\n", "\n", "\n")
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Git.GetCommit(ctx, "o", "r", "s")
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestGitService_GetCommit_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, _, err := client.Git.GetCommit(ctx, "%", "%", "%")
	testURLParseError(t, err)
}

func TestGitService_CreateCommit(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Commit{
		Message: String("Commit Message."),
		Tree:    &Tree{SHA: String("t")},
		Parents: []*Commit{{SHA: String("p")}},
	}

	mux.HandleFunc("/repos/o/r/git/commits", func(w http.ResponseWriter, r *http.Request) {
		v := new(createCommit)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")

		want := &createCommit{
			Message: input.Message,
			Tree:    String("t"),
			Parents: []string{"p"},
		}
		if !cmp.Equal(v, want) {
			t.Errorf("Request body = %+v, want %+v", v, want)
		}
		fmt.Fprint(w, `{"sha":"s"}`)
	})

	ctx := context.Background()
	commit, _, err := client.Git.CreateCommit(ctx, "o", "r", input)
	if err != nil {
		t.Errorf("Git.CreateCommit returned error: %v", err)
	}

	want := &Commit{SHA: String("s")}
	if !cmp.Equal(commit, want) {
		t.Errorf("Git.CreateCommit returned %+v, want %+v", commit, want)
	}

	const methodName = "CreateCommit"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Git.CreateCommit(ctx, "\n", "\n", input)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Git.CreateCommit(ctx, "o", "r", input)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestGitService_CreateSignedCommit(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	signature := "----- BEGIN PGP SIGNATURE -----\n\naaaa\naaaa\n----- END PGP SIGNATURE -----"

	input := &Commit{
		Message: String("Commit Message."),
		Tree:    &Tree{SHA: String("t")},
		Parents: []*Commit{{SHA: String("p")}},
		Verification: &SignatureVerification{
			Signature: String(signature),
		},
	}

	mux.HandleFunc("/repos/o/r/git/commits", func(w http.ResponseWriter, r *http.Request) {
		v := new(createCommit)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")

		want := &createCommit{
			Message:   input.Message,
			Tree:      String("t"),
			Parents:   []string{"p"},
			Signature: String(signature),
		}
		if !cmp.Equal(v, want) {
			t.Errorf("Request body = %+v, want %+v", v, want)
		}
		fmt.Fprint(w, `{"sha":"commitSha"}`)
	})

	ctx := context.Background()
	commit, _, err := client.Git.CreateCommit(ctx, "o", "r", input)
	if err != nil {
		t.Errorf("Git.CreateCommit returned error: %v", err)
	}

	want := &Commit{SHA: String("commitSha")}
	if !cmp.Equal(commit, want) {
		t.Errorf("Git.CreateCommit returned %+v, want %+v", commit, want)
	}

	const methodName = "CreateCommit"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Git.CreateCommit(ctx, "\n", "\n", input)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Git.CreateCommit(ctx, "o", "r", input)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestGitService_CreateSignedCommitWithInvalidParams(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	input := &Commit{
		SigningKey: &openpgp.Entity{},
	}

	ctx := context.Background()
	_, _, err := client.Git.CreateCommit(ctx, "o", "r", input)
	if err == nil {
		t.Errorf("Expected error to be returned because invalid params was passed")
	}
}

func TestGitService_CreateSignedCommitWithKey(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	s := strings.NewReader(testGPGKey)
	keyring, err := openpgp.ReadArmoredKeyRing(s)
	if err != nil {
		t.Errorf("Error reading keyring: %+v", err)
	}

	date, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", "Thu May 04 00:03:43 2017 +0200")
	author := CommitAuthor{
		Name:  String("go-github"),
		Email: String("go-github@github.com"),
		Date:  &date,
	}
	input := &Commit{
		Message:    String("Commit Message."),
		Tree:       &Tree{SHA: String("t")},
		Parents:    []*Commit{{SHA: String("p")}},
		SigningKey: keyring[0],
		Author:     &author,
	}

	messageReader := strings.NewReader(`tree t
parent p
author go-github <go-github@github.com> 1493849023 +0200
committer go-github <go-github@github.com> 1493849023 +0200

Commit Message.`)

	mux.HandleFunc("/repos/o/r/git/commits", func(w http.ResponseWriter, r *http.Request) {
		v := new(createCommit)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")

		want := &createCommit{
			Message: input.Message,
			Tree:    String("t"),
			Parents: []string{"p"},
			Author:  &author,
		}

		sigReader := strings.NewReader(*v.Signature)
		signer, err := openpgp.CheckArmoredDetachedSignature(keyring, messageReader, sigReader)
		if err != nil {
			t.Errorf("Error verifying signature: %+v", err)
		}
		if signer.Identities["go-github <go-github@github.com>"].Name != "go-github <go-github@github.com>" {
			t.Errorf("Signer is incorrect. got: %+v, want %+v", signer.Identities["go-github <go-github@github.com>"].Name, "go-github <go-github@github.com>")
		}
		// Nullify Signature since we checked it above
		v.Signature = nil
		if !cmp.Equal(v, want) {
			t.Errorf("Request body = %+v, want %+v", v, want)
		}
		fmt.Fprint(w, `{"sha":"commitSha"}`)
	})

	ctx := context.Background()
	commit, _, err := client.Git.CreateCommit(ctx, "o", "r", input)
	if err != nil {
		t.Errorf("Git.CreateCommit returned error: %v", err)
	}

	want := &Commit{SHA: String("commitSha")}
	if !cmp.Equal(commit, want) {
		t.Errorf("Git.CreateCommit returned %+v, want %+v", commit, want)
	}
}

func TestGitService_createSignature_nilSigningKey(t *testing.T) {
	a := &createCommit{
		Message: String("Commit Message."),
		Tree:    String("t"),
		Parents: []string{"p"},
	}

	_, err := createSignature(nil, a)

	if err == nil {
		t.Errorf("Expected error to be returned because no author was passed")
	}
}

func TestGitService_createSignature_nilCommit(t *testing.T) {
	_, err := createSignature(&openpgp.Entity{}, nil)

	if err == nil {
		t.Errorf("Expected error to be returned because no author was passed")
	}
}

func TestGitService_createSignature_noAuthor(t *testing.T) {
	a := &createCommit{
		Message: String("Commit Message."),
		Tree:    String("t"),
		Parents: []string{"p"},
	}

	_, err := createSignature(&openpgp.Entity{}, a)

	if err == nil {
		t.Errorf("Expected error to be returned because no author was passed")
	}
}

func TestGitService_createSignature_invalidKey(t *testing.T) {
	date, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", "Thu May 04 00:03:43 2017 +0200")

	_, err := createSignature(&openpgp.Entity{}, &createCommit{
		Message: String("Commit Message."),
		Tree:    String("t"),
		Parents: []string{"p"},
		Author: &CommitAuthor{
			Name:  String("go-github"),
			Email: String("go-github@github.com"),
			Date:  &date,
		},
	})

	if err == nil {
		t.Errorf("Expected error to be returned due to invalid key")
	}
}

func TestGitService_createSignatureMessage_nilCommit(t *testing.T) {
	_, err := createSignatureMessage(nil)
	if err == nil {
		t.Errorf("Expected error to be returned due to nil key")
	}
}

func TestGitService_createSignatureMessage_nilMessage(t *testing.T) {
	date, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", "Thu May 04 00:03:43 2017 +0200")

	_, err := createSignatureMessage(&createCommit{
		Message: nil,
		Parents: []string{"p"},
		Author: &CommitAuthor{
			Name:  String("go-github"),
			Email: String("go-github@github.com"),
			Date:  &date,
		},
	})
	if err == nil {
		t.Errorf("Expected error to be returned due to nil key")
	}
}

func TestGitService_createSignatureMessage_emptyMessage(t *testing.T) {
	date, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", "Thu May 04 00:03:43 2017 +0200")
	emptyString := ""
	_, err := createSignatureMessage(&createCommit{
		Message: &emptyString,
		Parents: []string{"p"},
		Author: &CommitAuthor{
			Name:  String("go-github"),
			Email: String("go-github@github.com"),
			Date:  &date,
		},
	})
	if err == nil {
		t.Errorf("Expected error to be returned due to nil key")
	}
}

func TestGitService_createSignatureMessage_nilAuthor(t *testing.T) {
	_, err := createSignatureMessage(&createCommit{
		Message: String("Commit Message."),
		Parents: []string{"p"},
		Author:  nil,
	})
	if err == nil {
		t.Errorf("Expected error to be returned due to nil key")
	}
}

func TestGitService_createSignatureMessage_withoutTree(t *testing.T) {
	date, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", "Thu May 04 00:03:43 2017 +0200")

	msg, _ := createSignatureMessage(&createCommit{
		Message: String("Commit Message."),
		Parents: []string{"p"},
		Author: &CommitAuthor{
			Name:  String("go-github"),
			Email: String("go-github@github.com"),
			Date:  &date,
		},
	})
	expected := `parent p
author go-github <go-github@github.com> 1493849023 +0200
committer go-github <go-github@github.com> 1493849023 +0200

Commit Message.`
	if msg != expected {
		t.Errorf("Returned message incorrect. returned %s, want %s", msg, expected)
	}
}

func TestGitService_createSignatureMessage_withoutCommitter(t *testing.T) {
	date, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", "Thu May 04 00:03:43 2017 +0200")

	msg, _ := createSignatureMessage(&createCommit{
		Message: String("Commit Message."),
		Parents: []string{"p"},
		Author: &CommitAuthor{
			Name:  String("go-github"),
			Email: String("go-github@github.com"),
			Date:  &date,
		},
		Committer: &CommitAuthor{
			Name:  String("foo"),
			Email: String("foo@bar.com"),
			Date:  &date,
		},
	})
	expected := `parent p
author go-github <go-github@github.com> 1493849023 +0200
committer foo <foo@bar.com> 1493849023 +0200

Commit Message.`
	if msg != expected {
		t.Errorf("Returned message incorrect. returned %s, want %s", msg, expected)
	}
}

func TestGitService_CreateCommit_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, _, err := client.Git.CreateCommit(ctx, "%", "%", &Commit{})
	testURLParseError(t, err)
}

const testGPGKey = `
-----BEGIN PGP PRIVATE KEY BLOCK-----

lQOYBFyi1qYBCAD3EPfLJzIt4qkAceUKkhdvfaIvOsBwXbfr5sSu/lkMqL0Wq47+
Uo+IHSWJ4Osg8FaQ8uiobK1DcxxSceZF2kgWOGP3kTUpYp0bcWe8shsC15PJ/EuL
6L34RvtxCHuvf35b7HLuJEbkr/z7bVveByasTTFSSbDxwhNYUmUJ0ZqdbOqv0k0e
GENq3AUgX/qVwNcgb0oY6Wa8gqiLfE4VtmxEVHqYF99FR3I/SCcVAXfWtA7D4QKv
FFWRwMFomfnkAtyI78pu0F9P4EM+vME8pmP2923GguWMxIDvkmHg+CpyTA687ab+
xvltGbh6OpWiTv91U8OX2502CxYso1xRQ35gnFupeXhAR/1xUC27/dAOjycTntpq
xgEAQWxGsgG5C7Vzsch5vss7xBBumYUIp0X/0avcCOJ+klAIxE37LIUrUvKOKU80
b53Z1OlSrVOWWDP5vZLQheBB9wZXmwhJ38f7v6S0Rj7+EyRB0JN2eKb5GhgM6Qxi
u+6t65rO0+Oz+MTSSASDw2m8kweHItuJGGQ/Caa1BCrBKO4RVb2ipYGWwQcXrnCF
cBpHQaag7bD3by1SVpKUm/lPhukjXcvPKeMN/aVs+hWTvkqSVt1r2ynxZ13ewn+H
TY13mgEbiUUdzU0Q3Xtn+ekUaQh8IQRGSW/jfCb5grbfWPtcDCj67RS3KfDgy/s5
w16OTrv02+SgFd9k6EIsaT/a+zSq4LqDqHtAhF1hgdVdOS1aj7y+hVzgHhEkY31c
6qaB+elUjP2C04Nxdaa6IhBGpaIu/dRARB8S6Db3e80Rd+Nyc6ueCehcNLtzo3Z+
Z4EIpg9hPBEFG1yNNLvsTOb4vhxc5idQkLITqJug7AYX09lyUDFj/vnly7tuE3Wp
RQR2VkD2ov0hQdl0IImXBXO4el8R5QvU+Msi3vcByf08jxI/BNbNUA/sYwGy6zop
fJvATlWKIqi5uoK3VuAoNF3emIWynpIlfBlbOXF7h/rD+e8bg7R8WhmNqQtoakq/
Gojzk0R2WQEKlTdVQIB+8tSPBb4NAjLwwcTrqOCP8Otb4CwVnGEDzlqmaT+5z+2F
exjAayU9u/bN9KDrsjZLCwqI8atRMAF9FoNxKNWvkoiJvfItlndp/s/EhrmvSLgM
FY4pAC//Q8dziqoMPx/6+yMOdDUTIXGTmmwW8Md3frYi1C3SFYNqLp6d75DKIlv6
pHE6EhIHKGpKmnO5ywrxdnPKSKj9rea1T26pv6g9mVEhreRFDtl9Lv3MF03+FUwh
IVzXBYf+TgbUvw1Txckrykgwbx1uqTSLKhL+YeKnBxDZ2bXyTQbnslBXPdqAslOQ
hLEETbYGLdZMcetSnOaEClXaT0YSZTWgApfyQvO+OAbVEtgX3jCiDB88CtgGyfdw
a3VvRh3fp3c+F1an6vga0FgzaRy/0ZeU6oFxut9llRACOTphqCzwncXhJLH2SmRx
VULiFYkHE7plZAEDynqDH0QaTfMDcldmHYr8q3l3rZsKr8NUwIucpMo5VjBeRGQ9
UbEDoFqpsDlHEgvpPjZW2n3WO53NqjjZRBnhFAkSgy3CVqjFd/0Cdl0g/vrOksVd
XltV7yNN3fpK3u10ozTuQCXyhKZUzS6hAR3AWESNKiKqb4UEtFjDuL1kr6lM/dPC
8Md7xwa4WZ4/zYA93BqnErSENRVTkZ/iqz+A4VS/+V5dFcT+uAZZlPiO8EmboSt1
XcFVJ94YzgSivmDIMNOzOoL46xSSTp0cylIPRoHlpsqkQT8hbMa5gQjYuXBrW0yH
7fzLl+DLtG6nWNtNzb0IAm1cfNqUuOcQg+akzN9Fan2c3Zfp8rIbEoTC0w0sXlyd
VWuw+/SckIyvdutU0fkhQ8W2ljutUY133RlJ14q8qgTy6BFtv/K+kDWiMltcd5g1
qU9wtIMbVYPYbfTuVyv1YXcTGcO+Am+EBtBwI3q0fkpMQKWscFvIGTVjefd9sgxN
ic7BjZ3rzC2HuzDEjMNZlbWz8QI/ebR7WKGKz01eaOLjPeCdfBHcw50ggd2z7+x7
KwdpDbGi0aHogREEh/ha9piSG15qhnFA1qlhRia6XcOKijzDxJK5wnGN9cgt+Ac8
GqTZ8GkMqUD85GDHTzFq1osOetsLh1iQhmTCnt7lbcMY5T3RlxUtuQ70G8hIFc4O
BtM4C0p6q2EEMgf5DrSlWG3KrBoucoPezYtG4nGblEiyjrOSXf+UgYSA9vItrk0O
4ltnHnocefD4kDPm9Him4V9AYRQ4MCu0busDhho4LwWxoOuEfBPTEgDaHA8GVzEm
n7DmaGHjO8xmGDceVcNBy1zT29h/2PBuSz5R5L1xSyzfu0LIx5kZRtRGPxa7Pkt9
2x3lZdL3zWI2ixE1OWxqZ+56LHxo8EVO599l6Uz/xetoaVCyHapQ7xULvzuUmgRw
hnkeF9dwAVYmor8YDNzmR+0vuJhBxBkjJarKCIRfbn6bCQ+/tjDIpXsTc70Me6wj
Dx5qQZb5FWGPgIfT8Hv0BWkLsr3tKcRTl64IgyNS+wUfApuDtcTI783YUSF3nSau
dcIUGA3ZUYqLupXRlHSqbAoSqRiosW0WR1ttUg9lW8Z/hxGrGlXqimFwEmO5V/2/
2OjEjqgirpcDkLZa+JxFYm2idMh5Bzme5RH4DgYFEyq+eouPkwFHFKn9HoQlvxit
MzFxze+8aFla4avfQTMdBxhmipw8I90N4GdL0qvDTo/higLE1+4V5zXsJH/6RqaY
E6OYYonyBi9DGAXTp/mMU6rijtJdkUUAJxIhpTaszHRYweWETi40lMjHBQASiQbU
wABTXSL7J0Ae7pqGKg+XfvdG5xOBu4ZFHNxYPlfPsa3t9koDy77fid6k/oqyLvru
c9y6qpb7qmO71NCgY/YeS3Vheoemgvm3TesO5LAy50eaCnBJsRBFQZ36zzHw4k4Q
UqeZuoPsn3GQnyHxwdxIVmA4EgrLpTlKeAqsVg0PTJUPHdIMMq1EscoLHry+TFdF
ltPH5DpMGbCSZIHHWMVWmIzduuneNFR+oXSBqbi/Bg2yCxbswfXu6gfBxbU+TiVz
3QCuYVFUh6tS52dvXsPdFnkFCK770SB9LoW+oIDGPXJ1XTaNw2BtV1xLYnf+rQsO
r6DDGU4lEMApvhakIWQFU+3Lm66FAqTpaXejq/7RNZFryX38gmC+k8f4LQwqj0sJ
ljw8Vjh8hfSOOabjAP9ja2vUWvPiVOfIleSZVBEuEXm0TSOJs5oby8uUx3GkCVyr
MrfwmYKNcAU+gXdyJZJVCnTYmwk+ZoLBbZOiBj5906u6f9daqPVcz1V2GYOGDuLi
D1/4PW6SeABwOsFYZUdOL+SAxUGMjmcdpffznQtYld+kY8/iOpp95jsQyXy9be==
=VFSn
-----END PGP PRIVATE KEY BLOCK-----`
