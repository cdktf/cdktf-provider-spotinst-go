// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceancdverificationtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdVerificationTemplateMetricsProviderJenkinsOutputReference interface {
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
	InternalValue() *OceancdVerificationTemplateMetricsProviderJenkins
	SetInternalValue(val *OceancdVerificationTemplateMetricsProviderJenkins)
	JenkinsInterval() *string
	SetJenkinsInterval(val *string)
	JenkinsIntervalInput() *string
	JenkinsParameters() OceancdVerificationTemplateMetricsProviderJenkinsJenkinsParametersOutputReference
	JenkinsParametersInput() *OceancdVerificationTemplateMetricsProviderJenkinsJenkinsParameters
	PipelineName() *string
	SetPipelineName(val *string)
	PipelineNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	TlsVerification() interface{}
	SetTlsVerification(val interface{})
	TlsVerificationInput() interface{}
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
	PutJenkinsParameters(value *OceancdVerificationTemplateMetricsProviderJenkinsJenkinsParameters)
	ResetJenkinsParameters()
	ResetTlsVerification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceancdVerificationTemplateMetricsProviderJenkinsOutputReference
type jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) InternalValue() *OceancdVerificationTemplateMetricsProviderJenkins {
	var returns *OceancdVerificationTemplateMetricsProviderJenkins
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) JenkinsInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jenkinsInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) JenkinsIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jenkinsIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) JenkinsParameters() OceancdVerificationTemplateMetricsProviderJenkinsJenkinsParametersOutputReference {
	var returns OceancdVerificationTemplateMetricsProviderJenkinsJenkinsParametersOutputReference
	_jsii_.Get(
		j,
		"jenkinsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) JenkinsParametersInput() *OceancdVerificationTemplateMetricsProviderJenkinsJenkinsParameters {
	var returns *OceancdVerificationTemplateMetricsProviderJenkinsJenkinsParameters
	_jsii_.Get(
		j,
		"jenkinsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) PipelineNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) TlsVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) TlsVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsVerificationInput",
		&returns,
	)
	return returns
}


func NewOceancdVerificationTemplateMetricsProviderJenkinsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceancdVerificationTemplateMetricsProviderJenkinsOutputReference {
	_init_.Initialize()

	if err := validateNewOceancdVerificationTemplateMetricsProviderJenkinsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdVerificationTemplate.OceancdVerificationTemplateMetricsProviderJenkinsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceancdVerificationTemplateMetricsProviderJenkinsOutputReference_Override(o OceancdVerificationTemplateMetricsProviderJenkinsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdVerificationTemplate.OceancdVerificationTemplateMetricsProviderJenkinsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference)SetInternalValue(val *OceancdVerificationTemplateMetricsProviderJenkins) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference)SetJenkinsInterval(val *string) {
	if err := j.validateSetJenkinsIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jenkinsInterval",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference)SetPipelineName(val *string) {
	if err := j.validateSetPipelineNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineName",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference)SetTlsVerification(val interface{}) {
	if err := j.validateSetTlsVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsVerification",
		val,
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) PutJenkinsParameters(value *OceancdVerificationTemplateMetricsProviderJenkinsJenkinsParameters) {
	if err := o.validatePutJenkinsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putJenkinsParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) ResetJenkinsParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetJenkinsParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) ResetTlsVerification() {
	_jsii_.InvokeVoid(
		o,
		"resetTlsVerification",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderJenkinsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

