package podsession

import (
	"time"
)

// Config provides the options used to configure the proxy.
type Config struct {
	// SessionCookie is the name of the cookie used to store the session identifier.
	SessionCookie string

	// Heartbeat is the interval after which the Proxy should mark the service as being used.
	Heartbeat time.Duration

	// NameLength is the number of characters to use of the hash that identifies Services.
	NameLength int
}
