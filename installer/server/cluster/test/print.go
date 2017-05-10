package main

import (
	"encoding/json"
	"fmt"

	"github.com/coreos/tectonic-installer/installer/server/cluster"
	"github.com/coreos/tectonic-installer/installer/server/terraform"
)

func main() {
	cfg := &cluster.Config{
		ClusterName:           "test",
		AdminEmail:            "admin@example.com",
		BaseDomain:            "example.com",
		ContainerLinuxChannel: "stable",
		ClusterCIDR:           "10.2.0.0/16",
	}

	fmt.Println("Go representation:")
	fmt.Printf("%#v\n", cfg)
	fmt.Println()

	// print hcl
	vars := cfg.Variables()
	str, _ := terraform.MapVarsToTFVars(vars)
	fmt.Println("HCL:")
	fmt.Println(str)
	fmt.Println()

	// print JSON
	data, _ := json.Marshal(cfg)
	fmt.Println("JSON:")
	fmt.Println(string(data))
}
