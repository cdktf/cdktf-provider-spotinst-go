// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeImportBackendServicesNamedPorts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#name OceanGkeImport#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#ports OceanGkeImport#ports}.
	Ports *[]*string `field:"required" json:"ports" yaml:"ports"`
}

