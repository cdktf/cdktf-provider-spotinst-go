// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGcpScheduledTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#task_type ElastigroupGcp#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#cron_expression ElastigroupGcp#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#is_enabled ElastigroupGcp#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#max_capacity ElastigroupGcp#max_capacity}.
	MaxCapacity *string `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#min_capacity ElastigroupGcp#min_capacity}.
	MinCapacity *string `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#target_capacity ElastigroupGcp#target_capacity}.
	TargetCapacity *string `field:"optional" json:"targetCapacity" yaml:"targetCapacity"`
}

