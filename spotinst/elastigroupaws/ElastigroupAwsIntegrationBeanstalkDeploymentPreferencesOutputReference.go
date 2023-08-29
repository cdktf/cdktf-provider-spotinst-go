// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference interface {
	cdktf.ComplexObject
	AutomaticRoll() interface{}
	SetAutomaticRoll(val interface{})
	AutomaticRollInput() interface{}
	BatchSizePercentage() *float64
	SetBatchSizePercentage(val *float64)
	BatchSizePercentageInput() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GracePeriod() *float64
	SetGracePeriod(val *float64)
	GracePeriodInput() *float64
	InternalValue() *ElastigroupAwsIntegrationBeanstalkDeploymentPreferences
	SetInternalValue(val *ElastigroupAwsIntegrationBeanstalkDeploymentPreferences)
	Strategy() ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategyOutputReference
	StrategyInput() *ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutStrategy(value *ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategy)
	ResetAutomaticRoll()
	ResetBatchSizePercentage()
	ResetGracePeriod()
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference
type jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) AutomaticRoll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticRoll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) AutomaticRollInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticRollInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) BatchSizePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) BatchSizePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) InternalValue() *ElastigroupAwsIntegrationBeanstalkDeploymentPreferences {
	var returns *ElastigroupAwsIntegrationBeanstalkDeploymentPreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) Strategy() ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategyOutputReference {
	var returns ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategyOutputReference
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) StrategyInput() *ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategy {
	var returns *ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategy
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference_Override(e ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference)SetAutomaticRoll(val interface{}) {
	if err := j.validateSetAutomaticRollParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticRoll",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference)SetBatchSizePercentage(val *float64) {
	if err := j.validateSetBatchSizePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSizePercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference)SetGracePeriod(val *float64) {
	if err := j.validateSetGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracePeriod",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference)SetInternalValue(val *ElastigroupAwsIntegrationBeanstalkDeploymentPreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) PutStrategy(value *ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategy) {
	if err := e.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStrategy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) ResetAutomaticRoll() {
	_jsii_.InvokeVoid(
		e,
		"resetAutomaticRoll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) ResetBatchSizePercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetBatchSizePercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) ResetGracePeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetGracePeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

