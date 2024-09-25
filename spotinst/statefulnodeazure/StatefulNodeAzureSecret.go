// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureSecret struct {
	// source_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/stateful_node_azure#source_vault StatefulNodeAzure#source_vault}
	SourceVault interface{} `field:"required" json:"sourceVault" yaml:"sourceVault"`
	// vault_certificates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/stateful_node_azure#vault_certificates StatefulNodeAzure#vault_certificates}
	VaultCertificates interface{} `field:"required" json:"vaultCertificates" yaml:"vaultCertificates"`
}

