package elastigroupaws


type ElastigroupAwsMetadataOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.123.0/docs/resources/elastigroup_aws#http_tokens ElastigroupAws#http_tokens}.
	HttpTokens *string `field:"required" json:"httpTokens" yaml:"httpTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.123.0/docs/resources/elastigroup_aws#http_put_response_hop_limit ElastigroupAws#http_put_response_hop_limit}.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.123.0/docs/resources/elastigroup_aws#instance_metadata_tags ElastigroupAws#instance_metadata_tags}.
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

