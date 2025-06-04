// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanspark/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanSparkIngressOutputReference interface {
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
	Controller() OceanSparkIngressControllerOutputReference
	ControllerInput() *OceanSparkIngressController
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomEndpoint() OceanSparkIngressCustomEndpointOutputReference
	CustomEndpointInput() *OceanSparkIngressCustomEndpoint
	// Experimental.
	Fqn() *string
	InternalValue() *OceanSparkIngress
	SetInternalValue(val *OceanSparkIngress)
	LoadBalancer() OceanSparkIngressLoadBalancerOutputReference
	LoadBalancerInput() *OceanSparkIngressLoadBalancer
	PrivateLink() OceanSparkIngressPrivateLinkOutputReference
	PrivateLinkInput() *OceanSparkIngressPrivateLink
	ServiceAnnotations() *map[string]*string
	SetServiceAnnotations(val *map[string]*string)
	ServiceAnnotationsInput() *map[string]*string
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
	PutController(value *OceanSparkIngressController)
	PutCustomEndpoint(value *OceanSparkIngressCustomEndpoint)
	PutLoadBalancer(value *OceanSparkIngressLoadBalancer)
	PutPrivateLink(value *OceanSparkIngressPrivateLink)
	ResetController()
	ResetCustomEndpoint()
	ResetLoadBalancer()
	ResetPrivateLink()
	ResetServiceAnnotations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanSparkIngressOutputReference
type jsiiProxy_OceanSparkIngressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) Controller() OceanSparkIngressControllerOutputReference {
	var returns OceanSparkIngressControllerOutputReference
	_jsii_.Get(
		j,
		"controller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) ControllerInput() *OceanSparkIngressController {
	var returns *OceanSparkIngressController
	_jsii_.Get(
		j,
		"controllerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) CustomEndpoint() OceanSparkIngressCustomEndpointOutputReference {
	var returns OceanSparkIngressCustomEndpointOutputReference
	_jsii_.Get(
		j,
		"customEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) CustomEndpointInput() *OceanSparkIngressCustomEndpoint {
	var returns *OceanSparkIngressCustomEndpoint
	_jsii_.Get(
		j,
		"customEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) InternalValue() *OceanSparkIngress {
	var returns *OceanSparkIngress
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) LoadBalancer() OceanSparkIngressLoadBalancerOutputReference {
	var returns OceanSparkIngressLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) LoadBalancerInput() *OceanSparkIngressLoadBalancer {
	var returns *OceanSparkIngressLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) PrivateLink() OceanSparkIngressPrivateLinkOutputReference {
	var returns OceanSparkIngressPrivateLinkOutputReference
	_jsii_.Get(
		j,
		"privateLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) PrivateLinkInput() *OceanSparkIngressPrivateLink {
	var returns *OceanSparkIngressPrivateLink
	_jsii_.Get(
		j,
		"privateLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) ServiceAnnotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"serviceAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) ServiceAnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"serviceAnnotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSparkIngressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanSparkIngressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanSparkIngressOutputReference {
	_init_.Initialize()

	if err := validateNewOceanSparkIngressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanSparkIngressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanSpark.OceanSparkIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanSparkIngressOutputReference_Override(o OceanSparkIngressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanSpark.OceanSparkIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanSparkIngressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanSparkIngressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanSparkIngressOutputReference)SetInternalValue(val *OceanSparkIngress) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanSparkIngressOutputReference)SetServiceAnnotations(val *map[string]*string) {
	if err := j.validateSetServiceAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAnnotations",
		val,
	)
}

func (j *jsiiProxy_OceanSparkIngressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanSparkIngressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) PutController(value *OceanSparkIngressController) {
	if err := o.validatePutControllerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putController",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) PutCustomEndpoint(value *OceanSparkIngressCustomEndpoint) {
	if err := o.validatePutCustomEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCustomEndpoint",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) PutLoadBalancer(value *OceanSparkIngressLoadBalancer) {
	if err := o.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) PutPrivateLink(value *OceanSparkIngressPrivateLink) {
	if err := o.validatePutPrivateLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPrivateLink",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) ResetController() {
	_jsii_.InvokeVoid(
		o,
		"resetController",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) ResetCustomEndpoint() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomEndpoint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) ResetPrivateLink() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateLink",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) ResetServiceAnnotations() {
	_jsii_.InvokeVoid(
		o,
		"resetServiceAnnotations",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSparkIngressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanSparkIngressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

