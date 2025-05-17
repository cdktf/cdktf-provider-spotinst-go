// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpVngTemplateScheduling struct {
	// vng_template_shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.0/docs/resources/ocean_aks_np#vng_template_shutdown_hours OceanAksNp#vng_template_shutdown_hours}
	VngTemplateShutdownHours *OceanAksNpVngTemplateSchedulingVngTemplateShutdownHours `field:"optional" json:"vngTemplateShutdownHours" yaml:"vngTemplateShutdownHours"`
}

