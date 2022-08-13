// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAksVirtualNodeGroupLaunchSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#max_pods OceanAksVirtualNodeGroup#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#os_disk OceanAksVirtualNodeGroup#os_disk}
	OsDisk *OceanAksVirtualNodeGroupLaunchSpecificationOsDisk `field:"optional" json:"osDisk" yaml:"osDisk"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#tag OceanAksVirtualNodeGroup#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

