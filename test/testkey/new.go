package testkey

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"sync"
	"testing"

	"github.com/spiffe/spire/pkg/common/pemutil"
	"github.com/stretchr/testify/require"
)

var (
	keysMtx sync.Mutex

	rsa1024Idx int
	rsa2048Idx int
	rsa4096Idx int
	ec256Idx   int
	ec384Idx   int
)

func NewRSA1024(tb testing.TB) *rsa.PrivateKey {
	keysMtx.Lock()
	defer keysMtx.Unlock()
	if rsa1024Idx >= len(RSA1024Keys) {
		tb.Fatalf("exhausted %d pregenerated RSA-1024 test keys in test; use generate.sh to increase amount or refactor test to use less keys", len(RSA1024Keys))
	}
	key, err := pemutil.ParseRSAPrivateKey([]byte(RSA1024Keys[rsa1024Idx]))
	require.NoError(tb, err)
	rsa1024Idx++
	return key
}

func NewRSA2048(tb testing.TB) *rsa.PrivateKey {
	keysMtx.Lock()
	defer keysMtx.Unlock()
	if rsa2048Idx >= len(RSA2048Keys) {
		tb.Fatalf("exhausted %d pregenerated RSA-2048 test keys in test; use generate.sh to increase amount or refactor test to use less keys", len(RSA2048Keys))
	}
	key, err := pemutil.ParseRSAPrivateKey([]byte(RSA2048Keys[rsa2048Idx]))
	require.NoError(tb, err)
	rsa2048Idx++
	return key
}

func NewRSA4096(tb testing.TB) *rsa.PrivateKey {
	keysMtx.Lock()
	defer keysMtx.Unlock()
	if rsa4096Idx >= len(RSA4096Keys) {
		tb.Fatalf("exhausted %d pregenerated RSA-4096 test keys in test; use generate.sh to increase amount or refactor test to use less keys", len(RSA4096Keys))
	}
	key, err := pemutil.ParseRSAPrivateKey([]byte(RSA4096Keys[rsa4096Idx]))
	require.NoError(tb, err)
	rsa4096Idx++
	return key
}

func NewEC256(tb testing.TB) *ecdsa.PrivateKey {
	key, err := newEC256()
	require.NoError(tb, err)
	return key
}

func MustEC256() *ecdsa.PrivateKey {
	key, err := newEC256()
	check(err)
	return key
}

func newEC256() (*ecdsa.PrivateKey, error) {
	keysMtx.Lock()
	defer keysMtx.Unlock()
	if ec256Idx >= len(EC256Keys) {
		return nil, fmt.Errorf("exhausted %d pregenerated EC-256 test keys in test; use generate.sh to increase amount or refactor test to use less keys", len(EC256Keys))
	}
	key, err := pemutil.ParseECPrivateKey([]byte(EC256Keys[ec256Idx]))
	if err != nil {
		return nil, err
	}
	ec256Idx++
	return key, nil
}

func NewEC384(tb testing.TB) *ecdsa.PrivateKey {
	keysMtx.Lock()
	defer keysMtx.Unlock()
	if ec384Idx >= len(EC384Keys) {
		tb.Fatalf("exhausted %d pregenerated EC-384 test keys in test; use generate.sh to increase amount or refactor test to use less keys", len(EC384Keys))
	}
	key, err := pemutil.ParseECPrivateKey([]byte(EC384Keys[ec384Idx]))
	require.NoError(tb, err)
	ec384Idx++
	return key
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
