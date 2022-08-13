// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGkeBackendServicesNamedPorts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#name ElastigroupGke#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#ports ElastigroupGke#ports}.
	Ports *[]*string `field:"required" json:"ports" yaml:"ports"`
}

