// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAwsLaunchSpecSchedulingTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#cron_expression OceanAwsLaunchSpec#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#is_enabled OceanAwsLaunchSpec#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#task_type OceanAwsLaunchSpec#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// task_headroom block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#task_headroom OceanAwsLaunchSpec#task_headroom}
	TaskHeadroom interface{} `field:"optional" json:"taskHeadroom" yaml:"taskHeadroom"`
}
