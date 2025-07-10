// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/statefulnodeazure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/stateful_node_azure spotinst_stateful_node_azure}.
type StatefulNodeAzure interface {
	cdktf.TerraformResource
	AttachDataDisk() StatefulNodeAzureAttachDataDiskList
	AttachDataDiskInput() interface{}
	BootDiagnostics() StatefulNodeAzureBootDiagnosticsList
	BootDiagnosticsInput() interface{}
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
	DataDisk() StatefulNodeAzureDataDiskList
	DataDiskInput() interface{}
	DataDisksPersistenceMode() *string
	SetDataDisksPersistenceMode(val *string)
	DataDisksPersistenceModeInput() *string
	Delete() StatefulNodeAzureDeleteList
	DeleteInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DetachDataDisk() StatefulNodeAzureDetachDataDiskList
	DetachDataDiskInput() interface{}
	Extension() StatefulNodeAzureExtensionList
	ExtensionInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Health() StatefulNodeAzureHealthOutputReference
	HealthInput() *StatefulNodeAzureHealth
	Id() *string
	SetId(val *string)
	IdInput() *string
	Image() StatefulNodeAzureImageOutputReference
	ImageInput() *StatefulNodeAzureImage
	ImportVm() StatefulNodeAzureImportVmList
	ImportVmInput() interface{}
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() StatefulNodeAzureLoadBalancerList
	LoadBalancerInput() interface{}
	Login() StatefulNodeAzureLoginOutputReference
	LoginInput() *StatefulNodeAzureLogin
	ManagedServiceIdentities() StatefulNodeAzureManagedServiceIdentitiesList
	ManagedServiceIdentitiesInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() StatefulNodeAzureNetworkOutputReference
	NetworkInput() *StatefulNodeAzureNetwork
	// The tree node.
	Node() constructs.Node
	Os() *string
	SetOs(val *string)
	OsDisk() StatefulNodeAzureOsDiskOutputReference
	OsDiskInput() *StatefulNodeAzureOsDisk
	OsDiskPersistenceMode() *string
	SetOsDiskPersistenceMode(val *string)
	OsDiskPersistenceModeInput() *string
	OsInput() *string
	PreferredZone() *string
	SetPreferredZone(val *string)
	PreferredZoneInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProximityPlacementGroups() StatefulNodeAzureProximityPlacementGroupsList
	ProximityPlacementGroupsInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SchedulingTask() StatefulNodeAzureSchedulingTaskList
	SchedulingTaskInput() interface{}
	Secret() StatefulNodeAzureSecretList
	SecretInput() interface{}
	Security() StatefulNodeAzureSecurityOutputReference
	SecurityInput() *StatefulNodeAzureSecurity
	ShouldPersistDataDisks() interface{}
	SetShouldPersistDataDisks(val interface{})
	ShouldPersistDataDisksInput() interface{}
	ShouldPersistNetwork() interface{}
	SetShouldPersistNetwork(val interface{})
	ShouldPersistNetworkInput() interface{}
	ShouldPersistOsDisk() interface{}
	SetShouldPersistOsDisk(val interface{})
	ShouldPersistOsDiskInput() interface{}
	ShutdownScript() *string
	SetShutdownScript(val *string)
	ShutdownScriptInput() *string
	Signal() StatefulNodeAzureSignalList
	SignalInput() interface{}
	Strategy() StatefulNodeAzureStrategyOutputReference
	StrategyInput() *StatefulNodeAzureStrategy
	Tag() StatefulNodeAzureTagList
	TagInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdateState() StatefulNodeAzureUpdateStateList
	UpdateStateInput() interface{}
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VmName() *string
	SetVmName(val *string)
	VmNameInput() *string
	VmNamePrefix() *string
	SetVmNamePrefix(val *string)
	VmNamePrefixInput() *string
	VmSizes() StatefulNodeAzureVmSizesOutputReference
	VmSizesInput() *StatefulNodeAzureVmSizes
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
	PutAttachDataDisk(value interface{})
	PutBootDiagnostics(value interface{})
	PutDataDisk(value interface{})
	PutDelete(value interface{})
	PutDetachDataDisk(value interface{})
	PutExtension(value interface{})
	PutHealth(value *StatefulNodeAzureHealth)
	PutImage(value *StatefulNodeAzureImage)
	PutImportVm(value interface{})
	PutLoadBalancer(value interface{})
	PutLogin(value *StatefulNodeAzureLogin)
	PutManagedServiceIdentities(value interface{})
	PutNetwork(value *StatefulNodeAzureNetwork)
	PutOsDisk(value *StatefulNodeAzureOsDisk)
	PutProximityPlacementGroups(value interface{})
	PutSchedulingTask(value interface{})
	PutSecret(value interface{})
	PutSecurity(value *StatefulNodeAzureSecurity)
	PutSignal(value interface{})
	PutStrategy(value *StatefulNodeAzureStrategy)
	PutTag(value interface{})
	PutUpdateState(value interface{})
	PutVmSizes(value *StatefulNodeAzureVmSizes)
	ResetAttachDataDisk()
	ResetBootDiagnostics()
	ResetCustomData()
	ResetDataDisk()
	ResetDataDisksPersistenceMode()
	ResetDelete()
	ResetDescription()
	ResetDetachDataDisk()
	ResetExtension()
	ResetHealth()
	ResetId()
	ResetImage()
	ResetImportVm()
	ResetLicenseType()
	ResetLoadBalancer()
	ResetLogin()
	ResetManagedServiceIdentities()
	ResetNetwork()
	ResetOsDisk()
	ResetOsDiskPersistenceMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreferredZone()
	ResetProximityPlacementGroups()
	ResetSchedulingTask()
	ResetSecret()
	ResetSecurity()
	ResetShutdownScript()
	ResetSignal()
	ResetTag()
	ResetUpdateState()
	ResetUserData()
	ResetVmName()
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

// The jsii proxy struct for StatefulNodeAzure
type jsiiProxy_StatefulNodeAzure struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StatefulNodeAzure) AttachDataDisk() StatefulNodeAzureAttachDataDiskList {
	var returns StatefulNodeAzureAttachDataDiskList
	_jsii_.Get(
		j,
		"attachDataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) AttachDataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachDataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) BootDiagnostics() StatefulNodeAzureBootDiagnosticsList {
	var returns StatefulNodeAzureBootDiagnosticsList
	_jsii_.Get(
		j,
		"bootDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) BootDiagnosticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) CustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) CustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) DataDisk() StatefulNodeAzureDataDiskList {
	var returns StatefulNodeAzureDataDiskList
	_jsii_.Get(
		j,
		"dataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) DataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) DataDisksPersistenceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDisksPersistenceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) DataDisksPersistenceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDisksPersistenceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Delete() StatefulNodeAzureDeleteList {
	var returns StatefulNodeAzureDeleteList
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) DeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) DetachDataDisk() StatefulNodeAzureDetachDataDiskList {
	var returns StatefulNodeAzureDetachDataDiskList
	_jsii_.Get(
		j,
		"detachDataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) DetachDataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detachDataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Extension() StatefulNodeAzureExtensionList {
	var returns StatefulNodeAzureExtensionList
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Health() StatefulNodeAzureHealthOutputReference {
	var returns StatefulNodeAzureHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) HealthInput() *StatefulNodeAzureHealth {
	var returns *StatefulNodeAzureHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Image() StatefulNodeAzureImageOutputReference {
	var returns StatefulNodeAzureImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ImageInput() *StatefulNodeAzureImage {
	var returns *StatefulNodeAzureImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ImportVm() StatefulNodeAzureImportVmList {
	var returns StatefulNodeAzureImportVmList
	_jsii_.Get(
		j,
		"importVm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ImportVmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importVmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) LoadBalancer() StatefulNodeAzureLoadBalancerList {
	var returns StatefulNodeAzureLoadBalancerList
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) LoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Login() StatefulNodeAzureLoginOutputReference {
	var returns StatefulNodeAzureLoginOutputReference
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) LoginInput() *StatefulNodeAzureLogin {
	var returns *StatefulNodeAzureLogin
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ManagedServiceIdentities() StatefulNodeAzureManagedServiceIdentitiesList {
	var returns StatefulNodeAzureManagedServiceIdentitiesList
	_jsii_.Get(
		j,
		"managedServiceIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ManagedServiceIdentitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedServiceIdentitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Network() StatefulNodeAzureNetworkOutputReference {
	var returns StatefulNodeAzureNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) NetworkInput() *StatefulNodeAzureNetwork {
	var returns *StatefulNodeAzureNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) OsDisk() StatefulNodeAzureOsDiskOutputReference {
	var returns StatefulNodeAzureOsDiskOutputReference
	_jsii_.Get(
		j,
		"osDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) OsDiskInput() *StatefulNodeAzureOsDisk {
	var returns *StatefulNodeAzureOsDisk
	_jsii_.Get(
		j,
		"osDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) OsDiskPersistenceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskPersistenceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) OsDiskPersistenceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskPersistenceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) OsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) PreferredZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) PreferredZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ProximityPlacementGroups() StatefulNodeAzureProximityPlacementGroupsList {
	var returns StatefulNodeAzureProximityPlacementGroupsList
	_jsii_.Get(
		j,
		"proximityPlacementGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ProximityPlacementGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proximityPlacementGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) SchedulingTask() StatefulNodeAzureSchedulingTaskList {
	var returns StatefulNodeAzureSchedulingTaskList
	_jsii_.Get(
		j,
		"schedulingTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) SchedulingTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedulingTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Secret() StatefulNodeAzureSecretList {
	var returns StatefulNodeAzureSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Security() StatefulNodeAzureSecurityOutputReference {
	var returns StatefulNodeAzureSecurityOutputReference
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) SecurityInput() *StatefulNodeAzureSecurity {
	var returns *StatefulNodeAzureSecurity
	_jsii_.Get(
		j,
		"securityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ShouldPersistDataDisks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldPersistDataDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ShouldPersistDataDisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldPersistDataDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ShouldPersistNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldPersistNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ShouldPersistNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldPersistNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ShouldPersistOsDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldPersistOsDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ShouldPersistOsDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldPersistOsDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ShutdownScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ShutdownScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Signal() StatefulNodeAzureSignalList {
	var returns StatefulNodeAzureSignalList
	_jsii_.Get(
		j,
		"signal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) SignalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Strategy() StatefulNodeAzureStrategyOutputReference {
	var returns StatefulNodeAzureStrategyOutputReference
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) StrategyInput() *StatefulNodeAzureStrategy {
	var returns *StatefulNodeAzureStrategy
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Tag() StatefulNodeAzureTagList {
	var returns StatefulNodeAzureTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) UpdateState() StatefulNodeAzureUpdateStateList {
	var returns StatefulNodeAzureUpdateStateList
	_jsii_.Get(
		j,
		"updateState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) UpdateStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) VmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) VmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) VmNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) VmNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) VmSizes() StatefulNodeAzureVmSizesOutputReference {
	var returns StatefulNodeAzureVmSizesOutputReference
	_jsii_.Get(
		j,
		"vmSizes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) VmSizesInput() *StatefulNodeAzureVmSizes {
	var returns *StatefulNodeAzureVmSizes
	_jsii_.Get(
		j,
		"vmSizesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzure) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/stateful_node_azure spotinst_stateful_node_azure} Resource.
func NewStatefulNodeAzure(scope constructs.Construct, id *string, config *StatefulNodeAzureConfig) StatefulNodeAzure {
	_init_.Initialize()

	if err := validateNewStatefulNodeAzureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulNodeAzure{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzure",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/stateful_node_azure spotinst_stateful_node_azure} Resource.
func NewStatefulNodeAzure_Override(s StatefulNodeAzure, scope constructs.Construct, id *string, config *StatefulNodeAzureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzure",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetCustomData(val *string) {
	if err := j.validateSetCustomDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customData",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetDataDisksPersistenceMode(val *string) {
	if err := j.validateSetDataDisksPersistenceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDisksPersistenceMode",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetOs(val *string) {
	if err := j.validateSetOsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetOsDiskPersistenceMode(val *string) {
	if err := j.validateSetOsDiskPersistenceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskPersistenceMode",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetPreferredZone(val *string) {
	if err := j.validateSetPreferredZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredZone",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetShouldPersistDataDisks(val interface{}) {
	if err := j.validateSetShouldPersistDataDisksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldPersistDataDisks",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetShouldPersistNetwork(val interface{}) {
	if err := j.validateSetShouldPersistNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldPersistNetwork",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetShouldPersistOsDisk(val interface{}) {
	if err := j.validateSetShouldPersistOsDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldPersistOsDisk",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetShutdownScript(val *string) {
	if err := j.validateSetShutdownScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownScript",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetVmName(val *string) {
	if err := j.validateSetVmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmName",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetVmNamePrefix(val *string) {
	if err := j.validateSetVmNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmNamePrefix",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzure)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Generates CDKTF code for importing a StatefulNodeAzure resource upon running "cdktf plan <stack-name>".
func StatefulNodeAzure_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStatefulNodeAzure_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzure",
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
func StatefulNodeAzure_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStatefulNodeAzure_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzure",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StatefulNodeAzure_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStatefulNodeAzure_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzure",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StatefulNodeAzure_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStatefulNodeAzure_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzure",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StatefulNodeAzure_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzure",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutAttachDataDisk(value interface{}) {
	if err := s.validatePutAttachDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAttachDataDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutBootDiagnostics(value interface{}) {
	if err := s.validatePutBootDiagnosticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBootDiagnostics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutDataDisk(value interface{}) {
	if err := s.validatePutDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutDelete(value interface{}) {
	if err := s.validatePutDeleteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDelete",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutDetachDataDisk(value interface{}) {
	if err := s.validatePutDetachDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDetachDataDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutExtension(value interface{}) {
	if err := s.validatePutExtensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExtension",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutHealth(value *StatefulNodeAzureHealth) {
	if err := s.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHealth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutImage(value *StatefulNodeAzureImage) {
	if err := s.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putImage",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutImportVm(value interface{}) {
	if err := s.validatePutImportVmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putImportVm",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutLoadBalancer(value interface{}) {
	if err := s.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutLogin(value *StatefulNodeAzureLogin) {
	if err := s.validatePutLoginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLogin",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutManagedServiceIdentities(value interface{}) {
	if err := s.validatePutManagedServiceIdentitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putManagedServiceIdentities",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutNetwork(value *StatefulNodeAzureNetwork) {
	if err := s.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetwork",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutOsDisk(value *StatefulNodeAzureOsDisk) {
	if err := s.validatePutOsDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOsDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutProximityPlacementGroups(value interface{}) {
	if err := s.validatePutProximityPlacementGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProximityPlacementGroups",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutSchedulingTask(value interface{}) {
	if err := s.validatePutSchedulingTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSchedulingTask",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutSecret(value interface{}) {
	if err := s.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSecret",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutSecurity(value *StatefulNodeAzureSecurity) {
	if err := s.validatePutSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSecurity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutSignal(value interface{}) {
	if err := s.validatePutSignalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSignal",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutStrategy(value *StatefulNodeAzureStrategy) {
	if err := s.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStrategy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutTag(value interface{}) {
	if err := s.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTag",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutUpdateState(value interface{}) {
	if err := s.validatePutUpdateStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUpdateState",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) PutVmSizes(value *StatefulNodeAzureVmSizes) {
	if err := s.validatePutVmSizesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVmSizes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetAttachDataDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetAttachDataDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetBootDiagnostics() {
	_jsii_.InvokeVoid(
		s,
		"resetBootDiagnostics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetCustomData() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomData",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetDataDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetDataDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetDataDisksPersistenceMode() {
	_jsii_.InvokeVoid(
		s,
		"resetDataDisksPersistenceMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetDetachDataDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetDetachDataDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetExtension() {
	_jsii_.InvokeVoid(
		s,
		"resetExtension",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetHealth() {
	_jsii_.InvokeVoid(
		s,
		"resetHealth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetImage() {
	_jsii_.InvokeVoid(
		s,
		"resetImage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetImportVm() {
	_jsii_.InvokeVoid(
		s,
		"resetImportVm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetLicenseType() {
	_jsii_.InvokeVoid(
		s,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		s,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetLogin() {
	_jsii_.InvokeVoid(
		s,
		"resetLogin",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetManagedServiceIdentities() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedServiceIdentities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetNetwork() {
	_jsii_.InvokeVoid(
		s,
		"resetNetwork",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetOsDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetOsDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetOsDiskPersistenceMode() {
	_jsii_.InvokeVoid(
		s,
		"resetOsDiskPersistenceMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetPreferredZone() {
	_jsii_.InvokeVoid(
		s,
		"resetPreferredZone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetProximityPlacementGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetProximityPlacementGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetSchedulingTask() {
	_jsii_.InvokeVoid(
		s,
		"resetSchedulingTask",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetSecret() {
	_jsii_.InvokeVoid(
		s,
		"resetSecret",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetSecurity() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetShutdownScript() {
	_jsii_.InvokeVoid(
		s,
		"resetShutdownScript",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetSignal() {
	_jsii_.InvokeVoid(
		s,
		"resetSignal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetTag() {
	_jsii_.InvokeVoid(
		s,
		"resetTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetUpdateState() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdateState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetUserData() {
	_jsii_.InvokeVoid(
		s,
		"resetUserData",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetVmName() {
	_jsii_.InvokeVoid(
		s,
		"resetVmName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetVmNamePrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetVmNamePrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) ResetZones() {
	_jsii_.InvokeVoid(
		s,
		"resetZones",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzure) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzure) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

