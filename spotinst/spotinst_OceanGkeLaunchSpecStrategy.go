// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeLaunchSpecStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#preemptible_percentage OceanGkeLaunchSpec#preemptible_percentage}.
	PreemptiblePercentage *float64 `field:"optional" json:"preemptiblePercentage" yaml:"preemptiblePercentage"`
}

