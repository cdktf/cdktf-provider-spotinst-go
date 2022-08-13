// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type MrscalerAwsApplications struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#name MrscalerAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#args MrscalerAws#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#version MrscalerAws#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

