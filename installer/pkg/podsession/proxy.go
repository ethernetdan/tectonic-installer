package podsession

import (
	"crypto/sha256"
	"net/http"
	"time"
)

// NewProxy returns a Proxy configured with the given Config.
func NewProxy(cfg Config) (*Proxy, error) {
	return &Proxy{
		config: cfg,
	}, nil
}

// Proxy provides a http.Handler to proxy to Pod specific to each session. Proxy is responsible for the creation of the
// replication constructs that create the Pod.
type Proxy struct {
	config   Config
	lastSeen map[string]time.Time
}

// Handler proxies requests to the Pod for the session.
func (p *Proxy) Handler() http.Handler {
	return http.HandlerFunc(p.ServeHTTP)
}

// createProxyService uses the Kubernetes API to create a new service for the given session.
func (p *Proxy) createProxyService(name string) {
	//				* Use label indicating managed by the installer
	//				* Record current time in annotation + local map of session => last seen
	//				* Ignore already exist failures but retry for all others until timeout below
	//				* Reattempt connection for 15 seconds, if timeout:
	//						* Return 502 with cookie
}

// heartbeatProxyService marks that the proxy Service has just been used.
func (p *Proxy) heartbeatProxyService(name string) {
	//		* Patch Pod with heartbeat annotation set to current time

}

// hashServiceName returns the hash used to identify the Service for the session.
func (p *Proxy) hashServiceName(sessionName string) string {
	hashData := sha256.Sum256([]byte(sessionName))
	var serviceName []byte
	copy(serviceName, hashData[:p.config.NameLength])
	return string(serviceName)
}
