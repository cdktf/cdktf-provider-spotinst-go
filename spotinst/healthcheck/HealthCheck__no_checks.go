//go:build no_runtime_type_checking

package healthcheck

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HealthCheck) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validatePutCheckParameters(value *HealthCheckCheck) error {
	return nil
}

func validateHealthCheck_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHealthCheck_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateHealthCheck_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_HealthCheck) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HealthCheck) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HealthCheck) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HealthCheck) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_HealthCheck) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HealthCheck) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_HealthCheck) validateSetProxyAddressParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HealthCheck) validateSetProxyPortParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_HealthCheck) validateSetResourceIdParameters(val *string) error {
	return nil
}

func validateNewHealthCheckParameters(scope constructs.Construct, id *string, config *HealthCheckConfig) error {
	return nil
}

