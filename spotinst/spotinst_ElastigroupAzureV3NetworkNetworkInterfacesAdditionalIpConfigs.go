// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAzureV3NetworkNetworkInterfacesAdditionalIpConfigs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure_v3#name ElastigroupAzureV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure_v3#private_ip_version ElastigroupAzureV3#private_ip_version}.
	PrivateIpVersion *string `field:"optional" json:"privateIpVersion" yaml:"privateIpVersion"`
}

