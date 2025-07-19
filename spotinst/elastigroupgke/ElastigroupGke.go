// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/elastigroupgke/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/elastigroup_gke spotinst_elastigroup_gke}.
type ElastigroupGke interface {
	cdktf.TerraformResource
	BackendServices() ElastigroupGkeBackendServicesList
	BackendServicesInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterZoneName() *string
	SetClusterZoneName(val *string)
	ClusterZoneNameInput() *string
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
	DesiredCapacity() *float64
	SetDesiredCapacity(val *float64)
	DesiredCapacityInput() *float64
	Disk() ElastigroupGkeDiskList
	DiskInput() interface{}
	DrainingTimeout() *float64
	SetDrainingTimeout(val *float64)
	DrainingTimeoutInput() *float64
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
	Gpu() ElastigroupGkeGpuList
	GpuInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceNamePrefix() *string
	SetInstanceNamePrefix(val *string)
	InstanceNamePrefixInput() *string
	InstanceTypesCustom() ElastigroupGkeInstanceTypesCustomList
	InstanceTypesCustomInput() interface{}
	InstanceTypesOndemand() *string
	SetInstanceTypesOndemand(val *string)
	InstanceTypesOndemandInput() *string
	InstanceTypesPreemptible() *[]*string
	SetInstanceTypesPreemptible(val *[]*string)
	InstanceTypesPreemptibleInput() *[]*string
	IntegrationDockerSwarm() ElastigroupGkeIntegrationDockerSwarmOutputReference
	IntegrationDockerSwarmInput() *ElastigroupGkeIntegrationDockerSwarm
	IntegrationGke() ElastigroupGkeIntegrationGkeOutputReference
	IntegrationGkeInput() *ElastigroupGkeIntegrationGke
	IpForwarding() interface{}
	SetIpForwarding(val interface{})
	IpForwardingInput() interface{}
	Labels() ElastigroupGkeLabelsList
	LabelsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	Metadata() ElastigroupGkeMetadataList
	MetadataInput() interface{}
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() ElastigroupGkeNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	NodeImage() *string
	SetNodeImage(val *string)
	NodeImageInput() *string
	OndemandCount() *float64
	SetOndemandCount(val *float64)
	OndemandCountInput() *float64
	OptimizationWindows() *[]*string
	SetOptimizationWindows(val *[]*string)
	OptimizationWindowsInput() *[]*string
	PreemptiblePercentage() *float64
	SetPreemptiblePercentage(val *float64)
	PreemptiblePercentageInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisioningModel() *string
	SetProvisioningModel(val *string)
	ProvisioningModelInput() *string
	// Experimental.
	RawOverrides() interface{}
	RevertToPreemptible() ElastigroupGkeRevertToPreemptibleList
	RevertToPreemptibleInput() interface{}
	ScalingDownPolicy() ElastigroupGkeScalingDownPolicyList
	ScalingDownPolicyInput() interface{}
	ScalingUpPolicy() ElastigroupGkeScalingUpPolicyList
	ScalingUpPolicyInput() interface{}
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ShieldedInstanceConfig() ElastigroupGkeShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *ElastigroupGkeShieldedInstanceConfig
	ShouldUtilizeCommitments() interface{}
	SetShouldUtilizeCommitments(val interface{})
	ShouldUtilizeCommitmentsInput() interface{}
	ShutdownScript() *string
	SetShutdownScript(val *string)
	ShutdownScriptInput() *string
	StartupScript() *string
	SetStartupScript(val *string)
	StartupScriptInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
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
	PutBackendServices(value interface{})
	PutDisk(value interface{})
	PutGpu(value interface{})
	PutInstanceTypesCustom(value interface{})
	PutIntegrationDockerSwarm(value *ElastigroupGkeIntegrationDockerSwarm)
	PutIntegrationGke(value *ElastigroupGkeIntegrationGke)
	PutLabels(value interface{})
	PutMetadata(value interface{})
	PutNetworkInterface(value interface{})
	PutRevertToPreemptible(value interface{})
	PutScalingDownPolicy(value interface{})
	PutScalingUpPolicy(value interface{})
	PutShieldedInstanceConfig(value *ElastigroupGkeShieldedInstanceConfig)
	ResetBackendServices()
	ResetClusterId()
	ResetDisk()
	ResetDrainingTimeout()
	ResetFallbackToOndemand()
	ResetGpu()
	ResetId()
	ResetInstanceNamePrefix()
	ResetInstanceTypesCustom()
	ResetInstanceTypesOndemand()
	ResetInstanceTypesPreemptible()
	ResetIntegrationDockerSwarm()
	ResetIntegrationGke()
	ResetIpForwarding()
	ResetLabels()
	ResetMaxSize()
	ResetMetadata()
	ResetMinCpuPlatform()
	ResetMinSize()
	ResetNetworkInterface()
	ResetNodeImage()
	ResetOndemandCount()
	ResetOptimizationWindows()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreemptiblePercentage()
	ResetProvisioningModel()
	ResetRevertToPreemptible()
	ResetScalingDownPolicy()
	ResetScalingUpPolicy()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
	ResetShouldUtilizeCommitments()
	ResetShutdownScript()
	ResetStartupScript()
	ResetTags()
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

// The jsii proxy struct for ElastigroupGke
type jsiiProxy_ElastigroupGke struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElastigroupGke) BackendServices() ElastigroupGkeBackendServicesList {
	var returns ElastigroupGkeBackendServicesList
	_jsii_.Get(
		j,
		"backendServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) BackendServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ClusterZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ClusterZoneNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterZoneNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Disk() ElastigroupGkeDiskList {
	var returns ElastigroupGkeDiskList
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) DiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) DrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) DrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) FallbackToOndemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) FallbackToOndemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Gpu() ElastigroupGkeGpuList {
	var returns ElastigroupGkeGpuList
	_jsii_.Get(
		j,
		"gpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) GpuInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) InstanceNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) InstanceNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) InstanceTypesCustom() ElastigroupGkeInstanceTypesCustomList {
	var returns ElastigroupGkeInstanceTypesCustomList
	_jsii_.Get(
		j,
		"instanceTypesCustom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) InstanceTypesCustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceTypesCustomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) InstanceTypesOndemand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypesOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) InstanceTypesOndemandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypesOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) InstanceTypesPreemptible() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesPreemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) InstanceTypesPreemptibleInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesPreemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) IntegrationDockerSwarm() ElastigroupGkeIntegrationDockerSwarmOutputReference {
	var returns ElastigroupGkeIntegrationDockerSwarmOutputReference
	_jsii_.Get(
		j,
		"integrationDockerSwarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) IntegrationDockerSwarmInput() *ElastigroupGkeIntegrationDockerSwarm {
	var returns *ElastigroupGkeIntegrationDockerSwarm
	_jsii_.Get(
		j,
		"integrationDockerSwarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) IntegrationGke() ElastigroupGkeIntegrationGkeOutputReference {
	var returns ElastigroupGkeIntegrationGkeOutputReference
	_jsii_.Get(
		j,
		"integrationGke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) IntegrationGkeInput() *ElastigroupGkeIntegrationGke {
	var returns *ElastigroupGkeIntegrationGke
	_jsii_.Get(
		j,
		"integrationGkeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) IpForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) IpForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Labels() ElastigroupGkeLabelsList {
	var returns ElastigroupGkeLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Metadata() ElastigroupGkeMetadataList {
	var returns ElastigroupGkeMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) NetworkInterface() ElastigroupGkeNetworkInterfaceList {
	var returns ElastigroupGkeNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) NodeImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) NodeImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) OndemandCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ondemandCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) OndemandCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ondemandCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) OptimizationWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optimizationWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) OptimizationWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optimizationWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) PreemptiblePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preemptiblePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) PreemptiblePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preemptiblePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ProvisioningModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ProvisioningModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) RevertToPreemptible() ElastigroupGkeRevertToPreemptibleList {
	var returns ElastigroupGkeRevertToPreemptibleList
	_jsii_.Get(
		j,
		"revertToPreemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) RevertToPreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revertToPreemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ScalingDownPolicy() ElastigroupGkeScalingDownPolicyList {
	var returns ElastigroupGkeScalingDownPolicyList
	_jsii_.Get(
		j,
		"scalingDownPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ScalingDownPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingDownPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ScalingUpPolicy() ElastigroupGkeScalingUpPolicyList {
	var returns ElastigroupGkeScalingUpPolicyList
	_jsii_.Get(
		j,
		"scalingUpPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ScalingUpPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingUpPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ShieldedInstanceConfig() ElastigroupGkeShieldedInstanceConfigOutputReference {
	var returns ElastigroupGkeShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ShieldedInstanceConfigInput() *ElastigroupGkeShieldedInstanceConfig {
	var returns *ElastigroupGkeShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ShouldUtilizeCommitments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldUtilizeCommitments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ShouldUtilizeCommitmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldUtilizeCommitmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ShutdownScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) ShutdownScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) StartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) StartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGke) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/elastigroup_gke spotinst_elastigroup_gke} Resource.
func NewElastigroupGke(scope constructs.Construct, id *string, config *ElastigroupGkeConfig) ElastigroupGke {
	_init_.Initialize()

	if err := validateNewElastigroupGkeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupGke{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGke",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/elastigroup_gke spotinst_elastigroup_gke} Resource.
func NewElastigroupGke_Override(e ElastigroupGke, scope constructs.Construct, id *string, config *ElastigroupGkeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGke",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetClusterZoneName(val *string) {
	if err := j.validateSetClusterZoneNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterZoneName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetDrainingTimeout(val *float64) {
	if err := j.validateSetDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainingTimeout",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetFallbackToOndemand(val interface{}) {
	if err := j.validateSetFallbackToOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackToOndemand",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetInstanceNamePrefix(val *string) {
	if err := j.validateSetInstanceNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceNamePrefix",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetInstanceTypesOndemand(val *string) {
	if err := j.validateSetInstanceTypesOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypesOndemand",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetInstanceTypesPreemptible(val *[]*string) {
	if err := j.validateSetInstanceTypesPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypesPreemptible",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetIpForwarding(val interface{}) {
	if err := j.validateSetIpForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipForwarding",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetNodeImage(val *string) {
	if err := j.validateSetNodeImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeImage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetOndemandCount(val *float64) {
	if err := j.validateSetOndemandCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ondemandCount",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetOptimizationWindows(val *[]*string) {
	if err := j.validateSetOptimizationWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optimizationWindows",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetPreemptiblePercentage(val *float64) {
	if err := j.validateSetPreemptiblePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptiblePercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetProvisioningModel(val *string) {
	if err := j.validateSetProvisioningModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningModel",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetShouldUtilizeCommitments(val interface{}) {
	if err := j.validateSetShouldUtilizeCommitmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldUtilizeCommitments",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetShutdownScript(val *string) {
	if err := j.validateSetShutdownScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownScript",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetStartupScript(val *string) {
	if err := j.validateSetStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupScript",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGke)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a ElastigroupGke resource upon running "cdktf plan <stack-name>".
func ElastigroupGke_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElastigroupGke_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGke",
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
func ElastigroupGke_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupGke_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGke",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupGke_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupGke_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGke",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupGke_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupGke_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGke",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElastigroupGke_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGke",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElastigroupGke) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElastigroupGke) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElastigroupGke) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupGke) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupGke) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupGke) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupGke) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupGke) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupGke) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupGke) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupGke) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupGke) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGke) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElastigroupGke) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupGke) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastigroupGke) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElastigroupGke) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastigroupGke) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutBackendServices(value interface{}) {
	if err := e.validatePutBackendServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBackendServices",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutDisk(value interface{}) {
	if err := e.validatePutDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDisk",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutGpu(value interface{}) {
	if err := e.validatePutGpuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putGpu",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutInstanceTypesCustom(value interface{}) {
	if err := e.validatePutInstanceTypesCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInstanceTypesCustom",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutIntegrationDockerSwarm(value *ElastigroupGkeIntegrationDockerSwarm) {
	if err := e.validatePutIntegrationDockerSwarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationDockerSwarm",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutIntegrationGke(value *ElastigroupGkeIntegrationGke) {
	if err := e.validatePutIntegrationGkeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationGke",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutLabels(value interface{}) {
	if err := e.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLabels",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutMetadata(value interface{}) {
	if err := e.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMetadata",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutNetworkInterface(value interface{}) {
	if err := e.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutRevertToPreemptible(value interface{}) {
	if err := e.validatePutRevertToPreemptibleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRevertToPreemptible",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutScalingDownPolicy(value interface{}) {
	if err := e.validatePutScalingDownPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingDownPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutScalingUpPolicy(value interface{}) {
	if err := e.validatePutScalingUpPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingUpPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) PutShieldedInstanceConfig(value *ElastigroupGkeShieldedInstanceConfig) {
	if err := e.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetBackendServices() {
	_jsii_.InvokeVoid(
		e,
		"resetBackendServices",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetClusterId() {
	_jsii_.InvokeVoid(
		e,
		"resetClusterId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetDisk() {
	_jsii_.InvokeVoid(
		e,
		"resetDisk",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetDrainingTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetDrainingTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetFallbackToOndemand() {
	_jsii_.InvokeVoid(
		e,
		"resetFallbackToOndemand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetGpu() {
	_jsii_.InvokeVoid(
		e,
		"resetGpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetInstanceNamePrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceNamePrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetInstanceTypesCustom() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesCustom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetInstanceTypesOndemand() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesOndemand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetInstanceTypesPreemptible() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesPreemptible",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetIntegrationDockerSwarm() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationDockerSwarm",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetIntegrationGke() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationGke",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetIpForwarding() {
	_jsii_.InvokeVoid(
		e,
		"resetIpForwarding",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetMaxSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetMetadata() {
	_jsii_.InvokeVoid(
		e,
		"resetMetadata",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		e,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetMinSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMinSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetNodeImage() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeImage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetOndemandCount() {
	_jsii_.InvokeVoid(
		e,
		"resetOndemandCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetOptimizationWindows() {
	_jsii_.InvokeVoid(
		e,
		"resetOptimizationWindows",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetPreemptiblePercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetPreemptiblePercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetProvisioningModel() {
	_jsii_.InvokeVoid(
		e,
		"resetProvisioningModel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetRevertToPreemptible() {
	_jsii_.InvokeVoid(
		e,
		"resetRevertToPreemptible",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetScalingDownPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingDownPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetScalingUpPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingUpPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetShouldUtilizeCommitments() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldUtilizeCommitments",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetShutdownScript() {
	_jsii_.InvokeVoid(
		e,
		"resetShutdownScript",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetStartupScript() {
	_jsii_.InvokeVoid(
		e,
		"resetStartupScript",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGke) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGke) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGke) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGke) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGke) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGke) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

