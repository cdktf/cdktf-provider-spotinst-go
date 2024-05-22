// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceancdrolloutspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdRolloutSpecTrafficOutputReference interface {
	cdktf.ComplexObject
	Alb() OceancdRolloutSpecTrafficAlbOutputReference
	AlbInput() *OceancdRolloutSpecTrafficAlb
	Ambassador() OceancdRolloutSpecTrafficAmbassadorOutputReference
	AmbassadorInput() *OceancdRolloutSpecTrafficAmbassador
	CanaryService() *string
	SetCanaryService(val *string)
	CanaryServiceInput() *string
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
	InternalValue() *OceancdRolloutSpecTraffic
	SetInternalValue(val *OceancdRolloutSpecTraffic)
	Istio() OceancdRolloutSpecTrafficIstioOutputReference
	IstioInput() *OceancdRolloutSpecTrafficIstio
	Nginx() OceancdRolloutSpecTrafficNginxOutputReference
	NginxInput() *OceancdRolloutSpecTrafficNginx
	PingPong() OceancdRolloutSpecTrafficPingPongOutputReference
	PingPongInput() *OceancdRolloutSpecTrafficPingPong
	Smi() OceancdRolloutSpecTrafficSmiOutputReference
	SmiInput() *OceancdRolloutSpecTrafficSmi
	StableService() *string
	SetStableService(val *string)
	StableServiceInput() *string
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
	PutAlb(value *OceancdRolloutSpecTrafficAlb)
	PutAmbassador(value *OceancdRolloutSpecTrafficAmbassador)
	PutIstio(value *OceancdRolloutSpecTrafficIstio)
	PutNginx(value *OceancdRolloutSpecTrafficNginx)
	PutPingPong(value *OceancdRolloutSpecTrafficPingPong)
	PutSmi(value *OceancdRolloutSpecTrafficSmi)
	ResetAlb()
	ResetAmbassador()
	ResetCanaryService()
	ResetIstio()
	ResetNginx()
	ResetPingPong()
	ResetSmi()
	ResetStableService()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceancdRolloutSpecTrafficOutputReference
type jsiiProxy_OceancdRolloutSpecTrafficOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) Alb() OceancdRolloutSpecTrafficAlbOutputReference {
	var returns OceancdRolloutSpecTrafficAlbOutputReference
	_jsii_.Get(
		j,
		"alb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) AlbInput() *OceancdRolloutSpecTrafficAlb {
	var returns *OceancdRolloutSpecTrafficAlb
	_jsii_.Get(
		j,
		"albInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) Ambassador() OceancdRolloutSpecTrafficAmbassadorOutputReference {
	var returns OceancdRolloutSpecTrafficAmbassadorOutputReference
	_jsii_.Get(
		j,
		"ambassador",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) AmbassadorInput() *OceancdRolloutSpecTrafficAmbassador {
	var returns *OceancdRolloutSpecTrafficAmbassador
	_jsii_.Get(
		j,
		"ambassadorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) CanaryService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) CanaryServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) InternalValue() *OceancdRolloutSpecTraffic {
	var returns *OceancdRolloutSpecTraffic
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) Istio() OceancdRolloutSpecTrafficIstioOutputReference {
	var returns OceancdRolloutSpecTrafficIstioOutputReference
	_jsii_.Get(
		j,
		"istio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) IstioInput() *OceancdRolloutSpecTrafficIstio {
	var returns *OceancdRolloutSpecTrafficIstio
	_jsii_.Get(
		j,
		"istioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) Nginx() OceancdRolloutSpecTrafficNginxOutputReference {
	var returns OceancdRolloutSpecTrafficNginxOutputReference
	_jsii_.Get(
		j,
		"nginx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) NginxInput() *OceancdRolloutSpecTrafficNginx {
	var returns *OceancdRolloutSpecTrafficNginx
	_jsii_.Get(
		j,
		"nginxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) PingPong() OceancdRolloutSpecTrafficPingPongOutputReference {
	var returns OceancdRolloutSpecTrafficPingPongOutputReference
	_jsii_.Get(
		j,
		"pingPong",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) PingPongInput() *OceancdRolloutSpecTrafficPingPong {
	var returns *OceancdRolloutSpecTrafficPingPong
	_jsii_.Get(
		j,
		"pingPongInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) Smi() OceancdRolloutSpecTrafficSmiOutputReference {
	var returns OceancdRolloutSpecTrafficSmiOutputReference
	_jsii_.Get(
		j,
		"smi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) SmiInput() *OceancdRolloutSpecTrafficSmi {
	var returns *OceancdRolloutSpecTrafficSmi
	_jsii_.Get(
		j,
		"smiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) StableService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stableService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) StableServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stableServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceancdRolloutSpecTrafficOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceancdRolloutSpecTrafficOutputReference {
	_init_.Initialize()

	if err := validateNewOceancdRolloutSpecTrafficOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdRolloutSpecTrafficOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecTrafficOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceancdRolloutSpecTrafficOutputReference_Override(o OceancdRolloutSpecTrafficOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecTrafficOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference)SetCanaryService(val *string) {
	if err := j.validateSetCanaryServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canaryService",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference)SetInternalValue(val *OceancdRolloutSpecTraffic) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference)SetStableService(val *string) {
	if err := j.validateSetStableServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stableService",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) PutAlb(value *OceancdRolloutSpecTrafficAlb) {
	if err := o.validatePutAlbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAlb",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) PutAmbassador(value *OceancdRolloutSpecTrafficAmbassador) {
	if err := o.validatePutAmbassadorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmbassador",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) PutIstio(value *OceancdRolloutSpecTrafficIstio) {
	if err := o.validatePutIstioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putIstio",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) PutNginx(value *OceancdRolloutSpecTrafficNginx) {
	if err := o.validatePutNginxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNginx",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) PutPingPong(value *OceancdRolloutSpecTrafficPingPong) {
	if err := o.validatePutPingPongParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPingPong",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) PutSmi(value *OceancdRolloutSpecTrafficSmi) {
	if err := o.validatePutSmiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSmi",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ResetAlb() {
	_jsii_.InvokeVoid(
		o,
		"resetAlb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ResetAmbassador() {
	_jsii_.InvokeVoid(
		o,
		"resetAmbassador",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ResetCanaryService() {
	_jsii_.InvokeVoid(
		o,
		"resetCanaryService",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ResetIstio() {
	_jsii_.InvokeVoid(
		o,
		"resetIstio",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ResetNginx() {
	_jsii_.InvokeVoid(
		o,
		"resetNginx",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ResetPingPong() {
	_jsii_.InvokeVoid(
		o,
		"resetPingPong",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ResetSmi() {
	_jsii_.InvokeVoid(
		o,
		"resetSmi",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ResetStableService() {
	_jsii_.InvokeVoid(
		o,
		"resetStableService",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

