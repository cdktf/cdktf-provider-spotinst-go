// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksVirtualNodeGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#name OceanAksVirtualNodeGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#ocean_id OceanAksVirtualNodeGroup#ocean_id}.
	OceanId *string `field:"required" json:"oceanId" yaml:"oceanId"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#autoscale OceanAksVirtualNodeGroup#autoscale}
	Autoscale interface{} `field:"optional" json:"autoscale" yaml:"autoscale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#id OceanAksVirtualNodeGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// label block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#label OceanAksVirtualNodeGroup#label}
	Label interface{} `field:"optional" json:"label" yaml:"label"`
	// launch_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#launch_specification OceanAksVirtualNodeGroup#launch_specification}
	LaunchSpecification interface{} `field:"optional" json:"launchSpecification" yaml:"launchSpecification"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#resource_limits OceanAksVirtualNodeGroup#resource_limits}
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
	// taint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#taint OceanAksVirtualNodeGroup#taint}
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks_virtual_node_group#zones OceanAksVirtualNodeGroup#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

