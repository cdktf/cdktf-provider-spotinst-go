//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElastigroupAwsIntegrationRoute53DomainsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElastigroupAwsIntegrationRoute53DomainsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsIntegrationRoute53DomainsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsIntegrationRoute53DomainsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsIntegrationRoute53DomainsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElastigroupAwsIntegrationRoute53DomainsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElastigroupAwsIntegrationRoute53DomainsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

