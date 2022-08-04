package httpsig

import (
	"crypto"
	"crypto/hmac"
	"errors"
)

// HMACSHA256 implements keyed HMAC over SHA256 digests
var HMACSHA256 Algorithm = hmac_sha256{}

type hmac_sha256 struct{}

func (hmac_sha256) Name() string {
	return "hmac-sha256"
}

func (a hmac_sha256) Sign(key interface{}, data []byte) ([]byte, error) {
	k := toHMACKey(key)
	if k == nil {
		return nil, unsupportedAlgorithm(a)
	}
	return HMACSign(k, crypto.SHA256, data)
}

func (a hmac_sha256) Verify(key interface{}, data, sig []byte) error {
	k := toHMACKey(key)
	if k == nil {
		return unsupportedAlgorithm(a)
	}
	return HMACVerify(k, crypto.SHA256, data, sig)
}

// HMACSign signs a digest of the data hashed using the provided hash and key.
func HMACSign(key []byte, hash crypto.Hash, data []byte) ([]byte, error) {
	h := hmac.New(hash.New, key)
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// HMACVerify verifies a signed digest of the data hashed using the provided
// hash and key.
func HMACVerify(key []byte, hash crypto.Hash, data, sig []byte) error {
	actual_sig, err := HMACSign(key, hash, data)
	if err != nil {
		return err
	}
	if !hmac.Equal(actual_sig, sig) {
		return errors.New("hmac signature mismatch")
	}
	return nil
}
