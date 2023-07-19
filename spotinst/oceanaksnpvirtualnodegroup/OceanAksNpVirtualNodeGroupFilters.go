package oceanaksnpvirtualnodegroup


type OceanAksNpVirtualNodeGroupFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_aks_np_virtual_node_group#architectures OceanAksNpVirtualNodeGroup#architectures}.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_aks_np_virtual_node_group#max_memory_gib OceanAksNpVirtualNodeGroup#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_aks_np_virtual_node_group#max_vcpu OceanAksNpVirtualNodeGroup#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_aks_np_virtual_node_group#min_memory_gib OceanAksNpVirtualNodeGroup#min_memory_gib}.
	MinMemoryGib *float64 `field:"optional" json:"minMemoryGib" yaml:"minMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_aks_np_virtual_node_group#min_vcpu OceanAksNpVirtualNodeGroup#min_vcpu}.
	MinVcpu *float64 `field:"optional" json:"minVcpu" yaml:"minVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_aks_np_virtual_node_group#series OceanAksNpVirtualNodeGroup#series}.
	Series *[]*string `field:"optional" json:"series" yaml:"series"`
}

