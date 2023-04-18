package mrscaleraws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v7/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v7/mrscaleraws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MrscalerAwsCoreScalingDownPolicyOutputReference interface {
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
	Dimensions() *map[string]*string
	SetDimensions(val *map[string]*string)
	DimensionsInput() *map[string]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MrscalerAwsCoreScalingDownPolicyOutputReference
type jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Adjustment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) AdjustmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Cooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) CooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Dimensions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) DimensionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) EvaluationPeriodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Maximum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) MaximumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) MaxTargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) MaxTargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Minimum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) MinimumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) MinTargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) MinTargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) PolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


func NewMrscalerAwsCoreScalingDownPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MrscalerAwsCoreScalingDownPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewMrscalerAwsCoreScalingDownPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.mrscalerAws.MrscalerAwsCoreScalingDownPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMrscalerAwsCoreScalingDownPolicyOutputReference_Override(m MrscalerAwsCoreScalingDownPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.mrscalerAws.MrscalerAwsCoreScalingDownPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetActionType(val *string) {
	if err := j.validateSetActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionType",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetAdjustment(val *string) {
	if err := j.validateSetAdjustmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adjustment",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetCooldown(val *float64) {
	if err := j.validateSetCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetDimensions(val *map[string]*string) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetEvaluationPeriods(val *float64) {
	if err := j.validateSetEvaluationPeriodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetMaximum(val *string) {
	if err := j.validateSetMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximum",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetMaxTargetCapacity(val *string) {
	if err := j.validateSetMaxTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetMinimum(val *string) {
	if err := j.validateSetMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimum",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetMinTargetCapacity(val *string) {
	if err := j.validateSetMinTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetPolicyName(val *string) {
	if err := j.validateSetPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetActionType() {
	_jsii_.InvokeVoid(
		m,
		"resetActionType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetAdjustment() {
	_jsii_.InvokeVoid(
		m,
		"resetAdjustment",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetCooldown() {
	_jsii_.InvokeVoid(
		m,
		"resetCooldown",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		m,
		"resetDimensions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetEvaluationPeriods() {
	_jsii_.InvokeVoid(
		m,
		"resetEvaluationPeriods",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetMaximum() {
	_jsii_.InvokeVoid(
		m,
		"resetMaximum",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetMaxTargetCapacity() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxTargetCapacity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetMinimum() {
	_jsii_.InvokeVoid(
		m,
		"resetMinimum",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetMinTargetCapacity() {
	_jsii_.InvokeVoid(
		m,
		"resetMinTargetCapacity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		m,
		"resetOperator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetPeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetPeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetStatistic() {
	_jsii_.InvokeVoid(
		m,
		"resetStatistic",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		m,
		"resetTarget",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAwsCoreScalingDownPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

