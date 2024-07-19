// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsItfTargetGroupConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#health_check_path ElastigroupAws#health_check_path}.
	HealthCheckPath *string `field:"required" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#port ElastigroupAws#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#protocol ElastigroupAws#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#vpc_id ElastigroupAws#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#health_check_interval_seconds ElastigroupAws#health_check_interval_seconds}.
	HealthCheckIntervalSeconds *float64 `field:"optional" json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#health_check_port ElastigroupAws#health_check_port}.
	HealthCheckPort *string `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#health_check_protocol ElastigroupAws#health_check_protocol}.
	HealthCheckProtocol *string `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#health_check_timeout_seconds ElastigroupAws#health_check_timeout_seconds}.
	HealthCheckTimeoutSeconds *float64 `field:"optional" json:"healthCheckTimeoutSeconds" yaml:"healthCheckTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#healthy_threshold_count ElastigroupAws#healthy_threshold_count}.
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#matcher ElastigroupAws#matcher}
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#protocol_version ElastigroupAws#protocol_version}.
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#tags ElastigroupAws#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/elastigroup_aws#unhealthy_threshold_count ElastigroupAws#unhealthy_threshold_count}.
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
}

