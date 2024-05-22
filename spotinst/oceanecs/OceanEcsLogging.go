// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs


type OceanEcsLogging struct {
	// export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/ocean_ecs#export OceanEcs#export}
	Export *OceanEcsLoggingExport `field:"optional" json:"export" yaml:"export"`
}

