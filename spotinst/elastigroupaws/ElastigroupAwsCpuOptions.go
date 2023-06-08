package elastigroupaws


type ElastigroupAwsCpuOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.1/docs/resources/elastigroup_aws#threads_per_core ElastigroupAws#threads_per_core}.
	ThreadsPerCore *float64 `field:"required" json:"threadsPerCore" yaml:"threadsPerCore"`
}

