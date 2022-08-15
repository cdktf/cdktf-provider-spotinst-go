// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ManagedInstanceAwsIntegrationRoute53Domains struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#hosted_zone_id ManagedInstanceAws#hosted_zone_id}.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// record_sets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#record_sets ManagedInstanceAws#record_sets}
	RecordSets interface{} `field:"required" json:"recordSets" yaml:"recordSets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#record_set_type ManagedInstanceAws#record_set_type}.
	RecordSetType *string `field:"optional" json:"recordSetType" yaml:"recordSetType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#spotinst_acct_id ManagedInstanceAws#spotinst_acct_id}.
	SpotinstAcctId *string `field:"optional" json:"spotinstAcctId" yaml:"spotinstAcctId"`
}
