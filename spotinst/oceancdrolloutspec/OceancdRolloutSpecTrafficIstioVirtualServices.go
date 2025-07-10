// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTrafficIstioVirtualServices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_rollout_spec#virtual_service_name OceancdRolloutSpec#virtual_service_name}.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
	// tls_routes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_rollout_spec#tls_routes OceancdRolloutSpec#tls_routes}
	TlsRoutes interface{} `field:"optional" json:"tlsRoutes" yaml:"tlsRoutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_rollout_spec#virtual_service_routes OceancdRolloutSpec#virtual_service_routes}.
	VirtualServiceRoutes *[]*string `field:"optional" json:"virtualServiceRoutes" yaml:"virtualServiceRoutes"`
}

