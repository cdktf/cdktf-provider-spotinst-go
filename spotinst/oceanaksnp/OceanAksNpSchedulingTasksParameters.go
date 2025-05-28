// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpSchedulingTasksParameters struct {
	// parameters_cluster_roll block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/ocean_aks_np#parameters_cluster_roll OceanAksNp#parameters_cluster_roll}
	ParametersClusterRoll *OceanAksNpSchedulingTasksParametersParametersClusterRoll `field:"optional" json:"parametersClusterRoll" yaml:"parametersClusterRoll"`
	// parameters_upgrade_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/ocean_aks_np#parameters_upgrade_config OceanAksNp#parameters_upgrade_config}
	ParametersUpgradeConfig *OceanAksNpSchedulingTasksParametersParametersUpgradeConfig `field:"optional" json:"parametersUpgradeConfig" yaml:"parametersUpgradeConfig"`
}

