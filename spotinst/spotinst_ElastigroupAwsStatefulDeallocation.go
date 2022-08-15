// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsStatefulDeallocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#should_delete_images ElastigroupAws#should_delete_images}.
	ShouldDeleteImages interface{} `field:"optional" json:"shouldDeleteImages" yaml:"shouldDeleteImages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#should_delete_network_interfaces ElastigroupAws#should_delete_network_interfaces}.
	ShouldDeleteNetworkInterfaces interface{} `field:"optional" json:"shouldDeleteNetworkInterfaces" yaml:"shouldDeleteNetworkInterfaces"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#should_delete_snapshots ElastigroupAws#should_delete_snapshots}.
	ShouldDeleteSnapshots interface{} `field:"optional" json:"shouldDeleteSnapshots" yaml:"shouldDeleteSnapshots"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#should_delete_volumes ElastigroupAws#should_delete_volumes}.
	ShouldDeleteVolumes interface{} `field:"optional" json:"shouldDeleteVolumes" yaml:"shouldDeleteVolumes"`
}
