// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package oceanaks

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OceanAksExtensionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OceanAksExtensionList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OceanAksExtensionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OceanAksExtensionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OceanAksExtensionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OceanAksExtensionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OceanAksExtensionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOceanAksExtensionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

