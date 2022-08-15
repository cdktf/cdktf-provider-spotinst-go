// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeImportStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#draining_timeout OceanGkeImport#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#preemptible_percentage OceanGkeImport#preemptible_percentage}.
	PreemptiblePercentage *float64 `field:"optional" json:"preemptiblePercentage" yaml:"preemptiblePercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#provisioning_model OceanGkeImport#provisioning_model}.
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
}
