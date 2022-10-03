package elastigroupgke


type ElastigroupGkeNetworkInterfaceAliasIpRanges struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#ip_cidr_range ElastigroupGke#ip_cidr_range}.
	IpCidrRange *string `field:"required" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#subnetwork_range_name ElastigroupGke#subnetwork_range_name}.
	SubnetworkRangeName *string `field:"required" json:"subnetworkRangeName" yaml:"subnetworkRangeName"`
}

