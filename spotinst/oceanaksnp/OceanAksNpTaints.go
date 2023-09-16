// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpTaints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/ocean_aks_np#effect OceanAksNp#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/ocean_aks_np#key OceanAksNp#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/ocean_aks_np#value OceanAksNp#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

