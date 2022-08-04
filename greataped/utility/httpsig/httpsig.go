package httpsig

// Algorithm provides methods used to sign/verify signatures.
type Algorithm interface {
	Name() string
	Sign(key interface{}, data []byte) (sig []byte, err error)
	Verify(key interface{}, data, sig []byte) error
}

// KeyGetter is an interface used by the verifier to retrieve a key stored
// by key id.
//
// The following types are supported for the specified algorithms:
// []byte            - HMAC signatures
// *rsa.PublicKey    - RSA signatures
// *rsa.PrivateKey   - RSA signatures
//
// Other types will treated as if no key was returned.
type KeyGetter interface {
	GetKey(id string) interface{}
}

// KeyGetterFunc is a convenience type for implementing a KeyGetter with a
// regular function
type KeyGetterFunc func(id string) interface{}

// GetKey calls fn(id)
func (fn KeyGetterFunc) GetKey(id string) interface{} {
	return fn(id)
}
