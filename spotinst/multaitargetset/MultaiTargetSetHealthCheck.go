package multaitargetset


type MultaiTargetSetHealthCheck struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/multai_target_set#healthy_threshold MultaiTargetSet#healthy_threshold}.
	HealthyThreshold *float64 `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/multai_target_set#interval MultaiTargetSet#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/multai_target_set#path MultaiTargetSet#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/multai_target_set#protocol MultaiTargetSet#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/multai_target_set#timeout MultaiTargetSet#timeout}.
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/multai_target_set#unhealthy_threshold MultaiTargetSet#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.132.0/docs/resources/multai_target_set#port MultaiTargetSet#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

