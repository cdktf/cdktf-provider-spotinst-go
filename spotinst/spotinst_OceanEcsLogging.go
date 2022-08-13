// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsLogging struct {
	// export block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#export OceanEcs#export}
	Export *OceanEcsLoggingExport `field:"optional" json:"export" yaml:"export"`
}

