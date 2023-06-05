package oceanaws


type OceanAwsScheduledTask struct {
	// shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/ocean_aws#shutdown_hours OceanAws#shutdown_hours}
	ShutdownHours *OceanAwsScheduledTaskShutdownHours `field:"optional" json:"shutdownHours" yaml:"shutdownHours"`
	// tasks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/ocean_aws#tasks OceanAws#tasks}
	Tasks interface{} `field:"optional" json:"tasks" yaml:"tasks"`
}

