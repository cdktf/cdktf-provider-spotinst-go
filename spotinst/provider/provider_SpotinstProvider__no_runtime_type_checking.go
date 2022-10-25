//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpotinstProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SpotinstProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSpotinstProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSpotinstProviderParameters(scope constructs.Construct, id *string, config *SpotinstProviderConfig) error {
	return nil
}

