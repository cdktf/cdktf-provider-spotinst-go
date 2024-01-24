// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport


type OceanGkeImportUpdatePolicyRollConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.160.0/docs/resources/ocean_gke_import#batch_size_percentage OceanGkeImport#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"required" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.160.0/docs/resources/ocean_gke_import#batch_min_healthy_percentage OceanGkeImport#batch_min_healthy_percentage}.
	BatchMinHealthyPercentage *float64 `field:"optional" json:"batchMinHealthyPercentage" yaml:"batchMinHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.160.0/docs/resources/ocean_gke_import#launch_spec_ids OceanGkeImport#launch_spec_ids}.
	LaunchSpecIds *[]*string `field:"optional" json:"launchSpecIds" yaml:"launchSpecIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.160.0/docs/resources/ocean_gke_import#respect_pdb OceanGkeImport#respect_pdb}.
	RespectPdb interface{} `field:"optional" json:"respectPdb" yaml:"respectPdb"`
}

