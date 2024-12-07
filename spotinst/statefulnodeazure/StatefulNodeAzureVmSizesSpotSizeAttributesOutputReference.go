// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/statefulnodeazure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference interface {
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
	InternalValue() *StatefulNodeAzureVmSizesSpotSizeAttributes
	SetInternalValue(val *StatefulNodeAzureVmSizesSpotSizeAttributes)
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetMaxCpu()
	ResetMaxMemory()
	ResetMaxStorage()
	ResetMinCpu()
	ResetMinMemory()
	ResetMinStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference
type jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) InternalValue() *StatefulNodeAzureVmSizesSpotSizeAttributes {
	var returns *StatefulNodeAzureVmSizesSpotSizeAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MaxCpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MaxCpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MaxMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MaxMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MaxStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MaxStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MinCpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MinCpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MinMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MinMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MinStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) MinStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStatefulNodeAzureVmSizesSpotSizeAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulNodeAzureVmSizesSpotSizeAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStatefulNodeAzureVmSizesSpotSizeAttributesOutputReference_Override(s StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetInternalValue(val *StatefulNodeAzureVmSizesSpotSizeAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetMaxCpu(val *float64) {
	if err := j.validateSetMaxCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCpu",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetMaxMemory(val *float64) {
	if err := j.validateSetMaxMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMemory",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetMaxStorage(val *float64) {
	if err := j.validateSetMaxStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStorage",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetMinCpu(val *float64) {
	if err := j.validateSetMinCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpu",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetMinMemory(val *float64) {
	if err := j.validateSetMinMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMemory",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetMinStorage(val *float64) {
	if err := j.validateSetMinStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minStorage",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ResetMaxCpu() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxCpu",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ResetMaxMemory() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxMemory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ResetMaxStorage() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxStorage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ResetMinCpu() {
	_jsii_.InvokeVoid(
		s,
		"resetMinCpu",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ResetMinMemory() {
	_jsii_.InvokeVoid(
		s,
		"resetMinMemory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ResetMinStorage() {
	_jsii_.InvokeVoid(
		s,
		"resetMinStorage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

