package cluster

import (
	"github.com/coreos/tectonic-installer/installer/server"
)

// Config holds the configuration needed to setup an individual cluster.
type Config struct {
	Name                              string `json:"name"`
	server.TerraformApplyHandlerInput `json:"input"`
}

// Apply overrides configuration using values from another config. Currently only works on variables.
func (c *Config) Apply(top Config) {
	for k, v := range top.Variables {
		c.Variables[k] = v
	}
}
