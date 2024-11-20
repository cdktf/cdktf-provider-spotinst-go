// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3Security struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/elastigroup_azure_v3#confidential_os_disk_encryption ElastigroupAzureV3#confidential_os_disk_encryption}.
	ConfidentialOsDiskEncryption interface{} `field:"optional" json:"confidentialOsDiskEncryption" yaml:"confidentialOsDiskEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/elastigroup_azure_v3#secure_boot_enabled ElastigroupAzureV3#secure_boot_enabled}.
	SecureBootEnabled interface{} `field:"optional" json:"secureBootEnabled" yaml:"secureBootEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/elastigroup_azure_v3#security_type ElastigroupAzureV3#security_type}.
	SecurityType *string `field:"optional" json:"securityType" yaml:"securityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/elastigroup_azure_v3#vtpm_enabled ElastigroupAzureV3#vtpm_enabled}.
	VtpmEnabled interface{} `field:"optional" json:"vtpmEnabled" yaml:"vtpmEnabled"`
}

