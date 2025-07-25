// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/statefulnodeazure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulNodeAzureVmSizesOutputReference interface {
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
	ExcludedVmSizes() *[]*string
	SetExcludedVmSizes(val *[]*string)
	ExcludedVmSizesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *StatefulNodeAzureVmSizes
	SetInternalValue(val *StatefulNodeAzureVmSizes)
	OdSizes() *[]*string
	SetOdSizes(val *[]*string)
	OdSizesInput() *[]*string
	PreferredSpotSizes() *[]*string
	SetPreferredSpotSizes(val *[]*string)
	PreferredSpotSizesInput() *[]*string
	SpotSizeAttributes() StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference
	SpotSizeAttributesInput() *StatefulNodeAzureVmSizesSpotSizeAttributes
	SpotSizes() *[]*string
	SetSpotSizes(val *[]*string)
	SpotSizesInput() *[]*string
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
	PutSpotSizeAttributes(value *StatefulNodeAzureVmSizesSpotSizeAttributes)
	ResetExcludedVmSizes()
	ResetPreferredSpotSizes()
	ResetSpotSizeAttributes()
	ResetSpotSizes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulNodeAzureVmSizesOutputReference
type jsiiProxy_StatefulNodeAzureVmSizesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ExcludedVmSizes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedVmSizes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ExcludedVmSizesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedVmSizesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) InternalValue() *StatefulNodeAzureVmSizes {
	var returns *StatefulNodeAzureVmSizes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) OdSizes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"odSizes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) OdSizesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"odSizesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) PreferredSpotSizes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredSpotSizes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) PreferredSpotSizesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredSpotSizesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) SpotSizeAttributes() StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference {
	var returns StatefulNodeAzureVmSizesSpotSizeAttributesOutputReference
	_jsii_.Get(
		j,
		"spotSizeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) SpotSizeAttributesInput() *StatefulNodeAzureVmSizesSpotSizeAttributes {
	var returns *StatefulNodeAzureVmSizesSpotSizeAttributes
	_jsii_.Get(
		j,
		"spotSizeAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) SpotSizes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"spotSizes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) SpotSizesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"spotSizesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStatefulNodeAzureVmSizesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StatefulNodeAzureVmSizesOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulNodeAzureVmSizesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulNodeAzureVmSizesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureVmSizesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStatefulNodeAzureVmSizesOutputReference_Override(s StatefulNodeAzureVmSizesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureVmSizesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference)SetExcludedVmSizes(val *[]*string) {
	if err := j.validateSetExcludedVmSizesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedVmSizes",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference)SetInternalValue(val *StatefulNodeAzureVmSizes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference)SetOdSizes(val *[]*string) {
	if err := j.validateSetOdSizesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odSizes",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference)SetPreferredSpotSizes(val *[]*string) {
	if err := j.validateSetPreferredSpotSizesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredSpotSizes",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference)SetSpotSizes(val *[]*string) {
	if err := j.validateSetSpotSizesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotSizes",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureVmSizesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) PutSpotSizeAttributes(value *StatefulNodeAzureVmSizesSpotSizeAttributes) {
	if err := s.validatePutSpotSizeAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSpotSizeAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ResetExcludedVmSizes() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludedVmSizes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ResetPreferredSpotSizes() {
	_jsii_.InvokeVoid(
		s,
		"resetPreferredSpotSizes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ResetSpotSizeAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetSpotSizeAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ResetSpotSizes() {
	_jsii_.InvokeVoid(
		s,
		"resetSpotSizes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StatefulNodeAzureVmSizesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

