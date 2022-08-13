// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAksHealth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#grace_period OceanAks#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
}

