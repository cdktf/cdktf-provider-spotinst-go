// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/subscription#endpoint Subscription#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/subscription#event_type Subscription#event_type}.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/subscription#protocol Subscription#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/subscription#resource_id Subscription#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/subscription#format Subscription#format}.
	Format *map[string]*string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/subscription#id Subscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

