// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsScheduledTaskOutputReference interface {
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
	Frequency() *string
	SetFrequency(val *string)
	FrequencyInput() *string
	GracePeriod() *string
	SetGracePeriod(val *string)
	GracePeriodInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	MaxCapacity() *string
	SetMaxCapacity(val *string)
	MaxCapacityInput() *string
	MinCapacity() *string
	SetMinCapacity(val *string)
	MinCapacityInput() *string
	ScaleMaxCapacity() *string
	SetScaleMaxCapacity(val *string)
	ScaleMaxCapacityInput() *string
	ScaleMinCapacity() *string
	SetScaleMinCapacity(val *string)
	ScaleMinCapacityInput() *string
	ScaleTargetCapacity() *string
	SetScaleTargetCapacity(val *string)
	ScaleTargetCapacityInput() *string
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	TargetCapacity() *string
	SetTargetCapacity(val *string)
	TargetCapacityInput() *string
	TaskType() *string
	SetTaskType(val *string)
	TaskTypeInput() *string
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
	ResetAdjustment()
	ResetAdjustmentPercentage()
	ResetBatchSizePercentage()
	ResetCronExpression()
	ResetFrequency()
	ResetGracePeriod()
	ResetIsEnabled()
	ResetMaxCapacity()
	ResetMinCapacity()
	ResetScaleMaxCapacity()
	ResetScaleMinCapacity()
	ResetScaleTargetCapacity()
	ResetStartTime()
	ResetTargetCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsScheduledTaskOutputReference
type jsiiProxy_ElastigroupAwsScheduledTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) Adjustment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) AdjustmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) AdjustmentPercentage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) AdjustmentPercentageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) BatchSizePercentage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) BatchSizePercentageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) CronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) CronExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GracePeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GracePeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) MaxCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) MaxCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) MinCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) MinCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ScaleMaxCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMaxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ScaleMaxCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMaxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ScaleMinCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMinCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ScaleMinCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMinCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ScaleTargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ScaleTargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) TargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) TargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) TaskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) TaskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsScheduledTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElastigroupAwsScheduledTaskOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastigroupAwsScheduledTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsScheduledTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElastigroupAwsScheduledTaskOutputReference_Override(e ElastigroupAwsScheduledTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsScheduledTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetAdjustment(val *string) {
	_jsii_.Set(
		j,
		"adjustment",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetAdjustmentPercentage(val *string) {
	_jsii_.Set(
		j,
		"adjustmentPercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetBatchSizePercentage(val *string) {
	_jsii_.Set(
		j,
		"batchSizePercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetCronExpression(val *string) {
	_jsii_.Set(
		j,
		"cronExpression",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetFrequency(val *string) {
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetGracePeriod(val *string) {
	_jsii_.Set(
		j,
		"gracePeriod",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetIsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetMaxCapacity(val *string) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetMinCapacity(val *string) {
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetScaleMaxCapacity(val *string) {
	_jsii_.Set(
		j,
		"scaleMaxCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetScaleMinCapacity(val *string) {
	_jsii_.Set(
		j,
		"scaleMinCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetScaleTargetCapacity(val *string) {
	_jsii_.Set(
		j,
		"scaleTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetTargetCapacity(val *string) {
	_jsii_.Set(
		j,
		"targetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetTaskType(val *string) {
	_jsii_.Set(
		j,
		"taskType",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetAdjustment() {
	_jsii_.InvokeVoid(
		e,
		"resetAdjustment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetAdjustmentPercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetAdjustmentPercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetBatchSizePercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetBatchSizePercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetCronExpression() {
	_jsii_.InvokeVoid(
		e,
		"resetCronExpression",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetFrequency() {
	_jsii_.InvokeVoid(
		e,
		"resetFrequency",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetGracePeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetGracePeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetMinCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMinCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetScaleMaxCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleMaxCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetScaleMinCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleMinCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetScaleTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		e,
		"resetStartTime",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ResetTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScheduledTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

