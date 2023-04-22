package elastigroupgke


type ElastigroupGkeBackendServicesNamedPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/elastigroup_gke#name ElastigroupGke#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/elastigroup_gke#ports ElastigroupGke#ports}.
	Ports *[]*string `field:"required" json:"ports" yaml:"ports"`
}

