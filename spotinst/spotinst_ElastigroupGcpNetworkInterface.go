// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGcpNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#network ElastigroupGcp#network}.
	Network *string `field:"required" json:"network" yaml:"network"`
	// access_configs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#access_configs ElastigroupGcp#access_configs}
	AccessConfigs interface{} `field:"optional" json:"accessConfigs" yaml:"accessConfigs"`
	// alias_ip_ranges block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#alias_ip_ranges ElastigroupGcp#alias_ip_ranges}
	AliasIpRanges interface{} `field:"optional" json:"aliasIpRanges" yaml:"aliasIpRanges"`
}

