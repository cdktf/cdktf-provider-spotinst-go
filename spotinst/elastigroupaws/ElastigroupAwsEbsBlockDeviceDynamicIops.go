// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsEbsBlockDeviceDynamicIops struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_aws#base_size ElastigroupAws#base_size}.
	BaseSize *float64 `field:"optional" json:"baseSize" yaml:"baseSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_aws#resource ElastigroupAws#resource}.
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_aws#size_per_resource_unit ElastigroupAws#size_per_resource_unit}.
	SizePerResourceUnit *float64 `field:"optional" json:"sizePerResourceUnit" yaml:"sizePerResourceUnit"`
}

