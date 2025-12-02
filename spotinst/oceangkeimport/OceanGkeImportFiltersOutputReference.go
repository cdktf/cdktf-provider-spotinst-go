// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceangkeimport/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanGkeImportFiltersOutputReference interface {
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
	ExcludeFamilies() *[]*string
	SetExcludeFamilies(val *[]*string)
	ExcludeFamiliesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeFamilies() *[]*string
	SetIncludeFamilies(val *[]*string)
	IncludeFamiliesInput() *[]*string
	InternalValue() *OceanGkeImportFilters
	SetInternalValue(val *OceanGkeImportFilters)
	MaxMemoryGib() *float64
	SetMaxMemoryGib(val *float64)
	MaxMemoryGibInput() *float64
	MaxVcpu() *float64
	SetMaxVcpu(val *float64)
	MaxVcpuInput() *float64
	MinMemoryGib() *float64
	SetMinMemoryGib(val *float64)
	MinMemoryGibInput() *float64
	MinVcpu() *float64
	SetMinVcpu(val *float64)
	MinVcpuInput() *float64
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
	ResetExcludeFamilies()
	ResetIncludeFamilies()
	ResetMaxMemoryGib()
	ResetMaxVcpu()
	ResetMinMemoryGib()
	ResetMinVcpu()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanGkeImportFiltersOutputReference
type jsiiProxy_OceanGkeImportFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) ExcludeFamilies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) ExcludeFamiliesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) IncludeFamilies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) IncludeFamiliesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) InternalValue() *OceanGkeImportFilters {
	var returns *OceanGkeImportFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) MaxMemoryGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) MaxMemoryGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) MaxVcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) MaxVcpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) MinMemoryGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) MinMemoryGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) MinVcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) MinVcpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanGkeImportFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanGkeImportFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewOceanGkeImportFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanGkeImportFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImportFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanGkeImportFiltersOutputReference_Override(o OceanGkeImportFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImportFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetExcludeFamilies(val *[]*string) {
	if err := j.validateSetExcludeFamiliesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeFamilies",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetIncludeFamilies(val *[]*string) {
	if err := j.validateSetIncludeFamiliesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeFamilies",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetInternalValue(val *OceanGkeImportFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetMaxMemoryGib(val *float64) {
	if err := j.validateSetMaxMemoryGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMemoryGib",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetMaxVcpu(val *float64) {
	if err := j.validateSetMaxVcpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVcpu",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetMinMemoryGib(val *float64) {
	if err := j.validateSetMinMemoryGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMemoryGib",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetMinVcpu(val *float64) {
	if err := j.validateSetMinVcpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minVcpu",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) ResetExcludeFamilies() {
	_jsii_.InvokeVoid(
		o,
		"resetExcludeFamilies",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) ResetIncludeFamilies() {
	_jsii_.InvokeVoid(
		o,
		"resetIncludeFamilies",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) ResetMaxMemoryGib() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxMemoryGib",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) ResetMaxVcpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxVcpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) ResetMinMemoryGib() {
	_jsii_.InvokeVoid(
		o,
		"resetMinMemoryGib",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) ResetMinVcpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMinVcpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanGkeImportFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

