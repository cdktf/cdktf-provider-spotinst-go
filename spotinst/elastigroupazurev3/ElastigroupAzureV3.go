// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/elastigroupazurev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/elastigroup_azure_v3 spotinst_elastigroup_azure_v3}.
type ElastigroupAzureV3 interface {
	cdktf.TerraformResource
	AvailabilityVsCost() *float64
	SetAvailabilityVsCost(val *float64)
	AvailabilityVsCostInput() *float64
	BootDiagnostics() ElastigroupAzureV3BootDiagnosticsList
	BootDiagnosticsInput() interface{}
	CapacityReservation() ElastigroupAzureV3CapacityReservationOutputReference
	CapacityReservationInput() *ElastigroupAzureV3CapacityReservation
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
	CustomData() *string
	SetCustomData(val *string)
	CustomDataInput() *string
	DataDisk() ElastigroupAzureV3DataDiskList
	DataDiskInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DesiredCapacity() *float64
	SetDesiredCapacity(val *float64)
	DesiredCapacityInput() *float64
	DrainingTimeout() *float64
	SetDrainingTimeout(val *float64)
	DrainingTimeoutInput() *float64
	Extensions() ElastigroupAzureV3ExtensionsList
	ExtensionsInput() interface{}
	FallbackToOnDemand() interface{}
	SetFallbackToOnDemand(val interface{})
	FallbackToOnDemandInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Health() ElastigroupAzureV3HealthOutputReference
	HealthInput() *ElastigroupAzureV3Health
	Id() *string
	SetId(val *string)
	IdInput() *string
	Image() ElastigroupAzureV3ImageList
	ImageInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() ElastigroupAzureV3LoadBalancerList
	LoadBalancerInput() interface{}
	Login() ElastigroupAzureV3LoginOutputReference
	LoginInput() *ElastigroupAzureV3Login
	ManagedServiceIdentity() ElastigroupAzureV3ManagedServiceIdentityList
	ManagedServiceIdentityInput() interface{}
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() ElastigroupAzureV3NetworkOutputReference
	NetworkInput() *ElastigroupAzureV3Network
	// The tree node.
	Node() constructs.Node
	OnDemandCount() *float64
	SetOnDemandCount(val *float64)
	OnDemandCountInput() *float64
	OptimizationWindows() *[]*string
	SetOptimizationWindows(val *[]*string)
	OptimizationWindowsInput() *[]*string
	Os() *string
	SetOs(val *string)
	OsDisk() ElastigroupAzureV3OsDiskOutputReference
	OsDiskInput() *ElastigroupAzureV3OsDisk
	OsInput() *string
	PreferredZones() *[]*string
	SetPreferredZones(val *[]*string)
	PreferredZonesInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProximityPlacementGroups() ElastigroupAzureV3ProximityPlacementGroupsList
	ProximityPlacementGroupsInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RevertToSpot() ElastigroupAzureV3RevertToSpotOutputReference
	RevertToSpotInput() *ElastigroupAzureV3RevertToSpot
	ScalingDownPolicy() ElastigroupAzureV3ScalingDownPolicyList
	ScalingDownPolicyInput() interface{}
	ScalingUpPolicy() ElastigroupAzureV3ScalingUpPolicyList
	ScalingUpPolicyInput() interface{}
	SchedulingTask() ElastigroupAzureV3SchedulingTaskList
	SchedulingTaskInput() interface{}
	Secret() ElastigroupAzureV3SecretList
	SecretInput() interface{}
	Security() ElastigroupAzureV3SecurityOutputReference
	SecurityInput() *ElastigroupAzureV3Security
	ShutdownScript() *string
	SetShutdownScript(val *string)
	ShutdownScriptInput() *string
	Signal() ElastigroupAzureV3SignalList
	SignalInput() interface{}
	SpotPercentage() *float64
	SetSpotPercentage(val *float64)
	SpotPercentageInput() *float64
	Tags() ElastigroupAzureV3TagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VmNamePrefix() *string
	SetVmNamePrefix(val *string)
	VmNamePrefixInput() *string
	VmSizes() ElastigroupAzureV3VmSizesOutputReference
	VmSizesInput() *ElastigroupAzureV3VmSizes
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	PutBootDiagnostics(value interface{})
	PutCapacityReservation(value *ElastigroupAzureV3CapacityReservation)
	PutDataDisk(value interface{})
	PutExtensions(value interface{})
	PutHealth(value *ElastigroupAzureV3Health)
	PutImage(value interface{})
	PutLoadBalancer(value interface{})
	PutLogin(value *ElastigroupAzureV3Login)
	PutManagedServiceIdentity(value interface{})
	PutNetwork(value *ElastigroupAzureV3Network)
	PutOsDisk(value *ElastigroupAzureV3OsDisk)
	PutProximityPlacementGroups(value interface{})
	PutRevertToSpot(value *ElastigroupAzureV3RevertToSpot)
	PutScalingDownPolicy(value interface{})
	PutScalingUpPolicy(value interface{})
	PutSchedulingTask(value interface{})
	PutSecret(value interface{})
	PutSecurity(value *ElastigroupAzureV3Security)
	PutSignal(value interface{})
	PutTags(value interface{})
	PutVmSizes(value *ElastigroupAzureV3VmSizes)
	ResetAvailabilityVsCost()
	ResetBootDiagnostics()
	ResetCapacityReservation()
	ResetCustomData()
	ResetDataDisk()
	ResetDescription()
	ResetDesiredCapacity()
	ResetDrainingTimeout()
	ResetExtensions()
	ResetHealth()
	ResetId()
	ResetImage()
	ResetLoadBalancer()
	ResetLogin()
	ResetManagedServiceIdentity()
	ResetMaxSize()
	ResetMinSize()
	ResetOnDemandCount()
	ResetOptimizationWindows()
	ResetOsDisk()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreferredZones()
	ResetProximityPlacementGroups()
	ResetRevertToSpot()
	ResetScalingDownPolicy()
	ResetScalingUpPolicy()
	ResetSchedulingTask()
	ResetSecret()
	ResetSecurity()
	ResetShutdownScript()
	ResetSignal()
	ResetSpotPercentage()
	ResetTags()
	ResetUserData()
	ResetVmNamePrefix()
	ResetZones()
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

// The jsii proxy struct for ElastigroupAzureV3
type jsiiProxy_ElastigroupAzureV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElastigroupAzureV3) AvailabilityVsCost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityVsCost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) AvailabilityVsCostInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityVsCostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) BootDiagnostics() ElastigroupAzureV3BootDiagnosticsList {
	var returns ElastigroupAzureV3BootDiagnosticsList
	_jsii_.Get(
		j,
		"bootDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) BootDiagnosticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) CapacityReservation() ElastigroupAzureV3CapacityReservationOutputReference {
	var returns ElastigroupAzureV3CapacityReservationOutputReference
	_jsii_.Get(
		j,
		"capacityReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) CapacityReservationInput() *ElastigroupAzureV3CapacityReservation {
	var returns *ElastigroupAzureV3CapacityReservation
	_jsii_.Get(
		j,
		"capacityReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) CustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) CustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) DataDisk() ElastigroupAzureV3DataDiskList {
	var returns ElastigroupAzureV3DataDiskList
	_jsii_.Get(
		j,
		"dataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) DataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) DrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) DrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Extensions() ElastigroupAzureV3ExtensionsList {
	var returns ElastigroupAzureV3ExtensionsList
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) FallbackToOnDemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOnDemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) FallbackToOnDemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOnDemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Health() ElastigroupAzureV3HealthOutputReference {
	var returns ElastigroupAzureV3HealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) HealthInput() *ElastigroupAzureV3Health {
	var returns *ElastigroupAzureV3Health
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Image() ElastigroupAzureV3ImageList {
	var returns ElastigroupAzureV3ImageList
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) LoadBalancer() ElastigroupAzureV3LoadBalancerList {
	var returns ElastigroupAzureV3LoadBalancerList
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) LoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Login() ElastigroupAzureV3LoginOutputReference {
	var returns ElastigroupAzureV3LoginOutputReference
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) LoginInput() *ElastigroupAzureV3Login {
	var returns *ElastigroupAzureV3Login
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ManagedServiceIdentity() ElastigroupAzureV3ManagedServiceIdentityList {
	var returns ElastigroupAzureV3ManagedServiceIdentityList
	_jsii_.Get(
		j,
		"managedServiceIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ManagedServiceIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedServiceIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Network() ElastigroupAzureV3NetworkOutputReference {
	var returns ElastigroupAzureV3NetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) NetworkInput() *ElastigroupAzureV3Network {
	var returns *ElastigroupAzureV3Network
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) OnDemandCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) OnDemandCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) OptimizationWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optimizationWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) OptimizationWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optimizationWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) OsDisk() ElastigroupAzureV3OsDiskOutputReference {
	var returns ElastigroupAzureV3OsDiskOutputReference
	_jsii_.Get(
		j,
		"osDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) OsDiskInput() *ElastigroupAzureV3OsDisk {
	var returns *ElastigroupAzureV3OsDisk
	_jsii_.Get(
		j,
		"osDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) OsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) PreferredZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) PreferredZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ProximityPlacementGroups() ElastigroupAzureV3ProximityPlacementGroupsList {
	var returns ElastigroupAzureV3ProximityPlacementGroupsList
	_jsii_.Get(
		j,
		"proximityPlacementGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ProximityPlacementGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proximityPlacementGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) RevertToSpot() ElastigroupAzureV3RevertToSpotOutputReference {
	var returns ElastigroupAzureV3RevertToSpotOutputReference
	_jsii_.Get(
		j,
		"revertToSpot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) RevertToSpotInput() *ElastigroupAzureV3RevertToSpot {
	var returns *ElastigroupAzureV3RevertToSpot
	_jsii_.Get(
		j,
		"revertToSpotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ScalingDownPolicy() ElastigroupAzureV3ScalingDownPolicyList {
	var returns ElastigroupAzureV3ScalingDownPolicyList
	_jsii_.Get(
		j,
		"scalingDownPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ScalingDownPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingDownPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ScalingUpPolicy() ElastigroupAzureV3ScalingUpPolicyList {
	var returns ElastigroupAzureV3ScalingUpPolicyList
	_jsii_.Get(
		j,
		"scalingUpPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ScalingUpPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingUpPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) SchedulingTask() ElastigroupAzureV3SchedulingTaskList {
	var returns ElastigroupAzureV3SchedulingTaskList
	_jsii_.Get(
		j,
		"schedulingTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) SchedulingTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedulingTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Secret() ElastigroupAzureV3SecretList {
	var returns ElastigroupAzureV3SecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Security() ElastigroupAzureV3SecurityOutputReference {
	var returns ElastigroupAzureV3SecurityOutputReference
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) SecurityInput() *ElastigroupAzureV3Security {
	var returns *ElastigroupAzureV3Security
	_jsii_.Get(
		j,
		"securityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ShutdownScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ShutdownScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Signal() ElastigroupAzureV3SignalList {
	var returns ElastigroupAzureV3SignalList
	_jsii_.Get(
		j,
		"signal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) SignalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) SpotPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) SpotPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Tags() ElastigroupAzureV3TagsList {
	var returns ElastigroupAzureV3TagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) VmNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) VmNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) VmSizes() ElastigroupAzureV3VmSizesOutputReference {
	var returns ElastigroupAzureV3VmSizesOutputReference
	_jsii_.Get(
		j,
		"vmSizes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) VmSizesInput() *ElastigroupAzureV3VmSizes {
	var returns *ElastigroupAzureV3VmSizes
	_jsii_.Get(
		j,
		"vmSizesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzureV3) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/elastigroup_azure_v3 spotinst_elastigroup_azure_v3} Resource.
func NewElastigroupAzureV3(scope constructs.Construct, id *string, config *ElastigroupAzureV3Config) ElastigroupAzureV3 {
	_init_.Initialize()

	if err := validateNewElastigroupAzureV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAzureV3{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/elastigroup_azure_v3 spotinst_elastigroup_azure_v3} Resource.
func NewElastigroupAzureV3_Override(e ElastigroupAzureV3, scope constructs.Construct, id *string, config *ElastigroupAzureV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetAvailabilityVsCost(val *float64) {
	if err := j.validateSetAvailabilityVsCostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityVsCost",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetCustomData(val *string) {
	if err := j.validateSetCustomDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customData",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetDrainingTimeout(val *float64) {
	if err := j.validateSetDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainingTimeout",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetFallbackToOnDemand(val interface{}) {
	if err := j.validateSetFallbackToOnDemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackToOnDemand",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetOnDemandCount(val *float64) {
	if err := j.validateSetOnDemandCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandCount",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetOptimizationWindows(val *[]*string) {
	if err := j.validateSetOptimizationWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optimizationWindows",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetOs(val *string) {
	if err := j.validateSetOsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetPreferredZones(val *[]*string) {
	if err := j.validateSetPreferredZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredZones",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetShutdownScript(val *string) {
	if err := j.validateSetShutdownScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownScript",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetSpotPercentage(val *float64) {
	if err := j.validateSetSpotPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetVmNamePrefix(val *string) {
	if err := j.validateSetVmNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmNamePrefix",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzureV3)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Generates CDKTF code for importing a ElastigroupAzureV3 resource upon running "cdktf plan <stack-name>".
func ElastigroupAzureV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElastigroupAzureV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3",
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
func ElastigroupAzureV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAzureV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupAzureV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAzureV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupAzureV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAzureV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElastigroupAzureV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.elastigroupAzureV3.ElastigroupAzureV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAzureV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzureV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAzureV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAzureV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAzureV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAzureV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAzureV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzureV3) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutBootDiagnostics(value interface{}) {
	if err := e.validatePutBootDiagnosticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBootDiagnostics",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutCapacityReservation(value *ElastigroupAzureV3CapacityReservation) {
	if err := e.validatePutCapacityReservationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCapacityReservation",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutDataDisk(value interface{}) {
	if err := e.validatePutDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDataDisk",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutExtensions(value interface{}) {
	if err := e.validatePutExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putExtensions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutHealth(value *ElastigroupAzureV3Health) {
	if err := e.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHealth",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutImage(value interface{}) {
	if err := e.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putImage",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutLoadBalancer(value interface{}) {
	if err := e.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutLogin(value *ElastigroupAzureV3Login) {
	if err := e.validatePutLoginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogin",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutManagedServiceIdentity(value interface{}) {
	if err := e.validatePutManagedServiceIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedServiceIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutNetwork(value *ElastigroupAzureV3Network) {
	if err := e.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetwork",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutOsDisk(value *ElastigroupAzureV3OsDisk) {
	if err := e.validatePutOsDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOsDisk",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutProximityPlacementGroups(value interface{}) {
	if err := e.validatePutProximityPlacementGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProximityPlacementGroups",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutRevertToSpot(value *ElastigroupAzureV3RevertToSpot) {
	if err := e.validatePutRevertToSpotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRevertToSpot",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutScalingDownPolicy(value interface{}) {
	if err := e.validatePutScalingDownPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingDownPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutScalingUpPolicy(value interface{}) {
	if err := e.validatePutScalingUpPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingUpPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutSchedulingTask(value interface{}) {
	if err := e.validatePutSchedulingTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSchedulingTask",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutSecret(value interface{}) {
	if err := e.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSecret",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutSecurity(value *ElastigroupAzureV3Security) {
	if err := e.validatePutSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSecurity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutSignal(value interface{}) {
	if err := e.validatePutSignalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSignal",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) PutVmSizes(value *ElastigroupAzureV3VmSizes) {
	if err := e.validatePutVmSizesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVmSizes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetAvailabilityVsCost() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityVsCost",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetBootDiagnostics() {
	_jsii_.InvokeVoid(
		e,
		"resetBootDiagnostics",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetCapacityReservation() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityReservation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetCustomData() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomData",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetDataDisk() {
	_jsii_.InvokeVoid(
		e,
		"resetDataDisk",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetDrainingTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetDrainingTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetExtensions() {
	_jsii_.InvokeVoid(
		e,
		"resetExtensions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetHealth() {
	_jsii_.InvokeVoid(
		e,
		"resetHealth",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetImage() {
	_jsii_.InvokeVoid(
		e,
		"resetImage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		e,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetLogin() {
	_jsii_.InvokeVoid(
		e,
		"resetLogin",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetManagedServiceIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedServiceIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetMaxSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetMinSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMinSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetOnDemandCount() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetOptimizationWindows() {
	_jsii_.InvokeVoid(
		e,
		"resetOptimizationWindows",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetOsDisk() {
	_jsii_.InvokeVoid(
		e,
		"resetOsDisk",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetPreferredZones() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredZones",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetProximityPlacementGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetProximityPlacementGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetRevertToSpot() {
	_jsii_.InvokeVoid(
		e,
		"resetRevertToSpot",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetScalingDownPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingDownPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetScalingUpPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingUpPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetSchedulingTask() {
	_jsii_.InvokeVoid(
		e,
		"resetSchedulingTask",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetSecret() {
	_jsii_.InvokeVoid(
		e,
		"resetSecret",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetSecurity() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetShutdownScript() {
	_jsii_.InvokeVoid(
		e,
		"resetShutdownScript",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetSignal() {
	_jsii_.InvokeVoid(
		e,
		"resetSignal",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetSpotPercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotPercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetUserData() {
	_jsii_.InvokeVoid(
		e,
		"resetUserData",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetVmNamePrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetVmNamePrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) ResetZones() {
	_jsii_.InvokeVoid(
		e,
		"resetZones",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzureV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzureV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

