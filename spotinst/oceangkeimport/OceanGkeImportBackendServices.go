// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport


type OceanGkeImportBackendServices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#service_name OceanGkeImport#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#location_type OceanGkeImport#location_type}.
	LocationType *string `field:"optional" json:"locationType" yaml:"locationType"`
	// named_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#named_ports OceanGkeImport#named_ports}
	NamedPorts interface{} `field:"optional" json:"namedPorts" yaml:"namedPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#scheme OceanGkeImport#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

