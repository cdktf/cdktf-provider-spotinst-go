// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAksVirtualNodeGroupAutoscale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#auto_headroom_percentage OceanAksVirtualNodeGroup#auto_headroom_percentage}.
	AutoHeadroomPercentage *float64 `field:"optional" json:"autoHeadroomPercentage" yaml:"autoHeadroomPercentage"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#autoscale_headroom OceanAksVirtualNodeGroup#autoscale_headroom}
	AutoscaleHeadroom interface{} `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
}
