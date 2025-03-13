// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceancdverificationprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.212.0/docs/resources/oceancd_verification_provider spotinst_oceancd_verification_provider}.
type OceancdVerificationProvider interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudWatch() OceancdVerificationProviderCloudWatchOutputReference
	CloudWatchInput() *OceancdVerificationProviderCloudWatch
	ClusterIds() *[]*string
	SetClusterIds(val *[]*string)
	ClusterIdsInput() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Datadog() OceancdVerificationProviderDatadogOutputReference
	DatadogInput() *OceancdVerificationProviderDatadog
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Jenkins() OceancdVerificationProviderJenkinsOutputReference
	JenkinsInput() *OceancdVerificationProviderJenkins
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewRelic() OceancdVerificationProviderNewRelicOutputReference
	NewRelicInput() *OceancdVerificationProviderNewRelic
	// The tree node.
	Node() constructs.Node
	Prometheus() OceancdVerificationProviderPrometheusOutputReference
	PrometheusInput() *OceancdVerificationProviderPrometheus
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCloudWatch(value *OceancdVerificationProviderCloudWatch)
	PutDatadog(value *OceancdVerificationProviderDatadog)
	PutJenkins(value *OceancdVerificationProviderJenkins)
	PutNewRelic(value *OceancdVerificationProviderNewRelic)
	PutPrometheus(value *OceancdVerificationProviderPrometheus)
	ResetCloudWatch()
	ResetDatadog()
	ResetId()
	ResetJenkins()
	ResetNewRelic()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrometheus()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OceancdVerificationProvider
type jsiiProxy_OceancdVerificationProvider struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceancdVerificationProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) CloudWatch() OceancdVerificationProviderCloudWatchOutputReference {
	var returns OceancdVerificationProviderCloudWatchOutputReference
	_jsii_.Get(
		j,
		"cloudWatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) CloudWatchInput() *OceancdVerificationProviderCloudWatch {
	var returns *OceancdVerificationProviderCloudWatch
	_jsii_.Get(
		j,
		"cloudWatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) ClusterIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) ClusterIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Datadog() OceancdVerificationProviderDatadogOutputReference {
	var returns OceancdVerificationProviderDatadogOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) DatadogInput() *OceancdVerificationProviderDatadog {
	var returns *OceancdVerificationProviderDatadog
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Jenkins() OceancdVerificationProviderJenkinsOutputReference {
	var returns OceancdVerificationProviderJenkinsOutputReference
	_jsii_.Get(
		j,
		"jenkins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) JenkinsInput() *OceancdVerificationProviderJenkins {
	var returns *OceancdVerificationProviderJenkins
	_jsii_.Get(
		j,
		"jenkinsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) NewRelic() OceancdVerificationProviderNewRelicOutputReference {
	var returns OceancdVerificationProviderNewRelicOutputReference
	_jsii_.Get(
		j,
		"newRelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) NewRelicInput() *OceancdVerificationProviderNewRelic {
	var returns *OceancdVerificationProviderNewRelic
	_jsii_.Get(
		j,
		"newRelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Prometheus() OceancdVerificationProviderPrometheusOutputReference {
	var returns OceancdVerificationProviderPrometheusOutputReference
	_jsii_.Get(
		j,
		"prometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) PrometheusInput() *OceancdVerificationProviderPrometheus {
	var returns *OceancdVerificationProviderPrometheus
	_jsii_.Get(
		j,
		"prometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdVerificationProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.212.0/docs/resources/oceancd_verification_provider spotinst_oceancd_verification_provider} Resource.
func NewOceancdVerificationProvider(scope constructs.Construct, id *string, config *OceancdVerificationProviderConfig) OceancdVerificationProvider {
	_init_.Initialize()

	if err := validateNewOceancdVerificationProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdVerificationProvider{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdVerificationProvider.OceancdVerificationProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.212.0/docs/resources/oceancd_verification_provider spotinst_oceancd_verification_provider} Resource.
func NewOceancdVerificationProvider_Override(o OceancdVerificationProvider, scope constructs.Construct, id *string, config *OceancdVerificationProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdVerificationProvider.OceancdVerificationProvider",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetClusterIds(val *[]*string) {
	if err := j.validateSetClusterIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIds",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceancdVerificationProvider)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a OceancdVerificationProvider resource upon running "cdktf plan <stack-name>".
func OceancdVerificationProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceancdVerificationProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceancdVerificationProvider.OceancdVerificationProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func OceancdVerificationProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceancdVerificationProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceancdVerificationProvider.OceancdVerificationProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceancdVerificationProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceancdVerificationProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceancdVerificationProvider.OceancdVerificationProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceancdVerificationProvider_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceancdVerificationProvider_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceancdVerificationProvider.OceancdVerificationProvider",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceancdVerificationProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceancdVerificationProvider.OceancdVerificationProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceancdVerificationProvider) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdVerificationProvider) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdVerificationProvider) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdVerificationProvider) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdVerificationProvider) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdVerificationProvider) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdVerificationProvider) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdVerificationProvider) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdVerificationProvider) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdVerificationProvider) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationProvider) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdVerificationProvider) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) PutCloudWatch(value *OceancdVerificationProviderCloudWatch) {
	if err := o.validatePutCloudWatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCloudWatch",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) PutDatadog(value *OceancdVerificationProviderDatadog) {
	if err := o.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatadog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) PutJenkins(value *OceancdVerificationProviderJenkins) {
	if err := o.validatePutJenkinsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putJenkins",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) PutNewRelic(value *OceancdVerificationProviderNewRelic) {
	if err := o.validatePutNewRelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNewRelic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) PutPrometheus(value *OceancdVerificationProviderPrometheus) {
	if err := o.validatePutPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPrometheus",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) ResetCloudWatch() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudWatch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) ResetDatadog() {
	_jsii_.InvokeVoid(
		o,
		"resetDatadog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) ResetJenkins() {
	_jsii_.InvokeVoid(
		o,
		"resetJenkins",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) ResetNewRelic() {
	_jsii_.InvokeVoid(
		o,
		"resetNewRelic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) ResetPrometheus() {
	_jsii_.InvokeVoid(
		o,
		"resetPrometheus",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdVerificationProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdVerificationProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

