// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureExtension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/stateful_node_azure#api_version StatefulNodeAzure#api_version}.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/stateful_node_azure#minor_version_auto_upgrade StatefulNodeAzure#minor_version_auto_upgrade}.
	MinorVersionAutoUpgrade interface{} `field:"required" json:"minorVersionAutoUpgrade" yaml:"minorVersionAutoUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/stateful_node_azure#name StatefulNodeAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/stateful_node_azure#publisher StatefulNodeAzure#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/stateful_node_azure#type StatefulNodeAzure#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/stateful_node_azure#protected_settings StatefulNodeAzure#protected_settings}.
	ProtectedSettings *map[string]*string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/stateful_node_azure#public_settings StatefulNodeAzure#public_settings}.
	PublicSettings *map[string]*string `field:"optional" json:"publicSettings" yaml:"publicSettings"`
}

