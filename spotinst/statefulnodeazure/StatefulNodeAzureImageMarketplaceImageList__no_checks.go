// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package statefulnodeazure

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulNodeAzureImageMarketplaceImageList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulNodeAzureImageMarketplaceImageList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureImageMarketplaceImageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureImageMarketplaceImageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureImageMarketplaceImageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureImageMarketplaceImageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulNodeAzureImageMarketplaceImageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

