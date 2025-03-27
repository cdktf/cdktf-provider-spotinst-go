// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsProviderJenkins struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#jenkins_interval OceancdVerificationTemplate#jenkins_interval}.
	JenkinsInterval *string `field:"required" json:"jenkinsInterval" yaml:"jenkinsInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#pipeline_name OceancdVerificationTemplate#pipeline_name}.
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#timeout OceancdVerificationTemplate#timeout}.
	Timeout *string `field:"required" json:"timeout" yaml:"timeout"`
	// jenkins_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#jenkins_parameters OceancdVerificationTemplate#jenkins_parameters}
	JenkinsParameters *OceancdVerificationTemplateMetricsProviderJenkinsJenkinsParameters `field:"optional" json:"jenkinsParameters" yaml:"jenkinsParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_verification_template#tls_verification OceancdVerificationTemplate#tls_verification}.
	TlsVerification interface{} `field:"optional" json:"tlsVerification" yaml:"tlsVerification"`
}

