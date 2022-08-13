// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsInstanceMetadataOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#http_tokens OceanEcs#http_tokens}.
	HttpTokens *string `field:"required" json:"httpTokens" yaml:"httpTokens"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#http_put_response_hop_limit OceanEcs#http_put_response_hop_limit}.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
}

