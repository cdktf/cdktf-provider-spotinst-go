package elastigroupgke


type ElastigroupGkeIntegrationDockerSwarm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#master_host ElastigroupGke#master_host}.
	MasterHost *string `field:"required" json:"masterHost" yaml:"masterHost"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#master_port ElastigroupGke#master_port}.
	MasterPort *float64 `field:"required" json:"masterPort" yaml:"masterPort"`
}

