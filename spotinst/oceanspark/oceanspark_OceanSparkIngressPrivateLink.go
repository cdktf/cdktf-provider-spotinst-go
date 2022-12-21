package oceanspark


type OceanSparkIngressPrivateLink struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#enabled OceanSpark#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#vpc_endpoint_service OceanSpark#vpc_endpoint_service}.
	VpcEndpointService *string `field:"optional" json:"vpcEndpointService" yaml:"vpcEndpointService"`
}

