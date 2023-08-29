// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkelaunchspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/oceangkelaunchspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanGkeLaunchSpecUpdatePolicyOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *OceanGkeLaunchSpecUpdatePolicy
	SetInternalValue(val *OceanGkeLaunchSpecUpdatePolicy)
	RollConfig() OceanGkeLaunchSpecUpdatePolicyRollConfigOutputReference
	RollConfigInput() *OceanGkeLaunchSpecUpdatePolicyRollConfig
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
	PutRollConfig(value *OceanGkeLaunchSpecUpdatePolicyRollConfig)
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

// The jsii proxy struct for OceanGkeLaunchSpecUpdatePolicyOutputReference
type jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) InternalValue() *OceanGkeLaunchSpecUpdatePolicy {
	var returns *OceanGkeLaunchSpecUpdatePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) RollConfig() OceanGkeLaunchSpecUpdatePolicyRollConfigOutputReference {
	var returns OceanGkeLaunchSpecUpdatePolicyRollConfigOutputReference
	_jsii_.Get(
		j,
		"rollConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) RollConfigInput() *OceanGkeLaunchSpecUpdatePolicyRollConfig {
	var returns *OceanGkeLaunchSpecUpdatePolicyRollConfig
	_jsii_.Get(
		j,
		"rollConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) ShouldRoll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldRoll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) ShouldRollInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldRollInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanGkeLaunchSpecUpdatePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanGkeLaunchSpecUpdatePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewOceanGkeLaunchSpecUpdatePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpecUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanGkeLaunchSpecUpdatePolicyOutputReference_Override(o OceanGkeLaunchSpecUpdatePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpecUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference)SetInternalValue(val *OceanGkeLaunchSpecUpdatePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference)SetShouldRoll(val interface{}) {
	if err := j.validateSetShouldRollParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldRoll",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) PutRollConfig(value *OceanGkeLaunchSpecUpdatePolicyRollConfig) {
	if err := o.validatePutRollConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRollConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) ResetRollConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetRollConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecUpdatePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

