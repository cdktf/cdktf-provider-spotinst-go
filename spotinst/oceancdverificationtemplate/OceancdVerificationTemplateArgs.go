// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateArgs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/oceancd_verification_template#arg_name OceancdVerificationTemplate#arg_name}.
	ArgName *string `field:"required" json:"argName" yaml:"argName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/oceancd_verification_template#value OceancdVerificationTemplate#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/oceancd_verification_template#value_from OceancdVerificationTemplate#value_from}
	ValueFrom *OceancdVerificationTemplateArgsValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

