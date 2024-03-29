// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataintegration

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataIntegration) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_DataIntegration) validatePutS3Parameters(value *DataIntegrationS3) error {
	return nil
}

func validateDataIntegration_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDataIntegration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDataIntegration_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDataIntegration_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DataIntegration) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataIntegration) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataIntegration) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataIntegration) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DataIntegration) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataIntegration) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_DataIntegration) validateSetStatusParameters(val *string) error {
	return nil
}

func validateNewDataIntegrationParameters(scope constructs.Construct, id *string, config *DataIntegrationConfig) error {
	return nil
}

