// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAwsAutoscalerOutputReference interface {
	cdktf.ComplexObject
	AutoHeadroomPercentage() *float64
	SetAutoHeadroomPercentage(val *float64)
	AutoHeadroomPercentageInput() *float64
	AutoscaleCooldown() *float64
	SetAutoscaleCooldown(val *float64)
	AutoscaleCooldownInput() *float64
	AutoscaleDown() OceanAwsAutoscalerAutoscaleDownOutputReference
	AutoscaleDownInput() *OceanAwsAutoscalerAutoscaleDown
	AutoscaleHeadroom() OceanAwsAutoscalerAutoscaleHeadroomOutputReference
	AutoscaleHeadroomInput() *OceanAwsAutoscalerAutoscaleHeadroom
	AutoscaleIsAutoConfig() interface{}
	SetAutoscaleIsAutoConfig(val interface{})
	AutoscaleIsAutoConfigInput() interface{}
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
	EnableAutomaticAndManualHeadroom() interface{}
	SetEnableAutomaticAndManualHeadroom(val interface{})
	EnableAutomaticAndManualHeadroomInput() interface{}
	ExtendedResourceDefinitions() *[]*string
	SetExtendedResourceDefinitions(val *[]*string)
	ExtendedResourceDefinitionsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *OceanAwsAutoscaler
	SetInternalValue(val *OceanAwsAutoscaler)
	ResourceLimits() OceanAwsAutoscalerResourceLimitsOutputReference
	ResourceLimitsInput() *OceanAwsAutoscalerResourceLimits
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
	PutAutoscaleDown(value *OceanAwsAutoscalerAutoscaleDown)
	PutAutoscaleHeadroom(value *OceanAwsAutoscalerAutoscaleHeadroom)
	PutResourceLimits(value *OceanAwsAutoscalerResourceLimits)
	ResetAutoHeadroomPercentage()
	ResetAutoscaleCooldown()
	ResetAutoscaleDown()
	ResetAutoscaleHeadroom()
	ResetAutoscaleIsAutoConfig()
	ResetAutoscaleIsEnabled()
	ResetEnableAutomaticAndManualHeadroom()
	ResetExtendedResourceDefinitions()
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

// The jsii proxy struct for OceanAwsAutoscalerOutputReference
type jsiiProxy_OceanAwsAutoscalerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoHeadroomPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoHeadroomPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoHeadroomPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoHeadroomPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleDown() OceanAwsAutoscalerAutoscaleDownOutputReference {
	var returns OceanAwsAutoscalerAutoscaleDownOutputReference
	_jsii_.Get(
		j,
		"autoscaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleDownInput() *OceanAwsAutoscalerAutoscaleDown {
	var returns *OceanAwsAutoscalerAutoscaleDown
	_jsii_.Get(
		j,
		"autoscaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleHeadroom() OceanAwsAutoscalerAutoscaleHeadroomOutputReference {
	var returns OceanAwsAutoscalerAutoscaleHeadroomOutputReference
	_jsii_.Get(
		j,
		"autoscaleHeadroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleHeadroomInput() *OceanAwsAutoscalerAutoscaleHeadroom {
	var returns *OceanAwsAutoscalerAutoscaleHeadroom
	_jsii_.Get(
		j,
		"autoscaleHeadroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleIsAutoConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleIsAutoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleIsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) AutoscaleIsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) EnableAutomaticAndManualHeadroom() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticAndManualHeadroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) EnableAutomaticAndManualHeadroomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticAndManualHeadroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) ExtendedResourceDefinitions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extendedResourceDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) ExtendedResourceDefinitionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extendedResourceDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) InternalValue() *OceanAwsAutoscaler {
	var returns *OceanAwsAutoscaler
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) ResourceLimits() OceanAwsAutoscalerResourceLimitsOutputReference {
	var returns OceanAwsAutoscalerResourceLimitsOutputReference
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) ResourceLimitsInput() *OceanAwsAutoscalerResourceLimits {
	var returns *OceanAwsAutoscalerResourceLimits
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanAwsAutoscalerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAwsAutoscalerOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAwsAutoscalerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAwsAutoscalerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAwsAutoscalerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAwsAutoscalerOutputReference_Override(o OceanAwsAutoscalerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAwsAutoscalerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetAutoHeadroomPercentage(val *float64) {
	if err := j.validateSetAutoHeadroomPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHeadroomPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetAutoscaleCooldown(val *float64) {
	if err := j.validateSetAutoscaleCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleCooldown",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetAutoscaleIsAutoConfig(val interface{}) {
	if err := j.validateSetAutoscaleIsAutoConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleIsAutoConfig",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetAutoscaleIsEnabled(val interface{}) {
	if err := j.validateSetAutoscaleIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleIsEnabled",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetEnableAutomaticAndManualHeadroom(val interface{}) {
	if err := j.validateSetEnableAutomaticAndManualHeadroomParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticAndManualHeadroom",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetExtendedResourceDefinitions(val *[]*string) {
	if err := j.validateSetExtendedResourceDefinitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedResourceDefinitions",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetInternalValue(val *OceanAwsAutoscaler) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAwsAutoscalerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) PutAutoscaleDown(value *OceanAwsAutoscalerAutoscaleDown) {
	if err := o.validatePutAutoscaleDownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleDown",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) PutAutoscaleHeadroom(value *OceanAwsAutoscalerAutoscaleHeadroom) {
	if err := o.validatePutAutoscaleHeadroomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleHeadroom",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) PutResourceLimits(value *OceanAwsAutoscalerResourceLimits) {
	if err := o.validatePutResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putResourceLimits",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ResetAutoHeadroomPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHeadroomPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ResetAutoscaleCooldown() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleCooldown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ResetAutoscaleDown() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleDown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ResetAutoscaleHeadroom() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleHeadroom",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ResetAutoscaleIsAutoConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleIsAutoConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ResetAutoscaleIsEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleIsEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ResetEnableAutomaticAndManualHeadroom() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableAutomaticAndManualHeadroom",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ResetExtendedResourceDefinitions() {
	_jsii_.InvokeVoid(
		o,
		"resetExtendedResourceDefinitions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		o,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAwsAutoscalerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

