// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationcenter


type NotificationCenterComputePolicyConfig struct {
	// events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/notification_center#events NotificationCenter#events}
	Events interface{} `field:"required" json:"events" yaml:"events"`
	// dynamic_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/notification_center#dynamic_rules NotificationCenter#dynamic_rules}
	DynamicRules interface{} `field:"optional" json:"dynamicRules" yaml:"dynamicRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/notification_center#resource_ids NotificationCenter#resource_ids}.
	ResourceIds *[]*string `field:"optional" json:"resourceIds" yaml:"resourceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/notification_center#should_include_all_resources NotificationCenter#should_include_all_resources}.
	ShouldIncludeAllResources interface{} `field:"optional" json:"shouldIncludeAllResources" yaml:"shouldIncludeAllResources"`
}

