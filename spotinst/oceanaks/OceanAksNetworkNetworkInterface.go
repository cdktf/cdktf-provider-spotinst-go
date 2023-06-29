package oceanaks


type OceanAksNetworkNetworkInterface struct {
	// additional_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/ocean_aks#additional_ip_config OceanAks#additional_ip_config}
	AdditionalIpConfig interface{} `field:"optional" json:"additionalIpConfig" yaml:"additionalIpConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/ocean_aks#assign_public_ip OceanAks#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/ocean_aks#is_primary OceanAks#is_primary}.
	IsPrimary interface{} `field:"optional" json:"isPrimary" yaml:"isPrimary"`
	// security_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/ocean_aks#security_group OceanAks#security_group}
	SecurityGroup *OceanAksNetworkNetworkInterfaceSecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/ocean_aks#subnet_name OceanAks#subnet_name}.
	SubnetName *string `field:"optional" json:"subnetName" yaml:"subnetName"`
}

