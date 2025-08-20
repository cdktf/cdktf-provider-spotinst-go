// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsLoggingExport struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_aws#s3 ElastigroupAws#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

