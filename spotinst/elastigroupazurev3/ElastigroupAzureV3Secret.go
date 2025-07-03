// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3Secret struct {
	// source_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.221.0/docs/resources/elastigroup_azure_v3#source_vault ElastigroupAzureV3#source_vault}
	SourceVault interface{} `field:"required" json:"sourceVault" yaml:"sourceVault"`
	// vault_certificates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.221.0/docs/resources/elastigroup_azure_v3#vault_certificates ElastigroupAzureV3#vault_certificates}
	VaultCertificates interface{} `field:"required" json:"vaultCertificates" yaml:"vaultCertificates"`
}

