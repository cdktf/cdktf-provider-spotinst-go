// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference interface {
	cdktf.ComplexObject
	AmiAutoUpdateClusterRoll() OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRollOutputReference
	AmiAutoUpdateClusterRollInput() *OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRoll
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
	InternalValue() *OceanAwsScheduledTaskTasksParametersAmiAutoUpdate
	SetInternalValue(val *OceanAwsScheduledTaskTasksParametersAmiAutoUpdate)
	MinorVersion() interface{}
	SetMinorVersion(val interface{})
	MinorVersionInput() interface{}
	Patch() interface{}
	SetPatch(val interface{})
	PatchInput() interface{}
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
	PutAmiAutoUpdateClusterRoll(value *OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRoll)
	ResetAmiAutoUpdateClusterRoll()
	ResetApplyRoll()
	ResetMinorVersion()
	ResetPatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference
type jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) AmiAutoUpdateClusterRoll() OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRollOutputReference {
	var returns OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRollOutputReference
	_jsii_.Get(
		j,
		"amiAutoUpdateClusterRoll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) AmiAutoUpdateClusterRollInput() *OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRoll {
	var returns *OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRoll
	_jsii_.Get(
		j,
		"amiAutoUpdateClusterRollInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ApplyRoll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyRoll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ApplyRollInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyRollInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) InternalValue() *OceanAwsScheduledTaskTasksParametersAmiAutoUpdate {
	var returns *OceanAwsScheduledTaskTasksParametersAmiAutoUpdate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) MinorVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) MinorVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) Patch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"patch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) PatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"patchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference_Override(o OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference)SetApplyRoll(val interface{}) {
	if err := j.validateSetApplyRollParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyRoll",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference)SetInternalValue(val *OceanAwsScheduledTaskTasksParametersAmiAutoUpdate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference)SetMinorVersion(val interface{}) {
	if err := j.validateSetMinorVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minorVersion",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference)SetPatch(val interface{}) {
	if err := j.validateSetPatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patch",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) PutAmiAutoUpdateClusterRoll(value *OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRoll) {
	if err := o.validatePutAmiAutoUpdateClusterRollParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmiAutoUpdateClusterRoll",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ResetAmiAutoUpdateClusterRoll() {
	_jsii_.InvokeVoid(
		o,
		"resetAmiAutoUpdateClusterRoll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ResetApplyRoll() {
	_jsii_.InvokeVoid(
		o,
		"resetApplyRoll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ResetMinorVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetMinorVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ResetPatch() {
	_jsii_.InvokeVoid(
		o,
		"resetPatch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersAmiAutoUpdateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

