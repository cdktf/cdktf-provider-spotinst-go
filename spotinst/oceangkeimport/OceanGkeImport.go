package oceangkeimport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v7/oceangkeimport/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/ocean_gke_import spotinst_ocean_gke_import}.
type OceanGkeImport interface {
	cdktf.TerraformResource
	Autoscaler() OceanGkeImportAutoscalerOutputReference
	AutoscalerInput() *OceanGkeImportAutoscaler
	BackendServices() OceanGkeImportBackendServicesList
	BackendServicesInput() interface{}
	Blacklist() *[]*string
	SetBlacklist(val *[]*string)
	BlacklistInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterControllerId() *string
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControllerClusterId() *string
	SetControllerClusterId(val *string)
	ControllerClusterIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
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
	RootVolumeType() *string
	SetRootVolumeType(val *string)
	RootVolumeTypeInput() *string
	ScheduledTask() OceanGkeImportScheduledTaskList
	ScheduledTaskInput() interface{}
	ShieldedInstanceConfig() OceanGkeImportShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *OceanGkeImportShieldedInstanceConfig
	Strategy() OceanGkeImportStrategyList
	StrategyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatePolicy() OceanGkeImportUpdatePolicyOutputReference
	UpdatePolicyInput() *OceanGkeImportUpdatePolicy
	UseAsTemplateOnly() interface{}
	SetUseAsTemplateOnly(val interface{})
	UseAsTemplateOnlyInput() interface{}
	Whitelist() *[]*string
	SetWhitelist(val *[]*string)
	WhitelistInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAutoscaler(value *OceanGkeImportAutoscaler)
	PutBackendServices(value interface{})
	PutScheduledTask(value interface{})
	PutShieldedInstanceConfig(value *OceanGkeImportShieldedInstanceConfig)
	PutStrategy(value interface{})
	PutUpdatePolicy(value *OceanGkeImportUpdatePolicy)
	ResetAutoscaler()
	ResetBackendServices()
	ResetBlacklist()
	ResetControllerClusterId()
	ResetDesiredCapacity()
	ResetId()
	ResetMaxSize()
	ResetMinSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRootVolumeType()
	ResetScheduledTask()
	ResetShieldedInstanceConfig()
	ResetStrategy()
	ResetUpdatePolicy()
	ResetUseAsTemplateOnly()
	ResetWhitelist()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OceanGkeImport
type jsiiProxy_OceanGkeImport struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanGkeImport) Autoscaler() OceanGkeImportAutoscalerOutputReference {
	var returns OceanGkeImportAutoscalerOutputReference
	_jsii_.Get(
		j,
		"autoscaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) AutoscalerInput() *OceanGkeImportAutoscaler {
	var returns *OceanGkeImportAutoscaler
	_jsii_.Get(
		j,
		"autoscalerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) BackendServices() OceanGkeImportBackendServicesList {
	var returns OceanGkeImportBackendServicesList
	_jsii_.Get(
		j,
		"backendServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) BackendServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Blacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) BlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ClusterControllerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterControllerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ControllerClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ControllerClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) RootVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) RootVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ScheduledTask() OceanGkeImportScheduledTaskList {
	var returns OceanGkeImportScheduledTaskList
	_jsii_.Get(
		j,
		"scheduledTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ScheduledTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ShieldedInstanceConfig() OceanGkeImportShieldedInstanceConfigOutputReference {
	var returns OceanGkeImportShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) ShieldedInstanceConfigInput() *OceanGkeImportShieldedInstanceConfig {
	var returns *OceanGkeImportShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Strategy() OceanGkeImportStrategyList {
	var returns OceanGkeImportStrategyList
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) StrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) UpdatePolicy() OceanGkeImportUpdatePolicyOutputReference {
	var returns OceanGkeImportUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) UpdatePolicyInput() *OceanGkeImportUpdatePolicy {
	var returns *OceanGkeImportUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) UseAsTemplateOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAsTemplateOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) UseAsTemplateOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAsTemplateOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) Whitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImport) WhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/ocean_gke_import spotinst_ocean_gke_import} Resource.
func NewOceanGkeImport(scope constructs.Construct, id *string, config *OceanGkeImportConfig) OceanGkeImport {
	_init_.Initialize()

	if err := validateNewOceanGkeImportParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanGkeImport{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImport",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/ocean_gke_import spotinst_ocean_gke_import} Resource.
func NewOceanGkeImport_Override(o OceanGkeImport, scope constructs.Construct, id *string, config *OceanGkeImportConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImport",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetBlacklist(val *[]*string) {
	if err := j.validateSetBlacklistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blacklist",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetControllerClusterId(val *string) {
	if err := j.validateSetControllerClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controllerClusterId",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetRootVolumeType(val *string) {
	if err := j.validateSetRootVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootVolumeType",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetUseAsTemplateOnly(val interface{}) {
	if err := j.validateSetUseAsTemplateOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAsTemplateOnly",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImport)SetWhitelist(val *[]*string) {
	if err := j.validateSetWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelist",
		val,
	)
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
func OceanGkeImport_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanGkeImport_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImport",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanGkeImport_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanGkeImport_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImport",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanGkeImport_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanGkeImport_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImport",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanGkeImport_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImport",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanGkeImport) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanGkeImport) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanGkeImport) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanGkeImport) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanGkeImport) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanGkeImport) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanGkeImport) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanGkeImport) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanGkeImport) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanGkeImport) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanGkeImport) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanGkeImport) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanGkeImport) PutAutoscaler(value *OceanGkeImportAutoscaler) {
	if err := o.validatePutAutoscalerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaler",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeImport) PutBackendServices(value interface{}) {
	if err := o.validatePutBackendServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBackendServices",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeImport) PutScheduledTask(value interface{}) {
	if err := o.validatePutScheduledTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putScheduledTask",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeImport) PutShieldedInstanceConfig(value *OceanGkeImportShieldedInstanceConfig) {
	if err := o.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeImport) PutStrategy(value interface{}) {
	if err := o.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStrategy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeImport) PutUpdatePolicy(value *OceanGkeImportUpdatePolicy) {
	if err := o.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetAutoscaler() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaler",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetBackendServices() {
	_jsii_.InvokeVoid(
		o,
		"resetBackendServices",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetBlacklist() {
	_jsii_.InvokeVoid(
		o,
		"resetBlacklist",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetControllerClusterId() {
	_jsii_.InvokeVoid(
		o,
		"resetControllerClusterId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		o,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetMaxSize() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetMinSize() {
	_jsii_.InvokeVoid(
		o,
		"resetMinSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetRootVolumeType() {
	_jsii_.InvokeVoid(
		o,
		"resetRootVolumeType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetScheduledTask() {
	_jsii_.InvokeVoid(
		o,
		"resetScheduledTask",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetStrategy() {
	_jsii_.InvokeVoid(
		o,
		"resetStrategy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetUseAsTemplateOnly() {
	_jsii_.InvokeVoid(
		o,
		"resetUseAsTemplateOnly",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) ResetWhitelist() {
	_jsii_.InvokeVoid(
		o,
		"resetWhitelist",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImport) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeImport) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeImport) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeImport) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

