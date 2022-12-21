package oceanspark


type OceanSparkIngressLoadBalancer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#managed OceanSpark#managed}.
	Managed interface{} `field:"optional" json:"managed" yaml:"managed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#service_annotations OceanSpark#service_annotations}.
	ServiceAnnotations *map[string]*string `field:"optional" json:"serviceAnnotations" yaml:"serviceAnnotations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#target_group_arn OceanSpark#target_group_arn}.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

