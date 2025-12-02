// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanecs/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanEcsUpdatePolicyOutputReference interface {
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
	ConditionedRoll() interface{}
	SetConditionedRoll(val interface{})
	ConditionedRollInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *OceanEcsUpdatePolicy
	SetInternalValue(val *OceanEcsUpdatePolicy)
	RollConfig() OceanEcsUpdatePolicyRollConfigOutputReference
	RollConfigInput() *OceanEcsUpdatePolicyRollConfig
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutRollConfig(value *OceanEcsUpdatePolicyRollConfig)
	ResetAutoApplyTags()
	ResetConditionedRoll()
	ResetRollConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanEcsUpdatePolicyOutputReference
type jsiiProxy_OceanEcsUpdatePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) AutoApplyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoApplyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) AutoApplyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoApplyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ConditionedRoll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionedRoll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ConditionedRollInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionedRollInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) InternalValue() *OceanEcsUpdatePolicy {
	var returns *OceanEcsUpdatePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) RollConfig() OceanEcsUpdatePolicyRollConfigOutputReference {
	var returns OceanEcsUpdatePolicyRollConfigOutputReference
	_jsii_.Get(
		j,
		"rollConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) RollConfigInput() *OceanEcsUpdatePolicyRollConfig {
	var returns *OceanEcsUpdatePolicyRollConfig
	_jsii_.Get(
		j,
		"rollConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ShouldRoll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldRoll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ShouldRollInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldRollInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanEcsUpdatePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanEcsUpdatePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewOceanEcsUpdatePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanEcsUpdatePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcsUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanEcsUpdatePolicyOutputReference_Override(o OceanEcsUpdatePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcsUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference)SetAutoApplyTags(val interface{}) {
	if err := j.validateSetAutoApplyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoApplyTags",
		val,
	)
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference)SetConditionedRoll(val interface{}) {
	if err := j.validateSetConditionedRollParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionedRoll",
		val,
	)
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference)SetInternalValue(val *OceanEcsUpdatePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference)SetShouldRoll(val interface{}) {
	if err := j.validateSetShouldRollParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldRoll",
		val,
	)
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanEcsUpdatePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) PutRollConfig(value *OceanEcsUpdatePolicyRollConfig) {
	if err := o.validatePutRollConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRollConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ResetAutoApplyTags() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoApplyTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ResetConditionedRoll() {
	_jsii_.InvokeVoid(
		o,
		"resetConditionedRoll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ResetRollConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetRollConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanEcsUpdatePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

