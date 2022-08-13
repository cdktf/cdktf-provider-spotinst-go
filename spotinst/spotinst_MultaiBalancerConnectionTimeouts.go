// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type MultaiBalancerConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_balancer#draining MultaiBalancer#draining}.
	Draining *float64 `field:"optional" json:"draining" yaml:"draining"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_balancer#idle MultaiBalancer#idle}.
	Idle *float64 `field:"optional" json:"idle" yaml:"idle"`
}

