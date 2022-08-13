// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeLaunchSpecSchedulingTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#cron_expression OceanGkeLaunchSpec#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#is_enabled OceanGkeLaunchSpec#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#task_type OceanGkeLaunchSpec#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// task_headroom block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#task_headroom OceanGkeLaunchSpec#task_headroom}
	TaskHeadroom interface{} `field:"optional" json:"taskHeadroom" yaml:"taskHeadroom"`
}

