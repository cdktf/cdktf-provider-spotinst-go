// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpBackendServices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/elastigroup_gcp#service_name ElastigroupGcp#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// backend_balancing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/elastigroup_gcp#backend_balancing ElastigroupGcp#backend_balancing}
	BackendBalancing *ElastigroupGcpBackendServicesBackendBalancing `field:"optional" json:"backendBalancing" yaml:"backendBalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/elastigroup_gcp#location_type ElastigroupGcp#location_type}.
	LocationType *string `field:"optional" json:"locationType" yaml:"locationType"`
	// named_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/elastigroup_gcp#named_ports ElastigroupGcp#named_ports}
	NamedPorts interface{} `field:"optional" json:"namedPorts" yaml:"namedPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/elastigroup_gcp#scheme ElastigroupGcp#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

