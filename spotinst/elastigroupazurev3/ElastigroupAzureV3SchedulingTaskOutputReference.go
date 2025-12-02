// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/elastigroupazurev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAzureV3SchedulingTaskOutputReference interface {
	cdktf.ComplexObject
	Adjustment() *string
	SetAdjustment(val *string)
	AdjustmentInput() *string
	AdjustmentPercentage() *string
	SetAdjustmentPercentage(val *string)
	AdjustmentPercentageInput() *string
	BatchSizePercentage() *string
	SetBatchSizePercentage(val *string)
	BatchSizePercentageInput() *string
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
	CronExpression() *string
	SetCronExpression(val *string)
	CronExpressionInput() *string
	// Experimental.
	Fqn() *string
	GracePeriod() *string
	SetGracePeriod(val *string)
	GracePeriodInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	ScaleMaxCapacity() *string
	SetScaleMaxCapacity(val *string)
	ScaleMaxCapacityInput() *string
	ScaleMinCapacity() *string
	SetScaleMinCapacity(val *string)
	ScaleMinCapacityInput() *string
	ScaleTargetCapacity() *string
	SetScaleTargetCapacity(val *string)
	ScaleTargetCapacityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAdjustment()
	ResetAdjustmentPercentage()
	ResetBatchSizePercentage()
	ResetGracePeriod()
	ResetScaleMaxCapacity()
	ResetScaleMinCapacity()
	ResetScaleTargetCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAzureV3SchedulingTaskOutputReference
type jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) Adjustment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) AdjustmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) AdjustmentPercentage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) AdjustmentPercentageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) BatchSizePercentage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) BatchSizePercentageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) CronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) CronExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GracePeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GracePeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ScaleMaxCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMaxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ScaleMaxCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMaxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ScaleMinCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMinCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ScaleMinCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMinCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ScaleTargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ScaleTargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewElastigroupAzureV3SchedulingTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElastigroupAzureV3SchedulingTaskOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAzureV3SchedulingTaskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3SchedulingTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElastigroupAzureV3SchedulingTaskOutputReference_Override(e ElastigroupAzureV3SchedulingTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3SchedulingTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetAdjustment(val *string) {
	if err := j.validateSetAdjustmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adjustment",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetAdjustmentPercentage(val *string) {
	if err := j.validateSetAdjustmentPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adjustmentPercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetBatchSizePercentage(val *string) {
	if err := j.validateSetBatchSizePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSizePercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetCronExpression(val *string) {
	if err := j.validateSetCronExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cronExpression",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetGracePeriod(val *string) {
	if err := j.validateSetGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracePeriod",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetScaleMaxCapacity(val *string) {
	if err := j.validateSetScaleMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleMaxCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetScaleMinCapacity(val *string) {
	if err := j.validateSetScaleMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleMinCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetScaleTargetCapacity(val *string) {
	if err := j.validateSetScaleTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ResetAdjustment() {
	_jsii_.InvokeVoid(
		e,
		"resetAdjustment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ResetAdjustmentPercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetAdjustmentPercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ResetBatchSizePercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetBatchSizePercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ResetGracePeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetGracePeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ResetScaleMaxCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleMaxCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ResetScaleMinCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleMinCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ResetScaleTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3SchedulingTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

