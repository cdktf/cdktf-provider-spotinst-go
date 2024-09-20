// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProviderJobSpec struct {
	// job_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/oceancd_verification_template#job_template OceancdVerificationTemplate#job_template}
	JobTemplate interface{} `field:"required" json:"jobTemplate" yaml:"jobTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/oceancd_verification_template#backoff_limit OceancdVerificationTemplate#backoff_limit}.
	BackoffLimit *float64 `field:"optional" json:"backoffLimit" yaml:"backoffLimit"`
}

