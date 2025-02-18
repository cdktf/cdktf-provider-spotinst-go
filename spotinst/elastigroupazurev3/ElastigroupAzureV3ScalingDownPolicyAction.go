// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3ScalingDownPolicyAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/elastigroup_azure_v3#adjustment ElastigroupAzureV3#adjustment}.
	Adjustment *string `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/elastigroup_azure_v3#maximum ElastigroupAzureV3#maximum}.
	Maximum *string `field:"optional" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/elastigroup_azure_v3#minimum ElastigroupAzureV3#minimum}.
	Minimum *string `field:"optional" json:"minimum" yaml:"minimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/elastigroup_azure_v3#target ElastigroupAzureV3#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/elastigroup_azure_v3#type ElastigroupAzureV3#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

