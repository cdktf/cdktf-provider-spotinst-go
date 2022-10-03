package oceanecs


type OceanEcsAutoscaler struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#auto_headroom_percentage OceanEcs#auto_headroom_percentage}.
	AutoHeadroomPercentage *float64 `field:"optional" json:"autoHeadroomPercentage" yaml:"autoHeadroomPercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#cooldown OceanEcs#cooldown}.
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// down block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#down OceanEcs#down}
	Down *OceanEcsAutoscalerDown `field:"optional" json:"down" yaml:"down"`
	// headroom block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#headroom OceanEcs#headroom}
	Headroom *OceanEcsAutoscalerHeadroom `field:"optional" json:"headroom" yaml:"headroom"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#is_auto_config OceanEcs#is_auto_config}.
	IsAutoConfig interface{} `field:"optional" json:"isAutoConfig" yaml:"isAutoConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#is_enabled OceanEcs#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#resource_limits OceanEcs#resource_limits}
	ResourceLimits *OceanEcsAutoscalerResourceLimits `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

