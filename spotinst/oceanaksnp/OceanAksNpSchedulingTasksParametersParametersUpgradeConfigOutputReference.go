// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanaksnp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference interface {
	cdktf.ComplexObject
	ApplyRoll() interface{}
	SetApplyRoll(val interface{})
	ApplyRollInput() interface{}
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
	InternalValue() *OceanAksNpSchedulingTasksParametersParametersUpgradeConfig
	SetInternalValue(val *OceanAksNpSchedulingTasksParametersParametersUpgradeConfig)
	RollParameters() OceanAksNpSchedulingTasksParametersParametersUpgradeConfigRollParametersOutputReference
	RollParametersInput() *OceanAksNpSchedulingTasksParametersParametersUpgradeConfigRollParameters
	ScopeVersion() *string
	SetScopeVersion(val *string)
	ScopeVersionInput() *string
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
	PutRollParameters(value *OceanAksNpSchedulingTasksParametersParametersUpgradeConfigRollParameters)
	ResetApplyRoll()
	ResetRollParameters()
	ResetScopeVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference
type jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ApplyRoll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyRoll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ApplyRollInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyRollInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) InternalValue() *OceanAksNpSchedulingTasksParametersParametersUpgradeConfig {
	var returns *OceanAksNpSchedulingTasksParametersParametersUpgradeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) RollParameters() OceanAksNpSchedulingTasksParametersParametersUpgradeConfigRollParametersOutputReference {
	var returns OceanAksNpSchedulingTasksParametersParametersUpgradeConfigRollParametersOutputReference
	_jsii_.Get(
		j,
		"rollParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) RollParametersInput() *OceanAksNpSchedulingTasksParametersParametersUpgradeConfigRollParameters {
	var returns *OceanAksNpSchedulingTasksParametersParametersUpgradeConfigRollParameters
	_jsii_.Get(
		j,
		"rollParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ScopeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ScopeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNp.OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference_Override(o OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNp.OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference)SetApplyRoll(val interface{}) {
	if err := j.validateSetApplyRollParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyRoll",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference)SetInternalValue(val *OceanAksNpSchedulingTasksParametersParametersUpgradeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference)SetScopeVersion(val *string) {
	if err := j.validateSetScopeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeVersion",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) PutRollParameters(value *OceanAksNpSchedulingTasksParametersParametersUpgradeConfigRollParameters) {
	if err := o.validatePutRollParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRollParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ResetApplyRoll() {
	_jsii_.InvokeVoid(
		o,
		"resetApplyRoll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ResetRollParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetRollParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ResetScopeVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetScopeVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAksNpSchedulingTasksParametersParametersUpgradeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

