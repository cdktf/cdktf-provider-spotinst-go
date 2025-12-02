// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/elastigroupazurev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference interface {
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
	InternalValue() *ElastigroupAzureV3VmSizesSpotSizeAttributes
	SetInternalValue(val *ElastigroupAzureV3VmSizesSpotSizeAttributes)
	MaxCpu() *float64
	SetMaxCpu(val *float64)
	MaxCpuInput() *float64
	MaxMemory() *float64
	SetMaxMemory(val *float64)
	MaxMemoryInput() *float64
	MaxStorage() *float64
	SetMaxStorage(val *float64)
	MaxStorageInput() *float64
	MinCpu() *float64
	SetMinCpu(val *float64)
	MinCpuInput() *float64
	MinMemory() *float64
	SetMinMemory(val *float64)
	MinMemoryInput() *float64
	MinStorage() *float64
	SetMinStorage(val *float64)
	MinStorageInput() *float64
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
	ResetMaxCpu()
	ResetMaxMemory()
	ResetMaxStorage()
	ResetMinCpu()
	ResetMinMemory()
	ResetMinStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference
type jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) InternalValue() *ElastigroupAzureV3VmSizesSpotSizeAttributes {
	var returns *ElastigroupAzureV3VmSizesSpotSizeAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MaxCpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MaxCpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MaxMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MaxMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MaxStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MaxStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MinCpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MinCpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MinMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MinMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MinStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) MinStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAzureV3VmSizesSpotSizeAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference_Override(e ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetInternalValue(val *ElastigroupAzureV3VmSizesSpotSizeAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetMaxCpu(val *float64) {
	if err := j.validateSetMaxCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCpu",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetMaxMemory(val *float64) {
	if err := j.validateSetMaxMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMemory",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetMaxStorage(val *float64) {
	if err := j.validateSetMaxStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStorage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetMinCpu(val *float64) {
	if err := j.validateSetMinCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpu",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetMinMemory(val *float64) {
	if err := j.validateSetMinMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMemory",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetMinStorage(val *float64) {
	if err := j.validateSetMinStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minStorage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ResetMaxCpu() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxCpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ResetMaxMemory() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxMemory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ResetMaxStorage() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxStorage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ResetMinCpu() {
	_jsii_.InvokeVoid(
		e,
		"resetMinCpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ResetMinMemory() {
	_jsii_.InvokeVoid(
		e,
		"resetMinMemory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ResetMinStorage() {
	_jsii_.InvokeVoid(
		e,
		"resetMinStorage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3VmSizesSpotSizeAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

