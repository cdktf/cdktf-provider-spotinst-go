// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark


type OceanSparkCompute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_spark#create_vngs OceanSpark#create_vngs}.
	CreateVngs interface{} `field:"optional" json:"createVngs" yaml:"createVngs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_spark#use_taints OceanSpark#use_taints}.
	UseTaints interface{} `field:"optional" json:"useTaints" yaml:"useTaints"`
}

