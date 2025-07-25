// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/ocean_aws spotinst_ocean_aws}.
type OceanAws interface {
	cdktf.TerraformResource
	AssociateIpv6Address() interface{}
	SetAssociateIpv6Address(val interface{})
	AssociateIpv6AddressInput() interface{}
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	AssociatePublicIpAddressInput() interface{}
	AttachLoadBalancer() OceanAwsAttachLoadBalancerList
	AttachLoadBalancerInput() interface{}
	Autoscaler() OceanAwsAutoscalerOutputReference
	AutoscalerInput() *OceanAwsAutoscaler
	Blacklist() *[]*string
	SetBlacklist(val *[]*string)
	BlacklistInput() *[]*string
	BlockDeviceMappings() OceanAwsBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterOrientation() OceanAwsClusterOrientationList
	ClusterOrientationInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControllerId() *string
	SetControllerId(val *string)
	ControllerIdInput() *string
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
	DetachLoadBalancer() OceanAwsDetachLoadBalancerList
	DetachLoadBalancerInput() interface{}
	DrainingTimeout() *float64
	SetDrainingTimeout(val *float64)
	DrainingTimeoutInput() *float64
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	FallbackToOndemand() interface{}
	SetFallbackToOndemand(val interface{})
	FallbackToOndemandInput() interface{}
	Filters() OceanAwsFiltersOutputReference
	FiltersInput() *OceanAwsFilters
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
	HealthCheckUnhealthyDurationBeforeReplacement() *float64
	SetHealthCheckUnhealthyDurationBeforeReplacement(val *float64)
	HealthCheckUnhealthyDurationBeforeReplacementInput() *float64
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceMetadataOptions() OceanAwsInstanceMetadataOptionsOutputReference
	InstanceMetadataOptionsInput() *OceanAwsInstanceMetadataOptions
	InstanceStorePolicy() OceanAwsInstanceStorePolicyOutputReference
	InstanceStorePolicyInput() *OceanAwsInstanceStorePolicy
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancers() OceanAwsLoadBalancersList
	LoadBalancersInput() interface{}
	Logging() OceanAwsLoggingOutputReference
	LoggingInput() *OceanAwsLogging
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	Monitoring() interface{}
	SetMonitoring(val interface{})
	MonitoringInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReservedEnis() *float64
	SetReservedEnis(val *float64)
	ReservedEnisInput() *float64
	ResourceTagSpecification() OceanAwsResourceTagSpecificationList
	ResourceTagSpecificationInput() interface{}
	RootVolumeSize() *float64
	SetRootVolumeSize(val *float64)
	RootVolumeSizeInput() *float64
	ScheduledTask() OceanAwsScheduledTaskOutputReference
	ScheduledTaskInput() *OceanAwsScheduledTask
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SpotPercentage() *float64
	SetSpotPercentage(val *float64)
	SpotPercentageInput() *float64
	SpreadNodesBy() *string
	SetSpreadNodesBy(val *string)
	SpreadNodesByInput() *string
	StartupTaints() OceanAwsStartupTaintsList
	StartupTaintsInput() interface{}
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() OceanAwsTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatePolicy() OceanAwsUpdatePolicyOutputReference
	UpdatePolicyInput() *OceanAwsUpdatePolicy
	UseAsTemplateOnly() interface{}
	SetUseAsTemplateOnly(val interface{})
	UseAsTemplateOnlyInput() interface{}
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	UtilizeCommitments() interface{}
	SetUtilizeCommitments(val interface{})
	UtilizeCommitmentsInput() interface{}
	UtilizeReservedInstances() interface{}
	SetUtilizeReservedInstances(val interface{})
	UtilizeReservedInstancesInput() interface{}
	Whitelist() *[]*string
	SetWhitelist(val *[]*string)
	WhitelistInput() *[]*string
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
	PutAttachLoadBalancer(value interface{})
	PutAutoscaler(value *OceanAwsAutoscaler)
	PutBlockDeviceMappings(value interface{})
	PutClusterOrientation(value interface{})
	PutDetachLoadBalancer(value interface{})
	PutFilters(value *OceanAwsFilters)
	PutInstanceMetadataOptions(value *OceanAwsInstanceMetadataOptions)
	PutInstanceStorePolicy(value *OceanAwsInstanceStorePolicy)
	PutLoadBalancers(value interface{})
	PutLogging(value *OceanAwsLogging)
	PutResourceTagSpecification(value interface{})
	PutScheduledTask(value *OceanAwsScheduledTask)
	PutStartupTaints(value interface{})
	PutTags(value interface{})
	PutUpdatePolicy(value *OceanAwsUpdatePolicy)
	ResetAssociateIpv6Address()
	ResetAssociatePublicIpAddress()
	ResetAttachLoadBalancer()
	ResetAutoscaler()
	ResetBlacklist()
	ResetBlockDeviceMappings()
	ResetClusterOrientation()
	ResetControllerId()
	ResetDesiredCapacity()
	ResetDetachLoadBalancer()
	ResetDrainingTimeout()
	ResetEbsOptimized()
	ResetFallbackToOndemand()
	ResetFilters()
	ResetGracePeriod()
	ResetHealthCheckUnhealthyDurationBeforeReplacement()
	ResetIamInstanceProfile()
	ResetId()
	ResetInstanceMetadataOptions()
	ResetInstanceStorePolicy()
	ResetKeyName()
	ResetLoadBalancers()
	ResetLogging()
	ResetMaxSize()
	ResetMinSize()
	ResetMonitoring()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetReservedEnis()
	ResetResourceTagSpecification()
	ResetRootVolumeSize()
	ResetScheduledTask()
	ResetSpotPercentage()
	ResetSpreadNodesBy()
	ResetStartupTaints()
	ResetTags()
	ResetUpdatePolicy()
	ResetUseAsTemplateOnly()
	ResetUserData()
	ResetUtilizeCommitments()
	ResetUtilizeReservedInstances()
	ResetWhitelist()
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

// The jsii proxy struct for OceanAws
type jsiiProxy_OceanAws struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanAws) AssociateIpv6Address() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateIpv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) AssociateIpv6AddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateIpv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) AttachLoadBalancer() OceanAwsAttachLoadBalancerList {
	var returns OceanAwsAttachLoadBalancerList
	_jsii_.Get(
		j,
		"attachLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) AttachLoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Autoscaler() OceanAwsAutoscalerOutputReference {
	var returns OceanAwsAutoscalerOutputReference
	_jsii_.Get(
		j,
		"autoscaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) AutoscalerInput() *OceanAwsAutoscaler {
	var returns *OceanAwsAutoscaler
	_jsii_.Get(
		j,
		"autoscalerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Blacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) BlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) BlockDeviceMappings() OceanAwsBlockDeviceMappingsList {
	var returns OceanAwsBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ClusterOrientation() OceanAwsClusterOrientationList {
	var returns OceanAwsClusterOrientationList
	_jsii_.Get(
		j,
		"clusterOrientation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ClusterOrientationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterOrientationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ControllerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ControllerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) DetachLoadBalancer() OceanAwsDetachLoadBalancerList {
	var returns OceanAwsDetachLoadBalancerList
	_jsii_.Get(
		j,
		"detachLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) DetachLoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detachLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) DrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) DrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) FallbackToOndemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) FallbackToOndemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Filters() OceanAwsFiltersOutputReference {
	var returns OceanAwsFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) FiltersInput() *OceanAwsFilters {
	var returns *OceanAwsFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) GracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) GracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) HealthCheckUnhealthyDurationBeforeReplacement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckUnhealthyDurationBeforeReplacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) HealthCheckUnhealthyDurationBeforeReplacementInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckUnhealthyDurationBeforeReplacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) InstanceMetadataOptions() OceanAwsInstanceMetadataOptionsOutputReference {
	var returns OceanAwsInstanceMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"instanceMetadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) InstanceMetadataOptionsInput() *OceanAwsInstanceMetadataOptions {
	var returns *OceanAwsInstanceMetadataOptions
	_jsii_.Get(
		j,
		"instanceMetadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) InstanceStorePolicy() OceanAwsInstanceStorePolicyOutputReference {
	var returns OceanAwsInstanceStorePolicyOutputReference
	_jsii_.Get(
		j,
		"instanceStorePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) InstanceStorePolicyInput() *OceanAwsInstanceStorePolicy {
	var returns *OceanAwsInstanceStorePolicy
	_jsii_.Get(
		j,
		"instanceStorePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) LoadBalancers() OceanAwsLoadBalancersList {
	var returns OceanAwsLoadBalancersList
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) LoadBalancersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Logging() OceanAwsLoggingOutputReference {
	var returns OceanAwsLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) LoggingInput() *OceanAwsLogging {
	var returns *OceanAwsLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ReservedEnis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedEnis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ReservedEnisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedEnisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ResourceTagSpecification() OceanAwsResourceTagSpecificationList {
	var returns OceanAwsResourceTagSpecificationList
	_jsii_.Get(
		j,
		"resourceTagSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ResourceTagSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) RootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) RootVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ScheduledTask() OceanAwsScheduledTaskOutputReference {
	var returns OceanAwsScheduledTaskOutputReference
	_jsii_.Get(
		j,
		"scheduledTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) ScheduledTaskInput() *OceanAwsScheduledTask {
	var returns *OceanAwsScheduledTask
	_jsii_.Get(
		j,
		"scheduledTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) SpotPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) SpotPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) SpreadNodesBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spreadNodesBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) SpreadNodesByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spreadNodesByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) StartupTaints() OceanAwsStartupTaintsList {
	var returns OceanAwsStartupTaintsList
	_jsii_.Get(
		j,
		"startupTaints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) StartupTaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startupTaintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Tags() OceanAwsTagsList {
	var returns OceanAwsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UpdatePolicy() OceanAwsUpdatePolicyOutputReference {
	var returns OceanAwsUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UpdatePolicyInput() *OceanAwsUpdatePolicy {
	var returns *OceanAwsUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UseAsTemplateOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAsTemplateOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UseAsTemplateOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAsTemplateOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UtilizeCommitments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeCommitments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UtilizeCommitmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeCommitmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UtilizeReservedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeReservedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) UtilizeReservedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeReservedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) Whitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAws) WhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/ocean_aws spotinst_ocean_aws} Resource.
func NewOceanAws(scope constructs.Construct, id *string, config *OceanAwsConfig) OceanAws {
	_init_.Initialize()

	if err := validateNewOceanAwsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAws{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAws",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/ocean_aws spotinst_ocean_aws} Resource.
func NewOceanAws_Override(o OceanAws, scope constructs.Construct, id *string, config *OceanAwsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAws.OceanAws",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanAws)SetAssociateIpv6Address(val interface{}) {
	if err := j.validateSetAssociateIpv6AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateIpv6Address",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetBlacklist(val *[]*string) {
	if err := j.validateSetBlacklistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blacklist",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetControllerId(val *string) {
	if err := j.validateSetControllerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controllerId",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetDrainingTimeout(val *float64) {
	if err := j.validateSetDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainingTimeout",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetFallbackToOndemand(val interface{}) {
	if err := j.validateSetFallbackToOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackToOndemand",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetGracePeriod(val *float64) {
	if err := j.validateSetGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracePeriod",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetHealthCheckUnhealthyDurationBeforeReplacement(val *float64) {
	if err := j.validateSetHealthCheckUnhealthyDurationBeforeReplacementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckUnhealthyDurationBeforeReplacement",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetReservedEnis(val *float64) {
	if err := j.validateSetReservedEnisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedEnis",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetRootVolumeSize(val *float64) {
	if err := j.validateSetRootVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetSpotPercentage(val *float64) {
	if err := j.validateSetSpotPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetSpreadNodesBy(val *string) {
	if err := j.validateSetSpreadNodesByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spreadNodesBy",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetUseAsTemplateOnly(val interface{}) {
	if err := j.validateSetUseAsTemplateOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAsTemplateOnly",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetUtilizeCommitments(val interface{}) {
	if err := j.validateSetUtilizeCommitmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utilizeCommitments",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetUtilizeReservedInstances(val interface{}) {
	if err := j.validateSetUtilizeReservedInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utilizeReservedInstances",
		val,
	)
}

func (j *jsiiProxy_OceanAws)SetWhitelist(val *[]*string) {
	if err := j.validateSetWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelist",
		val,
	)
}

// Generates CDKTF code for importing a OceanAws resource upon running "cdktf plan <stack-name>".
func OceanAws_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceanAws_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAws.OceanAws",
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
func OceanAws_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAws_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAws.OceanAws",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAws_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAws_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAws.OceanAws",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAws_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAws_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAws.OceanAws",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanAws_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanAws.OceanAws",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanAws) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceanAws) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanAws) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAws) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAws) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAws) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAws) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAws) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAws) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAws) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAws) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAws) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAws) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceanAws) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAws) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanAws) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceanAws) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanAws) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanAws) PutAttachLoadBalancer(value interface{}) {
	if err := o.validatePutAttachLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAttachLoadBalancer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutAutoscaler(value *OceanAwsAutoscaler) {
	if err := o.validatePutAutoscalerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaler",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutBlockDeviceMappings(value interface{}) {
	if err := o.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutClusterOrientation(value interface{}) {
	if err := o.validatePutClusterOrientationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putClusterOrientation",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutDetachLoadBalancer(value interface{}) {
	if err := o.validatePutDetachLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDetachLoadBalancer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutFilters(value *OceanAwsFilters) {
	if err := o.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFilters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutInstanceMetadataOptions(value *OceanAwsInstanceMetadataOptions) {
	if err := o.validatePutInstanceMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInstanceMetadataOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutInstanceStorePolicy(value *OceanAwsInstanceStorePolicy) {
	if err := o.validatePutInstanceStorePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInstanceStorePolicy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutLoadBalancers(value interface{}) {
	if err := o.validatePutLoadBalancersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLoadBalancers",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutLogging(value *OceanAwsLogging) {
	if err := o.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogging",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutResourceTagSpecification(value interface{}) {
	if err := o.validatePutResourceTagSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putResourceTagSpecification",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutScheduledTask(value *OceanAwsScheduledTask) {
	if err := o.validatePutScheduledTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putScheduledTask",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutStartupTaints(value interface{}) {
	if err := o.validatePutStartupTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStartupTaints",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutTags(value interface{}) {
	if err := o.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTags",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) PutUpdatePolicy(value *OceanAwsUpdatePolicy) {
	if err := o.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAws) ResetAssociateIpv6Address() {
	_jsii_.InvokeVoid(
		o,
		"resetAssociateIpv6Address",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		o,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetAttachLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetAttachLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetAutoscaler() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaler",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetBlacklist() {
	_jsii_.InvokeVoid(
		o,
		"resetBlacklist",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		o,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetClusterOrientation() {
	_jsii_.InvokeVoid(
		o,
		"resetClusterOrientation",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetControllerId() {
	_jsii_.InvokeVoid(
		o,
		"resetControllerId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		o,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetDetachLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetDetachLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetDrainingTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainingTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetFallbackToOndemand() {
	_jsii_.InvokeVoid(
		o,
		"resetFallbackToOndemand",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetFilters() {
	_jsii_.InvokeVoid(
		o,
		"resetFilters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetGracePeriod() {
	_jsii_.InvokeVoid(
		o,
		"resetGracePeriod",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetHealthCheckUnhealthyDurationBeforeReplacement() {
	_jsii_.InvokeVoid(
		o,
		"resetHealthCheckUnhealthyDurationBeforeReplacement",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		o,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetInstanceMetadataOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceMetadataOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetInstanceStorePolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceStorePolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetKeyName() {
	_jsii_.InvokeVoid(
		o,
		"resetKeyName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetLoadBalancers() {
	_jsii_.InvokeVoid(
		o,
		"resetLoadBalancers",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetLogging() {
	_jsii_.InvokeVoid(
		o,
		"resetLogging",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetMaxSize() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetMinSize() {
	_jsii_.InvokeVoid(
		o,
		"resetMinSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetMonitoring() {
	_jsii_.InvokeVoid(
		o,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetRegion() {
	_jsii_.InvokeVoid(
		o,
		"resetRegion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetReservedEnis() {
	_jsii_.InvokeVoid(
		o,
		"resetReservedEnis",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetResourceTagSpecification() {
	_jsii_.InvokeVoid(
		o,
		"resetResourceTagSpecification",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetRootVolumeSize() {
	_jsii_.InvokeVoid(
		o,
		"resetRootVolumeSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetScheduledTask() {
	_jsii_.InvokeVoid(
		o,
		"resetScheduledTask",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetSpotPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetSpotPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetSpreadNodesBy() {
	_jsii_.InvokeVoid(
		o,
		"resetSpreadNodesBy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetStartupTaints() {
	_jsii_.InvokeVoid(
		o,
		"resetStartupTaints",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetUseAsTemplateOnly() {
	_jsii_.InvokeVoid(
		o,
		"resetUseAsTemplateOnly",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetUserData() {
	_jsii_.InvokeVoid(
		o,
		"resetUserData",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetUtilizeCommitments() {
	_jsii_.InvokeVoid(
		o,
		"resetUtilizeCommitments",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetUtilizeReservedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUtilizeReservedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) ResetWhitelist() {
	_jsii_.InvokeVoid(
		o,
		"resetWhitelist",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAws) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAws) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAws) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAws) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAws) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAws) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

