// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupawsbeanstalk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/elastigroupawsbeanstalk/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.169.1/docs/resources/elastigroup_aws_beanstalk spotinst_elastigroup_aws_beanstalk}.
type ElastigroupAwsBeanstalk interface {
	cdktf.TerraformResource
	BeanstalkEnvironmentId() *string
	SetBeanstalkEnvironmentId(val *string)
	BeanstalkEnvironmentIdInput() *string
	BeanstalkEnvironmentName() *string
	SetBeanstalkEnvironmentName(val *string)
	BeanstalkEnvironmentNameInput() *string
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
	DeploymentPreferences() ElastigroupAwsBeanstalkDeploymentPreferencesOutputReference
	DeploymentPreferencesInput() *ElastigroupAwsBeanstalkDeploymentPreferences
	DesiredCapacity() *float64
	SetDesiredCapacity(val *float64)
	DesiredCapacityInput() *float64
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
	InstanceTypesSpot() *[]*string
	SetInstanceTypesSpot(val *[]*string)
	InstanceTypesSpotInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Maintenance() *string
	SetMaintenance(val *string)
	MaintenanceInput() *string
	ManagedActions() ElastigroupAwsBeanstalkManagedActionsOutputReference
	ManagedActionsInput() *ElastigroupAwsBeanstalkManagedActions
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Product() *string
	SetProduct(val *string)
	ProductInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ScheduledTask() ElastigroupAwsBeanstalkScheduledTaskList
	ScheduledTaskInput() interface{}
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
	PutDeploymentPreferences(value *ElastigroupAwsBeanstalkDeploymentPreferences)
	PutManagedActions(value *ElastigroupAwsBeanstalkManagedActions)
	PutScheduledTask(value interface{})
	ResetBeanstalkEnvironmentId()
	ResetBeanstalkEnvironmentName()
	ResetDeploymentPreferences()
	ResetId()
	ResetMaintenance()
	ResetManagedActions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScheduledTask()
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

// The jsii proxy struct for ElastigroupAwsBeanstalk
type jsiiProxy_ElastigroupAwsBeanstalk struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) BeanstalkEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beanstalkEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) BeanstalkEnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beanstalkEnvironmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) BeanstalkEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beanstalkEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) BeanstalkEnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beanstalkEnvironmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) DeploymentPreferences() ElastigroupAwsBeanstalkDeploymentPreferencesOutputReference {
	var returns ElastigroupAwsBeanstalkDeploymentPreferencesOutputReference
	_jsii_.Get(
		j,
		"deploymentPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) DeploymentPreferencesInput() *ElastigroupAwsBeanstalkDeploymentPreferences {
	var returns *ElastigroupAwsBeanstalkDeploymentPreferences
	_jsii_.Get(
		j,
		"deploymentPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) InstanceTypesSpot() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesSpot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) InstanceTypesSpotInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesSpotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Maintenance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) MaintenanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) ManagedActions() ElastigroupAwsBeanstalkManagedActionsOutputReference {
	var returns ElastigroupAwsBeanstalkManagedActionsOutputReference
	_jsii_.Get(
		j,
		"managedActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) ManagedActionsInput() *ElastigroupAwsBeanstalkManagedActions {
	var returns *ElastigroupAwsBeanstalkManagedActions
	_jsii_.Get(
		j,
		"managedActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Product() *string {
	var returns *string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) ProductInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) ScheduledTask() ElastigroupAwsBeanstalkScheduledTaskList {
	var returns ElastigroupAwsBeanstalkScheduledTaskList
	_jsii_.Get(
		j,
		"scheduledTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) ScheduledTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.169.1/docs/resources/elastigroup_aws_beanstalk spotinst_elastigroup_aws_beanstalk} Resource.
func NewElastigroupAwsBeanstalk(scope constructs.Construct, id *string, config *ElastigroupAwsBeanstalkConfig) ElastigroupAwsBeanstalk {
	_init_.Initialize()

	if err := validateNewElastigroupAwsBeanstalkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsBeanstalk{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAwsBeanstalk.ElastigroupAwsBeanstalk",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.169.1/docs/resources/elastigroup_aws_beanstalk spotinst_elastigroup_aws_beanstalk} Resource.
func NewElastigroupAwsBeanstalk_Override(e ElastigroupAwsBeanstalk, scope constructs.Construct, id *string, config *ElastigroupAwsBeanstalkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAwsBeanstalk.ElastigroupAwsBeanstalk",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetBeanstalkEnvironmentId(val *string) {
	if err := j.validateSetBeanstalkEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beanstalkEnvironmentId",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetBeanstalkEnvironmentName(val *string) {
	if err := j.validateSetBeanstalkEnvironmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beanstalkEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetInstanceTypesSpot(val *[]*string) {
	if err := j.validateSetInstanceTypesSpotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypesSpot",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetMaintenance(val *string) {
	if err := j.validateSetMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenance",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetProduct(val *string) {
	if err := j.validateSetProductParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"product",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsBeanstalk)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a ElastigroupAwsBeanstalk resource upon running "cdktf plan <stack-name>".
func ElastigroupAwsBeanstalk_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElastigroupAwsBeanstalk_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAwsBeanstalk.ElastigroupAwsBeanstalk",
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
func ElastigroupAwsBeanstalk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAwsBeanstalk_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAwsBeanstalk.ElastigroupAwsBeanstalk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupAwsBeanstalk_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAwsBeanstalk_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAwsBeanstalk.ElastigroupAwsBeanstalk",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupAwsBeanstalk_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAwsBeanstalk_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAwsBeanstalk.ElastigroupAwsBeanstalk",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElastigroupAwsBeanstalk_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.elastigroupAwsBeanstalk.ElastigroupAwsBeanstalk",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) PutDeploymentPreferences(value *ElastigroupAwsBeanstalkDeploymentPreferences) {
	if err := e.validatePutDeploymentPreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeploymentPreferences",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) PutManagedActions(value *ElastigroupAwsBeanstalkManagedActions) {
	if err := e.validatePutManagedActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedActions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) PutScheduledTask(value interface{}) {
	if err := e.validatePutScheduledTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScheduledTask",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ResetBeanstalkEnvironmentId() {
	_jsii_.InvokeVoid(
		e,
		"resetBeanstalkEnvironmentId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ResetBeanstalkEnvironmentName() {
	_jsii_.InvokeVoid(
		e,
		"resetBeanstalkEnvironmentName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ResetDeploymentPreferences() {
	_jsii_.InvokeVoid(
		e,
		"resetDeploymentPreferences",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ResetMaintenance() {
	_jsii_.InvokeVoid(
		e,
		"resetMaintenance",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ResetManagedActions() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedActions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ResetScheduledTask() {
	_jsii_.InvokeVoid(
		e,
		"resetScheduledTask",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsBeanstalk) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

