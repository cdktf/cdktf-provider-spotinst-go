// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdRolloutSpecConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.204.0/docs/resources/oceancd_rollout_spec#rollout_spec_name OceancdRolloutSpec#rollout_spec_name}.
	RolloutSpecName *string `field:"required" json:"rolloutSpecName" yaml:"rolloutSpecName"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.204.0/docs/resources/oceancd_rollout_spec#strategy OceancdRolloutSpec#strategy}
	Strategy *OceancdRolloutSpecStrategy `field:"required" json:"strategy" yaml:"strategy"`
	// failure_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.204.0/docs/resources/oceancd_rollout_spec#failure_policy OceancdRolloutSpec#failure_policy}
	FailurePolicy *OceancdRolloutSpecFailurePolicy `field:"optional" json:"failurePolicy" yaml:"failurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.204.0/docs/resources/oceancd_rollout_spec#id OceancdRolloutSpec#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// spot_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.204.0/docs/resources/oceancd_rollout_spec#spot_deployment OceancdRolloutSpec#spot_deployment}
	SpotDeployment *OceancdRolloutSpecSpotDeployment `field:"optional" json:"spotDeployment" yaml:"spotDeployment"`
	// spot_deployments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.204.0/docs/resources/oceancd_rollout_spec#spot_deployments OceancdRolloutSpec#spot_deployments}
	SpotDeployments interface{} `field:"optional" json:"spotDeployments" yaml:"spotDeployments"`
	// traffic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.204.0/docs/resources/oceancd_rollout_spec#traffic OceancdRolloutSpec#traffic}
	Traffic *OceancdRolloutSpecTraffic `field:"optional" json:"traffic" yaml:"traffic"`
}

