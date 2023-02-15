package elastigroupaws


type ElastigroupAwsIntegrationEcsBatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#job_queue_names ElastigroupAws#job_queue_names}.
	JobQueueNames *[]*string `field:"required" json:"jobQueueNames" yaml:"jobQueueNames"`
}

