// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs


type OceanEcsLoggingExport struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.184.0/docs/resources/ocean_ecs#s3 OceanEcs#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

