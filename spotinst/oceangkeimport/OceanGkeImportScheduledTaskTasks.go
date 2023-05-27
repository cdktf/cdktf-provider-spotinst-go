package oceangkeimport


type OceanGkeImportScheduledTaskTasks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_gke_import#cron_expression OceanGkeImport#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_gke_import#is_enabled OceanGkeImport#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_gke_import#task_type OceanGkeImport#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// task_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_gke_import#task_parameters OceanGkeImport#task_parameters}
	TaskParameters *OceanGkeImportScheduledTaskTasksTaskParameters `field:"optional" json:"taskParameters" yaml:"taskParameters"`
}

