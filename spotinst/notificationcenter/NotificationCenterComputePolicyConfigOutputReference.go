// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationcenter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/notificationcenter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationCenterComputePolicyConfigOutputReference interface {
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
	DynamicRules() NotificationCenterComputePolicyConfigDynamicRulesList
	DynamicRulesInput() interface{}
	Events() NotificationCenterComputePolicyConfigEventsList
	EventsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *NotificationCenterComputePolicyConfig
	SetInternalValue(val *NotificationCenterComputePolicyConfig)
	ResourceIds() *[]*string
	SetResourceIds(val *[]*string)
	ResourceIdsInput() *[]*string
	ShouldIncludeAllResources() interface{}
	SetShouldIncludeAllResources(val interface{})
	ShouldIncludeAllResourcesInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDynamicRules(value interface{})
	PutEvents(value interface{})
	ResetDynamicRules()
	ResetResourceIds()
	ResetShouldIncludeAllResources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotificationCenterComputePolicyConfigOutputReference
type jsiiProxy_NotificationCenterComputePolicyConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) DynamicRules() NotificationCenterComputePolicyConfigDynamicRulesList {
	var returns NotificationCenterComputePolicyConfigDynamicRulesList
	_jsii_.Get(
		j,
		"dynamicRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) DynamicRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) Events() NotificationCenterComputePolicyConfigEventsList {
	var returns NotificationCenterComputePolicyConfigEventsList
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) EventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) InternalValue() *NotificationCenterComputePolicyConfig {
	var returns *NotificationCenterComputePolicyConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ResourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ResourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ShouldIncludeAllResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldIncludeAllResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ShouldIncludeAllResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldIncludeAllResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNotificationCenterComputePolicyConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotificationCenterComputePolicyConfigOutputReference {
	_init_.Initialize()

	if err := validateNewNotificationCenterComputePolicyConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationCenterComputePolicyConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.notificationCenter.NotificationCenterComputePolicyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotificationCenterComputePolicyConfigOutputReference_Override(n NotificationCenterComputePolicyConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.notificationCenter.NotificationCenterComputePolicyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference)SetInternalValue(val *NotificationCenterComputePolicyConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference)SetResourceIds(val *[]*string) {
	if err := j.validateSetResourceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIds",
		val,
	)
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference)SetShouldIncludeAllResources(val interface{}) {
	if err := j.validateSetShouldIncludeAllResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldIncludeAllResources",
		val,
	)
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) PutDynamicRules(value interface{}) {
	if err := n.validatePutDynamicRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDynamicRules",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) PutEvents(value interface{}) {
	if err := n.validatePutEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putEvents",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ResetDynamicRules() {
	_jsii_.InvokeVoid(
		n,
		"resetDynamicRules",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ResetResourceIds() {
	_jsii_.InvokeVoid(
		n,
		"resetResourceIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ResetShouldIncludeAllResources() {
	_jsii_.InvokeVoid(
		n,
		"resetShouldIncludeAllResources",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationCenterComputePolicyConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

