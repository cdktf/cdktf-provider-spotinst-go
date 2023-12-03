// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkelaunchspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v12/oceangkelaunchspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_gke_launch_spec spotinst_ocean_gke_launch_spec}.
type OceanGkeLaunchSpec interface {
	cdktf.TerraformResource
	AutoscaleHeadrooms() OceanGkeLaunchSpecAutoscaleHeadroomsList
	AutoscaleHeadroomsAutomatic() OceanGkeLaunchSpecAutoscaleHeadroomsAutomaticList
	AutoscaleHeadroomsAutomaticInput() interface{}
	AutoscaleHeadroomsInput() interface{}
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
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	InstanceTypesInput() *[]*string
	Labels() OceanGkeLaunchSpecLabelsList
	LabelsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() OceanGkeLaunchSpecMetadataList
	MetadataInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterfaces() OceanGkeLaunchSpecNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	// The tree node.
	Node() constructs.Node
	NodePoolName() *string
	SetNodePoolName(val *string)
	NodePoolNameInput() *string
	OceanId() *string
	SetOceanId(val *string)
	OceanIdInput() *string
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
	ResourceLimits() OceanGkeLaunchSpecResourceLimitsOutputReference
	ResourceLimitsInput() *OceanGkeLaunchSpecResourceLimits
	RestrictScaleDown() interface{}
	SetRestrictScaleDown(val interface{})
	RestrictScaleDownInput() interface{}
	RootVolumeSize() *float64
	SetRootVolumeSize(val *float64)
	RootVolumeSizeInput() *float64
	RootVolumeType() *string
	SetRootVolumeType(val *string)
	RootVolumeTypeInput() *string
	SchedulingTask() OceanGkeLaunchSpecSchedulingTaskList
	SchedulingTaskInput() interface{}
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ShieldedInstanceConfig() OceanGkeLaunchSpecShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *OceanGkeLaunchSpecShieldedInstanceConfig
	SourceImage() *string
	SetSourceImage(val *string)
	SourceImageInput() *string
	Storage() OceanGkeLaunchSpecStorageOutputReference
	StorageInput() *OceanGkeLaunchSpecStorage
	Strategy() OceanGkeLaunchSpecStrategyList
	StrategyInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Taints() OceanGkeLaunchSpecTaintsList
	TaintsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatePolicy() OceanGkeLaunchSpecUpdatePolicyOutputReference
	UpdatePolicyInput() *OceanGkeLaunchSpecUpdatePolicy
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAutoscaleHeadrooms(value interface{})
	PutAutoscaleHeadroomsAutomatic(value interface{})
	PutLabels(value interface{})
	PutMetadata(value interface{})
	PutNetworkInterfaces(value interface{})
	PutResourceLimits(value *OceanGkeLaunchSpecResourceLimits)
	PutSchedulingTask(value interface{})
	PutShieldedInstanceConfig(value *OceanGkeLaunchSpecShieldedInstanceConfig)
	PutStorage(value *OceanGkeLaunchSpecStorage)
	PutStrategy(value interface{})
	PutTaints(value interface{})
	PutUpdatePolicy(value *OceanGkeLaunchSpecUpdatePolicy)
	ResetAutoscaleHeadrooms()
	ResetAutoscaleHeadroomsAutomatic()
	ResetId()
	ResetInstanceTypes()
	ResetLabels()
	ResetMetadata()
	ResetName()
	ResetNetworkInterfaces()
	ResetNodePoolName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceLimits()
	ResetRestrictScaleDown()
	ResetRootVolumeSize()
	ResetRootVolumeType()
	ResetSchedulingTask()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
	ResetSourceImage()
	ResetStorage()
	ResetStrategy()
	ResetTags()
	ResetTaints()
	ResetUpdatePolicy()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OceanGkeLaunchSpec
type jsiiProxy_OceanGkeLaunchSpec struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanGkeLaunchSpec) AutoscaleHeadrooms() OceanGkeLaunchSpecAutoscaleHeadroomsList {
	var returns OceanGkeLaunchSpecAutoscaleHeadroomsList
	_jsii_.Get(
		j,
		"autoscaleHeadrooms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) AutoscaleHeadroomsAutomatic() OceanGkeLaunchSpecAutoscaleHeadroomsAutomaticList {
	var returns OceanGkeLaunchSpecAutoscaleHeadroomsAutomaticList
	_jsii_.Get(
		j,
		"autoscaleHeadroomsAutomatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) AutoscaleHeadroomsAutomaticInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleHeadroomsAutomaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) AutoscaleHeadroomsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleHeadroomsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Labels() OceanGkeLaunchSpecLabelsList {
	var returns OceanGkeLaunchSpecLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Metadata() OceanGkeLaunchSpecMetadataList {
	var returns OceanGkeLaunchSpecMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) NetworkInterfaces() OceanGkeLaunchSpecNetworkInterfacesList {
	var returns OceanGkeLaunchSpecNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) NodePoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) NodePoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) OceanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) OceanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) ResourceLimits() OceanGkeLaunchSpecResourceLimitsOutputReference {
	var returns OceanGkeLaunchSpecResourceLimitsOutputReference
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) ResourceLimitsInput() *OceanGkeLaunchSpecResourceLimits {
	var returns *OceanGkeLaunchSpecResourceLimits
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) RestrictScaleDown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictScaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) RestrictScaleDownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictScaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) RootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) RootVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) RootVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) RootVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) SchedulingTask() OceanGkeLaunchSpecSchedulingTaskList {
	var returns OceanGkeLaunchSpecSchedulingTaskList
	_jsii_.Get(
		j,
		"schedulingTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) SchedulingTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedulingTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) ShieldedInstanceConfig() OceanGkeLaunchSpecShieldedInstanceConfigOutputReference {
	var returns OceanGkeLaunchSpecShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) ShieldedInstanceConfigInput() *OceanGkeLaunchSpecShieldedInstanceConfig {
	var returns *OceanGkeLaunchSpecShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) SourceImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) SourceImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Storage() OceanGkeLaunchSpecStorageOutputReference {
	var returns OceanGkeLaunchSpecStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) StorageInput() *OceanGkeLaunchSpecStorage {
	var returns *OceanGkeLaunchSpecStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Strategy() OceanGkeLaunchSpecStrategyList {
	var returns OceanGkeLaunchSpecStrategyList
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) StrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) Taints() OceanGkeLaunchSpecTaintsList {
	var returns OceanGkeLaunchSpecTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) UpdatePolicy() OceanGkeLaunchSpecUpdatePolicyOutputReference {
	var returns OceanGkeLaunchSpecUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpec) UpdatePolicyInput() *OceanGkeLaunchSpecUpdatePolicy {
	var returns *OceanGkeLaunchSpecUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_gke_launch_spec spotinst_ocean_gke_launch_spec} Resource.
func NewOceanGkeLaunchSpec(scope constructs.Construct, id *string, config *OceanGkeLaunchSpecConfig) OceanGkeLaunchSpec {
	_init_.Initialize()

	if err := validateNewOceanGkeLaunchSpecParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanGkeLaunchSpec{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpec",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_gke_launch_spec spotinst_ocean_gke_launch_spec} Resource.
func NewOceanGkeLaunchSpec_Override(o OceanGkeLaunchSpec, scope constructs.Construct, id *string, config *OceanGkeLaunchSpecConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpec",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetNodePoolName(val *string) {
	if err := j.validateSetNodePoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodePoolName",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetOceanId(val *string) {
	if err := j.validateSetOceanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oceanId",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetRestrictScaleDown(val interface{}) {
	if err := j.validateSetRestrictScaleDownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictScaleDown",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetRootVolumeSize(val *float64) {
	if err := j.validateSetRootVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetRootVolumeType(val *string) {
	if err := j.validateSetRootVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootVolumeType",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetSourceImage(val *string) {
	if err := j.validateSetSourceImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceImage",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpec)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a OceanGkeLaunchSpec resource upon running "cdktf plan <stack-name>".
func OceanGkeLaunchSpec_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceanGkeLaunchSpec_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpec",
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
func OceanGkeLaunchSpec_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanGkeLaunchSpec_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpec",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanGkeLaunchSpec_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanGkeLaunchSpec_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpec",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanGkeLaunchSpec_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanGkeLaunchSpec_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpec",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanGkeLaunchSpec_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpec",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpec) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanGkeLaunchSpec) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutAutoscaleHeadrooms(value interface{}) {
	if err := o.validatePutAutoscaleHeadroomsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleHeadrooms",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutAutoscaleHeadroomsAutomatic(value interface{}) {
	if err := o.validatePutAutoscaleHeadroomsAutomaticParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleHeadroomsAutomatic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutLabels(value interface{}) {
	if err := o.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLabels",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutMetadata(value interface{}) {
	if err := o.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMetadata",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutNetworkInterfaces(value interface{}) {
	if err := o.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutResourceLimits(value *OceanGkeLaunchSpecResourceLimits) {
	if err := o.validatePutResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putResourceLimits",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutSchedulingTask(value interface{}) {
	if err := o.validatePutSchedulingTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSchedulingTask",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutShieldedInstanceConfig(value *OceanGkeLaunchSpecShieldedInstanceConfig) {
	if err := o.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutStorage(value *OceanGkeLaunchSpecStorage) {
	if err := o.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStorage",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutStrategy(value interface{}) {
	if err := o.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStrategy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutTaints(value interface{}) {
	if err := o.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTaints",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) PutUpdatePolicy(value *OceanGkeLaunchSpecUpdatePolicy) {
	if err := o.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetAutoscaleHeadrooms() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleHeadrooms",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetAutoscaleHeadroomsAutomatic() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleHeadroomsAutomatic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetInstanceTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetLabels() {
	_jsii_.InvokeVoid(
		o,
		"resetLabels",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetMetadata() {
	_jsii_.InvokeVoid(
		o,
		"resetMetadata",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		o,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetNodePoolName() {
	_jsii_.InvokeVoid(
		o,
		"resetNodePoolName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		o,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetRestrictScaleDown() {
	_jsii_.InvokeVoid(
		o,
		"resetRestrictScaleDown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetRootVolumeSize() {
	_jsii_.InvokeVoid(
		o,
		"resetRootVolumeSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetRootVolumeType() {
	_jsii_.InvokeVoid(
		o,
		"resetRootVolumeType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetSchedulingTask() {
	_jsii_.InvokeVoid(
		o,
		"resetSchedulingTask",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		o,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetSourceImage() {
	_jsii_.InvokeVoid(
		o,
		"resetSourceImage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetStorage() {
	_jsii_.InvokeVoid(
		o,
		"resetStorage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetStrategy() {
	_jsii_.InvokeVoid(
		o,
		"resetStrategy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetTaints() {
	_jsii_.InvokeVoid(
		o,
		"resetTaints",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpec) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpec) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

