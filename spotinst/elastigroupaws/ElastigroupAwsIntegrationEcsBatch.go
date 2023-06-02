package elastigroupaws


type ElastigroupAwsIntegrationEcsBatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/elastigroup_aws#job_queue_names ElastigroupAws#job_queue_names}.
	JobQueueNames *[]*string `field:"required" json:"jobQueueNames" yaml:"jobQueueNames"`
}

