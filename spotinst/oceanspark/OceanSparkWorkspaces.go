// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark


type OceanSparkWorkspaces struct {
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/ocean_spark#storage OceanSpark#storage}
	Storage *OceanSparkWorkspacesStorage `field:"optional" json:"storage" yaml:"storage"`
}

