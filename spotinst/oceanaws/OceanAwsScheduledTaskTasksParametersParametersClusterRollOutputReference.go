// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference interface {
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
	InternalValue() *OceanAwsScheduledTaskTasksParametersParametersClusterRoll
	SetInternalValue(val *OceanAwsScheduledTaskTasksParametersParametersClusterRoll)
	RespectPdb() interface{}
	SetRespectPdb(val interface{})
	RespectPdbInput() interface{}
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
	ResetBatchMinHealthyPercentage()
	ResetBatchSizePercentage()
	ResetComment()
	ResetRespectPdb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference
type jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) BatchMinHealthyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchMinHealthyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) BatchMinHealthyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchMinHealthyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) BatchSizePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) BatchSizePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) InternalValue() *OceanAwsScheduledTaskTasksParametersParametersClusterRoll {
	var returns *OceanAwsScheduledTaskTasksParametersParametersClusterRoll
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) RespectPdb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectPdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) RespectPdbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectPdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference_Override(o OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference)SetBatchMinHealthyPercentage(val *float64) {
	if err := j.validateSetBatchMinHealthyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchMinHealthyPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference)SetBatchSizePercentage(val *float64) {
	if err := j.validateSetBatchSizePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSizePercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference)SetInternalValue(val *OceanAwsScheduledTaskTasksParametersParametersClusterRoll) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference)SetRespectPdb(val interface{}) {
	if err := j.validateSetRespectPdbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respectPdb",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) ResetBatchMinHealthyPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchMinHealthyPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) ResetBatchSizePercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchSizePercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		o,
		"resetComment",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) ResetRespectPdb() {
	_jsii_.InvokeVoid(
		o,
		"resetRespectPdb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAwsScheduledTaskTasksParametersParametersClusterRollOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

