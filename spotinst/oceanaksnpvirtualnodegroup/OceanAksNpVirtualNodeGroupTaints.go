package oceanaksnpvirtualnodegroup


type OceanAksNpVirtualNodeGroupTaints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_np_virtual_node_group#effect OceanAksNpVirtualNodeGroup#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_np_virtual_node_group#key OceanAksNpVirtualNodeGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_np_virtual_node_group#value OceanAksNpVirtualNodeGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

