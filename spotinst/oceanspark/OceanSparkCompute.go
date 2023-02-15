package oceanspark


type OceanSparkCompute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#create_vngs OceanSpark#create_vngs}.
	CreateVngs interface{} `field:"optional" json:"createVngs" yaml:"createVngs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#use_taints OceanSpark#use_taints}.
	UseTaints interface{} `field:"optional" json:"useTaints" yaml:"useTaints"`
}

