// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsResourceRequirements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/elastigroup_aws#required_memory_maximum ElastigroupAws#required_memory_maximum}.
	RequiredMemoryMaximum *float64 `field:"required" json:"requiredMemoryMaximum" yaml:"requiredMemoryMaximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/elastigroup_aws#required_memory_minimum ElastigroupAws#required_memory_minimum}.
	RequiredMemoryMinimum *float64 `field:"required" json:"requiredMemoryMinimum" yaml:"requiredMemoryMinimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/elastigroup_aws#required_vcpu_maximum ElastigroupAws#required_vcpu_maximum}.
	RequiredVcpuMaximum *float64 `field:"required" json:"requiredVcpuMaximum" yaml:"requiredVcpuMaximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/elastigroup_aws#required_vcpu_minimum ElastigroupAws#required_vcpu_minimum}.
	RequiredVcpuMinimum *float64 `field:"required" json:"requiredVcpuMinimum" yaml:"requiredVcpuMinimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/elastigroup_aws#excluded_instance_families ElastigroupAws#excluded_instance_families}.
	ExcludedInstanceFamilies *[]*string `field:"optional" json:"excludedInstanceFamilies" yaml:"excludedInstanceFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/elastigroup_aws#excluded_instance_generations ElastigroupAws#excluded_instance_generations}.
	ExcludedInstanceGenerations *[]*string `field:"optional" json:"excludedInstanceGenerations" yaml:"excludedInstanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/elastigroup_aws#excluded_instance_types ElastigroupAws#excluded_instance_types}.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/elastigroup_aws#required_gpu_maximum ElastigroupAws#required_gpu_maximum}.
	RequiredGpuMaximum *float64 `field:"optional" json:"requiredGpuMaximum" yaml:"requiredGpuMaximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/elastigroup_aws#required_gpu_minimum ElastigroupAws#required_gpu_minimum}.
	RequiredGpuMinimum *float64 `field:"optional" json:"requiredGpuMinimum" yaml:"requiredGpuMinimum"`
}

