// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTrafficIstioVirtualServicesTlsRoutes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/oceancd_rollout_spec#port OceancdRolloutSpec#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/oceancd_rollout_spec#sni_hosts OceancdRolloutSpec#sni_hosts}.
	SniHosts *[]*string `field:"optional" json:"sniHosts" yaml:"sniHosts"`
}

