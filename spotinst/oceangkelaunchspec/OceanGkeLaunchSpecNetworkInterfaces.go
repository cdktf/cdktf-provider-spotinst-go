// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkelaunchspec


type OceanGkeLaunchSpecNetworkInterfaces struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_gke_launch_spec#network OceanGkeLaunchSpec#network}.
	Network *string `field:"required" json:"network" yaml:"network"`
	// access_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_gke_launch_spec#access_configs OceanGkeLaunchSpec#access_configs}
	AccessConfigs interface{} `field:"optional" json:"accessConfigs" yaml:"accessConfigs"`
	// alias_ip_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_gke_launch_spec#alias_ip_ranges OceanGkeLaunchSpec#alias_ip_ranges}
	AliasIpRanges interface{} `field:"optional" json:"aliasIpRanges" yaml:"aliasIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_gke_launch_spec#project_id OceanGkeLaunchSpec#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

