// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsOptimizeImages struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#perform_at OceanEcs#perform_at}.
	PerformAt *string `field:"required" json:"performAt" yaml:"performAt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#should_optimize_ecs_ami OceanEcs#should_optimize_ecs_ami}.
	ShouldOptimizeEcsAmi interface{} `field:"required" json:"shouldOptimizeEcsAmi" yaml:"shouldOptimizeEcsAmi"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#time_windows OceanEcs#time_windows}.
	TimeWindows *[]*string `field:"optional" json:"timeWindows" yaml:"timeWindows"`
}
