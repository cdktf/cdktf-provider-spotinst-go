// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeImportScheduledTaskTasks struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#cron_expression OceanGkeImport#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#is_enabled OceanGkeImport#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#task_type OceanGkeImport#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#batch_size_percentage OceanGkeImport#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
}

