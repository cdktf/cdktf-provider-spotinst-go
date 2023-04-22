package elastigroupgcp


type ElastigroupGcpNetworkInterfaceAliasIpRanges struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/elastigroup_gcp#ip_cidr_range ElastigroupGcp#ip_cidr_range}.
	IpCidrRange *string `field:"required" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/elastigroup_gcp#subnetwork_range_name ElastigroupGcp#subnetwork_range_name}.
	SubnetworkRangeName *string `field:"required" json:"subnetworkRangeName" yaml:"subnetworkRangeName"`
}

