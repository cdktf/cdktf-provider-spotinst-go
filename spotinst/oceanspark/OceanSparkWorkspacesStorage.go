// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark


type OceanSparkWorkspacesStorage struct {
	// defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/ocean_spark#defaults OceanSpark#defaults}
	Defaults *OceanSparkWorkspacesStorageDefaults `field:"optional" json:"defaults" yaml:"defaults"`
}

