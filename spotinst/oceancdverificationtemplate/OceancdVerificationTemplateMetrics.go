// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#metrics_name OceancdVerificationTemplate#metrics_name}.
	MetricsName *string `field:"required" json:"metricsName" yaml:"metricsName"`
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#provider OceancdVerificationTemplate#provider}
	Provider interface{} `field:"required" json:"provider" yaml:"provider"`
	// baseline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#baseline OceancdVerificationTemplate#baseline}
	Baseline *OceancdVerificationTemplateMetricsBaseline `field:"optional" json:"baseline" yaml:"baseline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#consecutive_error_limit OceancdVerificationTemplate#consecutive_error_limit}.
	ConsecutiveErrorLimit *float64 `field:"optional" json:"consecutiveErrorLimit" yaml:"consecutiveErrorLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#count OceancdVerificationTemplate#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#dry_run OceancdVerificationTemplate#dry_run}.
	DryRun interface{} `field:"optional" json:"dryRun" yaml:"dryRun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#failure_condition OceancdVerificationTemplate#failure_condition}.
	FailureCondition *string `field:"optional" json:"failureCondition" yaml:"failureCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#failure_limit OceancdVerificationTemplate#failure_limit}.
	FailureLimit *float64 `field:"optional" json:"failureLimit" yaml:"failureLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#initial_delay OceancdVerificationTemplate#initial_delay}.
	InitialDelay *string `field:"optional" json:"initialDelay" yaml:"initialDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#interval OceancdVerificationTemplate#interval}.
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/oceancd_verification_template#success_condition OceancdVerificationTemplate#success_condition}.
	SuccessCondition *string `field:"optional" json:"successCondition" yaml:"successCondition"`
}

