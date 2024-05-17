// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanSparkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/ocean_spark#ocean_cluster_id OceanSpark#ocean_cluster_id}.
	OceanClusterId *string `field:"required" json:"oceanClusterId" yaml:"oceanClusterId"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/ocean_spark#compute OceanSpark#compute}
	Compute *OceanSparkCompute `field:"optional" json:"compute" yaml:"compute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/ocean_spark#id OceanSpark#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ingress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/ocean_spark#ingress OceanSpark#ingress}
	Ingress *OceanSparkIngress `field:"optional" json:"ingress" yaml:"ingress"`
	// log_collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/ocean_spark#log_collection OceanSpark#log_collection}
	LogCollection *OceanSparkLogCollection `field:"optional" json:"logCollection" yaml:"logCollection"`
	// spark block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/ocean_spark#spark OceanSpark#spark}
	Spark *OceanSparkSpark `field:"optional" json:"spark" yaml:"spark"`
	// webhook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/ocean_spark#webhook OceanSpark#webhook}
	Webhook *OceanSparkWebhook `field:"optional" json:"webhook" yaml:"webhook"`
}

