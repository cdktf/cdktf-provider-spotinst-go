// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpSchedulingTasksParameters struct {
	// parameters_cluster_roll block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.200.0/docs/resources/ocean_aks_np#parameters_cluster_roll OceanAksNp#parameters_cluster_roll}
	ParametersClusterRoll *OceanAksNpSchedulingTasksParametersParametersClusterRoll `field:"optional" json:"parametersClusterRoll" yaml:"parametersClusterRoll"`
}

