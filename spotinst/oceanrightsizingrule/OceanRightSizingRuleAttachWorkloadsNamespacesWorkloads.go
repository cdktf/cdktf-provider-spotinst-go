// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule


type OceanRightSizingRuleAttachWorkloadsNamespacesWorkloads struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.206.0/docs/resources/ocean_right_sizing_rule#workload_type OceanRightSizingRule#workload_type}.
	WorkloadType *string `field:"required" json:"workloadType" yaml:"workloadType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.206.0/docs/resources/ocean_right_sizing_rule#regex_name OceanRightSizingRule#regex_name}.
	RegexName *string `field:"optional" json:"regexName" yaml:"regexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.206.0/docs/resources/ocean_right_sizing_rule#workload_name OceanRightSizingRule#workload_name}.
	WorkloadName *string `field:"optional" json:"workloadName" yaml:"workloadName"`
}

