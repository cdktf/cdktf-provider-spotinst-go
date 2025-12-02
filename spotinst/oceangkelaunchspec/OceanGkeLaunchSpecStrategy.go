// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkelaunchspec


type OceanGkeLaunchSpecStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/ocean_gke_launch_spec#preemptible_percentage OceanGkeLaunchSpec#preemptible_percentage}.
	PreemptiblePercentage *float64 `field:"optional" json:"preemptiblePercentage" yaml:"preemptiblePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/ocean_gke_launch_spec#scaling_orientation OceanGkeLaunchSpec#scaling_orientation}.
	ScalingOrientation *string `field:"optional" json:"scalingOrientation" yaml:"scalingOrientation"`
}

