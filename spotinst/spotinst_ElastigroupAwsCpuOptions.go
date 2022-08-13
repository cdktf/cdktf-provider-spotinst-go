// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsCpuOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#threads_per_core ElastigroupAws#threads_per_core}.
	ThreadsPerCore *float64 `field:"required" json:"threadsPerCore" yaml:"threadsPerCore"`
}

