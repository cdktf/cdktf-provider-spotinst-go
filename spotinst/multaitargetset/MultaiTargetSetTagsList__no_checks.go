//go:build no_runtime_type_checking

package multaitargetset

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MultaiTargetSetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MultaiTargetSetTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MultaiTargetSetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MultaiTargetSetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MultaiTargetSetTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MultaiTargetSetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMultaiTargetSetTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
