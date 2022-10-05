package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v3/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v3/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ElastigroupAwsScalingUpPolicyStepAdjustmentsAction
	SetInternalValue(val *ElastigroupAwsScalingUpPolicyStepAdjustmentsAction)
	Maximum() *string
	SetMaximum(val *string)
	MaximumInput() *string
	MaxTargetCapacity() *string
	SetMaxTargetCapacity(val *string)
	MaxTargetCapacityInput() *string
	Minimum() *string
	SetMinimum(val *string)
	MinimumInput() *string
	MinTargetCapacity() *string
	SetMinTargetCapacity(val *string)
	MinTargetCapacityInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAdjustment()
	ResetMaximum()
	ResetMaxTargetCapacity()
	ResetMinimum()
	ResetMinTargetCapacity()
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

// The jsii proxy struct for ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference
type jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) Adjustment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) AdjustmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) InternalValue() *ElastigroupAwsScalingUpPolicyStepAdjustmentsAction {
	var returns *ElastigroupAwsScalingUpPolicyStepAdjustmentsAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) Maximum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) MaximumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) MaxTargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) MaxTargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) Minimum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) MinimumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) MinTargetCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) MinTargetCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference_Override(e ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetAdjustment(val *string) {
	if err := j.validateSetAdjustmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adjustment",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetInternalValue(val *ElastigroupAwsScalingUpPolicyStepAdjustmentsAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetMaximum(val *string) {
	if err := j.validateSetMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetMaxTargetCapacity(val *string) {
	if err := j.validateSetMaxTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetMinimum(val *string) {
	if err := j.validateSetMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetMinTargetCapacity(val *string) {
	if err := j.validateSetMinTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ResetAdjustment() {
	_jsii_.InvokeVoid(
		e,
		"resetAdjustment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ResetMaximum() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximum",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ResetMaxTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ResetMinimum() {
	_jsii_.InvokeVoid(
		e,
		"resetMinimum",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ResetMinTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMinTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		e,
		"resetTarget",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsScalingUpPolicyStepAdjustmentsActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

