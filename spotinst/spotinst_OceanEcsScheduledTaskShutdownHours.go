// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsScheduledTaskShutdownHours struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#time_windows OceanEcs#time_windows}.
	TimeWindows *[]*string `field:"required" json:"timeWindows" yaml:"timeWindows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#is_enabled OceanEcs#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

