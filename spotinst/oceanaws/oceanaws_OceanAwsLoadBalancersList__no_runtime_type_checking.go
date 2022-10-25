//go:build no_runtime_type_checking

package oceanaws

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OceanAwsLoadBalancersList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OceanAwsLoadBalancersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OceanAwsLoadBalancersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OceanAwsLoadBalancersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OceanAwsLoadBalancersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OceanAwsLoadBalancersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOceanAwsLoadBalancersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

