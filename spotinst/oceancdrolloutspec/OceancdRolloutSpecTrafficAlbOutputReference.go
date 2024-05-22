// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceancdrolloutspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdRolloutSpecTrafficAlbOutputReference interface {
	cdktf.ComplexObject
	AlbAnnotationPrefix() *string
	SetAlbAnnotationPrefix(val *string)
	AlbAnnotationPrefixInput() *string
	AlbIngress() *string
	SetAlbIngress(val *string)
	AlbIngressInput() *string
	AlbRootService() *string
	SetAlbRootService(val *string)
	AlbRootServiceInput() *string
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
	InternalValue() *OceancdRolloutSpecTrafficAlb
	SetInternalValue(val *OceancdRolloutSpecTrafficAlb)
	ServicePort() *float64
	SetServicePort(val *float64)
	ServicePortInput() *float64
	StickinessConfig() OceancdRolloutSpecTrafficAlbStickinessConfigOutputReference
	StickinessConfigInput() *OceancdRolloutSpecTrafficAlbStickinessConfig
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
	PutStickinessConfig(value *OceancdRolloutSpecTrafficAlbStickinessConfig)
	ResetAlbAnnotationPrefix()
	ResetStickinessConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceancdRolloutSpecTrafficAlbOutputReference
type jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) AlbAnnotationPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"albAnnotationPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) AlbAnnotationPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"albAnnotationPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) AlbIngress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"albIngress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) AlbIngressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"albIngressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) AlbRootService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"albRootService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) AlbRootServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"albRootServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) InternalValue() *OceancdRolloutSpecTrafficAlb {
	var returns *OceancdRolloutSpecTrafficAlb
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) ServicePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"servicePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) ServicePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"servicePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) StickinessConfig() OceancdRolloutSpecTrafficAlbStickinessConfigOutputReference {
	var returns OceancdRolloutSpecTrafficAlbStickinessConfigOutputReference
	_jsii_.Get(
		j,
		"stickinessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) StickinessConfigInput() *OceancdRolloutSpecTrafficAlbStickinessConfig {
	var returns *OceancdRolloutSpecTrafficAlbStickinessConfig
	_jsii_.Get(
		j,
		"stickinessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceancdRolloutSpecTrafficAlbOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceancdRolloutSpecTrafficAlbOutputReference {
	_init_.Initialize()

	if err := validateNewOceancdRolloutSpecTrafficAlbOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecTrafficAlbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceancdRolloutSpecTrafficAlbOutputReference_Override(o OceancdRolloutSpecTrafficAlbOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecTrafficAlbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference)SetAlbAnnotationPrefix(val *string) {
	if err := j.validateSetAlbAnnotationPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"albAnnotationPrefix",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference)SetAlbIngress(val *string) {
	if err := j.validateSetAlbIngressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"albIngress",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference)SetAlbRootService(val *string) {
	if err := j.validateSetAlbRootServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"albRootService",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference)SetInternalValue(val *OceancdRolloutSpecTrafficAlb) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference)SetServicePort(val *float64) {
	if err := j.validateSetServicePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePort",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) PutStickinessConfig(value *OceancdRolloutSpecTrafficAlbStickinessConfig) {
	if err := o.validatePutStickinessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStickinessConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) ResetAlbAnnotationPrefix() {
	_jsii_.InvokeVoid(
		o,
		"resetAlbAnnotationPrefix",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) ResetStickinessConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetStickinessConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficAlbOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

