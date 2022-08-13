// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type StatefulNodeAzureSchedulingTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#cron_expression StatefulNodeAzure#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#is_enabled StatefulNodeAzure#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#type StatefulNodeAzure#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

