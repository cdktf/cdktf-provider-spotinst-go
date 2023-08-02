package oceanaws


type OceanAwsAutoscaler struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/ocean_aws#auto_headroom_percentage OceanAws#auto_headroom_percentage}.
	AutoHeadroomPercentage *float64 `field:"optional" json:"autoHeadroomPercentage" yaml:"autoHeadroomPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/ocean_aws#autoscale_cooldown OceanAws#autoscale_cooldown}.
	AutoscaleCooldown *float64 `field:"optional" json:"autoscaleCooldown" yaml:"autoscaleCooldown"`
	// autoscale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/ocean_aws#autoscale_down OceanAws#autoscale_down}
	AutoscaleDown *OceanAwsAutoscalerAutoscaleDown `field:"optional" json:"autoscaleDown" yaml:"autoscaleDown"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/ocean_aws#autoscale_headroom OceanAws#autoscale_headroom}
	AutoscaleHeadroom *OceanAwsAutoscalerAutoscaleHeadroom `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/ocean_aws#autoscale_is_auto_config OceanAws#autoscale_is_auto_config}.
	AutoscaleIsAutoConfig interface{} `field:"optional" json:"autoscaleIsAutoConfig" yaml:"autoscaleIsAutoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/ocean_aws#autoscale_is_enabled OceanAws#autoscale_is_enabled}.
	AutoscaleIsEnabled interface{} `field:"optional" json:"autoscaleIsEnabled" yaml:"autoscaleIsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/ocean_aws#enable_automatic_and_manual_headroom OceanAws#enable_automatic_and_manual_headroom}.
	EnableAutomaticAndManualHeadroom interface{} `field:"optional" json:"enableAutomaticAndManualHeadroom" yaml:"enableAutomaticAndManualHeadroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/ocean_aws#extended_resource_definitions OceanAws#extended_resource_definitions}.
	ExtendedResourceDefinitions *[]*string `field:"optional" json:"extendedResourceDefinitions" yaml:"extendedResourceDefinitions"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/ocean_aws#resource_limits OceanAws#resource_limits}
	ResourceLimits *OceanAwsAutoscalerResourceLimits `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

