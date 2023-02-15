package oceanspark


type OceanSparkIngressCustomEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#address OceanSpark#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#enabled OceanSpark#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

