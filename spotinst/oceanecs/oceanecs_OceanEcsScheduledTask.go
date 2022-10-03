package oceanecs


type OceanEcsScheduledTask struct {
	// shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#shutdown_hours OceanEcs#shutdown_hours}
	ShutdownHours *OceanEcsScheduledTaskShutdownHours `field:"optional" json:"shutdownHours" yaml:"shutdownHours"`
	// tasks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#tasks OceanEcs#tasks}
	Tasks interface{} `field:"optional" json:"tasks" yaml:"tasks"`
}

