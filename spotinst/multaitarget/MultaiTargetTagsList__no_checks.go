// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package multaitarget

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MultaiTargetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MultaiTargetTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MultaiTargetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MultaiTargetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MultaiTargetTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MultaiTargetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMultaiTargetTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

