// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProvider struct {
	// cloud_watch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_verification_template#cloud_watch OceancdVerificationTemplate#cloud_watch}
	CloudWatch *OceancdVerificationTemplateMetricsProviderCloudWatch `field:"optional" json:"cloudWatch" yaml:"cloudWatch"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_verification_template#datadog OceancdVerificationTemplate#datadog}
	Datadog *OceancdVerificationTemplateMetricsProviderDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// jenkins block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_verification_template#jenkins OceancdVerificationTemplate#jenkins}
	Jenkins *OceancdVerificationTemplateMetricsProviderJenkins `field:"optional" json:"jenkins" yaml:"jenkins"`
	// job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_verification_template#job OceancdVerificationTemplate#job}
	Job *OceancdVerificationTemplateMetricsProviderJob `field:"optional" json:"job" yaml:"job"`
	// new_relic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_verification_template#new_relic OceancdVerificationTemplate#new_relic}
	NewRelic *OceancdVerificationTemplateMetricsProviderNewRelic `field:"optional" json:"newRelic" yaml:"newRelic"`
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_verification_template#prometheus OceancdVerificationTemplate#prometheus}
	Prometheus *OceancdVerificationTemplateMetricsProviderPrometheus `field:"optional" json:"prometheus" yaml:"prometheus"`
	// web block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/oceancd_verification_template#web OceancdVerificationTemplate#web}
	Web *OceancdVerificationTemplateMetricsProviderWeb `field:"optional" json:"web" yaml:"web"`
}

