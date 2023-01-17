package statefulnodeazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v5/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v5/statefulnodeazure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulNodeAzureNetworkNetworkInterfaceOutputReference interface {
	cdktf.ComplexObject
	AdditionalIpConfigurations() StatefulNodeAzureNetworkNetworkInterfaceAdditionalIpConfigurationsList
	AdditionalIpConfigurationsInput() interface{}
	ApplicationSecurityGroups() StatefulNodeAzureNetworkNetworkInterfaceApplicationSecurityGroupsList
	ApplicationSecurityGroupsInput() interface{}
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
	EnableIpForwarding() interface{}
	SetEnableIpForwarding(val interface{})
	EnableIpForwardingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsPrimary() interface{}
	SetIsPrimary(val interface{})
	IsPrimaryInput() interface{}
	NetworkSecurityGroup() StatefulNodeAzureNetworkNetworkInterfaceNetworkSecurityGroupList
	NetworkSecurityGroupInput() interface{}
	PrivateIpAddresses() *[]*string
	SetPrivateIpAddresses(val *[]*string)
	PrivateIpAddressesInput() *[]*string
	PublicIps() StatefulNodeAzureNetworkNetworkInterfacePublicIpsList
	PublicIpsInput() interface{}
	PublicIpSku() *string
	SetPublicIpSku(val *string)
	PublicIpSkuInput() *string
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
	PutAdditionalIpConfigurations(value interface{})
	PutApplicationSecurityGroups(value interface{})
	PutNetworkSecurityGroup(value interface{})
	PutPublicIps(value interface{})
	ResetAdditionalIpConfigurations()
	ResetApplicationSecurityGroups()
	ResetAssignPublicIp()
	ResetEnableIpForwarding()
	ResetNetworkSecurityGroup()
	ResetPrivateIpAddresses()
	ResetPublicIps()
	ResetPublicIpSku()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulNodeAzureNetworkNetworkInterfaceOutputReference
type jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) AdditionalIpConfigurations() StatefulNodeAzureNetworkNetworkInterfaceAdditionalIpConfigurationsList {
	var returns StatefulNodeAzureNetworkNetworkInterfaceAdditionalIpConfigurationsList
	_jsii_.Get(
		j,
		"additionalIpConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) AdditionalIpConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalIpConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ApplicationSecurityGroups() StatefulNodeAzureNetworkNetworkInterfaceApplicationSecurityGroupsList {
	var returns StatefulNodeAzureNetworkNetworkInterfaceApplicationSecurityGroupsList
	_jsii_.Get(
		j,
		"applicationSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ApplicationSecurityGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) AssignPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) AssignPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) EnableIpForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) EnableIpForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) IsPrimary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) IsPrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) NetworkSecurityGroup() StatefulNodeAzureNetworkNetworkInterfaceNetworkSecurityGroupList {
	var returns StatefulNodeAzureNetworkNetworkInterfaceNetworkSecurityGroupList
	_jsii_.Get(
		j,
		"networkSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) NetworkSecurityGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PrivateIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PublicIps() StatefulNodeAzureNetworkNetworkInterfacePublicIpsList {
	var returns StatefulNodeAzureNetworkNetworkInterfacePublicIpsList
	_jsii_.Get(
		j,
		"publicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PublicIpSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PublicIpSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) SubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) SubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStatefulNodeAzureNetworkNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StatefulNodeAzureNetworkNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulNodeAzureNetworkNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureNetworkNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStatefulNodeAzureNetworkNetworkInterfaceOutputReference_Override(s StatefulNodeAzureNetworkNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureNetworkNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetAssignPublicIp(val interface{}) {
	if err := j.validateSetAssignPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignPublicIp",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetEnableIpForwarding(val interface{}) {
	if err := j.validateSetEnableIpForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIpForwarding",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetIsPrimary(val interface{}) {
	if err := j.validateSetIsPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPrimary",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetPrivateIpAddresses(val *[]*string) {
	if err := j.validateSetPrivateIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddresses",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetPublicIpSku(val *string) {
	if err := j.validateSetPublicIpSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpSku",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetSubnetName(val *string) {
	if err := j.validateSetSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetName",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PutAdditionalIpConfigurations(value interface{}) {
	if err := s.validatePutAdditionalIpConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdditionalIpConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PutApplicationSecurityGroups(value interface{}) {
	if err := s.validatePutApplicationSecurityGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putApplicationSecurityGroups",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PutNetworkSecurityGroup(value interface{}) {
	if err := s.validatePutNetworkSecurityGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkSecurityGroup",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) PutPublicIps(value interface{}) {
	if err := s.validatePutPublicIpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPublicIps",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ResetAdditionalIpConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetAdditionalIpConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ResetApplicationSecurityGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationSecurityGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ResetAssignPublicIp() {
	_jsii_.InvokeVoid(
		s,
		"resetAssignPublicIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ResetEnableIpForwarding() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableIpForwarding",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ResetNetworkSecurityGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkSecurityGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ResetPrivateIpAddresses() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateIpAddresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ResetPublicIps() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicIps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ResetPublicIpSku() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicIpSku",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureNetworkNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

