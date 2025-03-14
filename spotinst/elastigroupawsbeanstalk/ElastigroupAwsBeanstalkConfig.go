// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupawsbeanstalk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsBeanstalkConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#desired_capacity ElastigroupAwsBeanstalk#desired_capacity}.
	DesiredCapacity *float64 `field:"required" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#instance_types_spot ElastigroupAwsBeanstalk#instance_types_spot}.
	InstanceTypesSpot *[]*string `field:"required" json:"instanceTypesSpot" yaml:"instanceTypesSpot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#max_size ElastigroupAwsBeanstalk#max_size}.
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#min_size ElastigroupAwsBeanstalk#min_size}.
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#name ElastigroupAwsBeanstalk#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#product ElastigroupAwsBeanstalk#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#region ElastigroupAwsBeanstalk#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#beanstalk_environment_id ElastigroupAwsBeanstalk#beanstalk_environment_id}.
	BeanstalkEnvironmentId *string `field:"optional" json:"beanstalkEnvironmentId" yaml:"beanstalkEnvironmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#beanstalk_environment_name ElastigroupAwsBeanstalk#beanstalk_environment_name}.
	BeanstalkEnvironmentName *string `field:"optional" json:"beanstalkEnvironmentName" yaml:"beanstalkEnvironmentName"`
	// deployment_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#deployment_preferences ElastigroupAwsBeanstalk#deployment_preferences}
	DeploymentPreferences *ElastigroupAwsBeanstalkDeploymentPreferences `field:"optional" json:"deploymentPreferences" yaml:"deploymentPreferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#id ElastigroupAwsBeanstalk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#maintenance ElastigroupAwsBeanstalk#maintenance}.
	Maintenance *string `field:"optional" json:"maintenance" yaml:"maintenance"`
	// managed_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#managed_actions ElastigroupAwsBeanstalk#managed_actions}
	ManagedActions *ElastigroupAwsBeanstalkManagedActions `field:"optional" json:"managedActions" yaml:"managedActions"`
	// scheduled_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/elastigroup_aws_beanstalk#scheduled_task ElastigroupAwsBeanstalk#scheduled_task}
	ScheduledTask interface{} `field:"optional" json:"scheduledTask" yaml:"scheduledTask"`
}

