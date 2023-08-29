// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsUpdatePolicyOutputReference interface {
	cdktf.ComplexObject
	AutoApplyTags() interface{}
	SetAutoApplyTags(val interface{})
	AutoApplyTagsInput() interface{}
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
	InternalValue() *ElastigroupAwsUpdatePolicy
	SetInternalValue(val *ElastigroupAwsUpdatePolicy)
	RollConfig() ElastigroupAwsUpdatePolicyRollConfigOutputReference
	RollConfigInput() *ElastigroupAwsUpdatePolicyRollConfig
	ShouldResumeStateful() interface{}
	SetShouldResumeStateful(val interface{})
	ShouldResumeStatefulInput() interface{}
	ShouldRoll() interface{}
	SetShouldRoll(val interface{})
	ShouldRollInput() interface{}
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
	PutRollConfig(value *ElastigroupAwsUpdatePolicyRollConfig)
	ResetAutoApplyTags()
	ResetRollConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsUpdatePolicyOutputReference
type jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) AutoApplyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoApplyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) AutoApplyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoApplyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) InternalValue() *ElastigroupAwsUpdatePolicy {
	var returns *ElastigroupAwsUpdatePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) RollConfig() ElastigroupAwsUpdatePolicyRollConfigOutputReference {
	var returns ElastigroupAwsUpdatePolicyRollConfigOutputReference
	_jsii_.Get(
		j,
		"rollConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) RollConfigInput() *ElastigroupAwsUpdatePolicyRollConfig {
	var returns *ElastigroupAwsUpdatePolicyRollConfig
	_jsii_.Get(
		j,
		"rollConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ShouldResumeStateful() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldResumeStateful",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ShouldResumeStatefulInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldResumeStatefulInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ShouldRoll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldRoll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ShouldRollInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldRollInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsUpdatePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsUpdatePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsUpdatePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsUpdatePolicyOutputReference_Override(e ElastigroupAwsUpdatePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference)SetAutoApplyTags(val interface{}) {
	if err := j.validateSetAutoApplyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoApplyTags",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference)SetInternalValue(val *ElastigroupAwsUpdatePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference)SetShouldResumeStateful(val interface{}) {
	if err := j.validateSetShouldResumeStatefulParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldResumeStateful",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference)SetShouldRoll(val interface{}) {
	if err := j.validateSetShouldRollParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldRoll",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) PutRollConfig(value *ElastigroupAwsUpdatePolicyRollConfig) {
	if err := e.validatePutRollConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRollConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ResetAutoApplyTags() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoApplyTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ResetRollConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetRollConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsUpdatePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

