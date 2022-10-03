//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package oceanaws

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OceanAwsScheduledTaskList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OceanAwsScheduledTaskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OceanAwsScheduledTaskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OceanAwsScheduledTaskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OceanAwsScheduledTaskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OceanAwsScheduledTaskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOceanAwsScheduledTaskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

