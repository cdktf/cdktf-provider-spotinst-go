// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecSpotDeployments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.0/docs/resources/oceancd_rollout_spec#spot_deployments_cluster_id OceancdRolloutSpec#spot_deployments_cluster_id}.
	SpotDeploymentsClusterId *string `field:"optional" json:"spotDeploymentsClusterId" yaml:"spotDeploymentsClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.0/docs/resources/oceancd_rollout_spec#spot_deployments_name OceancdRolloutSpec#spot_deployments_name}.
	SpotDeploymentsName *string `field:"optional" json:"spotDeploymentsName" yaml:"spotDeploymentsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.0/docs/resources/oceancd_rollout_spec#spot_deployments_namespace OceancdRolloutSpec#spot_deployments_namespace}.
	SpotDeploymentsNamespace *string `field:"optional" json:"spotDeploymentsNamespace" yaml:"spotDeploymentsNamespace"`
}

