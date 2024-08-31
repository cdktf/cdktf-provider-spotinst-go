// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanRightSizingRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// recommendation_application_intervals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.190.0/docs/resources/ocean_right_sizing_rule#recommendation_application_intervals OceanRightSizingRule#recommendation_application_intervals}
	RecommendationApplicationIntervals interface{} `field:"required" json:"recommendationApplicationIntervals" yaml:"recommendationApplicationIntervals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.190.0/docs/resources/ocean_right_sizing_rule#rule_name OceanRightSizingRule#rule_name}.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// attach_workloads block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.190.0/docs/resources/ocean_right_sizing_rule#attach_workloads OceanRightSizingRule#attach_workloads}
	AttachWorkloads interface{} `field:"optional" json:"attachWorkloads" yaml:"attachWorkloads"`
	// detach_workloads block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.190.0/docs/resources/ocean_right_sizing_rule#detach_workloads OceanRightSizingRule#detach_workloads}
	DetachWorkloads interface{} `field:"optional" json:"detachWorkloads" yaml:"detachWorkloads"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.190.0/docs/resources/ocean_right_sizing_rule#id OceanRightSizingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.190.0/docs/resources/ocean_right_sizing_rule#ocean_id OceanRightSizingRule#ocean_id}.
	OceanId *string `field:"optional" json:"oceanId" yaml:"oceanId"`
	// recommendation_application_boundaries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.190.0/docs/resources/ocean_right_sizing_rule#recommendation_application_boundaries OceanRightSizingRule#recommendation_application_boundaries}
	RecommendationApplicationBoundaries interface{} `field:"optional" json:"recommendationApplicationBoundaries" yaml:"recommendationApplicationBoundaries"`
	// recommendation_application_min_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.190.0/docs/resources/ocean_right_sizing_rule#recommendation_application_min_threshold OceanRightSizingRule#recommendation_application_min_threshold}
	RecommendationApplicationMinThreshold interface{} `field:"optional" json:"recommendationApplicationMinThreshold" yaml:"recommendationApplicationMinThreshold"`
	// recommendation_application_overhead_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.190.0/docs/resources/ocean_right_sizing_rule#recommendation_application_overhead_values OceanRightSizingRule#recommendation_application_overhead_values}
	RecommendationApplicationOverheadValues interface{} `field:"optional" json:"recommendationApplicationOverheadValues" yaml:"recommendationApplicationOverheadValues"`
}

