// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks


type OceanAksNetworkNetworkInterfaceAdditionalIpConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.1/docs/resources/ocean_aks#name OceanAks#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.1/docs/resources/ocean_aks#private_ip_version OceanAks#private_ip_version}.
	PrivateIpVersion *string `field:"optional" json:"privateIpVersion" yaml:"privateIpVersion"`
}

