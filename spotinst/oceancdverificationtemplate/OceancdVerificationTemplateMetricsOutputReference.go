// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceancdverificationtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdVerificationTemplateMetricsOutputReference interface {
	cdktf.ComplexObject
	Baseline() OceancdVerificationTemplateMetricsBaselineOutputReference
	BaselineInput() *OceancdVerificationTemplateMetricsBaseline
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
	ConsecutiveErrorLimit() *float64
	SetConsecutiveErrorLimit(val *float64)
	ConsecutiveErrorLimitInput() *float64
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DryRun() interface{}
	SetDryRun(val interface{})
	DryRunInput() interface{}
	FailureCondition() *string
	SetFailureCondition(val *string)
	FailureConditionInput() *string
	FailureLimit() *float64
	SetFailureLimit(val *float64)
	FailureLimitInput() *float64
	// Experimental.
	Fqn() *string
	InitialDelay() *string
	SetInitialDelay(val *string)
	InitialDelayInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	MetricsName() *string
	SetMetricsName(val *string)
	MetricsNameInput() *string
	Provider() OceancdVerificationTemplateMetricsProviderList
	ProviderInput() interface{}
	SuccessCondition() *string
	SetSuccessCondition(val *string)
	SuccessConditionInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutBaseline(value *OceancdVerificationTemplateMetricsBaseline)
	PutProvider(value interface{})
	ResetBaseline()
	ResetConsecutiveErrorLimit()
	ResetCount()
	ResetDryRun()
	ResetFailureCondition()
	ResetFailureLimit()
	ResetInitialDelay()
	ResetInterval()
	ResetSuccessCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceancdVerificationTemplateMetricsOutputReference
type jsiiProxy_OceancdVerificationTemplateMetricsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) Baseline() OceancdVerificationTemplateMetricsBaselineOutputReference {
	var returns OceancdVerificationTemplateMetricsBaselineOutputReference
	_jsii_.Get(
		j,
		"baseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) BaselineInput() *OceancdVerificationTemplateMetricsBaseline {
	var returns *OceancdVerificationTemplateMetricsBaseline
	_jsii_.Get(
		j,
		"baselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ConsecutiveErrorLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveErrorLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ConsecutiveErrorLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveErrorLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) DryRun() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dryRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) DryRunInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dryRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) FailureCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) FailureConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) FailureLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) FailureLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) InitialDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) InitialDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) MetricsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) MetricsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) Provider() OceancdVerificationTemplateMetricsProviderList {
	var returns OceancdVerificationTemplateMetricsProviderList
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) SuccessCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) SuccessConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceancdVerificationTemplateMetricsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OceancdVerificationTemplateMetricsOutputReference {
	_init_.Initialize()

	if err := validateNewOceancdVerificationTemplateMetricsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdVerificationTemplateMetricsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdVerificationTemplate.OceancdVerificationTemplateMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOceancdVerificationTemplateMetricsOutputReference_Override(o OceancdVerificationTemplateMetricsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdVerificationTemplate.OceancdVerificationTemplateMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetConsecutiveErrorLimit(val *float64) {
	if err := j.validateSetConsecutiveErrorLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consecutiveErrorLimit",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetDryRun(val interface{}) {
	if err := j.validateSetDryRunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dryRun",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetFailureCondition(val *string) {
	if err := j.validateSetFailureConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureCondition",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetFailureLimit(val *float64) {
	if err := j.validateSetFailureLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureLimit",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetInitialDelay(val *string) {
	if err := j.validateSetInitialDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDelay",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetInterval(val *string) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetMetricsName(val *string) {
	if err := j.validateSetMetricsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsName",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetSuccessCondition(val *string) {
	if err := j.validateSetSuccessConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successCondition",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) PutBaseline(value *OceancdVerificationTemplateMetricsBaseline) {
	if err := o.validatePutBaselineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBaseline",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) PutProvider(value interface{}) {
	if err := o.validatePutProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putProvider",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ResetBaseline() {
	_jsii_.InvokeVoid(
		o,
		"resetBaseline",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ResetConsecutiveErrorLimit() {
	_jsii_.InvokeVoid(
		o,
		"resetConsecutiveErrorLimit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		o,
		"resetCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ResetDryRun() {
	_jsii_.InvokeVoid(
		o,
		"resetDryRun",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ResetFailureCondition() {
	_jsii_.InvokeVoid(
		o,
		"resetFailureCondition",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ResetFailureLimit() {
	_jsii_.InvokeVoid(
		o,
		"resetFailureLimit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ResetInitialDelay() {
	_jsii_.InvokeVoid(
		o,
		"resetInitialDelay",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		o,
		"resetInterval",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ResetSuccessCondition() {
	_jsii_.InvokeVoid(
		o,
		"resetSuccessCondition",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

