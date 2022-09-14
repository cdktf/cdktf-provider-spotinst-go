// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanSparkIngress struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#service_annotations OceanSpark#service_annotations}.
	ServiceAnnotations *map[string]*string `field:"optional" json:"serviceAnnotations" yaml:"serviceAnnotations"`
}

