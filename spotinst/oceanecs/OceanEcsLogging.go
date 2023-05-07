package oceanecs


type OceanEcsLogging struct {
	// export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.116.0/docs/resources/ocean_ecs#export OceanEcs#export}
	Export *OceanEcsLoggingExport `field:"optional" json:"export" yaml:"export"`
}

