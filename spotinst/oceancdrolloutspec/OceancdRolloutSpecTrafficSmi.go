// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTrafficSmi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/oceancd_rollout_spec#smi_root_service OceancdRolloutSpec#smi_root_service}.
	SmiRootService *string `field:"optional" json:"smiRootService" yaml:"smiRootService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/oceancd_rollout_spec#traffic_split_name OceancdRolloutSpec#traffic_split_name}.
	TrafficSplitName *string `field:"optional" json:"trafficSplitName" yaml:"trafficSplitName"`
}

