package oceangkeimport


type OceanGkeImportScheduledTask struct {
	// shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/ocean_gke_import#shutdown_hours OceanGkeImport#shutdown_hours}
	ShutdownHours *OceanGkeImportScheduledTaskShutdownHours `field:"optional" json:"shutdownHours" yaml:"shutdownHours"`
	// tasks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/ocean_gke_import#tasks OceanGkeImport#tasks}
	Tasks interface{} `field:"optional" json:"tasks" yaml:"tasks"`
}

