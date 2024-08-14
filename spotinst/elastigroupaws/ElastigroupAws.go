// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.186.0/docs/resources/elastigroup_aws spotinst_elastigroup_aws}.
type ElastigroupAws interface {
	cdktf.TerraformResource
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	BlockDevicesMode() *string
	SetBlockDevicesMode(val *string)
	BlockDevicesModeInput() *string
	CapacityUnit() *string
	SetCapacityUnit(val *string)
	CapacityUnitInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConsiderOdPricing() interface{}
	SetConsiderOdPricing(val interface{})
	ConsiderOdPricingInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuCredits() *string
	SetCpuCredits(val *string)
	CpuCreditsInput() *string
	CpuOptions() ElastigroupAwsCpuOptionsOutputReference
	CpuOptionsInput() *ElastigroupAwsCpuOptions
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
	EbsBlockDevice() ElastigroupAwsEbsBlockDeviceList
	EbsBlockDeviceInput() interface{}
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	ElasticIps() *[]*string
	SetElasticIps(val *[]*string)
	ElasticIpsInput() *[]*string
	ElasticLoadBalancers() *[]*string
	SetElasticLoadBalancers(val *[]*string)
	ElasticLoadBalancersInput() *[]*string
	EnableMonitoring() interface{}
	SetEnableMonitoring(val interface{})
	EnableMonitoringInput() interface{}
	EphemeralBlockDevice() ElastigroupAwsEphemeralBlockDeviceList
	EphemeralBlockDeviceInput() interface{}
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
	HealthCheckGracePeriod() *float64
	SetHealthCheckGracePeriod(val *float64)
	HealthCheckGracePeriodInput() *float64
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	HealthCheckTypeInput() *string
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
	Images() ElastigroupAwsImagesList
	ImagesInput() interface{}
	ImmediateOdRecoverThreshold() *float64
	SetImmediateOdRecoverThreshold(val *float64)
	ImmediateOdRecoverThresholdInput() *float64
	InstanceTypesOndemand() *string
	SetInstanceTypesOndemand(val *string)
	InstanceTypesOndemandInput() *string
	InstanceTypesPreferredSpot() *[]*string
	SetInstanceTypesPreferredSpot(val *[]*string)
	InstanceTypesPreferredSpotInput() *[]*string
	InstanceTypesSpot() *[]*string
	SetInstanceTypesSpot(val *[]*string)
	InstanceTypesSpotInput() *[]*string
	InstanceTypesWeights() ElastigroupAwsInstanceTypesWeightsList
	InstanceTypesWeightsInput() interface{}
	IntegrationBeanstalk() ElastigroupAwsIntegrationBeanstalkOutputReference
	IntegrationBeanstalkInput() *ElastigroupAwsIntegrationBeanstalk
	IntegrationCodedeploy() ElastigroupAwsIntegrationCodedeployOutputReference
	IntegrationCodedeployInput() *ElastigroupAwsIntegrationCodedeploy
	IntegrationDockerSwarm() ElastigroupAwsIntegrationDockerSwarmOutputReference
	IntegrationDockerSwarmInput() *ElastigroupAwsIntegrationDockerSwarm
	IntegrationEcs() ElastigroupAwsIntegrationEcsOutputReference
	IntegrationEcsInput() *ElastigroupAwsIntegrationEcs
	IntegrationGitlab() ElastigroupAwsIntegrationGitlabOutputReference
	IntegrationGitlabInput() *ElastigroupAwsIntegrationGitlab
	IntegrationKubernetes() ElastigroupAwsIntegrationKubernetesOutputReference
	IntegrationKubernetesInput() *ElastigroupAwsIntegrationKubernetes
	IntegrationMesosphere() ElastigroupAwsIntegrationMesosphereOutputReference
	IntegrationMesosphereInput() *ElastigroupAwsIntegrationMesosphere
	IntegrationNomad() ElastigroupAwsIntegrationNomadOutputReference
	IntegrationNomadInput() *ElastigroupAwsIntegrationNomad
	IntegrationRancher() ElastigroupAwsIntegrationRancherOutputReference
	IntegrationRancherInput() *ElastigroupAwsIntegrationRancher
	IntegrationRoute53() ElastigroupAwsIntegrationRoute53OutputReference
	IntegrationRoute53Input() *ElastigroupAwsIntegrationRoute53
	Itf() ElastigroupAwsItfList
	ItfInput() interface{}
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifetimePeriod() *string
	SetLifetimePeriod(val *string)
	LifetimePeriodInput() *string
	Logging() ElastigroupAwsLoggingOutputReference
	LoggingInput() *ElastigroupAwsLogging
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MetadataOptions() ElastigroupAwsMetadataOptionsOutputReference
	MetadataOptionsInput() *ElastigroupAwsMetadataOptions
	MinimumInstanceLifetime() *float64
	SetMinimumInstanceLifetime(val *float64)
	MinimumInstanceLifetimeInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	MultipleMetrics() ElastigroupAwsMultipleMetricsOutputReference
	MultipleMetricsInput() *ElastigroupAwsMultipleMetrics
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() ElastigroupAwsNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	OndemandCount() *float64
	SetOndemandCount(val *float64)
	OndemandCountInput() *float64
	OnDemandTypes() *[]*string
	SetOnDemandTypes(val *[]*string)
	OnDemandTypesInput() *[]*string
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
	PreferredAvailabilityZones() *[]*string
	SetPreferredAvailabilityZones(val *[]*string)
	PreferredAvailabilityZonesInput() *[]*string
	PrivateIps() *[]*string
	SetPrivateIps(val *[]*string)
	PrivateIpsInput() *[]*string
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
	ResourceRequirements() ElastigroupAwsResourceRequirementsList
	ResourceRequirementsInput() interface{}
	ResourceTagSpecification() ElastigroupAwsResourceTagSpecificationList
	ResourceTagSpecificationInput() interface{}
	RevertToSpot() ElastigroupAwsRevertToSpotOutputReference
	RevertToSpotInput() *ElastigroupAwsRevertToSpot
	ScalingDownPolicy() ElastigroupAwsScalingDownPolicyList
	ScalingDownPolicyInput() interface{}
	ScalingStrategy() ElastigroupAwsScalingStrategyList
	ScalingStrategyInput() interface{}
	ScalingTargetPolicy() ElastigroupAwsScalingTargetPolicyList
	ScalingTargetPolicyInput() interface{}
	ScalingUpPolicy() ElastigroupAwsScalingUpPolicyList
	ScalingUpPolicyInput() interface{}
	ScheduledTask() ElastigroupAwsScheduledTaskList
	ScheduledTaskInput() interface{}
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	ShutdownScript() *string
	SetShutdownScript(val *string)
	ShutdownScriptInput() *string
	Signal() ElastigroupAwsSignalList
	SignalInput() interface{}
	SpotPercentage() *float64
	SetSpotPercentage(val *float64)
	SpotPercentageInput() *float64
	StatefulDeallocation() ElastigroupAwsStatefulDeallocationOutputReference
	StatefulDeallocationInput() *ElastigroupAwsStatefulDeallocation
	StatefulInstanceAction() ElastigroupAwsStatefulInstanceActionList
	StatefulInstanceActionInput() interface{}
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() ElastigroupAwsTagsList
	TagsInput() interface{}
	TargetGroupArns() *[]*string
	SetTargetGroupArns(val *[]*string)
	TargetGroupArnsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatePolicy() ElastigroupAwsUpdatePolicyOutputReference
	UpdatePolicyInput() *ElastigroupAwsUpdatePolicy
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	UtilizeCommitments() interface{}
	SetUtilizeCommitments(val interface{})
	UtilizeCommitmentsInput() interface{}
	UtilizeReservedInstances() interface{}
	SetUtilizeReservedInstances(val interface{})
	UtilizeReservedInstancesInput() interface{}
	WaitForCapacity() *float64
	SetWaitForCapacity(val *float64)
	WaitForCapacityInput() *float64
	WaitForCapacityTimeout() *float64
	SetWaitForCapacityTimeout(val *float64)
	WaitForCapacityTimeoutInput() *float64
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
	PutCpuOptions(value *ElastigroupAwsCpuOptions)
	PutEbsBlockDevice(value interface{})
	PutEphemeralBlockDevice(value interface{})
	PutImages(value interface{})
	PutInstanceTypesWeights(value interface{})
	PutIntegrationBeanstalk(value *ElastigroupAwsIntegrationBeanstalk)
	PutIntegrationCodedeploy(value *ElastigroupAwsIntegrationCodedeploy)
	PutIntegrationDockerSwarm(value *ElastigroupAwsIntegrationDockerSwarm)
	PutIntegrationEcs(value *ElastigroupAwsIntegrationEcs)
	PutIntegrationGitlab(value *ElastigroupAwsIntegrationGitlab)
	PutIntegrationKubernetes(value *ElastigroupAwsIntegrationKubernetes)
	PutIntegrationMesosphere(value *ElastigroupAwsIntegrationMesosphere)
	PutIntegrationNomad(value *ElastigroupAwsIntegrationNomad)
	PutIntegrationRancher(value *ElastigroupAwsIntegrationRancher)
	PutIntegrationRoute53(value *ElastigroupAwsIntegrationRoute53)
	PutItf(value interface{})
	PutLogging(value *ElastigroupAwsLogging)
	PutMetadataOptions(value *ElastigroupAwsMetadataOptions)
	PutMultipleMetrics(value *ElastigroupAwsMultipleMetrics)
	PutNetworkInterface(value interface{})
	PutResourceRequirements(value interface{})
	PutResourceTagSpecification(value interface{})
	PutRevertToSpot(value *ElastigroupAwsRevertToSpot)
	PutScalingDownPolicy(value interface{})
	PutScalingStrategy(value interface{})
	PutScalingTargetPolicy(value interface{})
	PutScalingUpPolicy(value interface{})
	PutScheduledTask(value interface{})
	PutSignal(value interface{})
	PutStatefulDeallocation(value *ElastigroupAwsStatefulDeallocation)
	PutStatefulInstanceAction(value interface{})
	PutTags(value interface{})
	PutUpdatePolicy(value *ElastigroupAwsUpdatePolicy)
	ResetAvailabilityZones()
	ResetBlockDevicesMode()
	ResetCapacityUnit()
	ResetConsiderOdPricing()
	ResetCpuCredits()
	ResetCpuOptions()
	ResetDescription()
	ResetDesiredCapacity()
	ResetDrainingTimeout()
	ResetEbsBlockDevice()
	ResetEbsOptimized()
	ResetElasticIps()
	ResetElasticLoadBalancers()
	ResetEnableMonitoring()
	ResetEphemeralBlockDevice()
	ResetHealthCheckGracePeriod()
	ResetHealthCheckType()
	ResetHealthCheckUnhealthyDurationBeforeReplacement()
	ResetIamInstanceProfile()
	ResetId()
	ResetImageId()
	ResetImages()
	ResetImmediateOdRecoverThreshold()
	ResetInstanceTypesOndemand()
	ResetInstanceTypesPreferredSpot()
	ResetInstanceTypesSpot()
	ResetInstanceTypesWeights()
	ResetIntegrationBeanstalk()
	ResetIntegrationCodedeploy()
	ResetIntegrationDockerSwarm()
	ResetIntegrationEcs()
	ResetIntegrationGitlab()
	ResetIntegrationKubernetes()
	ResetIntegrationMesosphere()
	ResetIntegrationNomad()
	ResetIntegrationRancher()
	ResetIntegrationRoute53()
	ResetItf()
	ResetKeyName()
	ResetLifetimePeriod()
	ResetLogging()
	ResetMaxSize()
	ResetMetadataOptions()
	ResetMinimumInstanceLifetime()
	ResetMinSize()
	ResetMultipleMetrics()
	ResetNetworkInterface()
	ResetOndemandCount()
	ResetOnDemandTypes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistBlockDevices()
	ResetPersistPrivateIp()
	ResetPersistRootDevice()
	ResetPlacementTenancy()
	ResetPreferredAvailabilityZones()
	ResetPrivateIps()
	ResetRegion()
	ResetResourceRequirements()
	ResetResourceTagSpecification()
	ResetRevertToSpot()
	ResetScalingDownPolicy()
	ResetScalingStrategy()
	ResetScalingTargetPolicy()
	ResetScalingUpPolicy()
	ResetScheduledTask()
	ResetShutdownScript()
	ResetSignal()
	ResetSpotPercentage()
	ResetStatefulDeallocation()
	ResetStatefulInstanceAction()
	ResetSubnetIds()
	ResetTags()
	ResetTargetGroupArns()
	ResetUpdatePolicy()
	ResetUserData()
	ResetUtilizeCommitments()
	ResetUtilizeReservedInstances()
	ResetWaitForCapacity()
	ResetWaitForCapacityTimeout()
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

// The jsii proxy struct for ElastigroupAws
type jsiiProxy_ElastigroupAws struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElastigroupAws) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) BlockDevicesMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockDevicesMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) BlockDevicesModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockDevicesModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) CapacityUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) CapacityUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ConsiderOdPricing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"considerOdPricing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ConsiderOdPricingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"considerOdPricingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) CpuCredits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCredits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) CpuCreditsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCreditsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) CpuOptions() ElastigroupAwsCpuOptionsOutputReference {
	var returns ElastigroupAwsCpuOptionsOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) CpuOptionsInput() *ElastigroupAwsCpuOptions {
	var returns *ElastigroupAwsCpuOptions
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) DrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) DrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) EbsBlockDevice() ElastigroupAwsEbsBlockDeviceList {
	var returns ElastigroupAwsEbsBlockDeviceList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ElasticIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ElasticIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ElasticLoadBalancers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticLoadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ElasticLoadBalancersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticLoadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) EnableMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) EnableMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) EphemeralBlockDevice() ElastigroupAwsEphemeralBlockDeviceList {
	var returns ElastigroupAwsEphemeralBlockDeviceList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) FallbackToOndemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) FallbackToOndemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) HealthCheckGracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) HealthCheckUnhealthyDurationBeforeReplacement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckUnhealthyDurationBeforeReplacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) HealthCheckUnhealthyDurationBeforeReplacementInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckUnhealthyDurationBeforeReplacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Images() ElastigroupAwsImagesList {
	var returns ElastigroupAwsImagesList
	_jsii_.Get(
		j,
		"images",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ImagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ImmediateOdRecoverThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"immediateOdRecoverThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ImmediateOdRecoverThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"immediateOdRecoverThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) InstanceTypesOndemand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypesOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) InstanceTypesOndemandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypesOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) InstanceTypesPreferredSpot() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesPreferredSpot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) InstanceTypesPreferredSpotInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesPreferredSpotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) InstanceTypesSpot() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesSpot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) InstanceTypesSpotInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesSpotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) InstanceTypesWeights() ElastigroupAwsInstanceTypesWeightsList {
	var returns ElastigroupAwsInstanceTypesWeightsList
	_jsii_.Get(
		j,
		"instanceTypesWeights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) InstanceTypesWeightsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceTypesWeightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationBeanstalk() ElastigroupAwsIntegrationBeanstalkOutputReference {
	var returns ElastigroupAwsIntegrationBeanstalkOutputReference
	_jsii_.Get(
		j,
		"integrationBeanstalk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationBeanstalkInput() *ElastigroupAwsIntegrationBeanstalk {
	var returns *ElastigroupAwsIntegrationBeanstalk
	_jsii_.Get(
		j,
		"integrationBeanstalkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationCodedeploy() ElastigroupAwsIntegrationCodedeployOutputReference {
	var returns ElastigroupAwsIntegrationCodedeployOutputReference
	_jsii_.Get(
		j,
		"integrationCodedeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationCodedeployInput() *ElastigroupAwsIntegrationCodedeploy {
	var returns *ElastigroupAwsIntegrationCodedeploy
	_jsii_.Get(
		j,
		"integrationCodedeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationDockerSwarm() ElastigroupAwsIntegrationDockerSwarmOutputReference {
	var returns ElastigroupAwsIntegrationDockerSwarmOutputReference
	_jsii_.Get(
		j,
		"integrationDockerSwarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationDockerSwarmInput() *ElastigroupAwsIntegrationDockerSwarm {
	var returns *ElastigroupAwsIntegrationDockerSwarm
	_jsii_.Get(
		j,
		"integrationDockerSwarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationEcs() ElastigroupAwsIntegrationEcsOutputReference {
	var returns ElastigroupAwsIntegrationEcsOutputReference
	_jsii_.Get(
		j,
		"integrationEcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationEcsInput() *ElastigroupAwsIntegrationEcs {
	var returns *ElastigroupAwsIntegrationEcs
	_jsii_.Get(
		j,
		"integrationEcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationGitlab() ElastigroupAwsIntegrationGitlabOutputReference {
	var returns ElastigroupAwsIntegrationGitlabOutputReference
	_jsii_.Get(
		j,
		"integrationGitlab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationGitlabInput() *ElastigroupAwsIntegrationGitlab {
	var returns *ElastigroupAwsIntegrationGitlab
	_jsii_.Get(
		j,
		"integrationGitlabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationKubernetes() ElastigroupAwsIntegrationKubernetesOutputReference {
	var returns ElastigroupAwsIntegrationKubernetesOutputReference
	_jsii_.Get(
		j,
		"integrationKubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationKubernetesInput() *ElastigroupAwsIntegrationKubernetes {
	var returns *ElastigroupAwsIntegrationKubernetes
	_jsii_.Get(
		j,
		"integrationKubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationMesosphere() ElastigroupAwsIntegrationMesosphereOutputReference {
	var returns ElastigroupAwsIntegrationMesosphereOutputReference
	_jsii_.Get(
		j,
		"integrationMesosphere",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationMesosphereInput() *ElastigroupAwsIntegrationMesosphere {
	var returns *ElastigroupAwsIntegrationMesosphere
	_jsii_.Get(
		j,
		"integrationMesosphereInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationNomad() ElastigroupAwsIntegrationNomadOutputReference {
	var returns ElastigroupAwsIntegrationNomadOutputReference
	_jsii_.Get(
		j,
		"integrationNomad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationNomadInput() *ElastigroupAwsIntegrationNomad {
	var returns *ElastigroupAwsIntegrationNomad
	_jsii_.Get(
		j,
		"integrationNomadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationRancher() ElastigroupAwsIntegrationRancherOutputReference {
	var returns ElastigroupAwsIntegrationRancherOutputReference
	_jsii_.Get(
		j,
		"integrationRancher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationRancherInput() *ElastigroupAwsIntegrationRancher {
	var returns *ElastigroupAwsIntegrationRancher
	_jsii_.Get(
		j,
		"integrationRancherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationRoute53() ElastigroupAwsIntegrationRoute53OutputReference {
	var returns ElastigroupAwsIntegrationRoute53OutputReference
	_jsii_.Get(
		j,
		"integrationRoute53",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) IntegrationRoute53Input() *ElastigroupAwsIntegrationRoute53 {
	var returns *ElastigroupAwsIntegrationRoute53
	_jsii_.Get(
		j,
		"integrationRoute53Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Itf() ElastigroupAwsItfList {
	var returns ElastigroupAwsItfList
	_jsii_.Get(
		j,
		"itf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ItfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) LifetimePeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifetimePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) LifetimePeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifetimePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Logging() ElastigroupAwsLoggingOutputReference {
	var returns ElastigroupAwsLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) LoggingInput() *ElastigroupAwsLogging {
	var returns *ElastigroupAwsLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MetadataOptions() ElastigroupAwsMetadataOptionsOutputReference {
	var returns ElastigroupAwsMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MetadataOptionsInput() *ElastigroupAwsMetadataOptions {
	var returns *ElastigroupAwsMetadataOptions
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MinimumInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MinimumInstanceLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInstanceLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MultipleMetrics() ElastigroupAwsMultipleMetricsOutputReference {
	var returns ElastigroupAwsMultipleMetricsOutputReference
	_jsii_.Get(
		j,
		"multipleMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) MultipleMetricsInput() *ElastigroupAwsMultipleMetrics {
	var returns *ElastigroupAwsMultipleMetrics
	_jsii_.Get(
		j,
		"multipleMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) NetworkInterface() ElastigroupAwsNetworkInterfaceList {
	var returns ElastigroupAwsNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) OndemandCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ondemandCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) OndemandCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ondemandCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) OnDemandTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onDemandTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) OnDemandTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onDemandTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Orientation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orientation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) OrientationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orientationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PersistBlockDevices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistBlockDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PersistBlockDevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistBlockDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PersistPrivateIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistPrivateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PersistPrivateIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistPrivateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PersistRootDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistRootDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PersistRootDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistRootDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PlacementTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PlacementTenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PreferredAvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredAvailabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PreferredAvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredAvailabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PrivateIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) PrivateIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Product() *string {
	var returns *string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ProductInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ResourceRequirements() ElastigroupAwsResourceRequirementsList {
	var returns ElastigroupAwsResourceRequirementsList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ResourceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ResourceTagSpecification() ElastigroupAwsResourceTagSpecificationList {
	var returns ElastigroupAwsResourceTagSpecificationList
	_jsii_.Get(
		j,
		"resourceTagSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ResourceTagSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) RevertToSpot() ElastigroupAwsRevertToSpotOutputReference {
	var returns ElastigroupAwsRevertToSpotOutputReference
	_jsii_.Get(
		j,
		"revertToSpot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) RevertToSpotInput() *ElastigroupAwsRevertToSpot {
	var returns *ElastigroupAwsRevertToSpot
	_jsii_.Get(
		j,
		"revertToSpotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScalingDownPolicy() ElastigroupAwsScalingDownPolicyList {
	var returns ElastigroupAwsScalingDownPolicyList
	_jsii_.Get(
		j,
		"scalingDownPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScalingDownPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingDownPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScalingStrategy() ElastigroupAwsScalingStrategyList {
	var returns ElastigroupAwsScalingStrategyList
	_jsii_.Get(
		j,
		"scalingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScalingStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScalingTargetPolicy() ElastigroupAwsScalingTargetPolicyList {
	var returns ElastigroupAwsScalingTargetPolicyList
	_jsii_.Get(
		j,
		"scalingTargetPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScalingTargetPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingTargetPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScalingUpPolicy() ElastigroupAwsScalingUpPolicyList {
	var returns ElastigroupAwsScalingUpPolicyList
	_jsii_.Get(
		j,
		"scalingUpPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScalingUpPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingUpPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScheduledTask() ElastigroupAwsScheduledTaskList {
	var returns ElastigroupAwsScheduledTaskList
	_jsii_.Get(
		j,
		"scheduledTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ScheduledTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ShutdownScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) ShutdownScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Signal() ElastigroupAwsSignalList {
	var returns ElastigroupAwsSignalList
	_jsii_.Get(
		j,
		"signal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) SignalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) SpotPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) SpotPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) StatefulDeallocation() ElastigroupAwsStatefulDeallocationOutputReference {
	var returns ElastigroupAwsStatefulDeallocationOutputReference
	_jsii_.Get(
		j,
		"statefulDeallocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) StatefulDeallocationInput() *ElastigroupAwsStatefulDeallocation {
	var returns *ElastigroupAwsStatefulDeallocation
	_jsii_.Get(
		j,
		"statefulDeallocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) StatefulInstanceAction() ElastigroupAwsStatefulInstanceActionList {
	var returns ElastigroupAwsStatefulInstanceActionList
	_jsii_.Get(
		j,
		"statefulInstanceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) StatefulInstanceActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulInstanceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) Tags() ElastigroupAwsTagsList {
	var returns ElastigroupAwsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) TargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) TargetGroupArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) UpdatePolicy() ElastigroupAwsUpdatePolicyOutputReference {
	var returns ElastigroupAwsUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) UpdatePolicyInput() *ElastigroupAwsUpdatePolicy {
	var returns *ElastigroupAwsUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) UtilizeCommitments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeCommitments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) UtilizeCommitmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeCommitmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) UtilizeReservedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeReservedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) UtilizeReservedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeReservedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) WaitForCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) WaitForCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) WaitForCapacityTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForCapacityTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAws) WaitForCapacityTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForCapacityTimeoutInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.186.0/docs/resources/elastigroup_aws spotinst_elastigroup_aws} Resource.
func NewElastigroupAws(scope constructs.Construct, id *string, config *ElastigroupAwsConfig) ElastigroupAws {
	_init_.Initialize()

	if err := validateNewElastigroupAwsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAws{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAws",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.186.0/docs/resources/elastigroup_aws spotinst_elastigroup_aws} Resource.
func NewElastigroupAws_Override(e ElastigroupAws, scope constructs.Construct, id *string, config *ElastigroupAwsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAws",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetBlockDevicesMode(val *string) {
	if err := j.validateSetBlockDevicesModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockDevicesMode",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetCapacityUnit(val *string) {
	if err := j.validateSetCapacityUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityUnit",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetConsiderOdPricing(val interface{}) {
	if err := j.validateSetConsiderOdPricingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"considerOdPricing",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetCpuCredits(val *string) {
	if err := j.validateSetCpuCreditsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCredits",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetDrainingTimeout(val *float64) {
	if err := j.validateSetDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainingTimeout",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetElasticIps(val *[]*string) {
	if err := j.validateSetElasticIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticIps",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetElasticLoadBalancers(val *[]*string) {
	if err := j.validateSetElasticLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticLoadBalancers",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetEnableMonitoring(val interface{}) {
	if err := j.validateSetEnableMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMonitoring",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetFallbackToOndemand(val interface{}) {
	if err := j.validateSetFallbackToOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackToOndemand",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetHealthCheckGracePeriod(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriod",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetHealthCheckUnhealthyDurationBeforeReplacement(val *float64) {
	if err := j.validateSetHealthCheckUnhealthyDurationBeforeReplacementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckUnhealthyDurationBeforeReplacement",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetImmediateOdRecoverThreshold(val *float64) {
	if err := j.validateSetImmediateOdRecoverThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"immediateOdRecoverThreshold",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetInstanceTypesOndemand(val *string) {
	if err := j.validateSetInstanceTypesOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypesOndemand",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetInstanceTypesPreferredSpot(val *[]*string) {
	if err := j.validateSetInstanceTypesPreferredSpotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypesPreferredSpot",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetInstanceTypesSpot(val *[]*string) {
	if err := j.validateSetInstanceTypesSpotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypesSpot",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetLifetimePeriod(val *string) {
	if err := j.validateSetLifetimePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetimePeriod",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetMinimumInstanceLifetime(val *float64) {
	if err := j.validateSetMinimumInstanceLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumInstanceLifetime",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetOndemandCount(val *float64) {
	if err := j.validateSetOndemandCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ondemandCount",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetOnDemandTypes(val *[]*string) {
	if err := j.validateSetOnDemandTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandTypes",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetOrientation(val *string) {
	if err := j.validateSetOrientationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orientation",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetPersistBlockDevices(val interface{}) {
	if err := j.validateSetPersistBlockDevicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistBlockDevices",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetPersistPrivateIp(val interface{}) {
	if err := j.validateSetPersistPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistPrivateIp",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetPersistRootDevice(val interface{}) {
	if err := j.validateSetPersistRootDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistRootDevice",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetPlacementTenancy(val *string) {
	if err := j.validateSetPlacementTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementTenancy",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetPreferredAvailabilityZones(val *[]*string) {
	if err := j.validateSetPreferredAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredAvailabilityZones",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetPrivateIps(val *[]*string) {
	if err := j.validateSetPrivateIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIps",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetProduct(val *string) {
	if err := j.validateSetProductParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"product",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetShutdownScript(val *string) {
	if err := j.validateSetShutdownScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownScript",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetSpotPercentage(val *float64) {
	if err := j.validateSetSpotPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetTargetGroupArns(val *[]*string) {
	if err := j.validateSetTargetGroupArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArns",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetUtilizeCommitments(val interface{}) {
	if err := j.validateSetUtilizeCommitmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utilizeCommitments",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetUtilizeReservedInstances(val interface{}) {
	if err := j.validateSetUtilizeReservedInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utilizeReservedInstances",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetWaitForCapacity(val *float64) {
	if err := j.validateSetWaitForCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAws)SetWaitForCapacityTimeout(val *float64) {
	if err := j.validateSetWaitForCapacityTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForCapacityTimeout",
		val,
	)
}

// Generates CDKTF code for importing a ElastigroupAws resource upon running "cdktf plan <stack-name>".
func ElastigroupAws_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElastigroupAws_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAws",
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
func ElastigroupAws_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAws_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAws",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupAws_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAws_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAws",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupAws_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAws_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAws",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElastigroupAws_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAws",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElastigroupAws) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElastigroupAws) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElastigroupAws) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAws) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAws) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAws) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAws) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAws) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAws) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAws) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAws) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAws) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAws) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElastigroupAws) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAws) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastigroupAws) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElastigroupAws) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastigroupAws) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutCpuOptions(value *ElastigroupAwsCpuOptions) {
	if err := e.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutEbsBlockDevice(value interface{}) {
	if err := e.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutEphemeralBlockDevice(value interface{}) {
	if err := e.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutImages(value interface{}) {
	if err := e.validatePutImagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putImages",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutInstanceTypesWeights(value interface{}) {
	if err := e.validatePutInstanceTypesWeightsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInstanceTypesWeights",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationBeanstalk(value *ElastigroupAwsIntegrationBeanstalk) {
	if err := e.validatePutIntegrationBeanstalkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationBeanstalk",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationCodedeploy(value *ElastigroupAwsIntegrationCodedeploy) {
	if err := e.validatePutIntegrationCodedeployParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationCodedeploy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationDockerSwarm(value *ElastigroupAwsIntegrationDockerSwarm) {
	if err := e.validatePutIntegrationDockerSwarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationDockerSwarm",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationEcs(value *ElastigroupAwsIntegrationEcs) {
	if err := e.validatePutIntegrationEcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationEcs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationGitlab(value *ElastigroupAwsIntegrationGitlab) {
	if err := e.validatePutIntegrationGitlabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationGitlab",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationKubernetes(value *ElastigroupAwsIntegrationKubernetes) {
	if err := e.validatePutIntegrationKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationKubernetes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationMesosphere(value *ElastigroupAwsIntegrationMesosphere) {
	if err := e.validatePutIntegrationMesosphereParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationMesosphere",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationNomad(value *ElastigroupAwsIntegrationNomad) {
	if err := e.validatePutIntegrationNomadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationNomad",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationRancher(value *ElastigroupAwsIntegrationRancher) {
	if err := e.validatePutIntegrationRancherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationRancher",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutIntegrationRoute53(value *ElastigroupAwsIntegrationRoute53) {
	if err := e.validatePutIntegrationRoute53Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationRoute53",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutItf(value interface{}) {
	if err := e.validatePutItfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putItf",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutLogging(value *ElastigroupAwsLogging) {
	if err := e.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogging",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutMetadataOptions(value *ElastigroupAwsMetadataOptions) {
	if err := e.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutMultipleMetrics(value *ElastigroupAwsMultipleMetrics) {
	if err := e.validatePutMultipleMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMultipleMetrics",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutNetworkInterface(value interface{}) {
	if err := e.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutResourceRequirements(value interface{}) {
	if err := e.validatePutResourceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putResourceRequirements",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutResourceTagSpecification(value interface{}) {
	if err := e.validatePutResourceTagSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putResourceTagSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutRevertToSpot(value *ElastigroupAwsRevertToSpot) {
	if err := e.validatePutRevertToSpotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRevertToSpot",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutScalingDownPolicy(value interface{}) {
	if err := e.validatePutScalingDownPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingDownPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutScalingStrategy(value interface{}) {
	if err := e.validatePutScalingStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingStrategy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutScalingTargetPolicy(value interface{}) {
	if err := e.validatePutScalingTargetPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingTargetPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutScalingUpPolicy(value interface{}) {
	if err := e.validatePutScalingUpPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingUpPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutScheduledTask(value interface{}) {
	if err := e.validatePutScheduledTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScheduledTask",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutSignal(value interface{}) {
	if err := e.validatePutSignalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSignal",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutStatefulDeallocation(value *ElastigroupAwsStatefulDeallocation) {
	if err := e.validatePutStatefulDeallocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStatefulDeallocation",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutStatefulInstanceAction(value interface{}) {
	if err := e.validatePutStatefulInstanceActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStatefulInstanceAction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) PutUpdatePolicy(value *ElastigroupAwsUpdatePolicy) {
	if err := e.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetBlockDevicesMode() {
	_jsii_.InvokeVoid(
		e,
		"resetBlockDevicesMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetCapacityUnit() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityUnit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetConsiderOdPricing() {
	_jsii_.InvokeVoid(
		e,
		"resetConsiderOdPricing",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetCpuCredits() {
	_jsii_.InvokeVoid(
		e,
		"resetCpuCredits",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetDrainingTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetDrainingTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetElasticIps() {
	_jsii_.InvokeVoid(
		e,
		"resetElasticIps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetElasticLoadBalancers() {
	_jsii_.InvokeVoid(
		e,
		"resetElasticLoadBalancers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetEnableMonitoring() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableMonitoring",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		e,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetHealthCheckGracePeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckGracePeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetHealthCheckUnhealthyDurationBeforeReplacement() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckUnhealthyDurationBeforeReplacement",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		e,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetImageId() {
	_jsii_.InvokeVoid(
		e,
		"resetImageId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetImages() {
	_jsii_.InvokeVoid(
		e,
		"resetImages",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetImmediateOdRecoverThreshold() {
	_jsii_.InvokeVoid(
		e,
		"resetImmediateOdRecoverThreshold",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetInstanceTypesOndemand() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesOndemand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetInstanceTypesPreferredSpot() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesPreferredSpot",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetInstanceTypesSpot() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesSpot",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetInstanceTypesWeights() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesWeights",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationBeanstalk() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationBeanstalk",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationCodedeploy() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationCodedeploy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationDockerSwarm() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationDockerSwarm",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationEcs() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationEcs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationGitlab() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationGitlab",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationKubernetes() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationKubernetes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationMesosphere() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationMesosphere",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationNomad() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationNomad",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationRancher() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationRancher",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetIntegrationRoute53() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationRoute53",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetItf() {
	_jsii_.InvokeVoid(
		e,
		"resetItf",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetKeyName() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetLifetimePeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetLifetimePeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetLogging() {
	_jsii_.InvokeVoid(
		e,
		"resetLogging",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetMaxSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetMinimumInstanceLifetime() {
	_jsii_.InvokeVoid(
		e,
		"resetMinimumInstanceLifetime",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetMinSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMinSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetMultipleMetrics() {
	_jsii_.InvokeVoid(
		e,
		"resetMultipleMetrics",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetOndemandCount() {
	_jsii_.InvokeVoid(
		e,
		"resetOndemandCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetOnDemandTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetPersistBlockDevices() {
	_jsii_.InvokeVoid(
		e,
		"resetPersistBlockDevices",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetPersistPrivateIp() {
	_jsii_.InvokeVoid(
		e,
		"resetPersistPrivateIp",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetPersistRootDevice() {
	_jsii_.InvokeVoid(
		e,
		"resetPersistRootDevice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetPlacementTenancy() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacementTenancy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetPreferredAvailabilityZones() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredAvailabilityZones",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetPrivateIps() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateIps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetResourceRequirements() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceRequirements",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetResourceTagSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceTagSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetRevertToSpot() {
	_jsii_.InvokeVoid(
		e,
		"resetRevertToSpot",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetScalingDownPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingDownPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetScalingStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetScalingTargetPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingTargetPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetScalingUpPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingUpPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetScheduledTask() {
	_jsii_.InvokeVoid(
		e,
		"resetScheduledTask",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetShutdownScript() {
	_jsii_.InvokeVoid(
		e,
		"resetShutdownScript",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetSignal() {
	_jsii_.InvokeVoid(
		e,
		"resetSignal",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetSpotPercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotPercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetStatefulDeallocation() {
	_jsii_.InvokeVoid(
		e,
		"resetStatefulDeallocation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetStatefulInstanceAction() {
	_jsii_.InvokeVoid(
		e,
		"resetStatefulInstanceAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetTargetGroupArns() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetGroupArns",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetUserData() {
	_jsii_.InvokeVoid(
		e,
		"resetUserData",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetUtilizeCommitments() {
	_jsii_.InvokeVoid(
		e,
		"resetUtilizeCommitments",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetUtilizeReservedInstances() {
	_jsii_.InvokeVoid(
		e,
		"resetUtilizeReservedInstances",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetWaitForCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetWaitForCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) ResetWaitForCapacityTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetWaitForCapacityTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAws) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAws) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAws) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAws) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAws) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAws) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

