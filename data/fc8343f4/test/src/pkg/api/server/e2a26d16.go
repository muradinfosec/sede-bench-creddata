/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	mathrand "math/rand"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	utilrand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/uuid"
	auditinternal "k8s.io/apiserver/pkg/apis/audit"
	"k8s.io/apiserver/pkg/audit"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/utils/clock"
	testingclock "k8s.io/utils/clock/testing"
)

func TestCachedTokenAuthenticator(t *testing.T) {
	var (
		calledWithToken []string

		resultUsers map[string]user.Info
		resultOk    bool
		resultErr   error
	)
	fakeAuth := authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
		calledWithToken = append(calledWithToken, token)
		return &authenticator.Response{User: resultUsers[token]}, resultOk, resultErr
	})
	fakeClock := testingclock.NewFakeClock(time.Now())

	a := newWithClock(fakeAuth, true, time.Minute, 0, fakeClock)

	calledWithToken, resultUsers, resultOk, resultErr = []string{}, nil, false, nil
	a.AuthenticateToken(context.Background(), "bad1")
	a.AuthenticateToken(context.Background(), "bad2")
	a.AuthenticateToken(context.Background(), "bad3")
	fakeClock.Step(2 * time.Microsecond)
	a.AuthenticateToken(context.Background(), "bad1")
	a.AuthenticateToken(context.Background(), "bad2")
	a.AuthenticateToken(context.Background(), "bad3")
	fakeClock.Step(2 * time.Microsecond)
	if !reflect.DeepEqual(calledWithToken, []string{"bad1", "bad2", "bad3", "bad1", "bad2", "bad3"}) {
		t.Errorf("Expected failing calls to not stay in the cache, got %v", calledWithToken)
	}

	// reset calls, make the backend return success for three user tokens
	calledWithToken = []string{}
	resultUsers, resultOk, resultErr = map[string]user.Info{}, true, nil
	resultUsers["usertoken1"] = &user.DefaultInfo{Name: "user1"}
	resultUsers["usertoken2"] = &user.DefaultInfo{Name: "user2"}
	resultUsers["usertoken3"] = &user.DefaultInfo{Name: "user3"}

	// populate cache
	if resp, ok, err := a.AuthenticateToken(context.Background(), "usertoken1"); err != nil || !ok || resp.User.GetName() != "user1" {
		t.Errorf("Expected user1")
	}
	if resp, ok, err := a.AuthenticateToken(context.Background(), "usertoken2"); err != nil || !ok || resp.User.GetName() != "user2" {
		t.Errorf("Expected user2")
	}
	if resp, ok, err := a.AuthenticateToken(context.Background(), "usertoken3"); err != nil || !ok || resp.User.GetName() != "user3" {
		t.Errorf("Expected user3")
	}
	if !reflect.DeepEqual(calledWithToken, []string{"usertoken1", "usertoken2", "usertoken3"}) {
		t.Errorf("Expected token calls, got %v", calledWithToken)
	}

	// reset calls, make the backend return failures
	calledWithToken = []string{}
	resultUsers, resultOk, resultErr = nil, false, nil

	// authenticate calls still succeed and backend is not hit
	if resp, ok, err := a.AuthenticateToken(context.Background(), "usertoken1"); err != nil || !ok || resp.User.GetName() != "user1" {
		t.Errorf("Expected user1")
	}
	if resp, ok, err := a.AuthenticateToken(context.Background(), "usertoken2"); err != nil || !ok || resp.User.GetName() != "user2" {
		t.Errorf("Expected user2")
	}
	if resp, ok, err := a.AuthenticateToken(context.Background(), "usertoken3"); err != nil || !ok || resp.User.GetName() != "user3" {
		t.Errorf("Expected user3")
	}
	if !reflect.DeepEqual(calledWithToken, []string{}) {
		t.Errorf("Expected no token calls, got %v", calledWithToken)
	}

	// skip forward in time
	fakeClock.Step(2 * time.Minute)

	// backend is consulted again and fails
	a.AuthenticateToken(context.Background(), "usertoken1")
	a.AuthenticateToken(context.Background(), "usertoken2")
	a.AuthenticateToken(context.Background(), "usertoken3")
	if !reflect.DeepEqual(calledWithToken, []string{"usertoken1", "usertoken2", "usertoken3"}) {
		t.Errorf("Expected token calls, got %v", calledWithToken)
	}
}

func TestCachedTokenAuthenticatorWithAudiences(t *testing.T) {
	resultUsers := make(map[string]user.Info)
	fakeAuth := authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
		auds, _ := authenticator.AudiencesFrom(ctx)
		return &authenticator.Response{User: resultUsers[auds[0]+token]}, true, nil
	})
	fakeClock := testingclock.NewFakeClock(time.Now())

	a := newWithClock(fakeAuth, true, time.Minute, 0, fakeClock)

	resultUsers["audAusertoken1"] = &user.DefaultInfo{Name: "user1"}
	resultUsers["audBusertoken1"] = &user.DefaultInfo{Name: "user1-different"}

	if u, ok, _ := a.AuthenticateToken(authenticator.WithAudiences(context.Background(), []string{"audA"}), "usertoken1"); !ok || u.User.GetName() != "user1" {
		t.Errorf("Expected user1")
	}
	if u, ok, _ := a.AuthenticateToken(authenticator.WithAudiences(context.Background(), []string{"audB"}), "usertoken1"); !ok || u.User.GetName() != "user1-different" {
		t.Errorf("Expected user1-different")
	}
}

var bKey string

// use a realistic token for benchmarking
const jwtToken = `eyJhbGciOiJpYzUxNSIsImtpZCI6IiJ9.eyJpc3MiOiJGS1pJeHNkUmJaL05PWHNRUHdMWE1RU2lQIiwiV3V0SWFrbm9kay5QaS9zaXVFZlFvRG9iaVRsYi9menVOdFRGTlgiOiJNZGJDdExTU04tRnNlIiwiUmpNeHRCaXFpUS5OeC9GWEJMVnBub0hiT3h5ci9hd2ZTV0YuTlVUdiI6IlJPRS1hV2xlYy1UM2d6QyIsIlRVQ1pZQXZHWlQuZFgvVWxobVRxTGNaUWxmZkcvcmVhSk9CYy1CeUR1d1NHLm5yc3YiOiJUT1YiLCJTakdTanJIU25ULlZLL1hvbU9zY3dMUVRyQURyL1huT3BuWGotR1l6dnJLVS5sY0UiOiIzdjg2SzZhVC1aNUI5LTY2czQtYzJxNi02NHQ4RjhHNTkxMnEiLCJzdWIiOiJ6TnNVTUU6RWxXS2J5ZE15R014Zlc6YVppRVBqZmt1LVNTaTpoaUMifQ.OPr1geIVmsc5HT9pRsGFfD4tJV29UW3FScX420smQfgAY0Cc_6TG2aggMRabIa715t9wZb9_1ACLo0Mv9oVoFF0NBPRLjijBi6K-5QTSWvipx-YB00zgTAE9X_-w3EbwnwaQgdefYvLZ7EWP6aWsQDOWFHfGia-EKQO4d2LBRa7FDNRT4gdh_9CwnW5DUjVzAz_wFG8xHcQ1nXGP8rcxO8XflrZizsgWFvWsetA9s2bwHvCfiTITjI0rZyI1K9H1izBAsDaD3ABurFUy6idCzWvJUiarvUfU8g8jYX1AxnYcqbdkc8DLB7wbSVoQs5lubzQgH3LQSyWeDMqytkMSaF`

func BenchmarkKeyFunc(b *testing.B) {
	randomCacheKey := make([]byte, 32)
	if _, err := rand.Read(randomCacheKey); err != nil {
		b.Fatal(err) // rand should never fail
	}
	hashPool := &sync.Pool{
		New: func() interface{} {
			return hmac.New(sha256.New, randomCacheKey)
		},
	}

	// use realistic audiences for benchmarking
	auds := []string{"3acf74e5-a62b-315f-3e75-d949cfcca373", "b70af713-afad-7aef-2895-984af5240957", "17b9a8b2-5706-2448-f79f-6576b62a3191"}

	b.Run("has audiences", func(b *testing.B) {
		var key string
		for n := 0; n < b.N; n++ {
			key = keyFunc(hashPool, auds, jwtToken)
		}
		bKey = key
	})

	b.Run("nil audiences", func(b *testing.B) {
		var key string
		for n := 0; n < b.N; n++ {
			key = keyFunc(hashPool, nil, jwtToken)
		}
		bKey = key
	})
}

func TestSharedLookup(t *testing.T) {
	var chewie = &authenticator.Response{User: &user.DefaultInfo{Name: "chewbacca"}}

	t.Run("actually shared", func(t *testing.T) {
		var lookups uint32
		c := make(chan struct{})
		a := New(authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
			<-c
			atomic.AddUint32(&lookups, 1)
			return chewie, true, nil
		}), true, time.Minute, 0)

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				a.AuthenticateToken(context.Background(), "")
			}()
		}

		// no good way to make sure that all the callers are queued so we sleep.
		time.Sleep(1 * time.Second)
		close(c)
		wg.Wait()

		if lookups > 3 {
			t.Fatalf("unexpected number of lookups: got=%d, wanted less than 3", lookups)
		}
	})

	t.Run("first caller bails, second caller gets result", func(t *testing.T) {
		c := make(chan struct{})
		a := New(authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
			<-c
			return chewie, true, nil
		}), true, time.Minute, 0)

		var wg sync.WaitGroup
		wg.Add(2)

		ctx1, cancel1 := context.WithCancel(context.Background())
		go func() {
			defer wg.Done()
			a.AuthenticateToken(ctx1, "")
		}()

		ctx2 := context.Background()

		var (
			resp *authenticator.Response
			ok   bool
			err  error
		)
		go func() {
			defer wg.Done()
			resp, ok, err = a.AuthenticateToken(ctx2, "")
		}()

		time.Sleep(1 * time.Second)
		cancel1()
		close(c)
		wg.Wait()

		if want := chewie; !cmp.Equal(resp, want) {
			t.Errorf("Unexpected diff: %v", cmp.Diff(resp, want))
		}
		if !ok {
			t.Errorf("Expected ok response")
		}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("lookup panics", func(t *testing.T) {
		a := New(authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
			panic("uh oh")
		}), true, time.Minute, 0)

		_, _, err := a.AuthenticateToken(context.Background(), "")
		if err != errAuthnCrash {
			t.Errorf("expected error: %v", err)
		}
	})

	t.Run("audiences are forwarded", func(t *testing.T) {
		ctx := authenticator.WithAudiences(context.Background(), []string{"a"})
		a := New(authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
			auds, _ := authenticator.AudiencesFrom(ctx)
			if got, want := auds, []string{"a"}; cmp.Equal(got, want) {
				t.Fatalf("unexpeced audiences: %v", cmp.Diff(got, want))
			}
			return nil, false, nil
		}), true, time.Minute, 0)

		a.AuthenticateToken(ctx, "")
	})
}

func TestCachedAuditAnnotations(t *testing.T) {
	snorlax := &authenticator.Response{User: &user.DefaultInfo{Name: "snorlax"}}

	t.Run("annotations from cache", func(t *testing.T) {
		var lookups uint32
		c := make(chan struct{})
		a := New(authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
			<-c
			atomic.AddUint32(&lookups, 1)
			audit.AddAuditAnnotation(ctx, "snorlax", "rocks")
			audit.AddAuditAnnotation(ctx, "pandas", "are amazing")
			return snorlax, true, nil
		}), false, time.Minute, 0)

		allAnnotations := make(chan map[string]string, 10)
		defer close(allAnnotations)

		var wg sync.WaitGroup
		for i := 0; i < cap(allAnnotations); i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				ctx := withAudit(context.Background())
				_, _, _ = a.AuthenticateToken(ctx, "token")

				allAnnotations <- audit.AuditEventFrom(ctx).Annotations
			}()
		}

		// no good way to make sure that all the callers are queued so we sleep.
		time.Sleep(1 * time.Second)
		close(c)
		wg.Wait()

		want := map[string]string{"snorlax": "rocks", "pandas": "are amazing"}
		for i := 0; i < cap(allAnnotations); i++ {
			annotations := <-allAnnotations
			if diff := cmp.Diff(want, annotations); diff != "" {
				t.Errorf("%d: unexpected annotations (-want +got): %s", i, diff)
			}
		}

		if queued := len(allAnnotations); queued != 0 {
			t.Errorf("expected all annoations to be processed: %d", queued)
		}

		if lookups > 3 {
			t.Errorf("unexpected number of lookups: got=%d, wanted less than 3", lookups)
		}
	})

	t.Run("annotations do not change during cache TTL", func(t *testing.T) {
		a := New(authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
			audit.AddAuditAnnotation(ctx, "timestamp", time.Now().String())
			return snorlax, true, nil
		}), false, time.Minute, 0)

		allAnnotations := make([]map[string]string, 0, 10)

		for i := 0; i < cap(allAnnotations); i++ {
			ctx := withAudit(context.Background())
			_, _, _ = a.AuthenticateToken(ctx, "token")
			allAnnotations = append(allAnnotations, audit.AuditEventFrom(ctx).Annotations)
		}

		if len(allAnnotations) != cap(allAnnotations) {
			t.Errorf("failed to process all annotations")
		}

		want := allAnnotations[0]
		if ok := len(want) == 1 && len(want["timestamp"]) > 0; !ok {
			t.Errorf("invalid annotations: %v", want)
		}

		for i, annotations := range allAnnotations[1:] {
			if diff := cmp.Diff(want, annotations); diff != "" {
				t.Errorf("%d: unexpected annotations (-want +got): %s", i, diff)
			}
		}
	})

	t.Run("different tokens can have different annotations", func(t *testing.T) {
		a := New(authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
			audit.AddAuditAnnotation(ctx, "timestamp", time.Now().String())
			return snorlax, true, nil
		}), false, time.Minute, 0)

		ctx1 := withAudit(context.Background())
		_, _, _ = a.AuthenticateToken(ctx1, "token1")
		annotations1 := audit.AuditEventFrom(ctx1).Annotations

		// guarantee different now times
		time.Sleep(time.Second)

		ctx2 := withAudit(context.Background())
		_, _, _ = a.AuthenticateToken(ctx2, "token2")
		annotations2 := audit.AuditEventFrom(ctx2).Annotations

		if ok := len(annotations1) == 1 && len(annotations1["timestamp"]) > 0; !ok {
			t.Errorf("invalid annotations 1: %v", annotations1)
		}
		if ok := len(annotations2) == 1 && len(annotations2["timestamp"]) > 0; !ok {
			t.Errorf("invalid annotations 2: %v", annotations2)
		}

		if annotations1["timestamp"] == annotations2["timestamp"] {
			t.Errorf("annotations should have different timestamp value: %v", annotations1)
		}
	})
}

func BenchmarkCachedTokenAuthenticator(b *testing.B) {
	tokenCount := []int{100, 500, 2500, 12500, 62500}
	threadCount := []int{1, 16, 256}
	for _, tokens := range tokenCount {
		for _, threads := range threadCount {
			newSingleBenchmark(tokens, threads).run(b)
		}
	}
}

func newSingleBenchmark(tokens, threads int) *singleBenchmark {
	s := &singleBenchmark{
		threadCount: threads,
		tokenCount:  tokens,
	}
	s.makeTokens()
	return s
}

// singleBenchmark collects all the state needed to run a benchmark. The
// question this benchmark answers is, "what's the average latency added by the
// cache for N concurrent tokens?"
//
// Given the size of the key range constructed by this test, the default go
// benchtime of 1 second is often inadequate to test caching and expiration
// behavior. A benchtime of 10 to 30 seconds is adequate to stress these
// code paths.
type singleBenchmark struct {
	threadCount int
	// These token.* variables are set by makeTokens()
	tokenCount int
	// pre-computed response for a token
	tokenToResponse map[string]*cacheRecord
	// include audiences for some
	tokenToAuds map[string]authenticator.Audiences
	// a list makes it easy to select a random one
	tokens []string
}

func (s *singleBenchmark) makeTokens() {
	s.tokenToResponse = map[string]*cacheRecord{}
	s.tokenToAuds = map[string]authenticator.Audiences{}
	s.tokens = []string{}

	rr := mathrand.New(mathrand.NewSource(mathrand.Int63()))

	for i := 0; i < s.tokenCount; i++ {
		tok := fmt.Sprintf("%v-%v", jwtToken, i)
		r := cacheRecord{
			resp: &authenticator.Response{
				User: &user.DefaultInfo{Name: fmt.Sprintf("holder of token %v", i)},
			},
		}
		// make different combinations of audience, failures, denies for the tokens.
		auds := []string{}
		for i := 0; i < rr.Intn(4); i++ {
			auds = append(auds, string(uuid.NewUUID()))
		}
		choice := rr.Float64()
		switch {
		case choice < 0.9:
			r.ok = true
			r.err = nil

			// add some realistic annotations on ~20% of successful authentications
			if f := rr.Float64(); f < 0.2 {
				r.annotations = map[string]string{
					"audience.authentication.kubernetes.io":  "c5443994-35e2-36be-ac03-7129ab271571",
					"namespace.authentication.kubernetes.io": "kube-system",
					"float.authentication.kubernetes.io":     fmt.Sprint(f),
				}
			}
		case choice < 0.99:
			r.ok = false
			r.err = nil
		default:
			r.ok = false
			r.err = errors.New("I can't think of a clever error name right now")
		}
		s.tokens = append(s.tokens, tok)
		s.tokenToResponse[tok] = &r
		if len(auds) > 0 {
			s.tokenToAuds[tok] = auds
		}
	}
}

func (s *singleBenchmark) lookup(ctx context.Context, token string) (*authenticator.Response, bool, error) {
	r, ok := s.tokenToResponse[token]
	if !ok {
		panic("test setup problem")
	}
	for key, val := range r.annotations {
		audit.AddAuditAnnotation(ctx, key, val)
	}
	return r.resp, r.ok, r.err
}

func (s *singleBenchmark) doAuthForTokenN(n int, a authenticator.Token) {
	tok := s.tokens[n]
	auds := s.tokenToAuds[tok]
	ctx := context.Background()
	ctx = authenticator.WithAudiences(ctx, auds)
	a.AuthenticateToken(ctx, tok)
}

func (s *singleBenchmark) run(b *testing.B) {
	b.Run(fmt.Sprintf("tokens=%d threads=%d", s.tokenCount, s.threadCount), s.bench)
}

func (s *singleBenchmark) bench(b *testing.B) {
	// Simulate slowness, qps limit, external service limitation, etc
	const maxInFlight = 40
	chokepoint := make(chan struct{}, maxInFlight)
	// lookup count
	var lookups uint64

	a := newWithClock(
		authenticator.TokenFunc(func(ctx context.Context, token string) (*authenticator.Response, bool, error) {
			atomic.AddUint64(&lookups, 1)

			chokepoint <- struct{}{}
			defer func() { <-chokepoint }()

			time.Sleep(1 * time.Millisecond)

			return s.lookup(ctx, token)
		}),
		true,
		4*time.Second,
		500*time.Millisecond,
		clock.RealClock{},
	)

	b.ResetTimer()
	b.SetParallelism(s.threadCount)
	b.RunParallel(func(pb *testing.PB) {
		r := mathrand.New(mathrand.NewSource(mathrand.Int63()))
		for pb.Next() {
			// some problems appear with random access, some appear with many
			// requests for a single entry, so we do both.
			s.doAuthForTokenN(r.Intn(s.tokenCount), a)
			s.doAuthForTokenN(0, a)
		}
	})
	b.StopTimer()

	b.ReportMetric(float64(lookups)/float64(b.N), "lookups/op")
}

// Add a test version of the audit context with a pre-populated event for easy annotation
// extraction.
func withAudit(ctx context.Context) context.Context {
	ctx = audit.WithAuditContext(ctx)
	ac := audit.AuditContextFrom(ctx)
	ac.Event.Level = auditinternal.LevelMetadata
	return ctx
}

func TestUnsafeConversions(t *testing.T) {

	// needs to be large to force allocations so we pick a random value between [1024, 2048]
	size := utilrand.IntnRange(1024, 2048+1)

	t.Run("toBytes semantics", func(t *testing.T) {
		t.Parallel()

		s := utilrand.String(size)
		b := toBytes(s)
		if len(b) != size {
			t.Errorf("unexpected length: %d", len(b))
		}
		if cap(b) != size {
			t.Errorf("unexpected capacity: %d", cap(b))
		}
		if !bytes.Equal(b, []byte(s)) {
			t.Errorf("unexpected equality failure: %#v", b)
		}
	})

	t.Run("toBytes allocations", func(t *testing.T) {

		s := utilrand.String(size)
		f := func() {
			b := toBytes(s)
			if len(b) != size {
				t.Errorf("invalid length: %d", len(b))
			}
		}
		allocs := testing.AllocsPerRun(100, f)
		if allocs > 0 {
			t.Errorf("expected zero allocations, got %v", allocs)
		}
	})

	t.Run("toString semantics", func(t *testing.T) {
		t.Parallel()

		b := make([]byte, size)
		if _, err := rand.Read(b); err != nil {
			t.Fatal(err)
		}
		s := toString(b)
		if len(s) != size {
			t.Errorf("unexpected length: %d", len(s))
		}
		if s != string(b) {
			t.Errorf("unexpected equality failure: %#v", s)
		}
	})

	t.Run("toString allocations", func(t *testing.T) {

		b := make([]byte, size)
		if _, err := rand.Read(b); err != nil {
			t.Fatal(err)
		}
		f := func() {
			s := toString(b)
			if len(s) != size {
				t.Errorf("invalid length: %d", len(s))
			}
		}
		allocs := testing.AllocsPerRun(100, f)
		if allocs > 0 {
			t.Errorf("expected zero allocations, got %v", allocs)
		}
	})
}

func TestKeyFunc(t *testing.T) {
	t.Parallel()

	hashPool := &sync.Pool{
		New: func() interface{} {
			return hmac.New(sha256.New, []byte("511e1e17-e2b7-9596-db8e-30de6390e923")) // deterministic HMAC for testing
		},
	}

	// use realistic audiences
	auds := []string{"7ada70b2-c23d-412f-2a75-f671eabbe526", "f79ac115-bfaf-7ffa-1715-476fe8392520", "68a5f9f4-4427-8901-e81a-9564a47a4772"}

	keyWithAuds := "\"\xf7\xac\xcd\x12\xf5\x83l\xa9;@\n\xa13a;\nd\x1f\xdelL\xd1\xe1!\x8a\xdahٛ\xbb\xf0"

	keyWithoutAuds := "\x810r \xb2\x6e\xdb\xa6?\x8e\x45\xf0,e\n1\xd3ȵ>\xabC\x9f\x73+\x75˿&\xa01\x4a"

	t.Run("has audiences", func(t *testing.T) {
		t.Parallel()

		key := keyFunc(hashPool, auds, jwtToken)
		if key != keyWithAuds {
			t.Errorf("unexpected equality failure: %#v", key)
		}
	})

	t.Run("nil audiences", func(t *testing.T) {
		t.Parallel()

		key := keyFunc(hashPool, nil, jwtToken)
		if key != keyWithoutAuds {
			t.Errorf("unexpected equality failure: %#v", key)
		}
	})
}
