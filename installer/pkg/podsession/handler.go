package podsession

import (
	"context"
	"net/http"
)

// Handler proxies requests to the Pod for the session.
func (p *Proxy) Handler() http.Handler {
	return http.HandlerFunc(p.proxyHTTP)
}

// proxyHTTP is wrapped to allow an http.Handler to be provided.
func (p *Proxy) proxyHTTP(w http.ResponseWriter, req *http.Request) {
	session := p.restoreSessionOrNew(req)
	p.setSession(w, session)

	// hash session to use as name of Service and ensure Service exists
	serviceName := hashServiceName(session, p.config.NameLength)
	go p.UpdateService(serviceName)

	// record Kubernetes service name for proxy
	req.Header.Set(KubernetesServiceHeader, serviceName)

	// timeout request if not successful
	c, cancel := context.WithTimeout(req.Context(), p.config.ProxyTimeout)
	req = req.WithContext(c)
	defer cancel()

	p.proxy.ServeHTTP(w, req)
}

// restoreSessionOrNew tries to restore the session name from a cookie and generates a new one if that fails.
func (p *Proxy) restoreSessionOrNew(req *http.Request) string {
	if c, err := req.Cookie(p.config.SessionCookie); err == nil {
		return c.Value
	}
	sessionData := GenerateRandomKey(32)
	return encode32(sessionData)
}

// setSession creates and sets a cookie containing the session name.
func (p *Proxy) setSession(w http.ResponseWriter, sessionName string) {
	cookie := &http.Cookie{
		Name:  p.config.SessionCookie,
		Value: sessionName,
	}
	http.SetCookie(w, cookie)
}
