// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule


type OceanRightSizingRuleDetachWorkloadsNamespaces struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_right_sizing_rule#namespace_name OceanRightSizingRule#namespace_name}.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_right_sizing_rule#labels OceanRightSizingRule#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// workloads block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_right_sizing_rule#workloads OceanRightSizingRule#workloads}
	Workloads interface{} `field:"optional" json:"workloads" yaml:"workloads"`
}

