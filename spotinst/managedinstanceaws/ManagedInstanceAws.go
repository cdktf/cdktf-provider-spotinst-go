// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/managedinstanceaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/managed_instance_aws spotinst_managed_instance_aws}.
type ManagedInstanceAws interface {
	cdktf.TerraformResource
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	BlockDeviceMappings() ManagedInstanceAwsBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
	BlockDevicesMode() *string
	SetBlockDevicesMode(val *string)
	BlockDevicesModeInput() *string
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
	CpuCredits() *string
	SetCpuCredits(val *string)
	CpuCreditsInput() *string
	Delete() ManagedInstanceAwsDeleteList
	DeleteInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DrainingTimeout() *float64
	SetDrainingTimeout(val *float64)
	DrainingTimeoutInput() *float64
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	ElasticIp() *string
	SetElasticIp(val *string)
	ElasticIpInput() *string
	EnableMonitoring() interface{}
	SetEnableMonitoring(val interface{})
	EnableMonitoringInput() interface{}
	FallbackToOndemand() interface{}
	SetFallbackToOndemand(val interface{})
	FallbackToOndemandInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GracePeriod() *float64
	SetGracePeriod(val *float64)
	GracePeriodInput() *float64
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	HealthCheckTypeInput() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	InstanceTypesInput() *[]*string
	IntegrationRoute53() ManagedInstanceAwsIntegrationRoute53OutputReference
	IntegrationRoute53Input() *ManagedInstanceAwsIntegrationRoute53
	KeyPair() *string
	SetKeyPair(val *string)
	KeyPairInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifeCycle() *string
	SetLifeCycle(val *string)
	LifeCycleInput() *string
	LoadBalancers() ManagedInstanceAwsLoadBalancersList
	LoadBalancersInput() interface{}
	ManagedInstanceAction() ManagedInstanceAwsManagedInstanceActionOutputReference
	ManagedInstanceActionInput() *ManagedInstanceAwsManagedInstanceAction
	MetadataOptions() ManagedInstanceAwsMetadataOptionsOutputReference
	MetadataOptionsInput() *ManagedInstanceAwsMetadataOptions
	MinimumInstanceLifetime() *float64
	SetMinimumInstanceLifetime(val *float64)
	MinimumInstanceLifetimeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() ManagedInstanceAwsNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	OptimizationWindows() *[]*string
	SetOptimizationWindows(val *[]*string)
	OptimizationWindowsInput() *[]*string
	Orientation() *string
	SetOrientation(val *string)
	OrientationInput() *string
	PersistBlockDevices() interface{}
	SetPersistBlockDevices(val interface{})
	PersistBlockDevicesInput() interface{}
	PersistPrivateIp() interface{}
	SetPersistPrivateIp(val interface{})
	PersistPrivateIpInput() interface{}
	PersistRootDevice() interface{}
	SetPersistRootDevice(val interface{})
	PersistRootDeviceInput() interface{}
	PlacementTenancy() *string
	SetPlacementTenancy(val *string)
	PlacementTenancyInput() *string
	PreferredType() *string
	SetPreferredType(val *string)
	PreferredTypeInput() *string
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
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
	ResourceTagSpecification() ManagedInstanceAwsResourceTagSpecificationList
	ResourceTagSpecificationInput() interface{}
	RevertToSpot() ManagedInstanceAwsRevertToSpotOutputReference
	RevertToSpotInput() *ManagedInstanceAwsRevertToSpot
	ScheduledTask() ManagedInstanceAwsScheduledTaskList
	ScheduledTaskInput() interface{}
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	ShutdownScript() *string
	SetShutdownScript(val *string)
	ShutdownScriptInput() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() ManagedInstanceAwsTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UnhealthyDuration() *float64
	SetUnhealthyDuration(val *float64)
	UnhealthyDurationInput() *float64
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	UtilizeReservedInstances() interface{}
	SetUtilizeReservedInstances(val interface{})
	UtilizeReservedInstancesInput() interface{}
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutBlockDeviceMappings(value interface{})
	PutDelete(value interface{})
	PutIntegrationRoute53(value *ManagedInstanceAwsIntegrationRoute53)
	PutLoadBalancers(value interface{})
	PutManagedInstanceAction(value *ManagedInstanceAwsManagedInstanceAction)
	PutMetadataOptions(value *ManagedInstanceAwsMetadataOptions)
	PutNetworkInterface(value interface{})
	PutResourceTagSpecification(value interface{})
	PutRevertToSpot(value *ManagedInstanceAwsRevertToSpot)
	PutScheduledTask(value interface{})
	PutTags(value interface{})
	ResetAutoHealing()
	ResetBlockDeviceMappings()
	ResetBlockDevicesMode()
	ResetCpuCredits()
	ResetDelete()
	ResetDescription()
	ResetDrainingTimeout()
	ResetEbsOptimized()
	ResetElasticIp()
	ResetEnableMonitoring()
	ResetFallbackToOndemand()
	ResetGracePeriod()
	ResetHealthCheckType()
	ResetIamInstanceProfile()
	ResetId()
	ResetIntegrationRoute53()
	ResetKeyPair()
	ResetLifeCycle()
	ResetLoadBalancers()
	ResetManagedInstanceAction()
	ResetMetadataOptions()
	ResetMinimumInstanceLifetime()
	ResetNetworkInterface()
	ResetOptimizationWindows()
	ResetOrientation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistPrivateIp()
	ResetPersistRootDevice()
	ResetPlacementTenancy()
	ResetPreferredType()
	ResetPrivateIp()
	ResetRegion()
	ResetResourceTagSpecification()
	ResetRevertToSpot()
	ResetScheduledTask()
	ResetSecurityGroupIds()
	ResetShutdownScript()
	ResetTags()
	ResetUnhealthyDuration()
	ResetUserData()
	ResetUtilizeReservedInstances()
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

// The jsii proxy struct for ManagedInstanceAws
type jsiiProxy_ManagedInstanceAws struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ManagedInstanceAws) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) BlockDeviceMappings() ManagedInstanceAwsBlockDeviceMappingsList {
	var returns ManagedInstanceAwsBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) BlockDevicesMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockDevicesMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) BlockDevicesModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockDevicesModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) CpuCredits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCredits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) CpuCreditsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCreditsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Delete() ManagedInstanceAwsDeleteList {
	var returns ManagedInstanceAwsDeleteList
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) DeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) DrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) DrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ElasticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ElasticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) EnableMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) EnableMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) FallbackToOndemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) FallbackToOndemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) GracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) GracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) IntegrationRoute53() ManagedInstanceAwsIntegrationRoute53OutputReference {
	var returns ManagedInstanceAwsIntegrationRoute53OutputReference
	_jsii_.Get(
		j,
		"integrationRoute53",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) IntegrationRoute53Input() *ManagedInstanceAwsIntegrationRoute53 {
	var returns *ManagedInstanceAwsIntegrationRoute53
	_jsii_.Get(
		j,
		"integrationRoute53Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) KeyPairInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) LifeCycle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifeCycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) LifeCycleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifeCycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) LoadBalancers() ManagedInstanceAwsLoadBalancersList {
	var returns ManagedInstanceAwsLoadBalancersList
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) LoadBalancersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ManagedInstanceAction() ManagedInstanceAwsManagedInstanceActionOutputReference {
	var returns ManagedInstanceAwsManagedInstanceActionOutputReference
	_jsii_.Get(
		j,
		"managedInstanceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ManagedInstanceActionInput() *ManagedInstanceAwsManagedInstanceAction {
	var returns *ManagedInstanceAwsManagedInstanceAction
	_jsii_.Get(
		j,
		"managedInstanceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) MetadataOptions() ManagedInstanceAwsMetadataOptionsOutputReference {
	var returns ManagedInstanceAwsMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) MetadataOptionsInput() *ManagedInstanceAwsMetadataOptions {
	var returns *ManagedInstanceAwsMetadataOptions
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) MinimumInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) MinimumInstanceLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInstanceLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) NetworkInterface() ManagedInstanceAwsNetworkInterfaceList {
	var returns ManagedInstanceAwsNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) OptimizationWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optimizationWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) OptimizationWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optimizationWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Orientation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orientation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) OrientationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orientationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PersistBlockDevices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistBlockDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PersistBlockDevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistBlockDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PersistPrivateIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistPrivateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PersistPrivateIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistPrivateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PersistRootDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistRootDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PersistRootDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistRootDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PlacementTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PlacementTenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PreferredType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PreferredTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Product() *string {
	var returns *string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ProductInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ResourceTagSpecification() ManagedInstanceAwsResourceTagSpecificationList {
	var returns ManagedInstanceAwsResourceTagSpecificationList
	_jsii_.Get(
		j,
		"resourceTagSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ResourceTagSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) RevertToSpot() ManagedInstanceAwsRevertToSpotOutputReference {
	var returns ManagedInstanceAwsRevertToSpotOutputReference
	_jsii_.Get(
		j,
		"revertToSpot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) RevertToSpotInput() *ManagedInstanceAwsRevertToSpot {
	var returns *ManagedInstanceAwsRevertToSpot
	_jsii_.Get(
		j,
		"revertToSpotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ScheduledTask() ManagedInstanceAwsScheduledTaskList {
	var returns ManagedInstanceAwsScheduledTaskList
	_jsii_.Get(
		j,
		"scheduledTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ScheduledTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ShutdownScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) ShutdownScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) Tags() ManagedInstanceAwsTagsList {
	var returns ManagedInstanceAwsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) UnhealthyDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) UnhealthyDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) UtilizeReservedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeReservedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) UtilizeReservedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeReservedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAws) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/managed_instance_aws spotinst_managed_instance_aws} Resource.
func NewManagedInstanceAws(scope constructs.Construct, id *string, config *ManagedInstanceAwsConfig) ManagedInstanceAws {
	_init_.Initialize()

	if err := validateNewManagedInstanceAwsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedInstanceAws{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAws",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/managed_instance_aws spotinst_managed_instance_aws} Resource.
func NewManagedInstanceAws_Override(m ManagedInstanceAws, scope constructs.Construct, id *string, config *ManagedInstanceAwsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAws",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetAutoHealing(val interface{}) {
	if err := j.validateSetAutoHealingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetBlockDevicesMode(val *string) {
	if err := j.validateSetBlockDevicesModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockDevicesMode",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetCpuCredits(val *string) {
	if err := j.validateSetCpuCreditsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCredits",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetDrainingTimeout(val *float64) {
	if err := j.validateSetDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainingTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetElasticIp(val *string) {
	if err := j.validateSetElasticIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticIp",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetEnableMonitoring(val interface{}) {
	if err := j.validateSetEnableMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMonitoring",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetFallbackToOndemand(val interface{}) {
	if err := j.validateSetFallbackToOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackToOndemand",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetGracePeriod(val *float64) {
	if err := j.validateSetGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracePeriod",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetKeyPair(val *string) {
	if err := j.validateSetKeyPairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPair",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetLifeCycle(val *string) {
	if err := j.validateSetLifeCycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifeCycle",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetMinimumInstanceLifetime(val *float64) {
	if err := j.validateSetMinimumInstanceLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumInstanceLifetime",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetOptimizationWindows(val *[]*string) {
	if err := j.validateSetOptimizationWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optimizationWindows",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetOrientation(val *string) {
	if err := j.validateSetOrientationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orientation",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetPersistBlockDevices(val interface{}) {
	if err := j.validateSetPersistBlockDevicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistBlockDevices",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetPersistPrivateIp(val interface{}) {
	if err := j.validateSetPersistPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistPrivateIp",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetPersistRootDevice(val interface{}) {
	if err := j.validateSetPersistRootDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistRootDevice",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetPlacementTenancy(val *string) {
	if err := j.validateSetPlacementTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementTenancy",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetPreferredType(val *string) {
	if err := j.validateSetPreferredTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredType",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetProduct(val *string) {
	if err := j.validateSetProductParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"product",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetShutdownScript(val *string) {
	if err := j.validateSetShutdownScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownScript",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetUnhealthyDuration(val *float64) {
	if err := j.validateSetUnhealthyDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyDuration",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetUtilizeReservedInstances(val interface{}) {
	if err := j.validateSetUtilizeReservedInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utilizeReservedInstances",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAws)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a ManagedInstanceAws resource upon running "cdktf plan <stack-name>".
func ManagedInstanceAws_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateManagedInstanceAws_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAws",
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
func ManagedInstanceAws_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedInstanceAws_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAws",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedInstanceAws_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedInstanceAws_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAws",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedInstanceAws_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedInstanceAws_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAws",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagedInstanceAws_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAws",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutBlockDeviceMappings(value interface{}) {
	if err := m.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutDelete(value interface{}) {
	if err := m.validatePutDeleteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDelete",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutIntegrationRoute53(value *ManagedInstanceAwsIntegrationRoute53) {
	if err := m.validatePutIntegrationRoute53Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIntegrationRoute53",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutLoadBalancers(value interface{}) {
	if err := m.validatePutLoadBalancersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLoadBalancers",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutManagedInstanceAction(value *ManagedInstanceAwsManagedInstanceAction) {
	if err := m.validatePutManagedInstanceActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putManagedInstanceAction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutMetadataOptions(value *ManagedInstanceAwsMetadataOptions) {
	if err := m.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutNetworkInterface(value interface{}) {
	if err := m.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutResourceTagSpecification(value interface{}) {
	if err := m.validatePutResourceTagSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putResourceTagSpecification",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutRevertToSpot(value *ManagedInstanceAwsRevertToSpot) {
	if err := m.validatePutRevertToSpotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRevertToSpot",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutScheduledTask(value interface{}) {
	if err := m.validatePutScheduledTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScheduledTask",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) PutTags(value interface{}) {
	if err := m.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTags",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		m,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetBlockDevicesMode() {
	_jsii_.InvokeVoid(
		m,
		"resetBlockDevicesMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetCpuCredits() {
	_jsii_.InvokeVoid(
		m,
		"resetCpuCredits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetDrainingTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetDrainingTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		m,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetElasticIp() {
	_jsii_.InvokeVoid(
		m,
		"resetElasticIp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetEnableMonitoring() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableMonitoring",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetFallbackToOndemand() {
	_jsii_.InvokeVoid(
		m,
		"resetFallbackToOndemand",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetGracePeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetGracePeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		m,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		m,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetIntegrationRoute53() {
	_jsii_.InvokeVoid(
		m,
		"resetIntegrationRoute53",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetKeyPair() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyPair",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetLifeCycle() {
	_jsii_.InvokeVoid(
		m,
		"resetLifeCycle",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetLoadBalancers() {
	_jsii_.InvokeVoid(
		m,
		"resetLoadBalancers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetManagedInstanceAction() {
	_jsii_.InvokeVoid(
		m,
		"resetManagedInstanceAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetMinimumInstanceLifetime() {
	_jsii_.InvokeVoid(
		m,
		"resetMinimumInstanceLifetime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		m,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetOptimizationWindows() {
	_jsii_.InvokeVoid(
		m,
		"resetOptimizationWindows",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetOrientation() {
	_jsii_.InvokeVoid(
		m,
		"resetOrientation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetPersistPrivateIp() {
	_jsii_.InvokeVoid(
		m,
		"resetPersistPrivateIp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetPersistRootDevice() {
	_jsii_.InvokeVoid(
		m,
		"resetPersistRootDevice",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetPlacementTenancy() {
	_jsii_.InvokeVoid(
		m,
		"resetPlacementTenancy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetPreferredType() {
	_jsii_.InvokeVoid(
		m,
		"resetPreferredType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		m,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetResourceTagSpecification() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceTagSpecification",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetRevertToSpot() {
	_jsii_.InvokeVoid(
		m,
		"resetRevertToSpot",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetScheduledTask() {
	_jsii_.InvokeVoid(
		m,
		"resetScheduledTask",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetShutdownScript() {
	_jsii_.InvokeVoid(
		m,
		"resetShutdownScript",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetUnhealthyDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetUnhealthyDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetUserData() {
	_jsii_.InvokeVoid(
		m,
		"resetUserData",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) ResetUtilizeReservedInstances() {
	_jsii_.InvokeVoid(
		m,
		"resetUtilizeReservedInstances",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAws) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAws) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

