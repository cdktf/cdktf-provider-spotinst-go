// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanawslaunchspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference interface {
	cdktf.ComplexObject
	BaseSize() *float64
	SetBaseSize(val *float64)
	BaseSizeInput() *float64
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
	InternalValue() *OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIops
	SetInternalValue(val *OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIops)
	Resource() *string
	SetResource(val *string)
	ResourceInput() *string
	SizePerResourceUnit() *float64
	SetSizePerResourceUnit(val *float64)
	SizePerResourceUnitInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference
type jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) BaseSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) BaseSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) InternalValue() *OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIops {
	var returns *OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIops
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) SizePerResourceUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizePerResourceUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) SizePerResourceUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizePerResourceUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAwsLaunchSpec.OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference_Override(o OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAwsLaunchSpec.OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference)SetBaseSize(val *float64) {
	if err := j.validateSetBaseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseSize",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference)SetInternalValue(val *OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIops) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference)SetResource(val *string) {
	if err := j.validateSetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference)SetSizePerResourceUnit(val *float64) {
	if err := j.validateSetSizePerResourceUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizePerResourceUnit",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicIopsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

