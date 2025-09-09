// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpSchedulingTasksParametersParametersUpgradeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_aks_np#apply_roll OceanAksNp#apply_roll}.
	ApplyRoll interface{} `field:"optional" json:"applyRoll" yaml:"applyRoll"`
	// roll_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_aks_np#roll_parameters OceanAksNp#roll_parameters}
	RollParameters *OceanAksNpSchedulingTasksParametersParametersUpgradeConfigRollParameters `field:"optional" json:"rollParameters" yaml:"rollParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_aks_np#scope_version OceanAksNp#scope_version}.
	ScopeVersion *string `field:"optional" json:"scopeVersion" yaml:"scopeVersion"`
}

