//go:build no_runtime_type_checking

package oceanecs

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OceanEcsScheduledTaskList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OceanEcsScheduledTaskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OceanEcsScheduledTaskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OceanEcsScheduledTaskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OceanEcsScheduledTaskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OceanEcsScheduledTaskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOceanEcsScheduledTaskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

