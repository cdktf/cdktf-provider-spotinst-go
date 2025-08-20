// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVault struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_azure_v3#secret_url ElastigroupAzureV3#secret_url}.
	SecretUrl *string `field:"required" json:"secretUrl" yaml:"secretUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_azure_v3#source_vault ElastigroupAzureV3#source_vault}.
	SourceVault *string `field:"required" json:"sourceVault" yaml:"sourceVault"`
}

