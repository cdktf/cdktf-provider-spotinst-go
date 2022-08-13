// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ManagedInstanceAwsIntegrationRoute53 struct {
	// domains block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#domains ManagedInstanceAws#domains}
	Domains interface{} `field:"required" json:"domains" yaml:"domains"`
}

