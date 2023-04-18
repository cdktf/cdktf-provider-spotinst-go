package oceanecslaunchspec


type OceanEcsLaunchSpecInstanceMetadataOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/ocean_ecs_launch_spec#http_tokens OceanEcsLaunchSpec#http_tokens}.
	HttpTokens *string `field:"required" json:"httpTokens" yaml:"httpTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/ocean_ecs_launch_spec#http_put_response_hop_limit OceanEcsLaunchSpec#http_put_response_hop_limit}.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
}

