package podsession

import (
	"crypto/sha256"
	"time"
)

// UpdateService creates/updates the given Kubernetes service.
//
// A cache with the last time the Service was used locally is checked. If the Service isn't found it is attempted to be
// created on the API server. If the Service was last seen a duration larger than specified in the proxy configuration,
// an annotation is updated on the Service with the current timestamp.
func (p *Proxy) UpdateService(serviceName string) {
	// check if session exists, update lastSeen
	if seen, exists := p.lastSeen[serviceName]; !exists {
		p.createProxyService(serviceName)
	} else if seen.Add(p.config.Heartbeat).Before(time.Now().UTC()) {
		p.heartbeatProxyService(serviceName)
	}
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
func hashServiceName(sessionName string, nameLength int) string {
	hashData := sha256.Sum256([]byte(sessionName))
	var serviceName []byte
	copy(serviceName, hashData[:nameLength])
	return string(serviceName)
}
