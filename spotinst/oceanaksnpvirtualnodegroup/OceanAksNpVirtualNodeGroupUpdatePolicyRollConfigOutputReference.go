// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanaksnpvirtualnodegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference interface {
	cdktf.ComplexObject
	BatchMinHealthyPercentage() *float64
	SetBatchMinHealthyPercentage(val *float64)
	BatchMinHealthyPercentageInput() *float64
	BatchSizePercentage() *float64
	SetBatchSizePercentage(val *float64)
	BatchSizePercentageInput() *float64
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	InternalValue() *OceanAksNpVirtualNodeGroupUpdatePolicyRollConfig
	SetInternalValue(val *OceanAksNpVirtualNodeGroupUpdatePolicyRollConfig)
	NodeNames() *[]*string
	SetNodeNames(val *[]*string)
	NodeNamesInput() *[]*string
	NodePoolNames() *[]*string
	SetNodePoolNames(val *[]*string)
	NodePoolNamesInput() *[]*string
	RespectPdb() interface{}
	SetRespectPdb(val interface{})
	RespectPdbInput() interface{}
	RespectRestrictScaleDown() interface{}
	SetRespectRestrictScaleDown(val interface{})
	RespectRestrictScaleDownInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VngIds() *[]*string
	SetVngIds(val *[]*string)
	VngIdsInput() *[]*string
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
	ResetBatchMinHealthyPercentage()
	ResetBatchSizePercentage()
	ResetComment()
	ResetNodeNames()
	ResetNodePoolNames()
	ResetRespectPdb()
	ResetRespectRestrictScaleDown()
	ResetVngIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference
type jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) BatchMinHealthyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchMinHealthyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) BatchMinHealthyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchMinHealthyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) BatchSizePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) BatchSizePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) InternalValue() *OceanAksNpVirtualNodeGroupUpdatePolicyRollConfig {
	var returns *OceanAksNpVirtualNodeGroupUpdatePolicyRollConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) NodeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) NodeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) NodePoolNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodePoolNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) NodePoolNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodePoolNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) RespectPdb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectPdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) RespectPdbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectPdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) RespectRestrictScaleDown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectRestrictScaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) RespectRestrictScaleDownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectRestrictScaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) VngIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vngIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) VngIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vngIdsInput",
		&returns,
	)
	return returns
}


func NewOceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference_Override(o OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetBatchMinHealthyPercentage(val *float64) {
	if err := j.validateSetBatchMinHealthyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchMinHealthyPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetBatchSizePercentage(val *float64) {
	if err := j.validateSetBatchSizePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSizePercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetInternalValue(val *OceanAksNpVirtualNodeGroupUpdatePolicyRollConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetNodeNames(val *[]*string) {
	if err := j.validateSetNodeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeNames",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetNodePoolNames(val *[]*string) {
	if err := j.validateSetNodePoolNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodePoolNames",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetRespectPdb(val interface{}) {
	if err := j.validateSetRespectPdbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respectPdb",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetRespectRestrictScaleDown(val interface{}) {
	if err := j.validateSetRespectRestrictScaleDownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respectRestrictScaleDown",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference)SetVngIds(val *[]*string) {
	if err := j.validateSetVngIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vngIds",
		val,
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ResetBatchMinHealthyPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchMinHealthyPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ResetBatchSizePercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchSizePercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		o,
		"resetComment",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ResetNodeNames() {
	_jsii_.InvokeVoid(
		o,
		"resetNodeNames",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ResetNodePoolNames() {
	_jsii_.InvokeVoid(
		o,
		"resetNodePoolNames",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ResetRespectPdb() {
	_jsii_.InvokeVoid(
		o,
		"resetRespectPdb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ResetRespectRestrictScaleDown() {
	_jsii_.InvokeVoid(
		o,
		"resetRespectRestrictScaleDown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ResetVngIds() {
	_jsii_.InvokeVoid(
		o,
		"resetVngIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupUpdatePolicyRollConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

