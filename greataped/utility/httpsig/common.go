package httpsig

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	// Rand is a hookable reader used as a random byte source.
	Rand io.Reader = rand.Reader
)

// requestPath returns the :path pseudo header according to the HTTP/2 spec.
func requestPath(req *http.Request) string {
	path := req.URL.Path
	if path == "" {
		path = "/"
	}
	if req.URL.RawQuery != "" {
		path += "?" + req.URL.RawQuery
	}
	return path
}

// BuildSignatureString constructs a signature string following section 2.3
func BuildSignatureString(req *http.Request, headers []string) string {
	if len(headers) == 0 {
		headers = []string{"date"}
	}
	values := make([]string, 0, len(headers))
	for _, h := range headers {
		switch h {
		case "(request-target)":
			values = append(values, fmt.Sprintf("%s: %s %s",
				h, strings.ToLower(req.Method), requestPath(req)))
		case "host":
			values = append(values, fmt.Sprintf("%s: %s", h, req.Host))
		case "date":
			if req.Header.Get(h) == "" {
				req.Header.Set(h, time.Now().UTC().Format(http.TimeFormat))
			}
			values = append(values, fmt.Sprintf("%s: %s", h, req.Header.Get(h)))
		default:
			for _, value := range req.Header[http.CanonicalHeaderKey(h)] {
				values = append(values,
					fmt.Sprintf("%s: %s", h, strings.TrimSpace(value)))
			}
		}
	}
	return strings.Join(values, "\n")
}

// BuildSignatureData is a convenience wrapper around BuildSignatureString that
// returns []byte instead of a string.
func BuildSignatureData(req *http.Request, headers []string) []byte {
	return []byte(BuildSignatureString(req, headers))
}

func toRSAPrivateKey(key interface{}) *rsa.PrivateKey {
	switch k := key.(type) {
	case *rsa.PrivateKey:
		return k
	default:
		return nil
	}
}

func toRSAPublicKey(key interface{}) *rsa.PublicKey {
	switch k := key.(type) {
	case *rsa.PublicKey:
		return k
	case *rsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

func toHMACKey(key interface{}) []byte {
	switch k := key.(type) {
	case []byte:
		return k
	default:
		return nil
	}
}

func unsupportedAlgorithm(a Algorithm) error {
	return fmt.Errorf("key does not support algorithm %q", a.Name())
}
