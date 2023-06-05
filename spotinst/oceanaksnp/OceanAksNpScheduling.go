package oceanaksnp


type OceanAksNpScheduling struct {
	// shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/ocean_aks_np#shutdown_hours OceanAksNp#shutdown_hours}
	ShutdownHours *OceanAksNpSchedulingShutdownHours `field:"optional" json:"shutdownHours" yaml:"shutdownHours"`
}

