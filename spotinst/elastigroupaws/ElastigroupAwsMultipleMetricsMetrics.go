// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsMultipleMetricsMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/elastigroup_aws#metric_name ElastigroupAws#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/elastigroup_aws#name ElastigroupAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/elastigroup_aws#namespace ElastigroupAws#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/elastigroup_aws#dimensions ElastigroupAws#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/elastigroup_aws#extended_statistic ElastigroupAws#extended_statistic}.
	ExtendedStatistic *string `field:"optional" json:"extendedStatistic" yaml:"extendedStatistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/elastigroup_aws#statistic ElastigroupAws#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/elastigroup_aws#unit ElastigroupAws#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

