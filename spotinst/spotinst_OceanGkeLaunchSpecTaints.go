// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeLaunchSpecTaints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#effect OceanGkeLaunchSpec#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#key OceanGkeLaunchSpec#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#value OceanGkeLaunchSpec#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}
