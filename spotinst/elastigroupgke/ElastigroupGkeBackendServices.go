// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke


type ElastigroupGkeBackendServices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.210.0/docs/resources/elastigroup_gke#service_name ElastigroupGke#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.210.0/docs/resources/elastigroup_gke#location_type ElastigroupGke#location_type}.
	LocationType *string `field:"optional" json:"locationType" yaml:"locationType"`
	// named_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.210.0/docs/resources/elastigroup_gke#named_ports ElastigroupGke#named_ports}
	NamedPorts interface{} `field:"optional" json:"namedPorts" yaml:"namedPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.210.0/docs/resources/elastigroup_gke#scheme ElastigroupGke#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

