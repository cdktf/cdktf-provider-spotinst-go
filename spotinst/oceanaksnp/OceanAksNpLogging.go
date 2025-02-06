// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpLogging struct {
	// export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.1/docs/resources/ocean_aks_np#export OceanAksNp#export}
	Export *OceanAksNpLoggingExport `field:"optional" json:"export" yaml:"export"`
}

