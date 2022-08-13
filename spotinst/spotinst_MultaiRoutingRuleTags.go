// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type MultaiRoutingRuleTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_routing_rule#key MultaiRoutingRule#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_routing_rule#value MultaiRoutingRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

