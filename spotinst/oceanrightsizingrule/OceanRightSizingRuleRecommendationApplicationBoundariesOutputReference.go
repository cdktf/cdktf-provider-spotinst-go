// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanrightsizingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference interface {
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
	CpuMax() *float64
	SetCpuMax(val *float64)
	CpuMaxInput() *float64
	CpuMin() *float64
	SetCpuMin(val *float64)
	CpuMinInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MemoryMax() *float64
	SetMemoryMax(val *float64)
	MemoryMaxInput() *float64
	MemoryMin() *float64
	SetMemoryMin(val *float64)
	MemoryMinInput() *float64
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
	ResetCpuMax()
	ResetCpuMin()
	ResetMemoryMax()
	ResetMemoryMin()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference
type jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) CpuMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) CpuMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) CpuMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) CpuMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) MemoryMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) MemoryMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) MemoryMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) MemoryMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanRightSizingRuleRecommendationApplicationBoundariesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference {
	_init_.Initialize()

	if err := validateNewOceanRightSizingRuleRecommendationApplicationBoundariesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOceanRightSizingRuleRecommendationApplicationBoundariesOutputReference_Override(o OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference)SetCpuMax(val *float64) {
	if err := j.validateSetCpuMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuMax",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference)SetCpuMin(val *float64) {
	if err := j.validateSetCpuMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuMin",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference)SetMemoryMax(val *float64) {
	if err := j.validateSetMemoryMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryMax",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference)SetMemoryMin(val *float64) {
	if err := j.validateSetMemoryMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryMin",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) ResetCpuMax() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuMax",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) ResetCpuMin() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuMin",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) ResetMemoryMax() {
	_jsii_.InvokeVoid(
		o,
		"resetMemoryMax",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) ResetMemoryMin() {
	_jsii_.InvokeVoid(
		o,
		"resetMemoryMin",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationBoundariesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

