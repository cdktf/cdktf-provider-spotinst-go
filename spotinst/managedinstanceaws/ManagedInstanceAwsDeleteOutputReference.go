// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/managedinstanceaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedInstanceAwsDeleteOutputReference interface {
	cdktf.ComplexObject
	AmiBackupShouldDeleteImages() interface{}
	SetAmiBackupShouldDeleteImages(val interface{})
	AmiBackupShouldDeleteImagesInput() interface{}
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
	DeallocationConfigShouldDeleteImages() interface{}
	SetDeallocationConfigShouldDeleteImages(val interface{})
	DeallocationConfigShouldDeleteImagesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ShouldDeleteNetworkInterfaces() interface{}
	SetShouldDeleteNetworkInterfaces(val interface{})
	ShouldDeleteNetworkInterfacesInput() interface{}
	ShouldDeleteSnapshots() interface{}
	SetShouldDeleteSnapshots(val interface{})
	ShouldDeleteSnapshotsInput() interface{}
	ShouldDeleteVolumes() interface{}
	SetShouldDeleteVolumes(val interface{})
	ShouldDeleteVolumesInput() interface{}
	ShouldTerminateInstance() interface{}
	SetShouldTerminateInstance(val interface{})
	ShouldTerminateInstanceInput() interface{}
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
	ResetAmiBackupShouldDeleteImages()
	ResetDeallocationConfigShouldDeleteImages()
	ResetShouldDeleteNetworkInterfaces()
	ResetShouldDeleteSnapshots()
	ResetShouldDeleteVolumes()
	ResetShouldTerminateInstance()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedInstanceAwsDeleteOutputReference
type jsiiProxy_ManagedInstanceAwsDeleteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) AmiBackupShouldDeleteImages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amiBackupShouldDeleteImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) AmiBackupShouldDeleteImagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amiBackupShouldDeleteImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) DeallocationConfigShouldDeleteImages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deallocationConfigShouldDeleteImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) DeallocationConfigShouldDeleteImagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deallocationConfigShouldDeleteImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ShouldDeleteNetworkInterfaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteNetworkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ShouldDeleteNetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteNetworkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ShouldDeleteSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ShouldDeleteSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ShouldDeleteVolumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ShouldDeleteVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ShouldTerminateInstance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTerminateInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ShouldTerminateInstanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTerminateInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedInstanceAwsDeleteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ManagedInstanceAwsDeleteOutputReference {
	_init_.Initialize()

	if err := validateNewManagedInstanceAwsDeleteOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedInstanceAwsDeleteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsDeleteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewManagedInstanceAwsDeleteOutputReference_Override(m ManagedInstanceAwsDeleteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsDeleteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetAmiBackupShouldDeleteImages(val interface{}) {
	if err := j.validateSetAmiBackupShouldDeleteImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amiBackupShouldDeleteImages",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetDeallocationConfigShouldDeleteImages(val interface{}) {
	if err := j.validateSetDeallocationConfigShouldDeleteImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deallocationConfigShouldDeleteImages",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetShouldDeleteNetworkInterfaces(val interface{}) {
	if err := j.validateSetShouldDeleteNetworkInterfacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDeleteNetworkInterfaces",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetShouldDeleteSnapshots(val interface{}) {
	if err := j.validateSetShouldDeleteSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDeleteSnapshots",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetShouldDeleteVolumes(val interface{}) {
	if err := j.validateSetShouldDeleteVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDeleteVolumes",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetShouldTerminateInstance(val interface{}) {
	if err := j.validateSetShouldTerminateInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTerminateInstance",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsDeleteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ResetAmiBackupShouldDeleteImages() {
	_jsii_.InvokeVoid(
		m,
		"resetAmiBackupShouldDeleteImages",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ResetDeallocationConfigShouldDeleteImages() {
	_jsii_.InvokeVoid(
		m,
		"resetDeallocationConfigShouldDeleteImages",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ResetShouldDeleteNetworkInterfaces() {
	_jsii_.InvokeVoid(
		m,
		"resetShouldDeleteNetworkInterfaces",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ResetShouldDeleteSnapshots() {
	_jsii_.InvokeVoid(
		m,
		"resetShouldDeleteSnapshots",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ResetShouldDeleteVolumes() {
	_jsii_.InvokeVoid(
		m,
		"resetShouldDeleteVolumes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ResetShouldTerminateInstance() {
	_jsii_.InvokeVoid(
		m,
		"resetShouldTerminateInstance",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedInstanceAwsDeleteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

