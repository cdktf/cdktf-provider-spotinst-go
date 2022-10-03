//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elastigroupazure

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElastigroupAzureLoadBalancersList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElastigroupAzureLoadBalancersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureLoadBalancersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureLoadBalancersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureLoadBalancersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAzureLoadBalancersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElastigroupAzureLoadBalancersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

