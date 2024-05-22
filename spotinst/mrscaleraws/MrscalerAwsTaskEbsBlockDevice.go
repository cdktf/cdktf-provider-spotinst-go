// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrscaleraws


type MrscalerAwsTaskEbsBlockDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/mrscaler_aws#size_in_gb MrscalerAws#size_in_gb}.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/mrscaler_aws#volume_type MrscalerAws#volume_type}.
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/mrscaler_aws#iops MrscalerAws#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/mrscaler_aws#volumes_per_instance MrscalerAws#volumes_per_instance}.
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}

