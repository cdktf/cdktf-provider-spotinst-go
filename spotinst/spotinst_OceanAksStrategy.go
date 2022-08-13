// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAksStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#fallback_to_ondemand OceanAks#fallback_to_ondemand}.
	FallbackToOndemand interface{} `field:"optional" json:"fallbackToOndemand" yaml:"fallbackToOndemand"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#spot_percentage OceanAks#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
}

