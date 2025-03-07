// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3SecretVaultCertificates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#certificate_store ElastigroupAzureV3#certificate_store}.
	CertificateStore *string `field:"required" json:"certificateStore" yaml:"certificateStore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#certificate_url ElastigroupAzureV3#certificate_url}.
	CertificateUrl *string `field:"required" json:"certificateUrl" yaml:"certificateUrl"`
}

