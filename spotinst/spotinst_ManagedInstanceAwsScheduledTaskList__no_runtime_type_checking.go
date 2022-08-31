//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedInstanceAwsScheduledTaskList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedInstanceAwsScheduledTaskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedInstanceAwsScheduledTaskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedInstanceAwsScheduledTaskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedInstanceAwsScheduledTaskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedInstanceAwsScheduledTaskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedInstanceAwsScheduledTaskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

