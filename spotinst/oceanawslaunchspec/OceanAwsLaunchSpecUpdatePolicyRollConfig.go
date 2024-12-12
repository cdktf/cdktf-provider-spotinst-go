// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecUpdatePolicyRollConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.201.0/docs/resources/ocean_aws_launch_spec#batch_size_percentage OceanAwsLaunchSpec#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"required" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.201.0/docs/resources/ocean_aws_launch_spec#respect_pdb OceanAwsLaunchSpec#respect_pdb}.
	RespectPdb interface{} `field:"optional" json:"respectPdb" yaml:"respectPdb"`
}

