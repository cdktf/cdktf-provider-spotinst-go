package elastigroupgcp


type ElastigroupGcpBackendServicesNamedPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.1/docs/resources/elastigroup_gcp#name ElastigroupGcp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.1/docs/resources/elastigroup_gcp#ports ElastigroupGcp#ports}.
	Ports *[]*string `field:"required" json:"ports" yaml:"ports"`
}

