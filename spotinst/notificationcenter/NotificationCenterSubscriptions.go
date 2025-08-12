// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationcenter


type NotificationCenterSubscriptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.0/docs/resources/notification_center#endpoint NotificationCenter#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.0/docs/resources/notification_center#subscription_type NotificationCenter#subscription_type}.
	SubscriptionType *string `field:"optional" json:"subscriptionType" yaml:"subscriptionType"`
}

