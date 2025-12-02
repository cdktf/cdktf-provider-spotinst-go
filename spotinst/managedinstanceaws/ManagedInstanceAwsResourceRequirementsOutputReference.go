// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/managedinstanceaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedInstanceAwsResourceRequirementsOutputReference interface {
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

// The jsii proxy struct for ManagedInstanceAwsResourceRequirementsOutputReference
type jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ExcludedInstanceFamilies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ExcludedInstanceFamiliesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ExcludedInstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ExcludedInstanceGenerationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceGenerationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ExcludedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredGpuMaximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredGpuMaximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredGpuMaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredGpuMaximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredGpuMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredGpuMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredGpuMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredGpuMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredMemoryMaximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredMemoryMaximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredMemoryMaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredMemoryMaximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredMemoryMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredMemoryMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredMemoryMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredMemoryMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredVcpuMaximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredVcpuMaximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredVcpuMaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredVcpuMaximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredVcpuMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredVcpuMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) RequiredVcpuMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredVcpuMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedInstanceAwsResourceRequirementsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ManagedInstanceAwsResourceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewManagedInstanceAwsResourceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsResourceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewManagedInstanceAwsResourceRequirementsOutputReference_Override(m ManagedInstanceAwsResourceRequirementsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsResourceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetExcludedInstanceFamilies(val *[]*string) {
	if err := j.validateSetExcludedInstanceFamiliesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceFamilies",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetExcludedInstanceGenerations(val *[]*string) {
	if err := j.validateSetExcludedInstanceGenerationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceGenerations",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetExcludedInstanceTypes(val *[]*string) {
	if err := j.validateSetExcludedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetRequiredGpuMaximum(val *float64) {
	if err := j.validateSetRequiredGpuMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredGpuMaximum",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetRequiredGpuMinimum(val *float64) {
	if err := j.validateSetRequiredGpuMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredGpuMinimum",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetRequiredMemoryMaximum(val *float64) {
	if err := j.validateSetRequiredMemoryMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredMemoryMaximum",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetRequiredMemoryMinimum(val *float64) {
	if err := j.validateSetRequiredMemoryMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredMemoryMinimum",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetRequiredVcpuMaximum(val *float64) {
	if err := j.validateSetRequiredVcpuMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredVcpuMaximum",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetRequiredVcpuMinimum(val *float64) {
	if err := j.validateSetRequiredVcpuMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredVcpuMinimum",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ResetExcludedInstanceFamilies() {
	_jsii_.InvokeVoid(
		m,
		"resetExcludedInstanceFamilies",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ResetExcludedInstanceGenerations() {
	_jsii_.InvokeVoid(
		m,
		"resetExcludedInstanceGenerations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ResetExcludedInstanceTypes() {
	_jsii_.InvokeVoid(
		m,
		"resetExcludedInstanceTypes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ResetRequiredGpuMaximum() {
	_jsii_.InvokeVoid(
		m,
		"resetRequiredGpuMaximum",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ResetRequiredGpuMinimum() {
	_jsii_.InvokeVoid(
		m,
		"resetRequiredGpuMinimum",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

