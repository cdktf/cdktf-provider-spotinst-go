package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v5/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v5/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	BatchMinHealthyPercentage() *float64
	SetBatchMinHealthyPercentage(val *float64)
	BatchMinHealthyPercentageInput() *float64
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
	InternalValue() *ElastigroupAwsUpdatePolicyRollConfigStrategy
	SetInternalValue(val *ElastigroupAwsUpdatePolicyRollConfigStrategy)
	OnFailure() ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference
	OnFailureInput() *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure
	ShouldDrainInstances() interface{}
	SetShouldDrainInstances(val interface{})
	ShouldDrainInstancesInput() interface{}
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
	PutOnFailure(value *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure)
	ResetBatchMinHealthyPercentage()
	ResetOnFailure()
	ResetShouldDrainInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference
type jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) BatchMinHealthyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchMinHealthyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) BatchMinHealthyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchMinHealthyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) InternalValue() *ElastigroupAwsUpdatePolicyRollConfigStrategy {
	var returns *ElastigroupAwsUpdatePolicyRollConfigStrategy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) OnFailure() ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference {
	var returns ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailureOutputReference
	_jsii_.Get(
		j,
		"onFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) OnFailureInput() *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure {
	var returns *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure
	_jsii_.Get(
		j,
		"onFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ShouldDrainInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDrainInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ShouldDrainInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDrainInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsUpdatePolicyRollConfigStrategyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference_Override(e ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference)SetBatchMinHealthyPercentage(val *float64) {
	if err := j.validateSetBatchMinHealthyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchMinHealthyPercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference)SetInternalValue(val *ElastigroupAwsUpdatePolicyRollConfigStrategy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference)SetShouldDrainInstances(val interface{}) {
	if err := j.validateSetShouldDrainInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDrainInstances",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) PutOnFailure(value *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure) {
	if err := e.validatePutOnFailureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOnFailure",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ResetBatchMinHealthyPercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetBatchMinHealthyPercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ResetOnFailure() {
	_jsii_.InvokeVoid(
		e,
		"resetOnFailure",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ResetShouldDrainInstances() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldDrainInstances",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyRollConfigStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

