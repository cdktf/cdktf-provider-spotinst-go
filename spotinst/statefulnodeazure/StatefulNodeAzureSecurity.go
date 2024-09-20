// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/stateful_node_azure#confidential_os_disk_encryption StatefulNodeAzure#confidential_os_disk_encryption}.
	ConfidentialOsDiskEncryption interface{} `field:"optional" json:"confidentialOsDiskEncryption" yaml:"confidentialOsDiskEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/stateful_node_azure#encryption_at_host StatefulNodeAzure#encryption_at_host}.
	EncryptionAtHost interface{} `field:"optional" json:"encryptionAtHost" yaml:"encryptionAtHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/stateful_node_azure#secure_boot_enabled StatefulNodeAzure#secure_boot_enabled}.
	SecureBootEnabled interface{} `field:"optional" json:"secureBootEnabled" yaml:"secureBootEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/stateful_node_azure#security_type StatefulNodeAzure#security_type}.
	SecurityType *string `field:"optional" json:"securityType" yaml:"securityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/stateful_node_azure#vtpm_enabled StatefulNodeAzure#vtpm_enabled}.
	VtpmEnabled interface{} `field:"optional" json:"vtpmEnabled" yaml:"vtpmEnabled"`
}

