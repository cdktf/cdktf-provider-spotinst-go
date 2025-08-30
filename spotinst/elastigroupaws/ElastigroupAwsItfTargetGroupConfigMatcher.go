// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsItfTargetGroupConfigMatcher struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/elastigroup_aws#grpc_code ElastigroupAws#grpc_code}.
	GrpcCode *string `field:"optional" json:"grpcCode" yaml:"grpcCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/elastigroup_aws#http_code ElastigroupAws#http_code}.
	HttpCode *string `field:"optional" json:"httpCode" yaml:"httpCode"`
}

