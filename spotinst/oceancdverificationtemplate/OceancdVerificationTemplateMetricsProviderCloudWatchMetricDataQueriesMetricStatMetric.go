// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProviderCloudWatchMetricDataQueriesMetricStatMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/oceancd_verification_template#metric_name OceancdVerificationTemplate#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/oceancd_verification_template#dimensions OceancdVerificationTemplate#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/oceancd_verification_template#namespace OceancdVerificationTemplate#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

