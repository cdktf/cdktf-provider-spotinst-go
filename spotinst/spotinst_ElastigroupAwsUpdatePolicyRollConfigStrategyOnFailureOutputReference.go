// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference interface {
	cdktf.ComplexObject
	ActionType() *string
	SetActionType(val *string)
	ActionTypeInput() *string
	BatchNum() *float64
	SetBatchNum(val *float64)
	BatchNumInput() *float64
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
	DrainingTimeout() *float64
	SetDrainingTimeout(val *float64)
	DrainingTimeoutInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure
	SetInternalValue(val *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure)
	ShouldDecrementTargetCapacity() interface{}
	SetShouldDecrementTargetCapacity(val interface{})
	ShouldDecrementTargetCapacityInput() interface{}
	ShouldHandleAllBatches() interface{}
	SetShouldHandleAllBatches(val interface{})
	ShouldHandleAllBatchesInput() interface{}
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
	ResetBatchNum()
	ResetDrainingTimeout()
	ResetShouldDecrementTargetCapacity()
	ResetShouldHandleAllBatches()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference
type jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) BatchNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) BatchNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) DrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) DrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) InternalValue() *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure {
	var returns *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ShouldDecrementTargetCapacity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDecrementTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ShouldDecrementTargetCapacityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDecrementTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ShouldHandleAllBatches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldHandleAllBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ShouldHandleAllBatchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldHandleAllBatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference_Override(e ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetActionType(val *string) {
	if err := j.validateSetActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionType",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetBatchNum(val *float64) {
	if err := j.validateSetBatchNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchNum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetDrainingTimeout(val *float64) {
	if err := j.validateSetDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainingTimeout",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetInternalValue(val *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetShouldDecrementTargetCapacity(val interface{}) {
	if err := j.validateSetShouldDecrementTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDecrementTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetShouldHandleAllBatches(val interface{}) {
	if err := j.validateSetShouldHandleAllBatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldHandleAllBatches",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ResetBatchNum() {
	_jsii_.InvokeVoid(
		e,
		"resetBatchNum",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ResetDrainingTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetDrainingTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ResetShouldDecrementTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldDecrementTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ResetShouldHandleAllBatches() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldHandleAllBatches",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

