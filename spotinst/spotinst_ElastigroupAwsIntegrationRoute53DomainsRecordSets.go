// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsIntegrationRoute53DomainsRecordSets struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#name ElastigroupAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#use_public_dns ElastigroupAws#use_public_dns}.
	UsePublicDns interface{} `field:"optional" json:"usePublicDns" yaml:"usePublicDns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#use_public_ip ElastigroupAws#use_public_ip}.
	UsePublicIp interface{} `field:"optional" json:"usePublicIp" yaml:"usePublicIp"`
}

