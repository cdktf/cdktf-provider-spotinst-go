// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3Extensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#api_version ElastigroupAzureV3#api_version}.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#minor_version_auto_upgrade ElastigroupAzureV3#minor_version_auto_upgrade}.
	MinorVersionAutoUpgrade interface{} `field:"required" json:"minorVersionAutoUpgrade" yaml:"minorVersionAutoUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#name ElastigroupAzureV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#publisher ElastigroupAzureV3#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#type ElastigroupAzureV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#enable_automatic_upgrade ElastigroupAzureV3#enable_automatic_upgrade}.
	EnableAutomaticUpgrade interface{} `field:"optional" json:"enableAutomaticUpgrade" yaml:"enableAutomaticUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#protected_settings ElastigroupAzureV3#protected_settings}.
	ProtectedSettings *map[string]*string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// protected_settings_from_key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#protected_settings_from_key_vault ElastigroupAzureV3#protected_settings_from_key_vault}
	ProtectedSettingsFromKeyVault *ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVault `field:"optional" json:"protectedSettingsFromKeyVault" yaml:"protectedSettingsFromKeyVault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/elastigroup_azure_v3#public_settings ElastigroupAzureV3#public_settings}.
	PublicSettings *map[string]*string `field:"optional" json:"publicSettings" yaml:"publicSettings"`
}

