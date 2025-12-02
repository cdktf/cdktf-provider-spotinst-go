// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsScheduledTaskTasksParameters struct {
	// ami_auto_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/ocean_aws#ami_auto_update OceanAws#ami_auto_update}
	AmiAutoUpdate *OceanAwsScheduledTaskTasksParametersAmiAutoUpdate `field:"optional" json:"amiAutoUpdate" yaml:"amiAutoUpdate"`
	// parameters_cluster_roll block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/ocean_aws#parameters_cluster_roll OceanAws#parameters_cluster_roll}
	ParametersClusterRoll *OceanAwsScheduledTaskTasksParametersParametersClusterRoll `field:"optional" json:"parametersClusterRoll" yaml:"parametersClusterRoll"`
}

