//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElastigroupAwsTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElastigroupAwsTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElastigroupAwsTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

