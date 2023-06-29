package oceanaksnp


type OceanAksNpAutoscaler struct {
	// autoscale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/ocean_aks_np#autoscale_down OceanAksNp#autoscale_down}
	AutoscaleDown *OceanAksNpAutoscalerAutoscaleDown `field:"optional" json:"autoscaleDown" yaml:"autoscaleDown"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/ocean_aks_np#autoscale_headroom OceanAksNp#autoscale_headroom}
	AutoscaleHeadroom *OceanAksNpAutoscalerAutoscaleHeadroom `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/ocean_aks_np#autoscale_is_enabled OceanAksNp#autoscale_is_enabled}.
	AutoscaleIsEnabled interface{} `field:"optional" json:"autoscaleIsEnabled" yaml:"autoscaleIsEnabled"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/ocean_aks_np#resource_limits OceanAksNp#resource_limits}
	ResourceLimits *OceanAksNpAutoscalerResourceLimits `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

