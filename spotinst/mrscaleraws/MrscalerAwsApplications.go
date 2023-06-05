package mrscaleraws


type MrscalerAwsApplications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/mrscaler_aws#name MrscalerAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/mrscaler_aws#args MrscalerAws#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/mrscaler_aws#version MrscalerAws#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

