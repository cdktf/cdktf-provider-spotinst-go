// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/elastigroupazure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAzureScalingUpPolicyOutputReference interface {
	cdktf.ComplexObject
	ActionType() *string
	SetActionType(val *string)
	ActionTypeInput() *string
	Adjustment() *string
	SetAdjustment(val *string)
	AdjustmentInput() *string
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
	Cooldown() *float64
	SetCooldown(val *float64)
	CooldownInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimensions() ElastigroupAzureScalingUpPolicyDimensionsList
	DimensionsInput() interface{}
	EvaluationPeriods() *float64
	SetEvaluationPeriods(val *float64)
	EvaluationPeriodsInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Maximum() *string
	SetMaximum(val *string)
	MaximumInput() *string
	MaxTargetCapacity() *string
	SetMaxTargetCapacity(val *string)
	MaxTargetCapacityInput() *string
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	Minimum() *string
	SetMinimum(val *string)
	MinimumInput() *string
	MinTargetCapacity() *string
	SetMinTargetCapacity(val *string)
	MinTargetCapacityInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Period() *float64
	SetPeriod(val *float64)
	PeriodInput() *float64
	PolicyName() *string
	SetPolicyName(val *string)
	PolicyNameInput() *string
	Statistic() *string
	SetStatistic(val *string)
	StatisticInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
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
	PutDimensions(value interface{})
	ResetActionType()
	ResetAdjustment()
	ResetCooldown()
	ResetDimensions()
	ResetEvaluationPeriods()
	ResetMaximum()
	ResetMaxTargetCapacity()
	ResetMinimum()
	ResetMinTargetCapacity()
	ResetOperator()
	ResetPeriod()
	ResetStatistic()
	ResetTarget()
	ResetUnit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAzureScalingUpPolicyOutputReference
type jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Adjustment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) AdjustmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Cooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) CooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Dimensions() ElastigroupAzureScalingUpPolicyDimensionsList {
	var returns ElastigroupAzureScalingUpPolicyDimensionsList
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) DimensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) EvaluationPeriodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Maximum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) MaximumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) MaxTargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) MaxTargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Minimum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) MinimumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) MinTargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) MinTargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) PolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


func NewElastigroupAzureScalingUpPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElastigroupAzureScalingUpPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAzureScalingUpPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzure.ElastigroupAzureScalingUpPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElastigroupAzureScalingUpPolicyOutputReference_Override(e ElastigroupAzureScalingUpPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzure.ElastigroupAzureScalingUpPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetActionType(val *string) {
	if err := j.validateSetActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionType",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetAdjustment(val *string) {
	if err := j.validateSetAdjustmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adjustment",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetCooldown(val *float64) {
	if err := j.validateSetCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetEvaluationPeriods(val *float64) {
	if err := j.validateSetEvaluationPeriodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetMaximum(val *string) {
	if err := j.validateSetMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetMaxTargetCapacity(val *string) {
	if err := j.validateSetMaxTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetMinimum(val *string) {
	if err := j.validateSetMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetMinTargetCapacity(val *string) {
	if err := j.validateSetMinTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetPolicyName(val *string) {
	if err := j.validateSetPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) PutDimensions(value interface{}) {
	if err := e.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDimensions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetActionType() {
	_jsii_.InvokeVoid(
		e,
		"resetActionType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetAdjustment() {
	_jsii_.InvokeVoid(
		e,
		"resetAdjustment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetCooldown() {
	_jsii_.InvokeVoid(
		e,
		"resetCooldown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		e,
		"resetDimensions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetEvaluationPeriods() {
	_jsii_.InvokeVoid(
		e,
		"resetEvaluationPeriods",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetMaximum() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximum",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetMaxTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetMinimum() {
	_jsii_.InvokeVoid(
		e,
		"resetMinimum",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetMinTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMinTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		e,
		"resetOperator",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetPeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetPeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetStatistic() {
	_jsii_.InvokeVoid(
		e,
		"resetStatistic",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		e,
		"resetTarget",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ResetUnit() {
	_jsii_.InvokeVoid(
		e,
		"resetUnit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAzureScalingUpPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

