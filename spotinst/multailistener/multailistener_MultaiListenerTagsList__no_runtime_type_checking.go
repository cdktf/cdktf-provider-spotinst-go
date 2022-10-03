//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package multailistener

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MultaiListenerTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MultaiListenerTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MultaiListenerTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MultaiListenerTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MultaiListenerTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MultaiListenerTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMultaiListenerTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

