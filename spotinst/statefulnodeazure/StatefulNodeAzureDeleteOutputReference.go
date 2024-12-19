// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/statefulnodeazure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulNodeAzureDeleteOutputReference interface {
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
	DiskShouldDeallocate() interface{}
	SetDiskShouldDeallocate(val interface{})
	DiskShouldDeallocateInput() interface{}
	DiskTtlInHours() *float64
	SetDiskTtlInHours(val *float64)
	DiskTtlInHoursInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NetworkShouldDeallocate() interface{}
	SetNetworkShouldDeallocate(val interface{})
	NetworkShouldDeallocateInput() interface{}
	NetworkTtlInHours() *float64
	SetNetworkTtlInHours(val *float64)
	NetworkTtlInHoursInput() *float64
	PublicIpShouldDeallocate() interface{}
	SetPublicIpShouldDeallocate(val interface{})
	PublicIpShouldDeallocateInput() interface{}
	PublicIpTtlInHours() *float64
	SetPublicIpTtlInHours(val *float64)
	PublicIpTtlInHoursInput() *float64
	ShouldDeregisterFromLb() interface{}
	SetShouldDeregisterFromLb(val interface{})
	ShouldDeregisterFromLbInput() interface{}
	ShouldRevertToOd() interface{}
	SetShouldRevertToOd(val interface{})
	ShouldRevertToOdInput() interface{}
	ShouldTerminateVm() interface{}
	SetShouldTerminateVm(val interface{})
	ShouldTerminateVmInput() interface{}
	SnapshotShouldDeallocate() interface{}
	SetSnapshotShouldDeallocate(val interface{})
	SnapshotShouldDeallocateInput() interface{}
	SnapshotTtlInHours() *float64
	SetSnapshotTtlInHours(val *float64)
	SnapshotTtlInHoursInput() *float64
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
	ResetDiskShouldDeallocate()
	ResetDiskTtlInHours()
	ResetNetworkShouldDeallocate()
	ResetNetworkTtlInHours()
	ResetPublicIpShouldDeallocate()
	ResetPublicIpTtlInHours()
	ResetShouldDeregisterFromLb()
	ResetShouldRevertToOd()
	ResetSnapshotShouldDeallocate()
	ResetSnapshotTtlInHours()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulNodeAzureDeleteOutputReference
type jsiiProxy_StatefulNodeAzureDeleteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) DiskShouldDeallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskShouldDeallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) DiskShouldDeallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskShouldDeallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) DiskTtlInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskTtlInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) DiskTtlInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskTtlInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) NetworkShouldDeallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkShouldDeallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) NetworkShouldDeallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkShouldDeallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) NetworkTtlInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkTtlInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) NetworkTtlInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkTtlInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) PublicIpShouldDeallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIpShouldDeallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) PublicIpShouldDeallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIpShouldDeallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) PublicIpTtlInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"publicIpTtlInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) PublicIpTtlInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"publicIpTtlInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ShouldDeregisterFromLb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeregisterFromLb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ShouldDeregisterFromLbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeregisterFromLbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ShouldRevertToOd() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldRevertToOd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ShouldRevertToOdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldRevertToOdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ShouldTerminateVm() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTerminateVm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ShouldTerminateVmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTerminateVmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) SnapshotShouldDeallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotShouldDeallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) SnapshotShouldDeallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotShouldDeallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) SnapshotTtlInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotTtlInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) SnapshotTtlInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotTtlInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStatefulNodeAzureDeleteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StatefulNodeAzureDeleteOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulNodeAzureDeleteOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulNodeAzureDeleteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureDeleteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStatefulNodeAzureDeleteOutputReference_Override(s StatefulNodeAzureDeleteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureDeleteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetDiskShouldDeallocate(val interface{}) {
	if err := j.validateSetDiskShouldDeallocateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskShouldDeallocate",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetDiskTtlInHours(val *float64) {
	if err := j.validateSetDiskTtlInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskTtlInHours",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetNetworkShouldDeallocate(val interface{}) {
	if err := j.validateSetNetworkShouldDeallocateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkShouldDeallocate",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetNetworkTtlInHours(val *float64) {
	if err := j.validateSetNetworkTtlInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkTtlInHours",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetPublicIpShouldDeallocate(val interface{}) {
	if err := j.validateSetPublicIpShouldDeallocateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpShouldDeallocate",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetPublicIpTtlInHours(val *float64) {
	if err := j.validateSetPublicIpTtlInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpTtlInHours",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetShouldDeregisterFromLb(val interface{}) {
	if err := j.validateSetShouldDeregisterFromLbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDeregisterFromLb",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetShouldRevertToOd(val interface{}) {
	if err := j.validateSetShouldRevertToOdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldRevertToOd",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetShouldTerminateVm(val interface{}) {
	if err := j.validateSetShouldTerminateVmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTerminateVm",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetSnapshotShouldDeallocate(val interface{}) {
	if err := j.validateSetSnapshotShouldDeallocateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotShouldDeallocate",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetSnapshotTtlInHours(val *float64) {
	if err := j.validateSetSnapshotTtlInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotTtlInHours",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureDeleteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetDiskShouldDeallocate() {
	_jsii_.InvokeVoid(
		s,
		"resetDiskShouldDeallocate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetDiskTtlInHours() {
	_jsii_.InvokeVoid(
		s,
		"resetDiskTtlInHours",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetNetworkShouldDeallocate() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkShouldDeallocate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetNetworkTtlInHours() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkTtlInHours",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetPublicIpShouldDeallocate() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicIpShouldDeallocate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetPublicIpTtlInHours() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicIpTtlInHours",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetShouldDeregisterFromLb() {
	_jsii_.InvokeVoid(
		s,
		"resetShouldDeregisterFromLb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetShouldRevertToOd() {
	_jsii_.InvokeVoid(
		s,
		"resetShouldRevertToOd",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetSnapshotShouldDeallocate() {
	_jsii_.InvokeVoid(
		s,
		"resetSnapshotShouldDeallocate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ResetSnapshotTtlInHours() {
	_jsii_.InvokeVoid(
		s,
		"resetSnapshotTtlInHours",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StatefulNodeAzureDeleteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

