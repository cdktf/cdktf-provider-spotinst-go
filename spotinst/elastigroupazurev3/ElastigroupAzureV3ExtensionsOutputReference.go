// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/elastigroupazurev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAzureV3ExtensionsOutputReference interface {
	cdktf.ComplexObject
	ApiVersion() *string
	SetApiVersion(val *string)
	ApiVersionInput() *string
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
	EnableAutomaticUpgrade() interface{}
	SetEnableAutomaticUpgrade(val interface{})
	EnableAutomaticUpgradeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinorVersionAutoUpgrade() interface{}
	SetMinorVersionAutoUpgrade(val interface{})
	MinorVersionAutoUpgradeInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	ProtectedSettings() *map[string]*string
	SetProtectedSettings(val *map[string]*string)
	ProtectedSettingsFromKeyVault() ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVaultOutputReference
	ProtectedSettingsFromKeyVaultInput() *ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVault
	ProtectedSettingsInput() *map[string]*string
	PublicSettings() *map[string]*string
	SetPublicSettings(val *map[string]*string)
	PublicSettingsInput() *map[string]*string
	Publisher() *string
	SetPublisher(val *string)
	PublisherInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutProtectedSettingsFromKeyVault(value *ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVault)
	ResetEnableAutomaticUpgrade()
	ResetProtectedSettings()
	ResetProtectedSettingsFromKeyVault()
	ResetPublicSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAzureV3ExtensionsOutputReference
type jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) EnableAutomaticUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) EnableAutomaticUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) MinorVersionAutoUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minorVersionAutoUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) MinorVersionAutoUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minorVersionAutoUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ProtectedSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"protectedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ProtectedSettingsFromKeyVault() ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVaultOutputReference {
	var returns ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVaultOutputReference
	_jsii_.Get(
		j,
		"protectedSettingsFromKeyVault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ProtectedSettingsFromKeyVaultInput() *ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVault {
	var returns *ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVault
	_jsii_.Get(
		j,
		"protectedSettingsFromKeyVaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ProtectedSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"protectedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) PublicSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"publicSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) PublicSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"publicSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewElastigroupAzureV3ExtensionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElastigroupAzureV3ExtensionsOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAzureV3ExtensionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3ExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElastigroupAzureV3ExtensionsOutputReference_Override(e ElastigroupAzureV3ExtensionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3ExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetApiVersion(val *string) {
	if err := j.validateSetApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiVersion",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetEnableAutomaticUpgrade(val interface{}) {
	if err := j.validateSetEnableAutomaticUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticUpgrade",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetMinorVersionAutoUpgrade(val interface{}) {
	if err := j.validateSetMinorVersionAutoUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minorVersionAutoUpgrade",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetProtectedSettings(val *map[string]*string) {
	if err := j.validateSetProtectedSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectedSettings",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetPublicSettings(val *map[string]*string) {
	if err := j.validateSetPublicSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicSettings",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetPublisher(val *string) {
	if err := j.validateSetPublisherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) PutProtectedSettingsFromKeyVault(value *ElastigroupAzureV3ExtensionsProtectedSettingsFromKeyVault) {
	if err := e.validatePutProtectedSettingsFromKeyVaultParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProtectedSettingsFromKeyVault",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ResetEnableAutomaticUpgrade() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableAutomaticUpgrade",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ResetProtectedSettings() {
	_jsii_.InvokeVoid(
		e,
		"resetProtectedSettings",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ResetProtectedSettingsFromKeyVault() {
	_jsii_.InvokeVoid(
		e,
		"resetProtectedSettingsFromKeyVault",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ResetPublicSettings() {
	_jsii_.InvokeVoid(
		e,
		"resetPublicSettings",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3ExtensionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

