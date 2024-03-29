// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpotinstProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SpotinstProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSpotinstProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSpotinstProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSpotinstProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSpotinstProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SpotinstProvider) validateSetEnabledParameters(val interface{}) error {
	return nil
}

func validateNewSpotinstProviderParameters(scope constructs.Construct, id *string, config *SpotinstProviderConfig) error {
	return nil
}

