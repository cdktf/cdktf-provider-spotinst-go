// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3VmSizesSpotSizeAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/elastigroup_azure_v3#max_cpu ElastigroupAzureV3#max_cpu}.
	MaxCpu *float64 `field:"optional" json:"maxCpu" yaml:"maxCpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/elastigroup_azure_v3#max_memory ElastigroupAzureV3#max_memory}.
	MaxMemory *float64 `field:"optional" json:"maxMemory" yaml:"maxMemory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/elastigroup_azure_v3#max_storage ElastigroupAzureV3#max_storage}.
	MaxStorage *float64 `field:"optional" json:"maxStorage" yaml:"maxStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/elastigroup_azure_v3#min_cpu ElastigroupAzureV3#min_cpu}.
	MinCpu *float64 `field:"optional" json:"minCpu" yaml:"minCpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/elastigroup_azure_v3#min_memory ElastigroupAzureV3#min_memory}.
	MinMemory *float64 `field:"optional" json:"minMemory" yaml:"minMemory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/elastigroup_azure_v3#min_storage ElastigroupAzureV3#min_storage}.
	MinStorage *float64 `field:"optional" json:"minStorage" yaml:"minStorage"`
}

