// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package statefulnodeazure

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulNodeAzureBootDiagnosticsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulNodeAzureBootDiagnosticsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureBootDiagnosticsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureBootDiagnosticsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureBootDiagnosticsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulNodeAzureBootDiagnosticsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulNodeAzureBootDiagnosticsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

