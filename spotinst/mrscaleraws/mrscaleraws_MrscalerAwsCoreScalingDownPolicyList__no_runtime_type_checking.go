//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package mrscaleraws

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMrscalerAwsCoreScalingDownPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
