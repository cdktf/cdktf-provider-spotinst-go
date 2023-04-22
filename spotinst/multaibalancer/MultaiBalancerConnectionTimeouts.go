package multaibalancer


type MultaiBalancerConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/multai_balancer#draining MultaiBalancer#draining}.
	Draining *float64 `field:"optional" json:"draining" yaml:"draining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/multai_balancer#idle MultaiBalancer#idle}.
	Idle *float64 `field:"optional" json:"idle" yaml:"idle"`
}

