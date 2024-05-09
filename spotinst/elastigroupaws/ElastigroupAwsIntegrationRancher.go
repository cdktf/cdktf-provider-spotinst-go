// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationRancher struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.0/docs/resources/elastigroup_aws#access_key ElastigroupAws#access_key}.
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.0/docs/resources/elastigroup_aws#master_host ElastigroupAws#master_host}.
	MasterHost *string `field:"required" json:"masterHost" yaml:"masterHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.0/docs/resources/elastigroup_aws#secret_key ElastigroupAws#secret_key}.
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.0/docs/resources/elastigroup_aws#version ElastigroupAws#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

