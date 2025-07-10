// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpBackendServicesBackendBalancing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gcp#backend_balancing_mode ElastigroupGcp#backend_balancing_mode}.
	BackendBalancingMode *string `field:"optional" json:"backendBalancingMode" yaml:"backendBalancingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gcp#max_rate_per_instance ElastigroupGcp#max_rate_per_instance}.
	MaxRatePerInstance *float64 `field:"optional" json:"maxRatePerInstance" yaml:"maxRatePerInstance"`
}

