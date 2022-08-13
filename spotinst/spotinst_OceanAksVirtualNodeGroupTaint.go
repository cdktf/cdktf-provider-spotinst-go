// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAksVirtualNodeGroupTaint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#effect OceanAksVirtualNodeGroup#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#key OceanAksVirtualNodeGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#value OceanAksVirtualNodeGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

