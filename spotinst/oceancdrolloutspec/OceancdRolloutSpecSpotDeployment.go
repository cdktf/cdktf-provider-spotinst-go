// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecSpotDeployment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.1/docs/resources/oceancd_rollout_spec#spot_deployment_cluster_id OceancdRolloutSpec#spot_deployment_cluster_id}.
	SpotDeploymentClusterId *string `field:"optional" json:"spotDeploymentClusterId" yaml:"spotDeploymentClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.1/docs/resources/oceancd_rollout_spec#spot_deployment_name OceancdRolloutSpec#spot_deployment_name}.
	SpotDeploymentName *string `field:"optional" json:"spotDeploymentName" yaml:"spotDeploymentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.1/docs/resources/oceancd_rollout_spec#spot_deployment_namespace OceancdRolloutSpec#spot_deployment_namespace}.
	SpotDeploymentNamespace *string `field:"optional" json:"spotDeploymentNamespace" yaml:"spotDeploymentNamespace"`
}

