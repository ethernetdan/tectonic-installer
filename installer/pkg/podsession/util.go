package podsession

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"strings"
)

// GenerateRandomKey creates a random key with the given length in bytes.
// On failure, returns nil. From github.com/gorilla/securecookie.
//
// Callers should explicitly check for the possibility of a nil return, treat
// it as a failure of the system random number generator, and not continue.
func GenerateRandomKey(length int) []byte {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}

// encode32 use base32 to encode into a string.
func encode32(in []byte) string {
	return strings.TrimRight(base32.StdEncoding.EncodeToString(in), "=")
}
