package multaitargetset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MultaiTargetSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/multai_target_set#balancer_id MultaiTargetSet#balancer_id}.
	BalancerId *string `field:"required" json:"balancerId" yaml:"balancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/multai_target_set#deployment_id MultaiTargetSet#deployment_id}.
	DeploymentId *string `field:"required" json:"deploymentId" yaml:"deploymentId"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/multai_target_set#health_check MultaiTargetSet#health_check}
	HealthCheck *MultaiTargetSetHealthCheck `field:"required" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/multai_target_set#protocol MultaiTargetSet#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/multai_target_set#weight MultaiTargetSet#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/multai_target_set#id MultaiTargetSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/multai_target_set#name MultaiTargetSet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/multai_target_set#port MultaiTargetSet#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/multai_target_set#tags MultaiTargetSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

