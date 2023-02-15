package oceanaws


type OceanAwsLogging struct {
	// export block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#export OceanAws#export}
	Export *OceanAwsLoggingExport `field:"optional" json:"export" yaml:"export"`
}

