// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package multaibalancer

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MultaiBalancerTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MultaiBalancerTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MultaiBalancerTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MultaiBalancerTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MultaiBalancerTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MultaiBalancerTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMultaiBalancerTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

