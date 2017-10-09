package podsession

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"net/http"
	"strings"
	"time"
)

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// attempt to retrieve session cookie
	var session string
	if c, err := r.Cookie(p.config.SessionCookie); err == nil {
		// use session token from cookie
		session = c.Value
	} else {
		// generate new session token
		session = encode32(GenerateRandomKey(32))
	}

	// set session cookie
	c := &http.Cookie{
		Name:  p.config.SessionCookie,
		Value: session,
	}
	http.SetCookie(w, c)

	// hash session to use as Service Name
	serviceName := p.hashServiceName(session)

	// check if session exists
	if seen, exists := p.lastSeen[serviceName]; !exists {
		go p.createProxyService(serviceName)
	} else if seen.Add(p.config.Heartbeat).Before(time.Now().UTC()) {
		go p.heartbeatProxyService(serviceName)
	}

	//* Proxy connection to Pod
}

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
