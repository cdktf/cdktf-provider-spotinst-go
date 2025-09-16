// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationcenter


type NotificationCenterComputePolicyConfigDynamicRules struct {
	// filter_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/notification_center#filter_conditions NotificationCenter#filter_conditions}
	FilterConditions interface{} `field:"optional" json:"filterConditions" yaml:"filterConditions"`
}

