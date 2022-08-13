// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeImportUpdatePolicyRollConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#batch_size_percentage OceanGkeImport#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"required" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#batch_min_healthy_percentage OceanGkeImport#batch_min_healthy_percentage}.
	BatchMinHealthyPercentage *float64 `field:"optional" json:"batchMinHealthyPercentage" yaml:"batchMinHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#launch_spec_ids OceanGkeImport#launch_spec_ids}.
	LaunchSpecIds *[]*string `field:"optional" json:"launchSpecIds" yaml:"launchSpecIds"`
}

