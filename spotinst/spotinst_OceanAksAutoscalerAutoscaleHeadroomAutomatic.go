// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAksAutoscalerAutoscaleHeadroomAutomatic struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#is_enabled OceanAks#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#percentage OceanAks#percentage}.
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

