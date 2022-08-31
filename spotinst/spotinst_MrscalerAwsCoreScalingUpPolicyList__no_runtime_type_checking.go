//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MrscalerAwsCoreScalingUpPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MrscalerAwsCoreScalingUpPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MrscalerAwsCoreScalingUpPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MrscalerAwsCoreScalingUpPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MrscalerAwsCoreScalingUpPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MrscalerAwsCoreScalingUpPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMrscalerAwsCoreScalingUpPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

