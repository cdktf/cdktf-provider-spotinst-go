// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanaks/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksAutoscalerOutputReference interface {
	cdktf.ComplexObject
	AutoscaleDown() OceanAksAutoscalerAutoscaleDownOutputReference
	AutoscaleDownInput() *OceanAksAutoscalerAutoscaleDown
	AutoscaleHeadroom() OceanAksAutoscalerAutoscaleHeadroomOutputReference
	AutoscaleHeadroomInput() *OceanAksAutoscalerAutoscaleHeadroom
	AutoscaleIsEnabled() interface{}
	SetAutoscaleIsEnabled(val interface{})
	AutoscaleIsEnabledInput() interface{}
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
	InternalValue() *OceanAksAutoscaler
	SetInternalValue(val *OceanAksAutoscaler)
	ResourceLimits() OceanAksAutoscalerResourceLimitsOutputReference
	ResourceLimitsInput() *OceanAksAutoscalerResourceLimits
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
	PutAutoscaleDown(value *OceanAksAutoscalerAutoscaleDown)
	PutAutoscaleHeadroom(value *OceanAksAutoscalerAutoscaleHeadroom)
	PutResourceLimits(value *OceanAksAutoscalerResourceLimits)
	ResetAutoscaleDown()
	ResetAutoscaleHeadroom()
	ResetAutoscaleIsEnabled()
	ResetResourceLimits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAksAutoscalerOutputReference
type jsiiProxy_OceanAksAutoscalerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) AutoscaleDown() OceanAksAutoscalerAutoscaleDownOutputReference {
	var returns OceanAksAutoscalerAutoscaleDownOutputReference
	_jsii_.Get(
		j,
		"autoscaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) AutoscaleDownInput() *OceanAksAutoscalerAutoscaleDown {
	var returns *OceanAksAutoscalerAutoscaleDown
	_jsii_.Get(
		j,
		"autoscaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) AutoscaleHeadroom() OceanAksAutoscalerAutoscaleHeadroomOutputReference {
	var returns OceanAksAutoscalerAutoscaleHeadroomOutputReference
	_jsii_.Get(
		j,
		"autoscaleHeadroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) AutoscaleHeadroomInput() *OceanAksAutoscalerAutoscaleHeadroom {
	var returns *OceanAksAutoscalerAutoscaleHeadroom
	_jsii_.Get(
		j,
		"autoscaleHeadroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) AutoscaleIsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) AutoscaleIsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) InternalValue() *OceanAksAutoscaler {
	var returns *OceanAksAutoscaler
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) ResourceLimits() OceanAksAutoscalerResourceLimitsOutputReference {
	var returns OceanAksAutoscalerResourceLimitsOutputReference
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) ResourceLimitsInput() *OceanAksAutoscalerResourceLimits {
	var returns *OceanAksAutoscalerResourceLimits
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanAksAutoscalerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAksAutoscalerOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAksAutoscalerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAksAutoscalerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAks.OceanAksAutoscalerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAksAutoscalerOutputReference_Override(o OceanAksAutoscalerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAks.OceanAksAutoscalerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference)SetAutoscaleIsEnabled(val interface{}) {
	if err := j.validateSetAutoscaleIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleIsEnabled",
		val,
	)
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference)SetInternalValue(val *OceanAksAutoscaler) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAksAutoscalerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) PutAutoscaleDown(value *OceanAksAutoscalerAutoscaleDown) {
	if err := o.validatePutAutoscaleDownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleDown",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) PutAutoscaleHeadroom(value *OceanAksAutoscalerAutoscaleHeadroom) {
	if err := o.validatePutAutoscaleHeadroomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleHeadroom",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) PutResourceLimits(value *OceanAksAutoscalerResourceLimits) {
	if err := o.validatePutResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putResourceLimits",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) ResetAutoscaleDown() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleDown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) ResetAutoscaleHeadroom() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleHeadroom",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) ResetAutoscaleIsEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleIsEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		o,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAksAutoscalerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

