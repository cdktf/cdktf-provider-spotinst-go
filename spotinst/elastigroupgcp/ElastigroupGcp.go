package elastigroupgcp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v7/elastigroupgcp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/elastigroup_gcp spotinst_elastigroup_gcp}.
type ElastigroupGcp interface {
	cdktf.TerraformResource
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	BackendServices() ElastigroupGcpBackendServicesList
	BackendServicesInput() interface{}
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DesiredCapacity() *float64
	SetDesiredCapacity(val *float64)
	DesiredCapacityInput() *float64
	Disk() ElastigroupGcpDiskList
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
	Gpu() ElastigroupGcpGpuList
	GpuInput() interface{}
	HealthCheckGracePeriod() *float64
	SetHealthCheckGracePeriod(val *float64)
	HealthCheckGracePeriodInput() *float64
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	HealthCheckTypeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceNamePrefix() *string
	SetInstanceNamePrefix(val *string)
	InstanceNamePrefixInput() *string
	InstanceTypesCustom() ElastigroupGcpInstanceTypesCustomList
	InstanceTypesCustomInput() interface{}
	InstanceTypesOndemand() *string
	SetInstanceTypesOndemand(val *string)
	InstanceTypesOndemandInput() *string
	InstanceTypesPreemptible() *[]*string
	SetInstanceTypesPreemptible(val *[]*string)
	InstanceTypesPreemptibleInput() *[]*string
	IntegrationDockerSwarm() ElastigroupGcpIntegrationDockerSwarmOutputReference
	IntegrationDockerSwarmInput() *ElastigroupGcpIntegrationDockerSwarm
	IntegrationGke() ElastigroupGcpIntegrationGkeOutputReference
	IntegrationGkeInput() *ElastigroupGcpIntegrationGke
	IpForwarding() interface{}
	SetIpForwarding(val interface{})
	IpForwardingInput() interface{}
	Labels() ElastigroupGcpLabelsList
	LabelsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	Metadata() ElastigroupGcpMetadataList
	MetadataInput() interface{}
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() ElastigroupGcpNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	OndemandCount() *float64
	SetOndemandCount(val *float64)
	OndemandCountInput() *float64
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
	ScalingDownPolicy() ElastigroupGcpScalingDownPolicyList
	ScalingDownPolicyInput() interface{}
	ScalingUpPolicy() ElastigroupGcpScalingUpPolicyList
	ScalingUpPolicyInput() interface{}
	ScheduledTask() ElastigroupGcpScheduledTaskList
	ScheduledTaskInput() interface{}
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ShutdownScript() *string
	SetShutdownScript(val *string)
	ShutdownScriptInput() *string
	StartupScript() *string
	SetStartupScript(val *string)
	StartupScriptInput() *string
	Subnets() ElastigroupGcpSubnetsList
	SubnetsInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UnhealthyDuration() *float64
	SetUnhealthyDuration(val *float64)
	UnhealthyDurationInput() *float64
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
	PutBackendServices(value interface{})
	PutDisk(value interface{})
	PutGpu(value interface{})
	PutInstanceTypesCustom(value interface{})
	PutIntegrationDockerSwarm(value *ElastigroupGcpIntegrationDockerSwarm)
	PutIntegrationGke(value *ElastigroupGcpIntegrationGke)
	PutLabels(value interface{})
	PutMetadata(value interface{})
	PutNetworkInterface(value interface{})
	PutScalingDownPolicy(value interface{})
	PutScalingUpPolicy(value interface{})
	PutScheduledTask(value interface{})
	PutSubnets(value interface{})
	ResetAutoHealing()
	ResetAvailabilityZones()
	ResetBackendServices()
	ResetDescription()
	ResetDisk()
	ResetDrainingTimeout()
	ResetFallbackToOndemand()
	ResetGpu()
	ResetHealthCheckGracePeriod()
	ResetHealthCheckType()
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
	ResetMinSize()
	ResetNetworkInterface()
	ResetOndemandCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreemptiblePercentage()
	ResetProvisioningModel()
	ResetScalingDownPolicy()
	ResetScalingUpPolicy()
	ResetScheduledTask()
	ResetServiceAccount()
	ResetShutdownScript()
	ResetStartupScript()
	ResetSubnets()
	ResetTags()
	ResetUnhealthyDuration()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ElastigroupGcp
type jsiiProxy_ElastigroupGcp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElastigroupGcp) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) BackendServices() ElastigroupGcpBackendServicesList {
	var returns ElastigroupGcpBackendServicesList
	_jsii_.Get(
		j,
		"backendServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) BackendServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Disk() ElastigroupGcpDiskList {
	var returns ElastigroupGcpDiskList
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) DiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) DrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) DrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) FallbackToOndemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) FallbackToOndemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Gpu() ElastigroupGcpGpuList {
	var returns ElastigroupGcpGpuList
	_jsii_.Get(
		j,
		"gpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) GpuInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) HealthCheckGracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) InstanceNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) InstanceNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) InstanceTypesCustom() ElastigroupGcpInstanceTypesCustomList {
	var returns ElastigroupGcpInstanceTypesCustomList
	_jsii_.Get(
		j,
		"instanceTypesCustom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) InstanceTypesCustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceTypesCustomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) InstanceTypesOndemand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypesOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) InstanceTypesOndemandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypesOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) InstanceTypesPreemptible() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesPreemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) InstanceTypesPreemptibleInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesPreemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) IntegrationDockerSwarm() ElastigroupGcpIntegrationDockerSwarmOutputReference {
	var returns ElastigroupGcpIntegrationDockerSwarmOutputReference
	_jsii_.Get(
		j,
		"integrationDockerSwarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) IntegrationDockerSwarmInput() *ElastigroupGcpIntegrationDockerSwarm {
	var returns *ElastigroupGcpIntegrationDockerSwarm
	_jsii_.Get(
		j,
		"integrationDockerSwarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) IntegrationGke() ElastigroupGcpIntegrationGkeOutputReference {
	var returns ElastigroupGcpIntegrationGkeOutputReference
	_jsii_.Get(
		j,
		"integrationGke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) IntegrationGkeInput() *ElastigroupGcpIntegrationGke {
	var returns *ElastigroupGcpIntegrationGke
	_jsii_.Get(
		j,
		"integrationGkeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) IpForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) IpForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Labels() ElastigroupGcpLabelsList {
	var returns ElastigroupGcpLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Metadata() ElastigroupGcpMetadataList {
	var returns ElastigroupGcpMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) NetworkInterface() ElastigroupGcpNetworkInterfaceList {
	var returns ElastigroupGcpNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) OndemandCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ondemandCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) OndemandCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ondemandCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) PreemptiblePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preemptiblePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) PreemptiblePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preemptiblePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ProvisioningModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ProvisioningModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ScalingDownPolicy() ElastigroupGcpScalingDownPolicyList {
	var returns ElastigroupGcpScalingDownPolicyList
	_jsii_.Get(
		j,
		"scalingDownPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ScalingDownPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingDownPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ScalingUpPolicy() ElastigroupGcpScalingUpPolicyList {
	var returns ElastigroupGcpScalingUpPolicyList
	_jsii_.Get(
		j,
		"scalingUpPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ScalingUpPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingUpPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ScheduledTask() ElastigroupGcpScheduledTaskList {
	var returns ElastigroupGcpScheduledTaskList
	_jsii_.Get(
		j,
		"scheduledTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ScheduledTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ShutdownScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) ShutdownScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) StartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) StartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Subnets() ElastigroupGcpSubnetsList {
	var returns ElastigroupGcpSubnetsList
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) SubnetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) UnhealthyDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcp) UnhealthyDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyDurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/elastigroup_gcp spotinst_elastigroup_gcp} Resource.
func NewElastigroupGcp(scope constructs.Construct, id *string, config *ElastigroupGcpConfig) ElastigroupGcp {
	_init_.Initialize()

	if err := validateNewElastigroupGcpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupGcp{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGcp.ElastigroupGcp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/elastigroup_gcp spotinst_elastigroup_gcp} Resource.
func NewElastigroupGcp_Override(e ElastigroupGcp, scope constructs.Construct, id *string, config *ElastigroupGcpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGcp.ElastigroupGcp",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetAutoHealing(val interface{}) {
	if err := j.validateSetAutoHealingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetDrainingTimeout(val *float64) {
	if err := j.validateSetDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainingTimeout",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetFallbackToOndemand(val interface{}) {
	if err := j.validateSetFallbackToOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackToOndemand",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetHealthCheckGracePeriod(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriod",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetInstanceNamePrefix(val *string) {
	if err := j.validateSetInstanceNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceNamePrefix",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetInstanceTypesOndemand(val *string) {
	if err := j.validateSetInstanceTypesOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypesOndemand",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetInstanceTypesPreemptible(val *[]*string) {
	if err := j.validateSetInstanceTypesPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypesPreemptible",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetIpForwarding(val interface{}) {
	if err := j.validateSetIpForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipForwarding",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetOndemandCount(val *float64) {
	if err := j.validateSetOndemandCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ondemandCount",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetPreemptiblePercentage(val *float64) {
	if err := j.validateSetPreemptiblePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptiblePercentage",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetProvisioningModel(val *string) {
	if err := j.validateSetProvisioningModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningModel",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetShutdownScript(val *string) {
	if err := j.validateSetShutdownScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownScript",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetStartupScript(val *string) {
	if err := j.validateSetStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupScript",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcp)SetUnhealthyDuration(val *float64) {
	if err := j.validateSetUnhealthyDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyDuration",
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
func ElastigroupGcp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupGcp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupGcp.ElastigroupGcp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupGcp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupGcp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupGcp.ElastigroupGcp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupGcp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupGcp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupGcp.ElastigroupGcp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElastigroupGcp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.elastigroupGcp.ElastigroupGcp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElastigroupGcp) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupGcp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupGcp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupGcp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupGcp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupGcp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupGcp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupGcp) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupGcp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupGcp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupGcp) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutBackendServices(value interface{}) {
	if err := e.validatePutBackendServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBackendServices",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutDisk(value interface{}) {
	if err := e.validatePutDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDisk",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutGpu(value interface{}) {
	if err := e.validatePutGpuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putGpu",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutInstanceTypesCustom(value interface{}) {
	if err := e.validatePutInstanceTypesCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInstanceTypesCustom",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutIntegrationDockerSwarm(value *ElastigroupGcpIntegrationDockerSwarm) {
	if err := e.validatePutIntegrationDockerSwarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationDockerSwarm",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutIntegrationGke(value *ElastigroupGcpIntegrationGke) {
	if err := e.validatePutIntegrationGkeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationGke",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutLabels(value interface{}) {
	if err := e.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLabels",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutMetadata(value interface{}) {
	if err := e.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMetadata",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutNetworkInterface(value interface{}) {
	if err := e.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutScalingDownPolicy(value interface{}) {
	if err := e.validatePutScalingDownPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingDownPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutScalingUpPolicy(value interface{}) {
	if err := e.validatePutScalingUpPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingUpPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutScheduledTask(value interface{}) {
	if err := e.validatePutScheduledTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScheduledTask",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) PutSubnets(value interface{}) {
	if err := e.validatePutSubnetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSubnets",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetBackendServices() {
	_jsii_.InvokeVoid(
		e,
		"resetBackendServices",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetDisk() {
	_jsii_.InvokeVoid(
		e,
		"resetDisk",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetDrainingTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetDrainingTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetFallbackToOndemand() {
	_jsii_.InvokeVoid(
		e,
		"resetFallbackToOndemand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetGpu() {
	_jsii_.InvokeVoid(
		e,
		"resetGpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetHealthCheckGracePeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckGracePeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetInstanceNamePrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceNamePrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetInstanceTypesCustom() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesCustom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetInstanceTypesOndemand() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesOndemand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetInstanceTypesPreemptible() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypesPreemptible",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetIntegrationDockerSwarm() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationDockerSwarm",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetIntegrationGke() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationGke",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetIpForwarding() {
	_jsii_.InvokeVoid(
		e,
		"resetIpForwarding",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetMaxSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetMetadata() {
	_jsii_.InvokeVoid(
		e,
		"resetMetadata",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetMinSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMinSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetOndemandCount() {
	_jsii_.InvokeVoid(
		e,
		"resetOndemandCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetPreemptiblePercentage() {
	_jsii_.InvokeVoid(
		e,
		"resetPreemptiblePercentage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetProvisioningModel() {
	_jsii_.InvokeVoid(
		e,
		"resetProvisioningModel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetScalingDownPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingDownPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetScalingUpPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingUpPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetScheduledTask() {
	_jsii_.InvokeVoid(
		e,
		"resetScheduledTask",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetShutdownScript() {
	_jsii_.InvokeVoid(
		e,
		"resetShutdownScript",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetStartupScript() {
	_jsii_.InvokeVoid(
		e,
		"resetStartupScript",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetSubnets() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) ResetUnhealthyDuration() {
	_jsii_.InvokeVoid(
		e,
		"resetUnhealthyDuration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

