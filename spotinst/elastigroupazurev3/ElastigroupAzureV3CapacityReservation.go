// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3CapacityReservation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/elastigroup_azure_v3#should_utilize ElastigroupAzureV3#should_utilize}.
	ShouldUtilize interface{} `field:"required" json:"shouldUtilize" yaml:"shouldUtilize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/elastigroup_azure_v3#utilization_strategy ElastigroupAzureV3#utilization_strategy}.
	UtilizationStrategy *string `field:"required" json:"utilizationStrategy" yaml:"utilizationStrategy"`
	// capacity_reservation_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/elastigroup_azure_v3#capacity_reservation_groups ElastigroupAzureV3#capacity_reservation_groups}
	CapacityReservationGroups *ElastigroupAzureV3CapacityReservationCapacityReservationGroups `field:"optional" json:"capacityReservationGroups" yaml:"capacityReservationGroups"`
}

