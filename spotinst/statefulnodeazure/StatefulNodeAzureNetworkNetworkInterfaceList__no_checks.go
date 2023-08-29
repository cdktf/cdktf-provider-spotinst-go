// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package statefulnodeazure

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulNodeAzureNetworkNetworkInterfaceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

