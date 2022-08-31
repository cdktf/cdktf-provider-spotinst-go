//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OceanAksManagedServiceIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OceanAksManagedServiceIdentityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OceanAksManagedServiceIdentityList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OceanAksManagedServiceIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OceanAksManagedServiceIdentityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OceanAksManagedServiceIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOceanAksManagedServiceIdentityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

