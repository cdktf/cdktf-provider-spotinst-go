// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanaksnpvirtualnodegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference interface {
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
	InternalValue() *OceanAksNpVirtualNodeGroupSchedulingShutdownHours
	SetInternalValue(val *OceanAksNpVirtualNodeGroupSchedulingShutdownHours)
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeWindows() *[]*string
	SetTimeWindows(val *[]*string)
	TimeWindowsInput() *[]*string
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
	ResetIsEnabled()
	ResetTimeWindows()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference
type jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) InternalValue() *OceanAksNpVirtualNodeGroupSchedulingShutdownHours {
	var returns *OceanAksNpVirtualNodeGroupSchedulingShutdownHours
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) TimeWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"timeWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) TimeWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"timeWindowsInput",
		&returns,
	)
	return returns
}


func NewOceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference_Override(o OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference)SetInternalValue(val *OceanAksNpVirtualNodeGroupSchedulingShutdownHours) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference)SetTimeWindows(val *[]*string) {
	if err := j.validateSetTimeWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeWindows",
		val,
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) ResetTimeWindows() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeWindows",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupSchedulingShutdownHoursOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

