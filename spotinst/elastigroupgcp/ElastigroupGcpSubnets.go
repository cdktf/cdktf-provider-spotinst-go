package elastigroupgcp


type ElastigroupGcpSubnets struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#region ElastigroupGcp#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#subnet_names ElastigroupGcp#subnet_names}.
	SubnetNames *[]*string `field:"required" json:"subnetNames" yaml:"subnetNames"`
}

