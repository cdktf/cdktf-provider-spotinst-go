// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProviderJobSpecJobTemplateTemplateSpec struct {
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.200.0/docs/resources/oceancd_verification_template#containers OceancdVerificationTemplate#containers}
	Containers interface{} `field:"required" json:"containers" yaml:"containers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.200.0/docs/resources/oceancd_verification_template#restart_policy OceancdVerificationTemplate#restart_policy}.
	RestartPolicy *string `field:"required" json:"restartPolicy" yaml:"restartPolicy"`
}

