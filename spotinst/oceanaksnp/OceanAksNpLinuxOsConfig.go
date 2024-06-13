// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpLinuxOsConfig struct {
	// sysctls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/ocean_aks_np#sysctls OceanAksNp#sysctls}
	Sysctls interface{} `field:"optional" json:"sysctls" yaml:"sysctls"`
}

