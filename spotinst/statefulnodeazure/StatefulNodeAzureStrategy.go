// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#fallback_to_on_demand StatefulNodeAzure#fallback_to_on_demand}.
	FallbackToOnDemand interface{} `field:"required" json:"fallbackToOnDemand" yaml:"fallbackToOnDemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#availability_vs_cost StatefulNodeAzure#availability_vs_cost}.
	AvailabilityVsCost *float64 `field:"optional" json:"availabilityVsCost" yaml:"availabilityVsCost"`
	// capacity_reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#capacity_reservation StatefulNodeAzure#capacity_reservation}
	CapacityReservation interface{} `field:"optional" json:"capacityReservation" yaml:"capacityReservation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#draining_timeout StatefulNodeAzure#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#od_windows StatefulNodeAzure#od_windows}.
	OdWindows *[]*string `field:"optional" json:"odWindows" yaml:"odWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#optimization_windows StatefulNodeAzure#optimization_windows}.
	OptimizationWindows *[]*string `field:"optional" json:"optimizationWindows" yaml:"optimizationWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#preferred_life_cycle StatefulNodeAzure#preferred_life_cycle}.
	PreferredLifeCycle *string `field:"optional" json:"preferredLifeCycle" yaml:"preferredLifeCycle"`
	// revert_to_spot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#revert_to_spot StatefulNodeAzure#revert_to_spot}
	RevertToSpot *StatefulNodeAzureStrategyRevertToSpot `field:"optional" json:"revertToSpot" yaml:"revertToSpot"`
}

