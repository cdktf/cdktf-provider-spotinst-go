// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package organizationuser

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OrganizationUserPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OrganizationUserPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OrganizationUserPoliciesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OrganizationUserPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OrganizationUserPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OrganizationUserPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OrganizationUserPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOrganizationUserPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

