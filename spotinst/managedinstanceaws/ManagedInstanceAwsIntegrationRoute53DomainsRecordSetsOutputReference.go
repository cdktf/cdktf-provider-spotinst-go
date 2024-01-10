// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/managedinstanceaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsePublicDns() interface{}
	SetUsePublicDns(val interface{})
	UsePublicDnsInput() interface{}
	UsePublicIp() interface{}
	SetUsePublicIp(val interface{})
	UsePublicIpInput() interface{}
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
	ResetUsePublicDns()
	ResetUsePublicIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference
type jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) UsePublicDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePublicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) UsePublicDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePublicDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) UsePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) UsePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePublicIpInput",
		&returns,
	)
	return returns
}


func NewManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference {
	_init_.Initialize()

	if err := validateNewManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference_Override(m ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference)SetUsePublicDns(val interface{}) {
	if err := j.validateSetUsePublicDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePublicDns",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference)SetUsePublicIp(val interface{}) {
	if err := j.validateSetUsePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePublicIp",
		val,
	)
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) ResetUsePublicDns() {
	_jsii_.InvokeVoid(
		m,
		"resetUsePublicDns",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) ResetUsePublicIp() {
	_jsii_.InvokeVoid(
		m,
		"resetUsePublicIp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

