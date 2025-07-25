// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanspark/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/ocean_spark spotinst_ocean_spark}.
type OceanSpark interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Compute() OceanSparkComputeOutputReference
	ComputeInput() *OceanSparkCompute
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
	Ingress() OceanSparkIngressOutputReference
	IngressInput() *OceanSparkIngress
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogCollection() OceanSparkLogCollectionOutputReference
	LogCollectionInput() *OceanSparkLogCollection
	// The tree node.
	Node() constructs.Node
	OceanClusterId() *string
	SetOceanClusterId(val *string)
	OceanClusterIdInput() *string
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
	Spark() OceanSparkSparkOutputReference
	SparkInput() *OceanSparkSpark
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Webhook() OceanSparkWebhookOutputReference
	WebhookInput() *OceanSparkWebhook
	Workspaces() OceanSparkWorkspacesOutputReference
	WorkspacesInput() *OceanSparkWorkspaces
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
	PutCompute(value *OceanSparkCompute)
	PutIngress(value *OceanSparkIngress)
	PutLogCollection(value *OceanSparkLogCollection)
	PutSpark(value *OceanSparkSpark)
	PutWebhook(value *OceanSparkWebhook)
	PutWorkspaces(value *OceanSparkWorkspaces)
	ResetCompute()
	ResetId()
	ResetIngress()
	ResetLogCollection()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSpark()
	ResetWebhook()
	ResetWorkspaces()
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

// The jsii proxy struct for OceanSpark
type jsiiProxy_OceanSpark struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanSpark) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Compute() OceanSparkComputeOutputReference {
	var returns OceanSparkComputeOutputReference
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) ComputeInput() *OceanSparkCompute {
	var returns *OceanSparkCompute
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Ingress() OceanSparkIngressOutputReference {
	var returns OceanSparkIngressOutputReference
	_jsii_.Get(
		j,
		"ingress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) IngressInput() *OceanSparkIngress {
	var returns *OceanSparkIngress
	_jsii_.Get(
		j,
		"ingressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) LogCollection() OceanSparkLogCollectionOutputReference {
	var returns OceanSparkLogCollectionOutputReference
	_jsii_.Get(
		j,
		"logCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) LogCollectionInput() *OceanSparkLogCollection {
	var returns *OceanSparkLogCollection
	_jsii_.Get(
		j,
		"logCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) OceanClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) OceanClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Spark() OceanSparkSparkOutputReference {
	var returns OceanSparkSparkOutputReference
	_jsii_.Get(
		j,
		"spark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) SparkInput() *OceanSparkSpark {
	var returns *OceanSparkSpark
	_jsii_.Get(
		j,
		"sparkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Webhook() OceanSparkWebhookOutputReference {
	var returns OceanSparkWebhookOutputReference
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) WebhookInput() *OceanSparkWebhook {
	var returns *OceanSparkWebhook
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) Workspaces() OceanSparkWorkspacesOutputReference {
	var returns OceanSparkWorkspacesOutputReference
	_jsii_.Get(
		j,
		"workspaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanSpark) WorkspacesInput() *OceanSparkWorkspaces {
	var returns *OceanSparkWorkspaces
	_jsii_.Get(
		j,
		"workspacesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/ocean_spark spotinst_ocean_spark} Resource.
func NewOceanSpark(scope constructs.Construct, id *string, config *OceanSparkConfig) OceanSpark {
	_init_.Initialize()

	if err := validateNewOceanSparkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanSpark{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanSpark.OceanSpark",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/ocean_spark spotinst_ocean_spark} Resource.
func NewOceanSpark_Override(o OceanSpark, scope constructs.Construct, id *string, config *OceanSparkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanSpark.OceanSpark",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanSpark)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanSpark)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanSpark)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanSpark)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanSpark)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanSpark)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanSpark)SetOceanClusterId(val *string) {
	if err := j.validateSetOceanClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oceanClusterId",
		val,
	)
}

func (j *jsiiProxy_OceanSpark)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanSpark)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a OceanSpark resource upon running "cdktf plan <stack-name>".
func OceanSpark_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceanSpark_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanSpark.OceanSpark",
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
func OceanSpark_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanSpark_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanSpark.OceanSpark",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanSpark_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanSpark_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanSpark.OceanSpark",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanSpark_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanSpark_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanSpark.OceanSpark",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanSpark_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanSpark.OceanSpark",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanSpark) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceanSpark) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanSpark) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanSpark) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanSpark) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanSpark) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanSpark) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanSpark) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanSpark) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanSpark) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanSpark) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanSpark) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanSpark) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceanSpark) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanSpark) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanSpark) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceanSpark) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanSpark) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanSpark) PutCompute(value *OceanSparkCompute) {
	if err := o.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCompute",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSpark) PutIngress(value *OceanSparkIngress) {
	if err := o.validatePutIngressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putIngress",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSpark) PutLogCollection(value *OceanSparkLogCollection) {
	if err := o.validatePutLogCollectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogCollection",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSpark) PutSpark(value *OceanSparkSpark) {
	if err := o.validatePutSparkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSpark",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSpark) PutWebhook(value *OceanSparkWebhook) {
	if err := o.validatePutWebhookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWebhook",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSpark) PutWorkspaces(value *OceanSparkWorkspaces) {
	if err := o.validatePutWorkspacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWorkspaces",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanSpark) ResetCompute() {
	_jsii_.InvokeVoid(
		o,
		"resetCompute",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSpark) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSpark) ResetIngress() {
	_jsii_.InvokeVoid(
		o,
		"resetIngress",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSpark) ResetLogCollection() {
	_jsii_.InvokeVoid(
		o,
		"resetLogCollection",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSpark) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSpark) ResetSpark() {
	_jsii_.InvokeVoid(
		o,
		"resetSpark",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSpark) ResetWebhook() {
	_jsii_.InvokeVoid(
		o,
		"resetWebhook",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSpark) ResetWorkspaces() {
	_jsii_.InvokeVoid(
		o,
		"resetWorkspaces",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanSpark) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanSpark) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanSpark) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanSpark) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanSpark) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanSpark) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

