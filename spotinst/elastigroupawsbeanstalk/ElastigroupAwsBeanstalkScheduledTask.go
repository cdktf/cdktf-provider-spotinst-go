package elastigroupawsbeanstalk


type ElastigroupAwsBeanstalkScheduledTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#task_type ElastigroupAwsBeanstalk#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#adjustment ElastigroupAwsBeanstalk#adjustment}.
	Adjustment *string `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#adjustment_percentage ElastigroupAwsBeanstalk#adjustment_percentage}.
	AdjustmentPercentage *string `field:"optional" json:"adjustmentPercentage" yaml:"adjustmentPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#batch_size_percentage ElastigroupAwsBeanstalk#batch_size_percentage}.
	BatchSizePercentage *string `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#cron_expression ElastigroupAwsBeanstalk#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#frequency ElastigroupAwsBeanstalk#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#grace_period ElastigroupAwsBeanstalk#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#is_enabled ElastigroupAwsBeanstalk#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#max_capacity ElastigroupAwsBeanstalk#max_capacity}.
	MaxCapacity *string `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#min_capacity ElastigroupAwsBeanstalk#min_capacity}.
	MinCapacity *string `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#scale_max_capacity ElastigroupAwsBeanstalk#scale_max_capacity}.
	ScaleMaxCapacity *string `field:"optional" json:"scaleMaxCapacity" yaml:"scaleMaxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#scale_min_capacity ElastigroupAwsBeanstalk#scale_min_capacity}.
	ScaleMinCapacity *string `field:"optional" json:"scaleMinCapacity" yaml:"scaleMinCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#scale_target_capacity ElastigroupAwsBeanstalk#scale_target_capacity}.
	ScaleTargetCapacity *string `field:"optional" json:"scaleTargetCapacity" yaml:"scaleTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#start_time ElastigroupAwsBeanstalk#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_aws_beanstalk#target_capacity ElastigroupAwsBeanstalk#target_capacity}.
	TargetCapacity *string `field:"optional" json:"targetCapacity" yaml:"targetCapacity"`
}

