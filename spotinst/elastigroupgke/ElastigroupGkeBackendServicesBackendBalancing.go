// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke


type ElastigroupGkeBackendServicesBackendBalancing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.0/docs/resources/elastigroup_gke#backend_balancing_mode ElastigroupGke#backend_balancing_mode}.
	BackendBalancingMode *string `field:"optional" json:"backendBalancingMode" yaml:"backendBalancingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.0/docs/resources/elastigroup_gke#max_rate_per_instance ElastigroupGke#max_rate_per_instance}.
	MaxRatePerInstance *float64 `field:"optional" json:"maxRatePerInstance" yaml:"maxRatePerInstance"`
}

