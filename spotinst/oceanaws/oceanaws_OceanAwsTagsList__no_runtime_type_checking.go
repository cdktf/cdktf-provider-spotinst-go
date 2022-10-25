//go:build no_runtime_type_checking

package oceanaws

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OceanAwsTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OceanAwsTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OceanAwsTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OceanAwsTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OceanAwsTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OceanAwsTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOceanAwsTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

