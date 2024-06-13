// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProviderCloudWatchMetricDataQueriesMetricStatMetricDimensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/oceancd_verification_template#dimension_name OceancdVerificationTemplate#dimension_name}.
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/oceancd_verification_template#dimension_value OceancdVerificationTemplate#dimension_value}.
	DimensionValue *string `field:"required" json:"dimensionValue" yaml:"dimensionValue"`
}

