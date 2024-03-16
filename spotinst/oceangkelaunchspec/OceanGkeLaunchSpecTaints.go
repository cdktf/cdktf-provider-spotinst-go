// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkelaunchspec


type OceanGkeLaunchSpecTaints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_gke_launch_spec#effect OceanGkeLaunchSpec#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_gke_launch_spec#key OceanGkeLaunchSpec#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_gke_launch_spec#value OceanGkeLaunchSpec#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

