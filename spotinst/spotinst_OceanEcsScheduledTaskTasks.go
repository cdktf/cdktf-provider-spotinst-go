// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsScheduledTaskTasks struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#cron_expression OceanEcs#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#is_enabled OceanEcs#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#task_type OceanEcs#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
}

