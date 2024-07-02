// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceancdrolloutspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/oceancd_rollout_spec spotinst_oceancd_rollout_spec}.
type OceancdRolloutSpec interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	FailurePolicy() OceancdRolloutSpecFailurePolicyOutputReference
	FailurePolicyInput() *OceancdRolloutSpecFailurePolicy
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	RolloutSpecName() *string
	SetRolloutSpecName(val *string)
	RolloutSpecNameInput() *string
	SpotDeployment() OceancdRolloutSpecSpotDeploymentOutputReference
	SpotDeploymentInput() *OceancdRolloutSpecSpotDeployment
	SpotDeployments() OceancdRolloutSpecSpotDeploymentsList
	SpotDeploymentsInput() interface{}
	Strategy() OceancdRolloutSpecStrategyOutputReference
	StrategyInput() *OceancdRolloutSpecStrategy
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Traffic() OceancdRolloutSpecTrafficOutputReference
	TrafficInput() *OceancdRolloutSpecTraffic
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
	PutFailurePolicy(value *OceancdRolloutSpecFailurePolicy)
	PutSpotDeployment(value *OceancdRolloutSpecSpotDeployment)
	PutSpotDeployments(value interface{})
	PutStrategy(value *OceancdRolloutSpecStrategy)
	PutTraffic(value *OceancdRolloutSpecTraffic)
	ResetFailurePolicy()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSpotDeployment()
	ResetSpotDeployments()
	ResetTraffic()
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

// The jsii proxy struct for OceancdRolloutSpec
type jsiiProxy_OceancdRolloutSpec struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceancdRolloutSpec) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) FailurePolicy() OceancdRolloutSpecFailurePolicyOutputReference {
	var returns OceancdRolloutSpecFailurePolicyOutputReference
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) FailurePolicyInput() *OceancdRolloutSpecFailurePolicy {
	var returns *OceancdRolloutSpecFailurePolicy
	_jsii_.Get(
		j,
		"failurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) RolloutSpecName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolloutSpecName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) RolloutSpecNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolloutSpecNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) SpotDeployment() OceancdRolloutSpecSpotDeploymentOutputReference {
	var returns OceancdRolloutSpecSpotDeploymentOutputReference
	_jsii_.Get(
		j,
		"spotDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) SpotDeploymentInput() *OceancdRolloutSpecSpotDeployment {
	var returns *OceancdRolloutSpecSpotDeployment
	_jsii_.Get(
		j,
		"spotDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) SpotDeployments() OceancdRolloutSpecSpotDeploymentsList {
	var returns OceancdRolloutSpecSpotDeploymentsList
	_jsii_.Get(
		j,
		"spotDeployments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) SpotDeploymentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotDeploymentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Strategy() OceancdRolloutSpecStrategyOutputReference {
	var returns OceancdRolloutSpecStrategyOutputReference
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) StrategyInput() *OceancdRolloutSpecStrategy {
	var returns *OceancdRolloutSpecStrategy
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) Traffic() OceancdRolloutSpecTrafficOutputReference {
	var returns OceancdRolloutSpecTrafficOutputReference
	_jsii_.Get(
		j,
		"traffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdRolloutSpec) TrafficInput() *OceancdRolloutSpecTraffic {
	var returns *OceancdRolloutSpecTraffic
	_jsii_.Get(
		j,
		"trafficInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/oceancd_rollout_spec spotinst_oceancd_rollout_spec} Resource.
func NewOceancdRolloutSpec(scope constructs.Construct, id *string, config *OceancdRolloutSpecConfig) OceancdRolloutSpec {
	_init_.Initialize()

	if err := validateNewOceancdRolloutSpecParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdRolloutSpec{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpec",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/oceancd_rollout_spec spotinst_oceancd_rollout_spec} Resource.
func NewOceancdRolloutSpec_Override(o OceancdRolloutSpec, scope constructs.Construct, id *string, config *OceancdRolloutSpecConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpec",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceancdRolloutSpec)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpec)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpec)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpec)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpec)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpec)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpec)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpec)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceancdRolloutSpec)SetRolloutSpecName(val *string) {
	if err := j.validateSetRolloutSpecNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolloutSpecName",
		val,
	)
}

// Generates CDKTF code for importing a OceancdRolloutSpec resource upon running "cdktf plan <stack-name>".
func OceancdRolloutSpec_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceancdRolloutSpec_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpec",
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
func OceancdRolloutSpec_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceancdRolloutSpec_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpec",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceancdRolloutSpec_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceancdRolloutSpec_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpec",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceancdRolloutSpec_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceancdRolloutSpec_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpec",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceancdRolloutSpec_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceancdRolloutSpec.OceancdRolloutSpec",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceancdRolloutSpec) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdRolloutSpec) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpec) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdRolloutSpec) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdRolloutSpec) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdRolloutSpec) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpec) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdRolloutSpec) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdRolloutSpec) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdRolloutSpec) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpec) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdRolloutSpec) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) PutFailurePolicy(value *OceancdRolloutSpecFailurePolicy) {
	if err := o.validatePutFailurePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFailurePolicy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) PutSpotDeployment(value *OceancdRolloutSpecSpotDeployment) {
	if err := o.validatePutSpotDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSpotDeployment",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) PutSpotDeployments(value interface{}) {
	if err := o.validatePutSpotDeploymentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSpotDeployments",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) PutStrategy(value *OceancdRolloutSpecStrategy) {
	if err := o.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStrategy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) PutTraffic(value *OceancdRolloutSpecTraffic) {
	if err := o.validatePutTrafficParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTraffic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) ResetFailurePolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetFailurePolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) ResetSpotDeployment() {
	_jsii_.InvokeVoid(
		o,
		"resetSpotDeployment",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) ResetSpotDeployments() {
	_jsii_.InvokeVoid(
		o,
		"resetSpotDeployments",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) ResetTraffic() {
	_jsii_.InvokeVoid(
		o,
		"resetTraffic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdRolloutSpec) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpec) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpec) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpec) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpec) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdRolloutSpec) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

