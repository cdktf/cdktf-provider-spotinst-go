package elastigroupazure


type ElastigroupAzureHealthCheck struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/elastigroup_azure#health_check_type ElastigroupAzure#health_check_type}.
	HealthCheckType *string `field:"required" json:"healthCheckType" yaml:"healthCheckType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/elastigroup_azure#auto_healing ElastigroupAzure#auto_healing}.
	AutoHealing interface{} `field:"optional" json:"autoHealing" yaml:"autoHealing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/elastigroup_azure#grace_period ElastigroupAzure#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
}

