package oceanaws


type OceanAwsScheduledTaskTasks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/ocean_aws#cron_expression OceanAws#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/ocean_aws#is_enabled OceanAws#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/ocean_aws#task_type OceanAws#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
}

