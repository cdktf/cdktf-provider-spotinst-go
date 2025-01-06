// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanrightsizingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_right_sizing_rule spotinst_ocean_right_sizing_rule}.
type OceanRightSizingRule interface {
	cdktf.TerraformResource
	AttachWorkloads() OceanRightSizingRuleAttachWorkloadsList
	AttachWorkloadsInput() interface{}
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
	DetachWorkloads() OceanRightSizingRuleDetachWorkloadsList
	DetachWorkloadsInput() interface{}
	ExcludePreliminaryRecommendations() interface{}
	SetExcludePreliminaryRecommendations(val interface{})
	ExcludePreliminaryRecommendationsInput() interface{}
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
	OceanId() *string
	SetOceanId(val *string)
	OceanIdInput() *string
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
	RecommendationApplicationBoundaries() OceanRightSizingRuleRecommendationApplicationBoundariesList
	RecommendationApplicationBoundariesInput() interface{}
	RecommendationApplicationHpa() OceanRightSizingRuleRecommendationApplicationHpaList
	RecommendationApplicationHpaInput() interface{}
	RecommendationApplicationIntervals() OceanRightSizingRuleRecommendationApplicationIntervalsList
	RecommendationApplicationIntervalsInput() interface{}
	RecommendationApplicationMinThreshold() OceanRightSizingRuleRecommendationApplicationMinThresholdList
	RecommendationApplicationMinThresholdInput() interface{}
	RecommendationApplicationOverheadValues() OceanRightSizingRuleRecommendationApplicationOverheadValuesList
	RecommendationApplicationOverheadValuesInput() interface{}
	RestartReplicas() *string
	SetRestartReplicas(val *string)
	RestartReplicasInput() *string
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
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
	PutAttachWorkloads(value interface{})
	PutDetachWorkloads(value interface{})
	PutRecommendationApplicationBoundaries(value interface{})
	PutRecommendationApplicationHpa(value interface{})
	PutRecommendationApplicationIntervals(value interface{})
	PutRecommendationApplicationMinThreshold(value interface{})
	PutRecommendationApplicationOverheadValues(value interface{})
	ResetAttachWorkloads()
	ResetDetachWorkloads()
	ResetExcludePreliminaryRecommendations()
	ResetId()
	ResetOceanId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecommendationApplicationBoundaries()
	ResetRecommendationApplicationHpa()
	ResetRecommendationApplicationMinThreshold()
	ResetRecommendationApplicationOverheadValues()
	ResetRestartReplicas()
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

// The jsii proxy struct for OceanRightSizingRule
type jsiiProxy_OceanRightSizingRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanRightSizingRule) AttachWorkloads() OceanRightSizingRuleAttachWorkloadsList {
	var returns OceanRightSizingRuleAttachWorkloadsList
	_jsii_.Get(
		j,
		"attachWorkloads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) AttachWorkloadsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachWorkloadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) DetachWorkloads() OceanRightSizingRuleDetachWorkloadsList {
	var returns OceanRightSizingRuleDetachWorkloadsList
	_jsii_.Get(
		j,
		"detachWorkloads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) DetachWorkloadsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detachWorkloadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) ExcludePreliminaryRecommendations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePreliminaryRecommendations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) ExcludePreliminaryRecommendationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePreliminaryRecommendationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) OceanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) OceanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationBoundaries() OceanRightSizingRuleRecommendationApplicationBoundariesList {
	var returns OceanRightSizingRuleRecommendationApplicationBoundariesList
	_jsii_.Get(
		j,
		"recommendationApplicationBoundaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationBoundariesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationApplicationBoundariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationHpa() OceanRightSizingRuleRecommendationApplicationHpaList {
	var returns OceanRightSizingRuleRecommendationApplicationHpaList
	_jsii_.Get(
		j,
		"recommendationApplicationHpa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationHpaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationApplicationHpaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationIntervals() OceanRightSizingRuleRecommendationApplicationIntervalsList {
	var returns OceanRightSizingRuleRecommendationApplicationIntervalsList
	_jsii_.Get(
		j,
		"recommendationApplicationIntervals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationIntervalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationApplicationIntervalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationMinThreshold() OceanRightSizingRuleRecommendationApplicationMinThresholdList {
	var returns OceanRightSizingRuleRecommendationApplicationMinThresholdList
	_jsii_.Get(
		j,
		"recommendationApplicationMinThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationMinThresholdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationApplicationMinThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationOverheadValues() OceanRightSizingRuleRecommendationApplicationOverheadValuesList {
	var returns OceanRightSizingRuleRecommendationApplicationOverheadValuesList
	_jsii_.Get(
		j,
		"recommendationApplicationOverheadValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RecommendationApplicationOverheadValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationApplicationOverheadValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RestartReplicas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RestartReplicasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_right_sizing_rule spotinst_ocean_right_sizing_rule} Resource.
func NewOceanRightSizingRule(scope constructs.Construct, id *string, config *OceanRightSizingRuleConfig) OceanRightSizingRule {
	_init_.Initialize()

	if err := validateNewOceanRightSizingRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanRightSizingRule{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_right_sizing_rule spotinst_ocean_right_sizing_rule} Resource.
func NewOceanRightSizingRule_Override(o OceanRightSizingRule, scope constructs.Construct, id *string, config *OceanRightSizingRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRule",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetExcludePreliminaryRecommendations(val interface{}) {
	if err := j.validateSetExcludePreliminaryRecommendationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePreliminaryRecommendations",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetOceanId(val *string) {
	if err := j.validateSetOceanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oceanId",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetRestartReplicas(val *string) {
	if err := j.validateSetRestartReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartReplicas",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRule)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

// Generates CDKTF code for importing a OceanRightSizingRule resource upon running "cdktf plan <stack-name>".
func OceanRightSizingRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceanRightSizingRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRule",
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
func OceanRightSizingRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanRightSizingRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanRightSizingRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanRightSizingRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanRightSizingRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanRightSizingRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanRightSizingRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanRightSizingRule) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanRightSizingRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanRightSizingRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanRightSizingRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanRightSizingRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanRightSizingRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanRightSizingRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanRightSizingRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanRightSizingRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanRightSizingRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanRightSizingRule) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) PutAttachWorkloads(value interface{}) {
	if err := o.validatePutAttachWorkloadsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAttachWorkloads",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) PutDetachWorkloads(value interface{}) {
	if err := o.validatePutDetachWorkloadsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDetachWorkloads",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) PutRecommendationApplicationBoundaries(value interface{}) {
	if err := o.validatePutRecommendationApplicationBoundariesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRecommendationApplicationBoundaries",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) PutRecommendationApplicationHpa(value interface{}) {
	if err := o.validatePutRecommendationApplicationHpaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRecommendationApplicationHpa",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) PutRecommendationApplicationIntervals(value interface{}) {
	if err := o.validatePutRecommendationApplicationIntervalsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRecommendationApplicationIntervals",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) PutRecommendationApplicationMinThreshold(value interface{}) {
	if err := o.validatePutRecommendationApplicationMinThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRecommendationApplicationMinThreshold",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) PutRecommendationApplicationOverheadValues(value interface{}) {
	if err := o.validatePutRecommendationApplicationOverheadValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRecommendationApplicationOverheadValues",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetAttachWorkloads() {
	_jsii_.InvokeVoid(
		o,
		"resetAttachWorkloads",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetDetachWorkloads() {
	_jsii_.InvokeVoid(
		o,
		"resetDetachWorkloads",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetExcludePreliminaryRecommendations() {
	_jsii_.InvokeVoid(
		o,
		"resetExcludePreliminaryRecommendations",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetOceanId() {
	_jsii_.InvokeVoid(
		o,
		"resetOceanId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetRecommendationApplicationBoundaries() {
	_jsii_.InvokeVoid(
		o,
		"resetRecommendationApplicationBoundaries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetRecommendationApplicationHpa() {
	_jsii_.InvokeVoid(
		o,
		"resetRecommendationApplicationHpa",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetRecommendationApplicationMinThreshold() {
	_jsii_.InvokeVoid(
		o,
		"resetRecommendationApplicationMinThreshold",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetRecommendationApplicationOverheadValues() {
	_jsii_.InvokeVoid(
		o,
		"resetRecommendationApplicationOverheadValues",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) ResetRestartReplicas() {
	_jsii_.InvokeVoid(
		o,
		"resetRestartReplicas",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

