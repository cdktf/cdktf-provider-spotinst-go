// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference interface {
	cdktf.ComplexObject
	BaseSize() *float64
	SetBaseSize(val *float64)
	BaseSizeInput() *float64
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
	InternalValue() *ElastigroupAwsEbsBlockDeviceDynamicVolumeSize
	SetInternalValue(val *ElastigroupAwsEbsBlockDeviceDynamicVolumeSize)
	Resource() *string
	SetResource(val *string)
	ResourceInput() *string
	SizePerResourceUnit() *float64
	SetSizePerResourceUnit(val *float64)
	SizePerResourceUnitInput() *float64
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
	ResetBaseSize()
	ResetResource()
	ResetSizePerResourceUnit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference
type jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) BaseSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) BaseSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) InternalValue() *ElastigroupAwsEbsBlockDeviceDynamicVolumeSize {
	var returns *ElastigroupAwsEbsBlockDeviceDynamicVolumeSize
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) SizePerResourceUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizePerResourceUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) SizePerResourceUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizePerResourceUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference_Override(e ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference)SetBaseSize(val *float64) {
	if err := j.validateSetBaseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference)SetInternalValue(val *ElastigroupAwsEbsBlockDeviceDynamicVolumeSize) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference)SetResource(val *string) {
	if err := j.validateSetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference)SetSizePerResourceUnit(val *float64) {
	if err := j.validateSetSizePerResourceUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizePerResourceUnit",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) ResetBaseSize() {
	_jsii_.InvokeVoid(
		e,
		"resetBaseSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) ResetResource() {
	_jsii_.InvokeVoid(
		e,
		"resetResource",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) ResetSizePerResourceUnit() {
	_jsii_.InvokeVoid(
		e,
		"resetSizePerResourceUnit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsEbsBlockDeviceDynamicVolumeSizeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

