package elastigroupgke


type ElastigroupGkeGpu struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/elastigroup_gke#count ElastigroupGke#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/elastigroup_gke#type ElastigroupGke#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

