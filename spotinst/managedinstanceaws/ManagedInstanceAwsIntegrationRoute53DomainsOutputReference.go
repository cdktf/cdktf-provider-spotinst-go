package managedinstanceaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v7/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v7/managedinstanceaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedInstanceAwsIntegrationRoute53DomainsOutputReference interface {
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
	HostedZoneId() *string
	SetHostedZoneId(val *string)
	HostedZoneIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RecordSets() ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsList
	RecordSetsInput() interface{}
	RecordSetType() *string
	SetRecordSetType(val *string)
	RecordSetTypeInput() *string
	SpotinstAcctId() *string
	SetSpotinstAcctId(val *string)
	SpotinstAcctIdInput() *string
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
	PutRecordSets(value interface{})
	ResetRecordSetType()
	ResetSpotinstAcctId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedInstanceAwsIntegrationRoute53DomainsOutputReference
type jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) HostedZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) RecordSets() ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsList {
	var returns ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsList
	_jsii_.Get(
		j,
		"recordSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) RecordSetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) RecordSetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) RecordSetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) SpotinstAcctId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotinstAcctId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) SpotinstAcctIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotinstAcctIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedInstanceAwsIntegrationRoute53DomainsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ManagedInstanceAwsIntegrationRoute53DomainsOutputReference {
	_init_.Initialize()

	if err := validateNewManagedInstanceAwsIntegrationRoute53DomainsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53DomainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewManagedInstanceAwsIntegrationRoute53DomainsOutputReference_Override(m ManagedInstanceAwsIntegrationRoute53DomainsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53DomainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference)SetHostedZoneId(val *string) {
	if err := j.validateSetHostedZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostedZoneId",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference)SetRecordSetType(val *string) {
	if err := j.validateSetRecordSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSetType",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference)SetSpotinstAcctId(val *string) {
	if err := j.validateSetSpotinstAcctIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotinstAcctId",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) PutRecordSets(value interface{}) {
	if err := m.validatePutRecordSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRecordSets",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) ResetRecordSetType() {
	_jsii_.InvokeVoid(
		m,
		"resetRecordSetType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) ResetSpotinstAcctId() {
	_jsii_.InvokeVoid(
		m,
		"resetSpotinstAcctId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

