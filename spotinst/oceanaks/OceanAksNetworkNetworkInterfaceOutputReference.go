// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v10/oceanaks/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksNetworkNetworkInterfaceOutputReference interface {
	cdktf.ComplexObject
	AdditionalIpConfig() OceanAksNetworkNetworkInterfaceAdditionalIpConfigList
	AdditionalIpConfigInput() interface{}
	AssignPublicIp() interface{}
	SetAssignPublicIp(val interface{})
	AssignPublicIpInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsPrimary() interface{}
	SetIsPrimary(val interface{})
	IsPrimaryInput() interface{}
	SecurityGroup() OceanAksNetworkNetworkInterfaceSecurityGroupOutputReference
	SecurityGroupInput() *OceanAksNetworkNetworkInterfaceSecurityGroup
	SubnetName() *string
	SetSubnetName(val *string)
	SubnetNameInput() *string
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
	PutAdditionalIpConfig(value interface{})
	PutSecurityGroup(value *OceanAksNetworkNetworkInterfaceSecurityGroup)
	ResetAdditionalIpConfig()
	ResetAssignPublicIp()
	ResetIsPrimary()
	ResetSecurityGroup()
	ResetSubnetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAksNetworkNetworkInterfaceOutputReference
type jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) AdditionalIpConfig() OceanAksNetworkNetworkInterfaceAdditionalIpConfigList {
	var returns OceanAksNetworkNetworkInterfaceAdditionalIpConfigList
	_jsii_.Get(
		j,
		"additionalIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) AdditionalIpConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) AssignPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) AssignPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) IsPrimary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) IsPrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) SecurityGroup() OceanAksNetworkNetworkInterfaceSecurityGroupOutputReference {
	var returns OceanAksNetworkNetworkInterfaceSecurityGroupOutputReference
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) SecurityGroupInput() *OceanAksNetworkNetworkInterfaceSecurityGroup {
	var returns *OceanAksNetworkNetworkInterfaceSecurityGroup
	_jsii_.Get(
		j,
		"securityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) SubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) SubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanAksNetworkNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OceanAksNetworkNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAksNetworkNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAks.OceanAksNetworkNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOceanAksNetworkNetworkInterfaceOutputReference_Override(o OceanAksNetworkNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAks.OceanAksNetworkNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference)SetAssignPublicIp(val interface{}) {
	if err := j.validateSetAssignPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignPublicIp",
		val,
	)
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference)SetIsPrimary(val interface{}) {
	if err := j.validateSetIsPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPrimary",
		val,
	)
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference)SetSubnetName(val *string) {
	if err := j.validateSetSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetName",
		val,
	)
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) PutAdditionalIpConfig(value interface{}) {
	if err := o.validatePutAdditionalIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAdditionalIpConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) PutSecurityGroup(value *OceanAksNetworkNetworkInterfaceSecurityGroup) {
	if err := o.validatePutSecurityGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSecurityGroup",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) ResetAdditionalIpConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetAdditionalIpConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) ResetAssignPublicIp() {
	_jsii_.InvokeVoid(
		o,
		"resetAssignPublicIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) ResetIsPrimary() {
	_jsii_.InvokeVoid(
		o,
		"resetIsPrimary",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) ResetSecurityGroup() {
	_jsii_.InvokeVoid(
		o,
		"resetSecurityGroup",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) ResetSubnetName() {
	_jsii_.InvokeVoid(
		o,
		"resetSubnetName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAksNetworkNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

