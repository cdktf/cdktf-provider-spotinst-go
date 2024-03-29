// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package statefulnodeazure

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulNodeAzureDeleteList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StatefulNodeAzureDeleteList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulNodeAzureDeleteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureDeleteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureDeleteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureDeleteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureDeleteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulNodeAzureDeleteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

