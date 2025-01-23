// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdVerificationProviderConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.208.0/docs/resources/oceancd_verification_provider#cluster_ids OceancdVerificationProvider#cluster_ids}.
	ClusterIds *[]*string `field:"required" json:"clusterIds" yaml:"clusterIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.208.0/docs/resources/oceancd_verification_provider#name OceancdVerificationProvider#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloud_watch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.208.0/docs/resources/oceancd_verification_provider#cloud_watch OceancdVerificationProvider#cloud_watch}
	CloudWatch *OceancdVerificationProviderCloudWatch `field:"optional" json:"cloudWatch" yaml:"cloudWatch"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.208.0/docs/resources/oceancd_verification_provider#datadog OceancdVerificationProvider#datadog}
	Datadog *OceancdVerificationProviderDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.208.0/docs/resources/oceancd_verification_provider#id OceancdVerificationProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// jenkins block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.208.0/docs/resources/oceancd_verification_provider#jenkins OceancdVerificationProvider#jenkins}
	Jenkins *OceancdVerificationProviderJenkins `field:"optional" json:"jenkins" yaml:"jenkins"`
	// new_relic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.208.0/docs/resources/oceancd_verification_provider#new_relic OceancdVerificationProvider#new_relic}
	NewRelic *OceancdVerificationProviderNewRelic `field:"optional" json:"newRelic" yaml:"newRelic"`
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.208.0/docs/resources/oceancd_verification_provider#prometheus OceancdVerificationProvider#prometheus}
	Prometheus *OceancdVerificationProviderPrometheus `field:"optional" json:"prometheus" yaml:"prometheus"`
}

