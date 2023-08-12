package managedinstanceaws


type ManagedInstanceAwsResourceTagSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/managed_instance_aws#should_tag_amis ManagedInstanceAws#should_tag_amis}.
	ShouldTagAmis interface{} `field:"optional" json:"shouldTagAmis" yaml:"shouldTagAmis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/managed_instance_aws#should_tag_enis ManagedInstanceAws#should_tag_enis}.
	ShouldTagEnis interface{} `field:"optional" json:"shouldTagEnis" yaml:"shouldTagEnis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/managed_instance_aws#should_tag_snapshots ManagedInstanceAws#should_tag_snapshots}.
	ShouldTagSnapshots interface{} `field:"optional" json:"shouldTagSnapshots" yaml:"shouldTagSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/managed_instance_aws#should_tag_volumes ManagedInstanceAws#should_tag_volumes}.
	ShouldTagVolumes interface{} `field:"optional" json:"shouldTagVolumes" yaml:"shouldTagVolumes"`
}

