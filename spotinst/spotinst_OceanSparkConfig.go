// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanSparkConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#ocean_cluster_id OceanSpark#ocean_cluster_id}.
	OceanClusterId *string `field:"required" json:"oceanClusterId" yaml:"oceanClusterId"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#compute OceanSpark#compute}
	Compute *OceanSparkCompute `field:"optional" json:"compute" yaml:"compute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#id OceanSpark#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ingress block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#ingress OceanSpark#ingress}
	Ingress *OceanSparkIngress `field:"optional" json:"ingress" yaml:"ingress"`
	// log_collection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#log_collection OceanSpark#log_collection}
	LogCollection *OceanSparkLogCollection `field:"optional" json:"logCollection" yaml:"logCollection"`
	// webhook block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_spark#webhook OceanSpark#webhook}
	Webhook *OceanSparkWebhook `field:"optional" json:"webhook" yaml:"webhook"`
}

