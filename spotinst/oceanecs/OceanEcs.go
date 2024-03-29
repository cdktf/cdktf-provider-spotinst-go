// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanecs/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.166.0/docs/resources/ocean_ecs spotinst_ocean_ecs}.
type OceanEcs interface {
	cdktf.TerraformResource
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	AssociatePublicIpAddressInput() interface{}
	Autoscaler() OceanEcsAutoscalerOutputReference
	AutoscalerInput() *OceanEcsAutoscaler
	Blacklist() *[]*string
	SetBlacklist(val *[]*string)
	BlacklistInput() *[]*string
	BlockDeviceMappings() OceanEcsBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	ClusterOrientation() OceanEcsClusterOrientationList
	ClusterOrientationInput() interface{}
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
	DrainingTimeout() *float64
	SetDrainingTimeout(val *float64)
	DrainingTimeoutInput() *float64
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	Filters() OceanEcsFiltersOutputReference
	FiltersInput() *OceanEcsFilters
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceMetadataOptions() OceanEcsInstanceMetadataOptionsOutputReference
	InstanceMetadataOptionsInput() *OceanEcsInstanceMetadataOptions
	KeyPair() *string
	SetKeyPair(val *string)
	KeyPairInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logging() OceanEcsLoggingOutputReference
	LoggingInput() *OceanEcsLogging
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
	OptimizeImages() OceanEcsOptimizeImagesOutputReference
	OptimizeImagesInput() *OceanEcsOptimizeImages
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
	ScheduledTask() OceanEcsScheduledTaskList
	ScheduledTaskInput() interface{}
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SpotPercentage() *float64
	SetSpotPercentage(val *float64)
	SpotPercentageInput() *float64
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() OceanEcsTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatePolicy() OceanEcsUpdatePolicyOutputReference
	UpdatePolicyInput() *OceanEcsUpdatePolicy
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
	PutAutoscaler(value *OceanEcsAutoscaler)
	PutBlockDeviceMappings(value interface{})
	PutClusterOrientation(value interface{})
	PutFilters(value *OceanEcsFilters)
	PutInstanceMetadataOptions(value *OceanEcsInstanceMetadataOptions)
	PutLogging(value *OceanEcsLogging)
	PutOptimizeImages(value *OceanEcsOptimizeImages)
	PutScheduledTask(value interface{})
	PutTags(value interface{})
	PutUpdatePolicy(value *OceanEcsUpdatePolicy)
	ResetAssociatePublicIpAddress()
	ResetAutoscaler()
	ResetBlacklist()
	ResetBlockDeviceMappings()
	ResetClusterOrientation()
	ResetDesiredCapacity()
	ResetDrainingTimeout()
	ResetEbsOptimized()
	ResetFilters()
	ResetIamInstanceProfile()
	ResetId()
	ResetImageId()
	ResetInstanceMetadataOptions()
	ResetKeyPair()
	ResetLogging()
	ResetMaxSize()
	ResetMinSize()
	ResetMonitoring()
	ResetOptimizeImages()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScheduledTask()
	ResetSpotPercentage()
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

// The jsii proxy struct for OceanEcs
type jsiiProxy_OceanEcs struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanEcs) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Autoscaler() OceanEcsAutoscalerOutputReference {
	var returns OceanEcsAutoscalerOutputReference
	_jsii_.Get(
		j,
		"autoscaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) AutoscalerInput() *OceanEcsAutoscaler {
	var returns *OceanEcsAutoscaler
	_jsii_.Get(
		j,
		"autoscalerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Blacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) BlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) BlockDeviceMappings() OceanEcsBlockDeviceMappingsList {
	var returns OceanEcsBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ClusterOrientation() OceanEcsClusterOrientationList {
	var returns OceanEcsClusterOrientationList
	_jsii_.Get(
		j,
		"clusterOrientation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ClusterOrientationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterOrientationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) DrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) DrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Filters() OceanEcsFiltersOutputReference {
	var returns OceanEcsFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) FiltersInput() *OceanEcsFilters {
	var returns *OceanEcsFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) InstanceMetadataOptions() OceanEcsInstanceMetadataOptionsOutputReference {
	var returns OceanEcsInstanceMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"instanceMetadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) InstanceMetadataOptionsInput() *OceanEcsInstanceMetadataOptions {
	var returns *OceanEcsInstanceMetadataOptions
	_jsii_.Get(
		j,
		"instanceMetadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) KeyPairInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Logging() OceanEcsLoggingOutputReference {
	var returns OceanEcsLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) LoggingInput() *OceanEcsLogging {
	var returns *OceanEcsLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) OptimizeImages() OceanEcsOptimizeImagesOutputReference {
	var returns OceanEcsOptimizeImagesOutputReference
	_jsii_.Get(
		j,
		"optimizeImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) OptimizeImagesInput() *OceanEcsOptimizeImages {
	var returns *OceanEcsOptimizeImages
	_jsii_.Get(
		j,
		"optimizeImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ScheduledTask() OceanEcsScheduledTaskList {
	var returns OceanEcsScheduledTaskList
	_jsii_.Get(
		j,
		"scheduledTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) ScheduledTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) SpotPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) SpotPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Tags() OceanEcsTagsList {
	var returns OceanEcsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UpdatePolicy() OceanEcsUpdatePolicyOutputReference {
	var returns OceanEcsUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UpdatePolicyInput() *OceanEcsUpdatePolicy {
	var returns *OceanEcsUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UseAsTemplateOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAsTemplateOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UseAsTemplateOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAsTemplateOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UtilizeCommitments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeCommitments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UtilizeCommitmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeCommitmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UtilizeReservedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeReservedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) UtilizeReservedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeReservedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) Whitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcs) WhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.166.0/docs/resources/ocean_ecs spotinst_ocean_ecs} Resource.
func NewOceanEcs(scope constructs.Construct, id *string, config *OceanEcsConfig) OceanEcs {
	_init_.Initialize()

	if err := validateNewOceanEcsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanEcs{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcs",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.166.0/docs/resources/ocean_ecs spotinst_ocean_ecs} Resource.
func NewOceanEcs_Override(o OceanEcs, scope constructs.Construct, id *string, config *OceanEcsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcs",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanEcs)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetBlacklist(val *[]*string) {
	if err := j.validateSetBlacklistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blacklist",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetDrainingTimeout(val *float64) {
	if err := j.validateSetDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainingTimeout",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetKeyPair(val *string) {
	if err := j.validateSetKeyPairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPair",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetSpotPercentage(val *float64) {
	if err := j.validateSetSpotPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetUseAsTemplateOnly(val interface{}) {
	if err := j.validateSetUseAsTemplateOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAsTemplateOnly",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetUtilizeCommitments(val interface{}) {
	if err := j.validateSetUtilizeCommitmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utilizeCommitments",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetUtilizeReservedInstances(val interface{}) {
	if err := j.validateSetUtilizeReservedInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utilizeReservedInstances",
		val,
	)
}

func (j *jsiiProxy_OceanEcs)SetWhitelist(val *[]*string) {
	if err := j.validateSetWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelist",
		val,
	)
}

// Generates CDKTF code for importing a OceanEcs resource upon running "cdktf plan <stack-name>".
func OceanEcs_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceanEcs_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcs",
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
func OceanEcs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanEcs_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanEcs_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanEcs_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcs",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanEcs_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanEcs_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcs",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanEcs_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcs",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanEcs) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceanEcs) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanEcs) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanEcs) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcs) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanEcs) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanEcs) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanEcs) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanEcs) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanEcs) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanEcs) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanEcs) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcs) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceanEcs) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcs) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanEcs) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceanEcs) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanEcs) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanEcs) PutAutoscaler(value *OceanEcsAutoscaler) {
	if err := o.validatePutAutoscalerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaler",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) PutBlockDeviceMappings(value interface{}) {
	if err := o.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) PutClusterOrientation(value interface{}) {
	if err := o.validatePutClusterOrientationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putClusterOrientation",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) PutFilters(value *OceanEcsFilters) {
	if err := o.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFilters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) PutInstanceMetadataOptions(value *OceanEcsInstanceMetadataOptions) {
	if err := o.validatePutInstanceMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInstanceMetadataOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) PutLogging(value *OceanEcsLogging) {
	if err := o.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogging",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) PutOptimizeImages(value *OceanEcsOptimizeImages) {
	if err := o.validatePutOptimizeImagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOptimizeImages",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) PutScheduledTask(value interface{}) {
	if err := o.validatePutScheduledTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putScheduledTask",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) PutTags(value interface{}) {
	if err := o.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTags",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) PutUpdatePolicy(value *OceanEcsUpdatePolicy) {
	if err := o.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcs) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		o,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetAutoscaler() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaler",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetBlacklist() {
	_jsii_.InvokeVoid(
		o,
		"resetBlacklist",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		o,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetClusterOrientation() {
	_jsii_.InvokeVoid(
		o,
		"resetClusterOrientation",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		o,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetDrainingTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainingTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetFilters() {
	_jsii_.InvokeVoid(
		o,
		"resetFilters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		o,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetImageId() {
	_jsii_.InvokeVoid(
		o,
		"resetImageId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetInstanceMetadataOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceMetadataOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetKeyPair() {
	_jsii_.InvokeVoid(
		o,
		"resetKeyPair",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetLogging() {
	_jsii_.InvokeVoid(
		o,
		"resetLogging",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetMaxSize() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetMinSize() {
	_jsii_.InvokeVoid(
		o,
		"resetMinSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetMonitoring() {
	_jsii_.InvokeVoid(
		o,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetOptimizeImages() {
	_jsii_.InvokeVoid(
		o,
		"resetOptimizeImages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetScheduledTask() {
	_jsii_.InvokeVoid(
		o,
		"resetScheduledTask",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetSpotPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetSpotPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetUseAsTemplateOnly() {
	_jsii_.InvokeVoid(
		o,
		"resetUseAsTemplateOnly",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetUserData() {
	_jsii_.InvokeVoid(
		o,
		"resetUserData",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetUtilizeCommitments() {
	_jsii_.InvokeVoid(
		o,
		"resetUtilizeCommitments",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetUtilizeReservedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUtilizeReservedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) ResetWhitelist() {
	_jsii_.InvokeVoid(
		o,
		"resetWhitelist",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcs) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcs) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcs) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcs) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcs) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

