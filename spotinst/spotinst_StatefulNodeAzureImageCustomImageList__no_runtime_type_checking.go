//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulNodeAzureImageCustomImageList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulNodeAzureImageCustomImageList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureImageCustomImageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureImageCustomImageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureImageCustomImageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureImageCustomImageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulNodeAzureImageCustomImageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
