// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceancdrolloutspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference interface {
	cdktf.ComplexObject
	CanaryByHeader() *string
	SetCanaryByHeader(val *string)
	CanaryByHeaderInput() *string
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
	InternalValue() *OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotation
	SetInternalValue(val *OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotation)
	Key1() *string
	SetKey1(val *string)
	Key1Input() *string
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
	ResetCanaryByHeader()
	ResetKey1()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference
type jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) CanaryByHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryByHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) CanaryByHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryByHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) InternalValue() *OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotation {
	var returns *OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) Key1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) Key1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference {
	_init_.Initialize()

	if err := validateNewOceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference_Override(o OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference)SetCanaryByHeader(val *string) {
	if err := j.validateSetCanaryByHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canaryByHeader",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference)SetInternalValue(val *OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference)SetKey1(val *string) {
	if err := j.validateSetKey1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key1",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) ResetCanaryByHeader() {
	_jsii_.InvokeVoid(
		o,
		"resetCanaryByHeader",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) ResetKey1() {
	_jsii_.InvokeVoid(
		o,
		"resetKey1",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

