package oceanaws


type OceanAwsBlockDeviceMappingsEbs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_aws#delete_on_termination OceanAws#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// dynamic_volume_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_aws#dynamic_volume_size OceanAws#dynamic_volume_size}
	DynamicVolumeSize *OceanAwsBlockDeviceMappingsEbsDynamicVolumeSize `field:"optional" json:"dynamicVolumeSize" yaml:"dynamicVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_aws#encrypted OceanAws#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_aws#iops OceanAws#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_aws#kms_key_id OceanAws#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_aws#snapshot_id OceanAws#snapshot_id}.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_aws#throughput OceanAws#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_aws#volume_size OceanAws#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.120.0/docs/resources/ocean_aws#volume_type OceanAws#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

