package oceanaws


type OceanAwsLogging struct {
	// export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/ocean_aws#export OceanAws#export}
	Export *OceanAwsLoggingExport `field:"optional" json:"export" yaml:"export"`
}

