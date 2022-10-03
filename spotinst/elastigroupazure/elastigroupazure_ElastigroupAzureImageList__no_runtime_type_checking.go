//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elastigroupazure

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElastigroupAzureImageList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElastigroupAzureImageList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureImageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureImageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureImageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureImageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElastigroupAzureImageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

