package cluster

// Config describes how a Tectonic cluster should be configured. It is possible to marshal Terraform configuration directly
// from Config.
type Config struct {
	// ClusterName
	ClusterName string `hcl:"tectonic_cluster_name"`

	// AdminEmail is the email used to login as the admin user to the Tectonic Console.
	AdminEmail string `hcl:"tectonic_admin_email"`

	// BaseDomain is the suffix of any cluster domain created.
	BaseDomain string `hcl:"tectonic_base_domain"`

	// ContainerLinuxChannel is the channel of CL that machines should use.
	ContainerLinuxChannel string `hcl:"tectonic_cl_channel"`

	// ClusterCIDR is the range to use for the IPs of Pods.
	ClusterCIDR string `hcl:"tectonic_cluster_cidr"`
}
