package elastigroupgcp


type ElastigroupGcpIntegrationDockerSwarm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/elastigroup_gcp#master_host ElastigroupGcp#master_host}.
	MasterHost *string `field:"required" json:"masterHost" yaml:"masterHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/elastigroup_gcp#master_port ElastigroupGcp#master_port}.
	MasterPort *float64 `field:"required" json:"masterPort" yaml:"masterPort"`
}

