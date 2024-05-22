// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceancdrolloutspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference interface {
	cdktf.ComplexObject
	CanarySubsetName() *string
	SetCanarySubsetName(val *string)
	CanarySubsetNameInput() *string
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
	DestinationRuleName() *string
	SetDestinationRuleName(val *string)
	DestinationRuleNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *OceancdRolloutSpecTrafficIstioDestinationRule
	SetInternalValue(val *OceancdRolloutSpecTrafficIstioDestinationRule)
	StableSubsetName() *string
	SetStableSubsetName(val *string)
	StableSubsetNameInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference
type jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) CanarySubsetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canarySubsetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) CanarySubsetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canarySubsetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) DestinationRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) DestinationRuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) InternalValue() *OceancdRolloutSpecTrafficIstioDestinationRule {
	var returns *OceancdRolloutSpecTrafficIstioDestinationRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) StableSubsetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stableSubsetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) StableSubsetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stableSubsetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceancdRolloutSpecTrafficIstioDestinationRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference {
	_init_.Initialize()

	if err := validateNewOceancdRolloutSpecTrafficIstioDestinationRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceancdRolloutSpecTrafficIstioDestinationRuleOutputReference_Override(o OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference)SetCanarySubsetName(val *string) {
	if err := j.validateSetCanarySubsetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canarySubsetName",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference)SetDestinationRuleName(val *string) {
	if err := j.validateSetDestinationRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationRuleName",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference)SetInternalValue(val *OceancdRolloutSpecTrafficIstioDestinationRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference)SetStableSubsetName(val *string) {
	if err := j.validateSetStableSubsetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stableSubsetName",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpecTrafficIstioDestinationRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

