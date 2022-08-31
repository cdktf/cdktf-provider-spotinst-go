//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElastigroupAwsItfTargetGroupConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElastigroupAwsItfTargetGroupConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsItfTargetGroupConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsItfTargetGroupConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsItfTargetGroupConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsItfTargetGroupConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElastigroupAwsItfTargetGroupConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

