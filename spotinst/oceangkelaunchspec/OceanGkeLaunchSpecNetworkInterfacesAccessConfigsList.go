// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkelaunchspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v11/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v11/oceangkelaunchspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) OceanGkeLaunchSpecNetworkInterfacesAccessConfigsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList
type jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOceanGkeLaunchSpecNetworkInterfacesAccessConfigsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList {
	_init_.Initialize()

	if err := validateNewOceanGkeLaunchSpecNetworkInterfacesAccessConfigsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOceanGkeLaunchSpecNetworkInterfacesAccessConfigsList_Override(o OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) Get(index *float64) OceanGkeLaunchSpecNetworkInterfacesAccessConfigsOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OceanGkeLaunchSpecNetworkInterfacesAccessConfigsOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanGkeLaunchSpecNetworkInterfacesAccessConfigsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

