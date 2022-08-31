//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElastigroupGkeDiskInitializeParamsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElastigroupGkeDiskInitializeParamsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElastigroupGkeDiskInitializeParamsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ElastigroupGkeDiskInitializeParamsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElastigroupGkeDiskInitializeParamsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElastigroupGkeDiskInitializeParamsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElastigroupGkeDiskInitializeParamsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

