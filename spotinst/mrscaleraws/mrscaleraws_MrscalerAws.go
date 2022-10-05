package mrscaleraws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v3/mrscaleraws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws spotinst_mrscaler_aws}.
type MrscalerAws interface {
	cdktf.TerraformResource
	AdditionalInfo() *string
	SetAdditionalInfo(val *string)
	AdditionalInfoInput() *string
	AdditionalPrimarySecurityGroups() *[]*string
	SetAdditionalPrimarySecurityGroups(val *[]*string)
	AdditionalPrimarySecurityGroupsInput() *[]*string
	AdditionalReplicaSecurityGroups() *[]*string
	SetAdditionalReplicaSecurityGroups(val *[]*string)
	AdditionalReplicaSecurityGroupsInput() *[]*string
	Applications() MrscalerAwsApplicationsList
	ApplicationsInput() interface{}
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	BootstrapActionsFile() MrscalerAwsBootstrapActionsFileList
	BootstrapActionsFileInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ConfigurationsFile() MrscalerAwsConfigurationsFileList
	ConfigurationsFileInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoreDesiredCapacity() *float64
	SetCoreDesiredCapacity(val *float64)
	CoreDesiredCapacityInput() *float64
	CoreEbsBlockDevice() MrscalerAwsCoreEbsBlockDeviceList
	CoreEbsBlockDeviceInput() interface{}
	CoreEbsOptimized() interface{}
	SetCoreEbsOptimized(val interface{})
	CoreEbsOptimizedInput() interface{}
	CoreInstanceTypes() *[]*string
	SetCoreInstanceTypes(val *[]*string)
	CoreInstanceTypesInput() *[]*string
	CoreLifecycle() *string
	SetCoreLifecycle(val *string)
	CoreLifecycleInput() *string
	CoreMaxSize() *float64
	SetCoreMaxSize(val *float64)
	CoreMaxSizeInput() *float64
	CoreMinSize() *float64
	SetCoreMinSize(val *float64)
	CoreMinSizeInput() *float64
	CoreScalingDownPolicy() MrscalerAwsCoreScalingDownPolicyList
	CoreScalingDownPolicyInput() interface{}
	CoreScalingUpPolicy() MrscalerAwsCoreScalingUpPolicyList
	CoreScalingUpPolicyInput() interface{}
	CoreUnit() *string
	SetCoreUnit(val *string)
	CoreUnitInput() *string
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CustomAmiId() *string
	SetCustomAmiId(val *string)
	CustomAmiIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EbsRootVolumeSize() *float64
	SetEbsRootVolumeSize(val *float64)
	EbsRootVolumeSizeInput() *float64
	Ec2KeyName() *string
	SetEc2KeyName(val *string)
	Ec2KeyNameInput() *string
	ExposeClusterId() interface{}
	SetExposeClusterId(val interface{})
	ExposeClusterIdInput() interface{}
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
	InstanceWeights() MrscalerAwsInstanceWeightsList
	InstanceWeightsInput() interface{}
	JobFlowRole() *string
	SetJobFlowRole(val *string)
	JobFlowRoleInput() *string
	KeepJobFlowAlive() interface{}
	SetKeepJobFlowAlive(val interface{})
	KeepJobFlowAliveInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogUri() *string
	SetLogUri(val *string)
	LogUriInput() *string
	ManagedPrimarySecurityGroup() *string
	SetManagedPrimarySecurityGroup(val *string)
	ManagedPrimarySecurityGroupInput() *string
	ManagedReplicaSecurityGroup() *string
	SetManagedReplicaSecurityGroup(val *string)
	ManagedReplicaSecurityGroupInput() *string
	MasterEbsBlockDevice() MrscalerAwsMasterEbsBlockDeviceList
	MasterEbsBlockDeviceInput() interface{}
	MasterEbsOptimized() interface{}
	SetMasterEbsOptimized(val interface{})
	MasterEbsOptimizedInput() interface{}
	MasterInstanceTypes() *[]*string
	SetMasterInstanceTypes(val *[]*string)
	MasterInstanceTypesInput() *[]*string
	MasterLifecycle() *string
	SetMasterLifecycle(val *string)
	MasterLifecycleInput() *string
	MasterTarget() *float64
	SetMasterTarget(val *float64)
	MasterTargetInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputClusterId() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisioningTimeout() MrscalerAwsProvisioningTimeoutOutputReference
	ProvisioningTimeoutInput() *MrscalerAwsProvisioningTimeout
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	ReleaseLabelInput() *string
	RepoUpgradeOnBoot() *string
	SetRepoUpgradeOnBoot(val *string)
	RepoUpgradeOnBootInput() *string
	Retries() *float64
	SetRetries(val *float64)
	RetriesInput() *float64
	ScheduledTask() MrscalerAwsScheduledTaskList
	ScheduledTaskInput() interface{}
	SecurityConfig() *string
	SetSecurityConfig(val *string)
	SecurityConfigInput() *string
	ServiceAccessSecurityGroup() *string
	SetServiceAccessSecurityGroup(val *string)
	ServiceAccessSecurityGroupInput() *string
	ServiceRole() *string
	SetServiceRole(val *string)
	ServiceRoleInput() *string
	StepsFile() MrscalerAwsStepsFileList
	StepsFileInput() interface{}
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
	Tags() MrscalerAwsTagsList
	TagsInput() interface{}
	TaskDesiredCapacity() *float64
	SetTaskDesiredCapacity(val *float64)
	TaskDesiredCapacityInput() *float64
	TaskEbsBlockDevice() MrscalerAwsTaskEbsBlockDeviceList
	TaskEbsBlockDeviceInput() interface{}
	TaskEbsOptimized() interface{}
	SetTaskEbsOptimized(val interface{})
	TaskEbsOptimizedInput() interface{}
	TaskInstanceTypes() *[]*string
	SetTaskInstanceTypes(val *[]*string)
	TaskInstanceTypesInput() *[]*string
	TaskLifecycle() *string
	SetTaskLifecycle(val *string)
	TaskLifecycleInput() *string
	TaskMaxSize() *float64
	SetTaskMaxSize(val *float64)
	TaskMaxSizeInput() *float64
	TaskMinSize() *float64
	SetTaskMinSize(val *float64)
	TaskMinSizeInput() *float64
	TaskScalingDownPolicy() MrscalerAwsTaskScalingDownPolicyList
	TaskScalingDownPolicyInput() interface{}
	TaskScalingUpPolicy() MrscalerAwsTaskScalingUpPolicyList
	TaskScalingUpPolicyInput() interface{}
	TaskUnit() *string
	SetTaskUnit(val *string)
	TaskUnitInput() *string
	TerminationPolicies() MrscalerAwsTerminationPoliciesList
	TerminationPoliciesInput() interface{}
	TerminationProtected() interface{}
	SetTerminationProtected(val interface{})
	TerminationProtectedInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VisibleToAllUsers() interface{}
	SetVisibleToAllUsers(val interface{})
	VisibleToAllUsersInput() interface{}
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
	PutApplications(value interface{})
	PutBootstrapActionsFile(value interface{})
	PutConfigurationsFile(value interface{})
	PutCoreEbsBlockDevice(value interface{})
	PutCoreScalingDownPolicy(value interface{})
	PutCoreScalingUpPolicy(value interface{})
	PutInstanceWeights(value interface{})
	PutMasterEbsBlockDevice(value interface{})
	PutProvisioningTimeout(value *MrscalerAwsProvisioningTimeout)
	PutScheduledTask(value interface{})
	PutStepsFile(value interface{})
	PutTags(value interface{})
	PutTaskEbsBlockDevice(value interface{})
	PutTaskScalingDownPolicy(value interface{})
	PutTaskScalingUpPolicy(value interface{})
	PutTerminationPolicies(value interface{})
	ResetAdditionalInfo()
	ResetAdditionalPrimarySecurityGroups()
	ResetAdditionalReplicaSecurityGroups()
	ResetApplications()
	ResetAvailabilityZones()
	ResetBootstrapActionsFile()
	ResetClusterId()
	ResetConfigurationsFile()
	ResetCoreDesiredCapacity()
	ResetCoreEbsBlockDevice()
	ResetCoreEbsOptimized()
	ResetCoreInstanceTypes()
	ResetCoreLifecycle()
	ResetCoreMaxSize()
	ResetCoreMinSize()
	ResetCoreScalingDownPolicy()
	ResetCoreScalingUpPolicy()
	ResetCoreUnit()
	ResetCustomAmiId()
	ResetDescription()
	ResetEbsRootVolumeSize()
	ResetEc2KeyName()
	ResetExposeClusterId()
	ResetId()
	ResetInstanceWeights()
	ResetJobFlowRole()
	ResetKeepJobFlowAlive()
	ResetLogUri()
	ResetManagedPrimarySecurityGroup()
	ResetManagedReplicaSecurityGroup()
	ResetMasterEbsBlockDevice()
	ResetMasterEbsOptimized()
	ResetMasterInstanceTypes()
	ResetMasterLifecycle()
	ResetMasterTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProvisioningTimeout()
	ResetRegion()
	ResetReleaseLabel()
	ResetRepoUpgradeOnBoot()
	ResetRetries()
	ResetScheduledTask()
	ResetSecurityConfig()
	ResetServiceAccessSecurityGroup()
	ResetServiceRole()
	ResetStepsFile()
	ResetTags()
	ResetTaskDesiredCapacity()
	ResetTaskEbsBlockDevice()
	ResetTaskEbsOptimized()
	ResetTaskInstanceTypes()
	ResetTaskLifecycle()
	ResetTaskMaxSize()
	ResetTaskMinSize()
	ResetTaskScalingDownPolicy()
	ResetTaskScalingUpPolicy()
	ResetTaskUnit()
	ResetTerminationPolicies()
	ResetTerminationProtected()
	ResetVisibleToAllUsers()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MrscalerAws
type jsiiProxy_MrscalerAws struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MrscalerAws) AdditionalInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) AdditionalInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) AdditionalPrimarySecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalPrimarySecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) AdditionalPrimarySecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalPrimarySecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) AdditionalReplicaSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalReplicaSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) AdditionalReplicaSecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalReplicaSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Applications() MrscalerAwsApplicationsList {
	var returns MrscalerAwsApplicationsList
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ApplicationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) BootstrapActionsFile() MrscalerAwsBootstrapActionsFileList {
	var returns MrscalerAwsBootstrapActionsFileList
	_jsii_.Get(
		j,
		"bootstrapActionsFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) BootstrapActionsFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapActionsFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ConfigurationsFile() MrscalerAwsConfigurationsFileList {
	var returns MrscalerAwsConfigurationsFileList
	_jsii_.Get(
		j,
		"configurationsFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ConfigurationsFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationsFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreDesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreDesiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreDesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreDesiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreEbsBlockDevice() MrscalerAwsCoreEbsBlockDeviceList {
	var returns MrscalerAwsCoreEbsBlockDeviceList
	_jsii_.Get(
		j,
		"coreEbsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreEbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coreEbsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreEbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coreEbsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreEbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coreEbsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"coreInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"coreInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreLifecycle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreLifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreLifecycleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreLifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreMaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreMaxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreMaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreMaxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreMinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreMinSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreMinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreMinSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreScalingDownPolicy() MrscalerAwsCoreScalingDownPolicyList {
	var returns MrscalerAwsCoreScalingDownPolicyList
	_jsii_.Get(
		j,
		"coreScalingDownPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreScalingDownPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coreScalingDownPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreScalingUpPolicy() MrscalerAwsCoreScalingUpPolicyList {
	var returns MrscalerAwsCoreScalingUpPolicyList
	_jsii_.Get(
		j,
		"coreScalingUpPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreScalingUpPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coreScalingUpPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CoreUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) CustomAmiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) EbsRootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) EbsRootVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Ec2KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2KeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Ec2KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2KeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ExposeClusterId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exposeClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ExposeClusterIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exposeClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) InstanceWeights() MrscalerAwsInstanceWeightsList {
	var returns MrscalerAwsInstanceWeightsList
	_jsii_.Get(
		j,
		"instanceWeights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) InstanceWeightsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceWeightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) JobFlowRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobFlowRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) JobFlowRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobFlowRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) KeepJobFlowAlive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) KeepJobFlowAliveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) LogUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ManagedPrimarySecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPrimarySecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ManagedPrimarySecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPrimarySecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ManagedReplicaSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedReplicaSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ManagedReplicaSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedReplicaSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterEbsBlockDevice() MrscalerAwsMasterEbsBlockDeviceList {
	var returns MrscalerAwsMasterEbsBlockDeviceList
	_jsii_.Get(
		j,
		"masterEbsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterEbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterEbsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterEbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterEbsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterEbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterEbsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"masterInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"masterInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterLifecycle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterLifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterLifecycleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterLifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) MasterTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) OutputClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ProvisioningTimeout() MrscalerAwsProvisioningTimeoutOutputReference {
	var returns MrscalerAwsProvisioningTimeoutOutputReference
	_jsii_.Get(
		j,
		"provisioningTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ProvisioningTimeoutInput() *MrscalerAwsProvisioningTimeout {
	var returns *MrscalerAwsProvisioningTimeout
	_jsii_.Get(
		j,
		"provisioningTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) RepoUpgradeOnBoot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoUpgradeOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) RepoUpgradeOnBootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoUpgradeOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Retries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) RetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ScheduledTask() MrscalerAwsScheduledTaskList {
	var returns MrscalerAwsScheduledTaskList
	_jsii_.Get(
		j,
		"scheduledTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ScheduledTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) SecurityConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) SecurityConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ServiceAccessSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ServiceAccessSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) StepsFile() MrscalerAwsStepsFileList {
	var returns MrscalerAwsStepsFileList
	_jsii_.Get(
		j,
		"stepsFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) StepsFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepsFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) Tags() MrscalerAwsTagsList {
	var returns MrscalerAwsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskDesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskDesiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskDesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskDesiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskEbsBlockDevice() MrscalerAwsTaskEbsBlockDeviceList {
	var returns MrscalerAwsTaskEbsBlockDeviceList
	_jsii_.Get(
		j,
		"taskEbsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskEbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskEbsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskEbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskEbsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskEbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskEbsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskLifecycle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskLifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskLifecycleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskLifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskMaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskMaxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskMaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskMaxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskMinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskMinSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskMinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskMinSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskScalingDownPolicy() MrscalerAwsTaskScalingDownPolicyList {
	var returns MrscalerAwsTaskScalingDownPolicyList
	_jsii_.Get(
		j,
		"taskScalingDownPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskScalingDownPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskScalingDownPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskScalingUpPolicy() MrscalerAwsTaskScalingUpPolicyList {
	var returns MrscalerAwsTaskScalingUpPolicyList
	_jsii_.Get(
		j,
		"taskScalingUpPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskScalingUpPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskScalingUpPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TaskUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TerminationPolicies() MrscalerAwsTerminationPoliciesList {
	var returns MrscalerAwsTerminationPoliciesList
	_jsii_.Get(
		j,
		"terminationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TerminationPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TerminationProtected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TerminationProtectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) VisibleToAllUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrscalerAws) VisibleToAllUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws spotinst_mrscaler_aws} Resource.
func NewMrscalerAws(scope constructs.Construct, id *string, config *MrscalerAwsConfig) MrscalerAws {
	_init_.Initialize()

	if err := validateNewMrscalerAwsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MrscalerAws{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.mrscalerAws.MrscalerAws",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws spotinst_mrscaler_aws} Resource.
func NewMrscalerAws_Override(m MrscalerAws, scope constructs.Construct, id *string, config *MrscalerAwsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.mrscalerAws.MrscalerAws",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MrscalerAws)SetAdditionalInfo(val *string) {
	if err := j.validateSetAdditionalInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetAdditionalPrimarySecurityGroups(val *[]*string) {
	if err := j.validateSetAdditionalPrimarySecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalPrimarySecurityGroups",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetAdditionalReplicaSecurityGroups(val *[]*string) {
	if err := j.validateSetAdditionalReplicaSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalReplicaSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetCoreDesiredCapacity(val *float64) {
	if err := j.validateSetCoreDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreDesiredCapacity",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetCoreEbsOptimized(val interface{}) {
	if err := j.validateSetCoreEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreEbsOptimized",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetCoreInstanceTypes(val *[]*string) {
	if err := j.validateSetCoreInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetCoreLifecycle(val *string) {
	if err := j.validateSetCoreLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreLifecycle",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetCoreMaxSize(val *float64) {
	if err := j.validateSetCoreMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreMaxSize",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetCoreMinSize(val *float64) {
	if err := j.validateSetCoreMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreMinSize",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetCoreUnit(val *string) {
	if err := j.validateSetCoreUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreUnit",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetCustomAmiId(val *string) {
	if err := j.validateSetCustomAmiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAmiId",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetEbsRootVolumeSize(val *float64) {
	if err := j.validateSetEbsRootVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsRootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetEc2KeyName(val *string) {
	if err := j.validateSetEc2KeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2KeyName",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetExposeClusterId(val interface{}) {
	if err := j.validateSetExposeClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exposeClusterId",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetJobFlowRole(val *string) {
	if err := j.validateSetJobFlowRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobFlowRole",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetKeepJobFlowAlive(val interface{}) {
	if err := j.validateSetKeepJobFlowAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepJobFlowAlive",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetLogUri(val *string) {
	if err := j.validateSetLogUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logUri",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetManagedPrimarySecurityGroup(val *string) {
	if err := j.validateSetManagedPrimarySecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedPrimarySecurityGroup",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetManagedReplicaSecurityGroup(val *string) {
	if err := j.validateSetManagedReplicaSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedReplicaSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetMasterEbsOptimized(val interface{}) {
	if err := j.validateSetMasterEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterEbsOptimized",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetMasterInstanceTypes(val *[]*string) {
	if err := j.validateSetMasterInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetMasterLifecycle(val *string) {
	if err := j.validateSetMasterLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterLifecycle",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetMasterTarget(val *float64) {
	if err := j.validateSetMasterTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterTarget",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetRepoUpgradeOnBoot(val *string) {
	if err := j.validateSetRepoUpgradeOnBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repoUpgradeOnBoot",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetRetries(val *float64) {
	if err := j.validateSetRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retries",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetSecurityConfig(val *string) {
	if err := j.validateSetSecurityConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfig",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetServiceAccessSecurityGroup(val *string) {
	if err := j.validateSetServiceAccessSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetTaskDesiredCapacity(val *float64) {
	if err := j.validateSetTaskDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDesiredCapacity",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetTaskEbsOptimized(val interface{}) {
	if err := j.validateSetTaskEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskEbsOptimized",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetTaskInstanceTypes(val *[]*string) {
	if err := j.validateSetTaskInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetTaskLifecycle(val *string) {
	if err := j.validateSetTaskLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskLifecycle",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetTaskMaxSize(val *float64) {
	if err := j.validateSetTaskMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskMaxSize",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetTaskMinSize(val *float64) {
	if err := j.validateSetTaskMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskMinSize",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetTaskUnit(val *string) {
	if err := j.validateSetTaskUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskUnit",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetTerminationProtected(val interface{}) {
	if err := j.validateSetTerminationProtectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtected",
		val,
	)
}

func (j *jsiiProxy_MrscalerAws)SetVisibleToAllUsers(val interface{}) {
	if err := j.validateSetVisibleToAllUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleToAllUsers",
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
func MrscalerAws_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMrscalerAws_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.mrscalerAws.MrscalerAws",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MrscalerAws_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.mrscalerAws.MrscalerAws",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MrscalerAws) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MrscalerAws) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MrscalerAws) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MrscalerAws) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MrscalerAws) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MrscalerAws) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MrscalerAws) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MrscalerAws) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MrscalerAws) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MrscalerAws) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MrscalerAws) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MrscalerAws) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MrscalerAws) PutApplications(value interface{}) {
	if err := m.validatePutApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putApplications",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutBootstrapActionsFile(value interface{}) {
	if err := m.validatePutBootstrapActionsFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBootstrapActionsFile",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutConfigurationsFile(value interface{}) {
	if err := m.validatePutConfigurationsFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putConfigurationsFile",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutCoreEbsBlockDevice(value interface{}) {
	if err := m.validatePutCoreEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCoreEbsBlockDevice",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutCoreScalingDownPolicy(value interface{}) {
	if err := m.validatePutCoreScalingDownPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCoreScalingDownPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutCoreScalingUpPolicy(value interface{}) {
	if err := m.validatePutCoreScalingUpPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCoreScalingUpPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutInstanceWeights(value interface{}) {
	if err := m.validatePutInstanceWeightsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putInstanceWeights",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutMasterEbsBlockDevice(value interface{}) {
	if err := m.validatePutMasterEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMasterEbsBlockDevice",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutProvisioningTimeout(value *MrscalerAwsProvisioningTimeout) {
	if err := m.validatePutProvisioningTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putProvisioningTimeout",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutScheduledTask(value interface{}) {
	if err := m.validatePutScheduledTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScheduledTask",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutStepsFile(value interface{}) {
	if err := m.validatePutStepsFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStepsFile",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutTags(value interface{}) {
	if err := m.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTags",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutTaskEbsBlockDevice(value interface{}) {
	if err := m.validatePutTaskEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTaskEbsBlockDevice",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutTaskScalingDownPolicy(value interface{}) {
	if err := m.validatePutTaskScalingDownPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTaskScalingDownPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutTaskScalingUpPolicy(value interface{}) {
	if err := m.validatePutTaskScalingUpPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTaskScalingUpPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) PutTerminationPolicies(value interface{}) {
	if err := m.validatePutTerminationPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTerminationPolicies",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrscalerAws) ResetAdditionalInfo() {
	_jsii_.InvokeVoid(
		m,
		"resetAdditionalInfo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetAdditionalPrimarySecurityGroups() {
	_jsii_.InvokeVoid(
		m,
		"resetAdditionalPrimarySecurityGroups",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetAdditionalReplicaSecurityGroups() {
	_jsii_.InvokeVoid(
		m,
		"resetAdditionalReplicaSecurityGroups",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetApplications() {
	_jsii_.InvokeVoid(
		m,
		"resetApplications",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		m,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetBootstrapActionsFile() {
	_jsii_.InvokeVoid(
		m,
		"resetBootstrapActionsFile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetClusterId() {
	_jsii_.InvokeVoid(
		m,
		"resetClusterId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetConfigurationsFile() {
	_jsii_.InvokeVoid(
		m,
		"resetConfigurationsFile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreDesiredCapacity() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreDesiredCapacity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreEbsBlockDevice() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreEbsBlockDevice",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreEbsOptimized() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreEbsOptimized",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreInstanceTypes() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreInstanceTypes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreLifecycle() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreLifecycle",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreMaxSize() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreMaxSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreMinSize() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreMinSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreScalingDownPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreScalingDownPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreScalingUpPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreScalingUpPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCoreUnit() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreUnit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetCustomAmiId() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomAmiId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetEbsRootVolumeSize() {
	_jsii_.InvokeVoid(
		m,
		"resetEbsRootVolumeSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetEc2KeyName() {
	_jsii_.InvokeVoid(
		m,
		"resetEc2KeyName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetExposeClusterId() {
	_jsii_.InvokeVoid(
		m,
		"resetExposeClusterId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetInstanceWeights() {
	_jsii_.InvokeVoid(
		m,
		"resetInstanceWeights",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetJobFlowRole() {
	_jsii_.InvokeVoid(
		m,
		"resetJobFlowRole",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetKeepJobFlowAlive() {
	_jsii_.InvokeVoid(
		m,
		"resetKeepJobFlowAlive",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetLogUri() {
	_jsii_.InvokeVoid(
		m,
		"resetLogUri",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetManagedPrimarySecurityGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetManagedPrimarySecurityGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetManagedReplicaSecurityGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetManagedReplicaSecurityGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetMasterEbsBlockDevice() {
	_jsii_.InvokeVoid(
		m,
		"resetMasterEbsBlockDevice",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetMasterEbsOptimized() {
	_jsii_.InvokeVoid(
		m,
		"resetMasterEbsOptimized",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetMasterInstanceTypes() {
	_jsii_.InvokeVoid(
		m,
		"resetMasterInstanceTypes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetMasterLifecycle() {
	_jsii_.InvokeVoid(
		m,
		"resetMasterLifecycle",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetMasterTarget() {
	_jsii_.InvokeVoid(
		m,
		"resetMasterTarget",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetProvisioningTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetProvisioningTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetReleaseLabel() {
	_jsii_.InvokeVoid(
		m,
		"resetReleaseLabel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetRepoUpgradeOnBoot() {
	_jsii_.InvokeVoid(
		m,
		"resetRepoUpgradeOnBoot",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetRetries() {
	_jsii_.InvokeVoid(
		m,
		"resetRetries",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetScheduledTask() {
	_jsii_.InvokeVoid(
		m,
		"resetScheduledTask",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetServiceAccessSecurityGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceAccessSecurityGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetServiceRole() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceRole",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetStepsFile() {
	_jsii_.InvokeVoid(
		m,
		"resetStepsFile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskDesiredCapacity() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskDesiredCapacity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskEbsBlockDevice() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskEbsBlockDevice",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskEbsOptimized() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskEbsOptimized",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskInstanceTypes() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskInstanceTypes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskLifecycle() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskLifecycle",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskMaxSize() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskMaxSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskMinSize() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskMinSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskScalingDownPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskScalingDownPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskScalingUpPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskScalingUpPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTaskUnit() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskUnit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTerminationPolicies() {
	_jsii_.InvokeVoid(
		m,
		"resetTerminationPolicies",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetTerminationProtected() {
	_jsii_.InvokeVoid(
		m,
		"resetTerminationProtected",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) ResetVisibleToAllUsers() {
	_jsii_.InvokeVoid(
		m,
		"resetVisibleToAllUsers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrscalerAws) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAws) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAws) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrscalerAws) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

