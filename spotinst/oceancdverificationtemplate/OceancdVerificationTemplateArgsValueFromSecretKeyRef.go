// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateArgsValueFromSecretKeyRef struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/oceancd_verification_template#key OceancdVerificationTemplate#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/oceancd_verification_template#name OceancdVerificationTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

