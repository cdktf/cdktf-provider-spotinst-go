// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package multaitargetset


type MultaiTargetSetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/multai_target_set#key MultaiTargetSet#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/multai_target_set#value MultaiTargetSet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

