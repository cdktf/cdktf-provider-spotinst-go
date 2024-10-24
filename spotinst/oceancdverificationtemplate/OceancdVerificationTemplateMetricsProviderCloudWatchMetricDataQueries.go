// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProviderCloudWatchMetricDataQueries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.1/docs/resources/oceancd_verification_template#id OceancdVerificationTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.1/docs/resources/oceancd_verification_template#expression OceancdVerificationTemplate#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.1/docs/resources/oceancd_verification_template#label OceancdVerificationTemplate#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// metric_stat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.1/docs/resources/oceancd_verification_template#metric_stat OceancdVerificationTemplate#metric_stat}
	MetricStat *OceancdVerificationTemplateMetricsProviderCloudWatchMetricDataQueriesMetricStat `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.1/docs/resources/oceancd_verification_template#period OceancdVerificationTemplate#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.1/docs/resources/oceancd_verification_template#return_data OceancdVerificationTemplate#return_data}.
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

