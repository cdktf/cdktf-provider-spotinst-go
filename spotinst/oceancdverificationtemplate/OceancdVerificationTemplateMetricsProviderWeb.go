// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProviderWeb struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#url OceancdVerificationTemplate#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#body OceancdVerificationTemplate#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#insecure OceancdVerificationTemplate#insecure}.
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#json_path OceancdVerificationTemplate#json_path}.
	JsonPath *string `field:"optional" json:"jsonPath" yaml:"jsonPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#method OceancdVerificationTemplate#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#timeout_seconds OceancdVerificationTemplate#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// web_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#web_header OceancdVerificationTemplate#web_header}
	WebHeader interface{} `field:"optional" json:"webHeader" yaml:"webHeader"`
}

