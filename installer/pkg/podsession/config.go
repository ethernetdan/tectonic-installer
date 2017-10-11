package podsession

import (
	"time"
)

// Config provides the options used to configure the proxy.
type Config struct {
	// - Session
	// SessionCookie is the name of the cookie used to store the session identifier.
	SessionCookie string

	// NameLength is the number of characters to use of the hash that identifies Services.
	NameLength int

	// Heartbeat is the interval after which the Proxy should mark the service as being used on the API server.
	Heartbeat time.Duration

	// - Proxy
	// ProxyTimeout is the amount of time to attempt to make a connection to the Pod before giving up.
	ProxyTimeout time.Duration

	// ProxyScheme is the network protocol used to connect to the target Service.
	ProxyScheme string
}
