// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationcenter


type NotificationCenterRegisteredUsers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/notification_center#subscription_types NotificationCenter#subscription_types}.
	SubscriptionTypes *[]*string `field:"optional" json:"subscriptionTypes" yaml:"subscriptionTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/notification_center#user_email NotificationCenter#user_email}.
	UserEmail *string `field:"optional" json:"userEmail" yaml:"userEmail"`
}

