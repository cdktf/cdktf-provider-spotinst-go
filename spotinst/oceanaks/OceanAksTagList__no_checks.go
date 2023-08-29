// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package oceanaks

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OceanAksTagList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OceanAksTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OceanAksTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OceanAksTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OceanAksTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OceanAksTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOceanAksTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

