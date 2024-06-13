// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProviderCloudWatchMetricDataQueriesMetricStat struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/oceancd_verification_template#metric OceancdVerificationTemplate#metric}
	Metric *OceancdVerificationTemplateMetricsProviderCloudWatchMetricDataQueriesMetricStatMetric `field:"optional" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/oceancd_verification_template#metric_period OceancdVerificationTemplate#metric_period}.
	MetricPeriod *float64 `field:"optional" json:"metricPeriod" yaml:"metricPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/oceancd_verification_template#stat OceancdVerificationTemplate#stat}.
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/oceancd_verification_template#unit OceancdVerificationTemplate#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

