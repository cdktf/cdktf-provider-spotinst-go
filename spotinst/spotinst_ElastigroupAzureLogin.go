// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAzureLogin struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#user_name ElastigroupAzure#user_name}.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#password ElastigroupAzure#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#ssh_public_key ElastigroupAzure#ssh_public_key}.
	SshPublicKey *string `field:"optional" json:"sshPublicKey" yaml:"sshPublicKey"`
}

