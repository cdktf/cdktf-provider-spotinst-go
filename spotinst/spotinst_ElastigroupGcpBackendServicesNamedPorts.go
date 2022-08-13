// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGcpBackendServicesNamedPorts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#name ElastigroupGcp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#ports ElastigroupGcp#ports}.
	Ports *[]*string `field:"required" json:"ports" yaml:"ports"`
}

