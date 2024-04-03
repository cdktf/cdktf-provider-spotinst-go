// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureStrategyCapacityReservation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.1/docs/resources/stateful_node_azure#should_utilize StatefulNodeAzure#should_utilize}.
	ShouldUtilize interface{} `field:"required" json:"shouldUtilize" yaml:"shouldUtilize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.1/docs/resources/stateful_node_azure#utilization_strategy StatefulNodeAzure#utilization_strategy}.
	UtilizationStrategy *string `field:"required" json:"utilizationStrategy" yaml:"utilizationStrategy"`
	// capacity_reservation_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.1/docs/resources/stateful_node_azure#capacity_reservation_groups StatefulNodeAzure#capacity_reservation_groups}
	CapacityReservationGroups interface{} `field:"optional" json:"capacityReservationGroups" yaml:"capacityReservationGroups"`
}

