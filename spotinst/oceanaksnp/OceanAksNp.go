// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanaksnp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np spotinst_ocean_aks_np}.
type OceanAksNp interface {
	cdktf.TerraformResource
	AksClusterName() *string
	SetAksClusterName(val *string)
	AksClusterNameInput() *string
	AksInfrastructureResourceGroupName() *string
	SetAksInfrastructureResourceGroupName(val *string)
	AksInfrastructureResourceGroupNameInput() *string
	AksRegion() *string
	SetAksRegion(val *string)
	AksRegionInput() *string
	AksResourceGroupName() *string
	SetAksResourceGroupName(val *string)
	AksResourceGroupNameInput() *string
	Autoscaler() OceanAksNpAutoscalerOutputReference
	AutoscalerInput() *OceanAksNpAutoscaler
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	EnableNodePublicIp() interface{}
	SetEnableNodePublicIp(val interface{})
	EnableNodePublicIpInput() interface{}
	FallbackToOndemand() interface{}
	SetFallbackToOndemand(val interface{})
	FallbackToOndemandInput() interface{}
	Filters() OceanAksNpFiltersOutputReference
	FiltersInput() *OceanAksNpFilters
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Headrooms() OceanAksNpHeadroomsList
	HeadroomsInput() interface{}
	Health() OceanAksNpHealthOutputReference
	HealthInput() *OceanAksNpHealth
	Id() *string
	SetId(val *string)
	IdInput() *string
	KubernetesVersion() *string
	SetKubernetesVersion(val *string)
	KubernetesVersionInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinuxOsConfig() OceanAksNpLinuxOsConfigList
	LinuxOsConfigInput() interface{}
	Logging() OceanAksNpLoggingOutputReference
	LoggingInput() *OceanAksNpLogging
	MaxCount() *float64
	SetMaxCount(val *float64)
	MaxCountInput() *float64
	MaxPodsPerNode() *float64
	SetMaxPodsPerNode(val *float64)
	MaxPodsPerNodeInput() *float64
	MinCount() *float64
	SetMinCount(val *float64)
	MinCountInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OsDiskSizeGb() *float64
	SetOsDiskSizeGb(val *float64)
	OsDiskSizeGbInput() *float64
	OsDiskType() *string
	SetOsDiskType(val *string)
	OsDiskTypeInput() *string
	OsSku() *string
	SetOsSku(val *string)
	OsSkuInput() *string
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	PodSubnetIds() *[]*string
	SetPodSubnetIds(val *[]*string)
	PodSubnetIdsInput() *[]*string
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
	Scheduling() OceanAksNpSchedulingOutputReference
	SchedulingInput() *OceanAksNpScheduling
	SpotPercentage() *float64
	SetSpotPercentage(val *float64)
	SpotPercentageInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Taints() OceanAksNpTaintsList
	TaintsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatePolicy() OceanAksNpUpdatePolicyOutputReference
	UpdatePolicyInput() *OceanAksNpUpdatePolicy
	VnetSubnetIds() *[]*string
	SetVnetSubnetIds(val *[]*string)
	VnetSubnetIdsInput() *[]*string
	VngTemplateScheduling() OceanAksNpVngTemplateSchedulingOutputReference
	VngTemplateSchedulingInput() *OceanAksNpVngTemplateScheduling
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
	PutAutoscaler(value *OceanAksNpAutoscaler)
	PutFilters(value *OceanAksNpFilters)
	PutHeadrooms(value interface{})
	PutHealth(value *OceanAksNpHealth)
	PutLinuxOsConfig(value interface{})
	PutLogging(value *OceanAksNpLogging)
	PutScheduling(value *OceanAksNpScheduling)
	PutTaints(value interface{})
	PutUpdatePolicy(value *OceanAksNpUpdatePolicy)
	PutVngTemplateScheduling(value *OceanAksNpVngTemplateScheduling)
	ResetAutoscaler()
	ResetEnableNodePublicIp()
	ResetFallbackToOndemand()
	ResetFilters()
	ResetHeadrooms()
	ResetHealth()
	ResetId()
	ResetKubernetesVersion()
	ResetLabels()
	ResetLinuxOsConfig()
	ResetLogging()
	ResetMaxCount()
	ResetMaxPodsPerNode()
	ResetMinCount()
	ResetOsDiskSizeGb()
	ResetOsDiskType()
	ResetOsSku()
	ResetOsType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPodSubnetIds()
	ResetScheduling()
	ResetSpotPercentage()
	ResetTags()
	ResetTaints()
	ResetUpdatePolicy()
	ResetVnetSubnetIds()
	ResetVngTemplateScheduling()
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

// The jsii proxy struct for OceanAksNp
type jsiiProxy_OceanAksNp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanAksNp) AksClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AksClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksClusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AksInfrastructureResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksInfrastructureResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AksInfrastructureResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksInfrastructureResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AksRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AksRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AksResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AksResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Autoscaler() OceanAksNpAutoscalerOutputReference {
	var returns OceanAksNpAutoscalerOutputReference
	_jsii_.Get(
		j,
		"autoscaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AutoscalerInput() *OceanAksNpAutoscaler {
	var returns *OceanAksNpAutoscaler
	_jsii_.Get(
		j,
		"autoscalerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) ControllerClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) ControllerClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) EnableNodePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) EnableNodePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) FallbackToOndemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) FallbackToOndemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Filters() OceanAksNpFiltersOutputReference {
	var returns OceanAksNpFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) FiltersInput() *OceanAksNpFilters {
	var returns *OceanAksNpFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Headrooms() OceanAksNpHeadroomsList {
	var returns OceanAksNpHeadroomsList
	_jsii_.Get(
		j,
		"headrooms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) HeadroomsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headroomsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Health() OceanAksNpHealthOutputReference {
	var returns OceanAksNpHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) HealthInput() *OceanAksNpHealth {
	var returns *OceanAksNpHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) KubernetesVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) KubernetesVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) LinuxOsConfig() OceanAksNpLinuxOsConfigList {
	var returns OceanAksNpLinuxOsConfigList
	_jsii_.Get(
		j,
		"linuxOsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) LinuxOsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxOsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Logging() OceanAksNpLoggingOutputReference {
	var returns OceanAksNpLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) LoggingInput() *OceanAksNpLogging {
	var returns *OceanAksNpLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) MaxCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) MaxCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) MaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) MaxPodsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) MinCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) MinCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) OsDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) OsDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) OsDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) OsDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) OsSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) OsSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) PodSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"podSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) PodSubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"podSubnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Scheduling() OceanAksNpSchedulingOutputReference {
	var returns OceanAksNpSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) SchedulingInput() *OceanAksNpScheduling {
	var returns *OceanAksNpScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) SpotPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) SpotPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) Taints() OceanAksNpTaintsList {
	var returns OceanAksNpTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) UpdatePolicy() OceanAksNpUpdatePolicyOutputReference {
	var returns OceanAksNpUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) UpdatePolicyInput() *OceanAksNpUpdatePolicy {
	var returns *OceanAksNpUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) VnetSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vnetSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) VnetSubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vnetSubnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) VngTemplateScheduling() OceanAksNpVngTemplateSchedulingOutputReference {
	var returns OceanAksNpVngTemplateSchedulingOutputReference
	_jsii_.Get(
		j,
		"vngTemplateScheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNp) VngTemplateSchedulingInput() *OceanAksNpVngTemplateScheduling {
	var returns *OceanAksNpVngTemplateScheduling
	_jsii_.Get(
		j,
		"vngTemplateSchedulingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np spotinst_ocean_aks_np} Resource.
func NewOceanAksNp(scope constructs.Construct, id *string, config *OceanAksNpConfig) OceanAksNp {
	_init_.Initialize()

	if err := validateNewOceanAksNpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAksNp{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNp.OceanAksNp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np spotinst_ocean_aks_np} Resource.
func NewOceanAksNp_Override(o OceanAksNp, scope constructs.Construct, id *string, config *OceanAksNpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNp.OceanAksNp",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanAksNp)SetAksClusterName(val *string) {
	if err := j.validateSetAksClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aksClusterName",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetAksInfrastructureResourceGroupName(val *string) {
	if err := j.validateSetAksInfrastructureResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aksInfrastructureResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetAksRegion(val *string) {
	if err := j.validateSetAksRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aksRegion",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetAksResourceGroupName(val *string) {
	if err := j.validateSetAksResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aksResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetControllerClusterId(val *string) {
	if err := j.validateSetControllerClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controllerClusterId",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetEnableNodePublicIp(val interface{}) {
	if err := j.validateSetEnableNodePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNodePublicIp",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetFallbackToOndemand(val interface{}) {
	if err := j.validateSetFallbackToOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackToOndemand",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetKubernetesVersion(val *string) {
	if err := j.validateSetKubernetesVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesVersion",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetMaxCount(val *float64) {
	if err := j.validateSetMaxCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCount",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetMaxPodsPerNode(val *float64) {
	if err := j.validateSetMaxPodsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPodsPerNode",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetMinCount(val *float64) {
	if err := j.validateSetMinCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCount",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetOsDiskSizeGb(val *float64) {
	if err := j.validateSetOsDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetOsDiskType(val *string) {
	if err := j.validateSetOsDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskType",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetOsSku(val *string) {
	if err := j.validateSetOsSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osSku",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetPodSubnetIds(val *[]*string) {
	if err := j.validateSetPodSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podSubnetIds",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetSpotPercentage(val *float64) {
	if err := j.validateSetSpotPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OceanAksNp)SetVnetSubnetIds(val *[]*string) {
	if err := j.validateSetVnetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetSubnetIds",
		val,
	)
}

// Generates CDKTF code for importing a OceanAksNp resource upon running "cdktf plan <stack-name>".
func OceanAksNp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceanAksNp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAksNp.OceanAksNp",
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
func OceanAksNp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAksNp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAksNp.OceanAksNp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAksNp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAksNp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAksNp.OceanAksNp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAksNp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAksNp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAksNp.OceanAksNp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanAksNp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanAksNp.OceanAksNp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanAksNp) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceanAksNp) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanAksNp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAksNp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAksNp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAksNp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAksNp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAksNp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAksNp) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAksNp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAksNp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNp) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceanAksNp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNp) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanAksNp) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceanAksNp) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanAksNp) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanAksNp) PutAutoscaler(value *OceanAksNpAutoscaler) {
	if err := o.validatePutAutoscalerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaler",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) PutFilters(value *OceanAksNpFilters) {
	if err := o.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFilters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) PutHeadrooms(value interface{}) {
	if err := o.validatePutHeadroomsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHeadrooms",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) PutHealth(value *OceanAksNpHealth) {
	if err := o.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHealth",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) PutLinuxOsConfig(value interface{}) {
	if err := o.validatePutLinuxOsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLinuxOsConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) PutLogging(value *OceanAksNpLogging) {
	if err := o.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogging",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) PutScheduling(value *OceanAksNpScheduling) {
	if err := o.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putScheduling",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) PutTaints(value interface{}) {
	if err := o.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTaints",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) PutUpdatePolicy(value *OceanAksNpUpdatePolicy) {
	if err := o.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) PutVngTemplateScheduling(value *OceanAksNpVngTemplateScheduling) {
	if err := o.validatePutVngTemplateSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVngTemplateScheduling",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNp) ResetAutoscaler() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaler",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetEnableNodePublicIp() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableNodePublicIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetFallbackToOndemand() {
	_jsii_.InvokeVoid(
		o,
		"resetFallbackToOndemand",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetFilters() {
	_jsii_.InvokeVoid(
		o,
		"resetFilters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetHeadrooms() {
	_jsii_.InvokeVoid(
		o,
		"resetHeadrooms",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetHealth() {
	_jsii_.InvokeVoid(
		o,
		"resetHealth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetKubernetesVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetKubernetesVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetLabels() {
	_jsii_.InvokeVoid(
		o,
		"resetLabels",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetLinuxOsConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetLinuxOsConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetLogging() {
	_jsii_.InvokeVoid(
		o,
		"resetLogging",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetMaxCount() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetMaxPodsPerNode() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxPodsPerNode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetMinCount() {
	_jsii_.InvokeVoid(
		o,
		"resetMinCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetOsDiskSizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetOsDiskSizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetOsDiskType() {
	_jsii_.InvokeVoid(
		o,
		"resetOsDiskType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetOsSku() {
	_jsii_.InvokeVoid(
		o,
		"resetOsSku",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetOsType() {
	_jsii_.InvokeVoid(
		o,
		"resetOsType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetPodSubnetIds() {
	_jsii_.InvokeVoid(
		o,
		"resetPodSubnetIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetScheduling() {
	_jsii_.InvokeVoid(
		o,
		"resetScheduling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetSpotPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetSpotPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetTaints() {
	_jsii_.InvokeVoid(
		o,
		"resetTaints",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetVnetSubnetIds() {
	_jsii_.InvokeVoid(
		o,
		"resetVnetSubnetIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) ResetVngTemplateScheduling() {
	_jsii_.InvokeVoid(
		o,
		"resetVngTemplateScheduling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

