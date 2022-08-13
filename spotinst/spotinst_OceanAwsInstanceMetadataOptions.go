// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAwsInstanceMetadataOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#http_tokens OceanAws#http_tokens}.
	HttpTokens *string `field:"required" json:"httpTokens" yaml:"httpTokens"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#http_put_response_hop_limit OceanAws#http_put_response_hop_limit}.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
}

