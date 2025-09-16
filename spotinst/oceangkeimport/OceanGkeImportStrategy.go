// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport


type OceanGkeImportStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_import#draining_timeout OceanGkeImport#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_import#preemptible_percentage OceanGkeImport#preemptible_percentage}.
	PreemptiblePercentage *float64 `field:"optional" json:"preemptiblePercentage" yaml:"preemptiblePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_import#provisioning_model OceanGkeImport#provisioning_model}.
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_import#scaling_orientation OceanGkeImport#scaling_orientation}.
	ScalingOrientation *string `field:"optional" json:"scalingOrientation" yaml:"scalingOrientation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_import#should_utilize_commitments OceanGkeImport#should_utilize_commitments}.
	ShouldUtilizeCommitments interface{} `field:"optional" json:"shouldUtilizeCommitments" yaml:"shouldUtilizeCommitments"`
}

