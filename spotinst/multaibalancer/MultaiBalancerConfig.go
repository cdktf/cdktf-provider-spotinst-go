package multaibalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MultaiBalancerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/multai_balancer#name MultaiBalancer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// connection_timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/multai_balancer#connection_timeouts MultaiBalancer#connection_timeouts}
	ConnectionTimeouts *MultaiBalancerConnectionTimeouts `field:"optional" json:"connectionTimeouts" yaml:"connectionTimeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/multai_balancer#dns_cname_aliases MultaiBalancer#dns_cname_aliases}.
	DnsCnameAliases *[]*string `field:"optional" json:"dnsCnameAliases" yaml:"dnsCnameAliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/multai_balancer#id MultaiBalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/multai_balancer#scheme MultaiBalancer#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/multai_balancer#tags MultaiBalancer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

