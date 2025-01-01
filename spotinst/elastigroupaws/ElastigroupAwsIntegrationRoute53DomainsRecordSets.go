// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationRoute53DomainsRecordSets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/elastigroup_aws#name ElastigroupAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/elastigroup_aws#use_public_dns ElastigroupAws#use_public_dns}.
	UsePublicDns interface{} `field:"optional" json:"usePublicDns" yaml:"usePublicDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/elastigroup_aws#use_public_ip ElastigroupAws#use_public_ip}.
	UsePublicIp interface{} `field:"optional" json:"usePublicIp" yaml:"usePublicIp"`
}

