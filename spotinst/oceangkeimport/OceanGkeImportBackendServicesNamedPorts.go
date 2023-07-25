package oceangkeimport


type OceanGkeImportBackendServicesNamedPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs/resources/ocean_gke_import#name OceanGkeImport#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs/resources/ocean_gke_import#ports OceanGkeImport#ports}.
	Ports *[]*string `field:"required" json:"ports" yaml:"ports"`
}

