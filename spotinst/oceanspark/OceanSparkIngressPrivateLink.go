package oceanspark


type OceanSparkIngressPrivateLink struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/ocean_spark#enabled OceanSpark#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/ocean_spark#vpc_endpoint_service OceanSpark#vpc_endpoint_service}.
	VpcEndpointService *string `field:"optional" json:"vpcEndpointService" yaml:"vpcEndpointService"`
}

