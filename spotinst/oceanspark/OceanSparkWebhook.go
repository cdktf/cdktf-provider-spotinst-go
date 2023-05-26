package oceanspark


type OceanSparkWebhook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/ocean_spark#host_network_ports OceanSpark#host_network_ports}.
	HostNetworkPorts *[]*float64 `field:"optional" json:"hostNetworkPorts" yaml:"hostNetworkPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/ocean_spark#use_host_network OceanSpark#use_host_network}.
	UseHostNetwork interface{} `field:"optional" json:"useHostNetwork" yaml:"useHostNetwork"`
}

