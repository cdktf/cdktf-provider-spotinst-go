// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsResourceRequirementsOutputReference interface {
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
	ExcludedInstanceFamilies() *[]*string
	SetExcludedInstanceFamilies(val *[]*string)
	ExcludedInstanceFamiliesInput() *[]*string
	ExcludedInstanceGenerations() *[]*string
	SetExcludedInstanceGenerations(val *[]*string)
	ExcludedInstanceGenerationsInput() *[]*string
	ExcludedInstanceTypes() *[]*string
	SetExcludedInstanceTypes(val *[]*string)
	ExcludedInstanceTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RequiredGpuMaximum() *float64
	SetRequiredGpuMaximum(val *float64)
	RequiredGpuMaximumInput() *float64
	RequiredGpuMinimum() *float64
	SetRequiredGpuMinimum(val *float64)
	RequiredGpuMinimumInput() *float64
	RequiredMemoryMaximum() *float64
	SetRequiredMemoryMaximum(val *float64)
	RequiredMemoryMaximumInput() *float64
	RequiredMemoryMinimum() *float64
	SetRequiredMemoryMinimum(val *float64)
	RequiredMemoryMinimumInput() *float64
	RequiredVcpuMaximum() *float64
	SetRequiredVcpuMaximum(val *float64)
	RequiredVcpuMaximumInput() *float64
	RequiredVcpuMinimum() *float64
	SetRequiredVcpuMinimum(val *float64)
	RequiredVcpuMinimumInput() *float64
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
	ResetExcludedInstanceFamilies()
	ResetExcludedInstanceGenerations()
	ResetExcludedInstanceTypes()
	ResetRequiredGpuMaximum()
	ResetRequiredGpuMinimum()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsResourceRequirementsOutputReference
type jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ExcludedInstanceFamilies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ExcludedInstanceFamiliesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ExcludedInstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ExcludedInstanceGenerationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceGenerationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ExcludedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredGpuMaximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredGpuMaximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredGpuMaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredGpuMaximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredGpuMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredGpuMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredGpuMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredGpuMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredMemoryMaximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredMemoryMaximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredMemoryMaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredMemoryMaximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredMemoryMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredMemoryMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredMemoryMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredMemoryMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredVcpuMaximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredVcpuMaximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredVcpuMaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredVcpuMaximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredVcpuMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredVcpuMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) RequiredVcpuMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredVcpuMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsResourceRequirementsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElastigroupAwsResourceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsResourceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsResourceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElastigroupAwsResourceRequirementsOutputReference_Override(e ElastigroupAwsResourceRequirementsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsResourceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetExcludedInstanceFamilies(val *[]*string) {
	if err := j.validateSetExcludedInstanceFamiliesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceFamilies",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetExcludedInstanceGenerations(val *[]*string) {
	if err := j.validateSetExcludedInstanceGenerationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceGenerations",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetExcludedInstanceTypes(val *[]*string) {
	if err := j.validateSetExcludedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetRequiredGpuMaximum(val *float64) {
	if err := j.validateSetRequiredGpuMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredGpuMaximum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetRequiredGpuMinimum(val *float64) {
	if err := j.validateSetRequiredGpuMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredGpuMinimum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetRequiredMemoryMaximum(val *float64) {
	if err := j.validateSetRequiredMemoryMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredMemoryMaximum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetRequiredMemoryMinimum(val *float64) {
	if err := j.validateSetRequiredMemoryMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredMemoryMinimum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetRequiredVcpuMaximum(val *float64) {
	if err := j.validateSetRequiredVcpuMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredVcpuMaximum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetRequiredVcpuMinimum(val *float64) {
	if err := j.validateSetRequiredVcpuMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredVcpuMinimum",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ResetExcludedInstanceFamilies() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludedInstanceFamilies",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ResetExcludedInstanceGenerations() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludedInstanceGenerations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ResetExcludedInstanceTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludedInstanceTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ResetRequiredGpuMaximum() {
	_jsii_.InvokeVoid(
		e,
		"resetRequiredGpuMaximum",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ResetRequiredGpuMinimum() {
	_jsii_.InvokeVoid(
		e,
		"resetRequiredGpuMinimum",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsResourceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

