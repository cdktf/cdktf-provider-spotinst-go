// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceancdrolloutspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdRolloutSpecSpotDeploymentsOutputReference interface {
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
	SpotDeploymentsClusterId() *string
	SetSpotDeploymentsClusterId(val *string)
	SpotDeploymentsClusterIdInput() *string
	SpotDeploymentsName() *string
	SetSpotDeploymentsName(val *string)
	SpotDeploymentsNameInput() *string
	SpotDeploymentsNamespace() *string
	SetSpotDeploymentsNamespace(val *string)
	SpotDeploymentsNamespaceInput() *string
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
	ResetSpotDeploymentsClusterId()
	ResetSpotDeploymentsName()
	ResetSpotDeploymentsNamespace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceancdRolloutSpecSpotDeploymentsOutputReference
type jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) SpotDeploymentsClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotDeploymentsClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) SpotDeploymentsClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotDeploymentsClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) SpotDeploymentsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotDeploymentsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) SpotDeploymentsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotDeploymentsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) SpotDeploymentsNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotDeploymentsNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) SpotDeploymentsNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotDeploymentsNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceancdRolloutSpecSpotDeploymentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OceancdRolloutSpecSpotDeploymentsOutputReference {
	_init_.Initialize()

	if err := validateNewOceancdRolloutSpecSpotDeploymentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecSpotDeploymentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOceancdRolloutSpecSpotDeploymentsOutputReference_Override(o OceancdRolloutSpecSpotDeploymentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpecSpotDeploymentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference)SetSpotDeploymentsClusterId(val *string) {
	if err := j.validateSetSpotDeploymentsClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotDeploymentsClusterId",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference)SetSpotDeploymentsName(val *string) {
	if err := j.validateSetSpotDeploymentsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotDeploymentsName",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference)SetSpotDeploymentsNamespace(val *string) {
	if err := j.validateSetSpotDeploymentsNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotDeploymentsNamespace",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) ResetSpotDeploymentsClusterId() {
	_jsii_.InvokeVoid(
		o,
		"resetSpotDeploymentsClusterId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) ResetSpotDeploymentsName() {
	_jsii_.InvokeVoid(
		o,
		"resetSpotDeploymentsName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) ResetSpotDeploymentsNamespace() {
	_jsii_.InvokeVoid(
		o,
		"resetSpotDeploymentsNamespace",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpecSpotDeploymentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

