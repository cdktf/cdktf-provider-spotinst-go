// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws


type ManagedInstanceAwsDelete struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.0/docs/resources/managed_instance_aws#ami_backup_should_delete_images ManagedInstanceAws#ami_backup_should_delete_images}.
	AmiBackupShouldDeleteImages interface{} `field:"optional" json:"amiBackupShouldDeleteImages" yaml:"amiBackupShouldDeleteImages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.0/docs/resources/managed_instance_aws#deallocation_config_should_delete_images ManagedInstanceAws#deallocation_config_should_delete_images}.
	DeallocationConfigShouldDeleteImages interface{} `field:"optional" json:"deallocationConfigShouldDeleteImages" yaml:"deallocationConfigShouldDeleteImages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.0/docs/resources/managed_instance_aws#should_delete_network_interfaces ManagedInstanceAws#should_delete_network_interfaces}.
	ShouldDeleteNetworkInterfaces interface{} `field:"optional" json:"shouldDeleteNetworkInterfaces" yaml:"shouldDeleteNetworkInterfaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.0/docs/resources/managed_instance_aws#should_delete_snapshots ManagedInstanceAws#should_delete_snapshots}.
	ShouldDeleteSnapshots interface{} `field:"optional" json:"shouldDeleteSnapshots" yaml:"shouldDeleteSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.0/docs/resources/managed_instance_aws#should_delete_volumes ManagedInstanceAws#should_delete_volumes}.
	ShouldDeleteVolumes interface{} `field:"optional" json:"shouldDeleteVolumes" yaml:"shouldDeleteVolumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.0/docs/resources/managed_instance_aws#should_terminate_instance ManagedInstanceAws#should_terminate_instance}.
	ShouldTerminateInstance interface{} `field:"optional" json:"shouldTerminateInstance" yaml:"shouldTerminateInstance"`
}

