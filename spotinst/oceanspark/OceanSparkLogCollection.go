package oceanspark


type OceanSparkLogCollection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/ocean_spark#collect_app_logs OceanSpark#collect_app_logs}.
	CollectAppLogs interface{} `field:"optional" json:"collectAppLogs" yaml:"collectAppLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/ocean_spark#collect_driver_logs OceanSpark#collect_driver_logs}.
	CollectDriverLogs interface{} `field:"optional" json:"collectDriverLogs" yaml:"collectDriverLogs"`
}

