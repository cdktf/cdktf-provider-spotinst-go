// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanSparkLogCollection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#collect_driver_logs OceanSpark#collect_driver_logs}.
	CollectDriverLogs interface{} `field:"optional" json:"collectDriverLogs" yaml:"collectDriverLogs"`
}

