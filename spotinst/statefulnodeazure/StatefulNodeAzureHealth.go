package statefulnodeazure


type StatefulNodeAzureHealth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.116.0/docs/resources/stateful_node_azure#auto_healing StatefulNodeAzure#auto_healing}.
	AutoHealing interface{} `field:"required" json:"autoHealing" yaml:"autoHealing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.116.0/docs/resources/stateful_node_azure#health_check_types StatefulNodeAzure#health_check_types}.
	HealthCheckTypes *[]*string `field:"required" json:"healthCheckTypes" yaml:"healthCheckTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.116.0/docs/resources/stateful_node_azure#grace_period StatefulNodeAzure#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.116.0/docs/resources/stateful_node_azure#unhealthy_duration StatefulNodeAzure#unhealthy_duration}.
	UnhealthyDuration *float64 `field:"optional" json:"unhealthyDuration" yaml:"unhealthyDuration"`
}

