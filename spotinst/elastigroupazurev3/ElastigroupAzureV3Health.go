// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3Health struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/elastigroup_azure_v3#auto_healing ElastigroupAzureV3#auto_healing}.
	AutoHealing interface{} `field:"optional" json:"autoHealing" yaml:"autoHealing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/elastigroup_azure_v3#grace_period ElastigroupAzureV3#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/elastigroup_azure_v3#health_check_types ElastigroupAzureV3#health_check_types}.
	HealthCheckTypes *[]*string `field:"optional" json:"healthCheckTypes" yaml:"healthCheckTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/elastigroup_azure_v3#unhealthy_duration ElastigroupAzureV3#unhealthy_duration}.
	UnhealthyDuration *float64 `field:"optional" json:"unhealthyDuration" yaml:"unhealthyDuration"`
}

