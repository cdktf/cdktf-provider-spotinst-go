// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/elastigroupazurev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference interface {
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
	CrgName() *string
	SetCrgName(val *string)
	CrgNameInput() *string
	CrgResourceGroupName() *string
	SetCrgResourceGroupName(val *string)
	CrgResourceGroupNameInput() *string
	CrgShouldPrioritize() interface{}
	SetCrgShouldPrioritize(val interface{})
	CrgShouldPrioritizeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ElastigroupAzureV3CapacityReservationCapacityReservationGroups
	SetInternalValue(val *ElastigroupAzureV3CapacityReservationCapacityReservationGroups)
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
	ResetCrgShouldPrioritize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference
type jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) CrgName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crgName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) CrgNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crgNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) CrgResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crgResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) CrgResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crgResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) CrgShouldPrioritize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crgShouldPrioritize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) CrgShouldPrioritizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crgShouldPrioritizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) InternalValue() *ElastigroupAzureV3CapacityReservationCapacityReservationGroups {
	var returns *ElastigroupAzureV3CapacityReservationCapacityReservationGroups
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference_Override(e ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference)SetCrgName(val *string) {
	if err := j.validateSetCrgNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crgName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference)SetCrgResourceGroupName(val *string) {
	if err := j.validateSetCrgResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crgResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference)SetCrgShouldPrioritize(val interface{}) {
	if err := j.validateSetCrgShouldPrioritizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crgShouldPrioritize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference)SetInternalValue(val *ElastigroupAzureV3CapacityReservationCapacityReservationGroups) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) ResetCrgShouldPrioritize() {
	_jsii_.InvokeVoid(
		e,
		"resetCrgShouldPrioritize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAzureV3CapacityReservationCapacityReservationGroupsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

