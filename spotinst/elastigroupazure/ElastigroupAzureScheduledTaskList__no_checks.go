// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package elastigroupazure

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElastigroupAzureScheduledTaskList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElastigroupAzureScheduledTaskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureScheduledTaskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureScheduledTaskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureScheduledTaskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureScheduledTaskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElastigroupAzureScheduledTaskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

