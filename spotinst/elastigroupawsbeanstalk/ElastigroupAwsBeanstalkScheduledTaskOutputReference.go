package elastigroupawsbeanstalk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/elastigroupawsbeanstalk/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsBeanstalkScheduledTaskOutputReference interface {
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

// The jsii proxy struct for ElastigroupAwsBeanstalkScheduledTaskOutputReference
type jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) Adjustment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) AdjustmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) AdjustmentPercentage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) AdjustmentPercentageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) BatchSizePercentage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) BatchSizePercentageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) CronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) CronExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GracePeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GracePeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) MaxCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) MaxCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) MinCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) MinCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ScaleMaxCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMaxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ScaleMaxCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMaxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ScaleMinCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMinCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ScaleMinCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleMinCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ScaleTargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ScaleTargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) TargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) TargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) TaskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) TaskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsBeanstalkScheduledTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElastigroupAwsBeanstalkScheduledTaskOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsBeanstalkScheduledTaskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAwsBeanstalk.ElastigroupAwsBeanstalkScheduledTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElastigroupAwsBeanstalkScheduledTaskOutputReference_Override(e ElastigroupAwsBeanstalkScheduledTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAwsBeanstalk.ElastigroupAwsBeanstalkScheduledTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetAdjustment(val *string) {
	if err := j.validateSetAdjustmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adjustment",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetAdjustmentPercentage(val *string) {
	if err := j.validateSetAdjustmentPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adjustmentPercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetBatchSizePercentage(val *string) {
	if err := j.validateSetBatchSizePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSizePercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetCronExpression(val *string) {
	if err := j.validateSetCronExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cronExpression",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetFrequency(val *string) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetGracePeriod(val *string) {
	if err := j.validateSetGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracePeriod",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetMaxCapacity(val *string) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetMinCapacity(val *string) {
	if err := j.validateSetMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetScaleMaxCapacity(val *string) {
	if err := j.validateSetScaleMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleMaxCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetScaleMinCapacity(val *string) {
	if err := j.validateSetScaleMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleMinCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetScaleTargetCapacity(val *string) {
	if err := j.validateSetScaleTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetTargetCapacity(val *string) {
	if err := j.validateSetTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetTaskType(val *string) {
	if err := j.validateSetTaskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskType",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetAdjustment() {
	_jsii_.InvokeVoid(
		e,
		"resetAdjustment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetAdjustmentPercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetAdjustmentPercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetBatchSizePercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetBatchSizePercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetCronExpression() {
	_jsii_.InvokeVoid(
		e,
		"resetCronExpression",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetFrequency() {
	_jsii_.InvokeVoid(
		e,
		"resetFrequency",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetGracePeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetGracePeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetMinCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMinCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetScaleMaxCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleMaxCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetScaleMinCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleMinCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetScaleTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		e,
		"resetStartTime",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ResetTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsBeanstalkScheduledTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

