// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsScheduledTaskTasksParametersAmiAutoUpdate struct {
	// ami_auto_update_cluster_roll block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/ocean_aws#ami_auto_update_cluster_roll OceanAws#ami_auto_update_cluster_roll}
	AmiAutoUpdateClusterRoll *OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRoll `field:"optional" json:"amiAutoUpdateClusterRoll" yaml:"amiAutoUpdateClusterRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/ocean_aws#apply_roll OceanAws#apply_roll}.
	ApplyRoll interface{} `field:"optional" json:"applyRoll" yaml:"applyRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/ocean_aws#minor_version OceanAws#minor_version}.
	MinorVersion interface{} `field:"optional" json:"minorVersion" yaml:"minorVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/ocean_aws#patch OceanAws#patch}.
	Patch interface{} `field:"optional" json:"patch" yaml:"patch"`
}

