// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationcenter


type NotificationCenterComputePolicyConfigEvents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.0/docs/resources/notification_center#event NotificationCenter#event}.
	Event *string `field:"optional" json:"event" yaml:"event"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.0/docs/resources/notification_center#event_type NotificationCenter#event_type}.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
}

