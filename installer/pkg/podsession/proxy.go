package podsession

import (
	"net/http"
	"net/http/httputil"
	"time"
)

const (
	// ForwardedHostHeader contains the host of the original request.
	ForwardedHostHeader = "X-Forwarded-Host"

	// KubernetesServiceHeader contains the Kubernetes Service to be used for the request.
	KubernetesServiceHeader = "X-Kubernetes-Service"
)

// NewProxy returns a Proxy configured with the given Config.
func NewProxy(cfg Config) (*Proxy, error) {
	// setup reverse proxy
	director := func(req *http.Request) {
		// record original host
		req.Header.Set(ForwardedHostHeader, req.Host)

		// direct to Service using scheme from config and service determined in ServeHTTP
		req.URL.Scheme = cfg.ProxyScheme
		req.URL.Host = req.Header.Get(KubernetesServiceHeader)

		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	return &Proxy{
		config:   cfg,
		lastSeen: map[string]time.Time{},
		proxy:    &httputil.ReverseProxy{Director: director},
	}, nil
}

// Proxy provides a http.Handler to proxy to Pod specific to each session. Proxy is responsible for the creation of the
// replication constructs that create the Pod.
type Proxy struct {
	config   Config
	lastSeen map[string]time.Time
	proxy    *httputil.ReverseProxy
}
