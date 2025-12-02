// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceancdverificationtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdVerificationTemplateMetricsProviderOutputReference interface {
	cdktf.ComplexObject
	CloudWatch() OceancdVerificationTemplateMetricsProviderCloudWatchOutputReference
	CloudWatchInput() *OceancdVerificationTemplateMetricsProviderCloudWatch
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
	Datadog() OceancdVerificationTemplateMetricsProviderDatadogOutputReference
	DatadogInput() *OceancdVerificationTemplateMetricsProviderDatadog
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Jenkins() OceancdVerificationTemplateMetricsProviderJenkinsOutputReference
	JenkinsInput() *OceancdVerificationTemplateMetricsProviderJenkins
	Job() OceancdVerificationTemplateMetricsProviderJobOutputReference
	JobInput() *OceancdVerificationTemplateMetricsProviderJob
	NewRelic() OceancdVerificationTemplateMetricsProviderNewRelicOutputReference
	NewRelicInput() *OceancdVerificationTemplateMetricsProviderNewRelic
	Prometheus() OceancdVerificationTemplateMetricsProviderPrometheusOutputReference
	PrometheusInput() *OceancdVerificationTemplateMetricsProviderPrometheus
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Web() OceancdVerificationTemplateMetricsProviderWebOutputReference
	WebInput() *OceancdVerificationTemplateMetricsProviderWeb
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
	PutCloudWatch(value *OceancdVerificationTemplateMetricsProviderCloudWatch)
	PutDatadog(value *OceancdVerificationTemplateMetricsProviderDatadog)
	PutJenkins(value *OceancdVerificationTemplateMetricsProviderJenkins)
	PutJob(value *OceancdVerificationTemplateMetricsProviderJob)
	PutNewRelic(value *OceancdVerificationTemplateMetricsProviderNewRelic)
	PutPrometheus(value *OceancdVerificationTemplateMetricsProviderPrometheus)
	PutWeb(value *OceancdVerificationTemplateMetricsProviderWeb)
	ResetCloudWatch()
	ResetDatadog()
	ResetJenkins()
	ResetJob()
	ResetNewRelic()
	ResetPrometheus()
	ResetWeb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceancdVerificationTemplateMetricsProviderOutputReference
type jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) CloudWatch() OceancdVerificationTemplateMetricsProviderCloudWatchOutputReference {
	var returns OceancdVerificationTemplateMetricsProviderCloudWatchOutputReference
	_jsii_.Get(
		j,
		"cloudWatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) CloudWatchInput() *OceancdVerificationTemplateMetricsProviderCloudWatch {
	var returns *OceancdVerificationTemplateMetricsProviderCloudWatch
	_jsii_.Get(
		j,
		"cloudWatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) Datadog() OceancdVerificationTemplateMetricsProviderDatadogOutputReference {
	var returns OceancdVerificationTemplateMetricsProviderDatadogOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) DatadogInput() *OceancdVerificationTemplateMetricsProviderDatadog {
	var returns *OceancdVerificationTemplateMetricsProviderDatadog
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) Jenkins() OceancdVerificationTemplateMetricsProviderJenkinsOutputReference {
	var returns OceancdVerificationTemplateMetricsProviderJenkinsOutputReference
	_jsii_.Get(
		j,
		"jenkins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) JenkinsInput() *OceancdVerificationTemplateMetricsProviderJenkins {
	var returns *OceancdVerificationTemplateMetricsProviderJenkins
	_jsii_.Get(
		j,
		"jenkinsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) Job() OceancdVerificationTemplateMetricsProviderJobOutputReference {
	var returns OceancdVerificationTemplateMetricsProviderJobOutputReference
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) JobInput() *OceancdVerificationTemplateMetricsProviderJob {
	var returns *OceancdVerificationTemplateMetricsProviderJob
	_jsii_.Get(
		j,
		"jobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) NewRelic() OceancdVerificationTemplateMetricsProviderNewRelicOutputReference {
	var returns OceancdVerificationTemplateMetricsProviderNewRelicOutputReference
	_jsii_.Get(
		j,
		"newRelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) NewRelicInput() *OceancdVerificationTemplateMetricsProviderNewRelic {
	var returns *OceancdVerificationTemplateMetricsProviderNewRelic
	_jsii_.Get(
		j,
		"newRelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) Prometheus() OceancdVerificationTemplateMetricsProviderPrometheusOutputReference {
	var returns OceancdVerificationTemplateMetricsProviderPrometheusOutputReference
	_jsii_.Get(
		j,
		"prometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) PrometheusInput() *OceancdVerificationTemplateMetricsProviderPrometheus {
	var returns *OceancdVerificationTemplateMetricsProviderPrometheus
	_jsii_.Get(
		j,
		"prometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) Web() OceancdVerificationTemplateMetricsProviderWebOutputReference {
	var returns OceancdVerificationTemplateMetricsProviderWebOutputReference
	_jsii_.Get(
		j,
		"web",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) WebInput() *OceancdVerificationTemplateMetricsProviderWeb {
	var returns *OceancdVerificationTemplateMetricsProviderWeb
	_jsii_.Get(
		j,
		"webInput",
		&returns,
	)
	return returns
}


func NewOceancdVerificationTemplateMetricsProviderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OceancdVerificationTemplateMetricsProviderOutputReference {
	_init_.Initialize()

	if err := validateNewOceancdVerificationTemplateMetricsProviderOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdVerificationTemplate.OceancdVerificationTemplateMetricsProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOceancdVerificationTemplateMetricsProviderOutputReference_Override(o OceancdVerificationTemplateMetricsProviderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdVerificationTemplate.OceancdVerificationTemplateMetricsProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) PutCloudWatch(value *OceancdVerificationTemplateMetricsProviderCloudWatch) {
	if err := o.validatePutCloudWatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCloudWatch",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) PutDatadog(value *OceancdVerificationTemplateMetricsProviderDatadog) {
	if err := o.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatadog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) PutJenkins(value *OceancdVerificationTemplateMetricsProviderJenkins) {
	if err := o.validatePutJenkinsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putJenkins",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) PutJob(value *OceancdVerificationTemplateMetricsProviderJob) {
	if err := o.validatePutJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putJob",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) PutNewRelic(value *OceancdVerificationTemplateMetricsProviderNewRelic) {
	if err := o.validatePutNewRelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNewRelic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) PutPrometheus(value *OceancdVerificationTemplateMetricsProviderPrometheus) {
	if err := o.validatePutPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPrometheus",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) PutWeb(value *OceancdVerificationTemplateMetricsProviderWeb) {
	if err := o.validatePutWebParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWeb",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ResetCloudWatch() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudWatch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		o,
		"resetDatadog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ResetJenkins() {
	_jsii_.InvokeVoid(
		o,
		"resetJenkins",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ResetJob() {
	_jsii_.InvokeVoid(
		o,
		"resetJob",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ResetNewRelic() {
	_jsii_.InvokeVoid(
		o,
		"resetNewRelic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ResetPrometheus() {
	_jsii_.InvokeVoid(
		o,
		"resetPrometheus",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ResetWeb() {
	_jsii_.InvokeVoid(
		o,
		"resetWeb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceancdVerificationTemplateMetricsProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

