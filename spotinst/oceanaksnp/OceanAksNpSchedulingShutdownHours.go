package oceanaksnp


type OceanAksNpSchedulingShutdownHours struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/ocean_aks_np#time_windows OceanAksNp#time_windows}.
	TimeWindows *[]*string `field:"required" json:"timeWindows" yaml:"timeWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/ocean_aks_np#is_enabled OceanAksNp#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

