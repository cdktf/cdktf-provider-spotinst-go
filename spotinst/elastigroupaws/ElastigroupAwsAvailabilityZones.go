// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsAvailabilityZones struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.212.0/docs/resources/elastigroup_aws#availability_zones_name ElastigroupAws#availability_zones_name}.
	AvailabilityZonesName *string `field:"required" json:"availabilityZonesName" yaml:"availabilityZonesName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.212.0/docs/resources/elastigroup_aws#placement_group_name ElastigroupAws#placement_group_name}.
	PlacementGroupName *string `field:"optional" json:"placementGroupName" yaml:"placementGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.212.0/docs/resources/elastigroup_aws#subnet_ids ElastigroupAws#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

