// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProviderJobSpecJobTemplateTemplateSpecContainers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.197.1/docs/resources/oceancd_verification_template#command OceancdVerificationTemplate#command}.
	Command *[]*string `field:"required" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.197.1/docs/resources/oceancd_verification_template#container_name OceancdVerificationTemplate#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.197.1/docs/resources/oceancd_verification_template#image OceancdVerificationTemplate#image}.
	Image *string `field:"required" json:"image" yaml:"image"`
}

