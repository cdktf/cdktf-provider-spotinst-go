// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/oceanaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAwsAutoscalerAutoscaleHeadroomOutputReference interface {
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
	CpuPerUnit() *float64
	SetCpuPerUnit(val *float64)
	CpuPerUnitInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GpuPerUnit() *float64
	SetGpuPerUnit(val *float64)
	GpuPerUnitInput() *float64
	InternalValue() *OceanAwsAutoscalerAutoscaleHeadroom
	SetInternalValue(val *OceanAwsAutoscalerAutoscaleHeadroom)
	MemoryPerUnit() *float64
	SetMemoryPerUnit(val *float64)
	MemoryPerUnitInput() *float64
	NumOfUnits() *float64
	SetNumOfUnits(val *float64)
	NumOfUnitsInput() *float64
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
	ResetCpuPerUnit()
	ResetGpuPerUnit()
	ResetMemoryPerUnit()
	ResetNumOfUnits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAwsAutoscalerAutoscaleHeadroomOutputReference
type jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) CpuPerUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPerUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) CpuPerUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPerUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GpuPerUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuPerUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GpuPerUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuPerUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) InternalValue() *OceanAwsAutoscalerAutoscaleHeadroom {
	var returns *OceanAwsAutoscalerAutoscaleHeadroom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) MemoryPerUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) MemoryPerUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) NumOfUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numOfUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) NumOfUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numOfUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanAwsAutoscalerAutoscaleHeadroomOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAwsAutoscalerAutoscaleHeadroomOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAwsAutoscalerAutoscaleHeadroomOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAwsAutoscalerAutoscaleHeadroomOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAwsAutoscalerAutoscaleHeadroomOutputReference_Override(o OceanAwsAutoscalerAutoscaleHeadroomOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAwsAutoscalerAutoscaleHeadroomOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference)SetCpuPerUnit(val *float64) {
	if err := j.validateSetCpuPerUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuPerUnit",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference)SetGpuPerUnit(val *float64) {
	if err := j.validateSetGpuPerUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuPerUnit",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference)SetInternalValue(val *OceanAwsAutoscalerAutoscaleHeadroom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference)SetMemoryPerUnit(val *float64) {
	if err := j.validateSetMemoryPerUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryPerUnit",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference)SetNumOfUnits(val *float64) {
	if err := j.validateSetNumOfUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numOfUnits",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) ResetCpuPerUnit() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuPerUnit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) ResetGpuPerUnit() {
	_jsii_.InvokeVoid(
		o,
		"resetGpuPerUnit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) ResetMemoryPerUnit() {
	_jsii_.InvokeVoid(
		o,
		"resetMemoryPerUnit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) ResetNumOfUnits() {
	_jsii_.InvokeVoid(
		o,
		"resetNumOfUnits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAwsAutoscalerAutoscaleHeadroomOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

