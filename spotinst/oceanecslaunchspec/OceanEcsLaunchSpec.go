// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecslaunchspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanecslaunchspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.166.0/docs/resources/ocean_ecs_launch_spec spotinst_ocean_ecs_launch_spec}.
type OceanEcsLaunchSpec interface {
	cdktf.TerraformResource
	Attributes() OceanEcsLaunchSpecAttributesList
	AttributesInput() interface{}
	AutoscaleHeadrooms() OceanEcsLaunchSpecAutoscaleHeadroomsList
	AutoscaleHeadroomsInput() interface{}
	BlockDeviceMappings() OceanEcsLaunchSpecBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
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
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	Images() OceanEcsLaunchSpecImagesList
	ImagesInput() interface{}
	InstanceMetadataOptions() OceanEcsLaunchSpecInstanceMetadataOptionsOutputReference
	InstanceMetadataOptionsInput() *OceanEcsLaunchSpecInstanceMetadataOptions
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	InstanceTypesInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OceanId() *string
	SetOceanId(val *string)
	OceanIdInput() *string
	PreferredSpotTypes() *[]*string
	SetPreferredSpotTypes(val *[]*string)
	PreferredSpotTypesInput() *[]*string
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
	RestrictScaleDown() interface{}
	SetRestrictScaleDown(val interface{})
	RestrictScaleDownInput() interface{}
	SchedulingTask() OceanEcsLaunchSpecSchedulingTaskList
	SchedulingTaskInput() interface{}
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	Strategy() OceanEcsLaunchSpecStrategyList
	StrategyInput() interface{}
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() OceanEcsLaunchSpecTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
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
	PutAttributes(value interface{})
	PutAutoscaleHeadrooms(value interface{})
	PutBlockDeviceMappings(value interface{})
	PutImages(value interface{})
	PutInstanceMetadataOptions(value *OceanEcsLaunchSpecInstanceMetadataOptions)
	PutSchedulingTask(value interface{})
	PutStrategy(value interface{})
	PutTags(value interface{})
	ResetAttributes()
	ResetAutoscaleHeadrooms()
	ResetBlockDeviceMappings()
	ResetIamInstanceProfile()
	ResetId()
	ResetImageId()
	ResetImages()
	ResetInstanceMetadataOptions()
	ResetInstanceTypes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreferredSpotTypes()
	ResetRestrictScaleDown()
	ResetSchedulingTask()
	ResetSecurityGroupIds()
	ResetStrategy()
	ResetSubnetIds()
	ResetTags()
	ResetUserData()
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

// The jsii proxy struct for OceanEcsLaunchSpec
type jsiiProxy_OceanEcsLaunchSpec struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Attributes() OceanEcsLaunchSpecAttributesList {
	var returns OceanEcsLaunchSpecAttributesList
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) AutoscaleHeadrooms() OceanEcsLaunchSpecAutoscaleHeadroomsList {
	var returns OceanEcsLaunchSpecAutoscaleHeadroomsList
	_jsii_.Get(
		j,
		"autoscaleHeadrooms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) AutoscaleHeadroomsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleHeadroomsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) BlockDeviceMappings() OceanEcsLaunchSpecBlockDeviceMappingsList {
	var returns OceanEcsLaunchSpecBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Images() OceanEcsLaunchSpecImagesList {
	var returns OceanEcsLaunchSpecImagesList
	_jsii_.Get(
		j,
		"images",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) ImagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) InstanceMetadataOptions() OceanEcsLaunchSpecInstanceMetadataOptionsOutputReference {
	var returns OceanEcsLaunchSpecInstanceMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"instanceMetadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) InstanceMetadataOptionsInput() *OceanEcsLaunchSpecInstanceMetadataOptions {
	var returns *OceanEcsLaunchSpecInstanceMetadataOptions
	_jsii_.Get(
		j,
		"instanceMetadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) OceanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) OceanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) PreferredSpotTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredSpotTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) PreferredSpotTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredSpotTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) RestrictScaleDown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictScaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) RestrictScaleDownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictScaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) SchedulingTask() OceanEcsLaunchSpecSchedulingTaskList {
	var returns OceanEcsLaunchSpecSchedulingTaskList
	_jsii_.Get(
		j,
		"schedulingTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) SchedulingTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedulingTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Strategy() OceanEcsLaunchSpecStrategyList {
	var returns OceanEcsLaunchSpecStrategyList
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) StrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) Tags() OceanEcsLaunchSpecTagsList {
	var returns OceanEcsLaunchSpecTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpec) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.166.0/docs/resources/ocean_ecs_launch_spec spotinst_ocean_ecs_launch_spec} Resource.
func NewOceanEcsLaunchSpec(scope constructs.Construct, id *string, config *OceanEcsLaunchSpecConfig) OceanEcsLaunchSpec {
	_init_.Initialize()

	if err := validateNewOceanEcsLaunchSpecParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanEcsLaunchSpec{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcsLaunchSpec.OceanEcsLaunchSpec",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.166.0/docs/resources/ocean_ecs_launch_spec spotinst_ocean_ecs_launch_spec} Resource.
func NewOceanEcsLaunchSpec_Override(o OceanEcsLaunchSpec, scope constructs.Construct, id *string, config *OceanEcsLaunchSpecConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcsLaunchSpec.OceanEcsLaunchSpec",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetOceanId(val *string) {
	if err := j.validateSetOceanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oceanId",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetPreferredSpotTypes(val *[]*string) {
	if err := j.validateSetPreferredSpotTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredSpotTypes",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetRestrictScaleDown(val interface{}) {
	if err := j.validateSetRestrictScaleDownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictScaleDown",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpec)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

// Generates CDKTF code for importing a OceanEcsLaunchSpec resource upon running "cdktf plan <stack-name>".
func OceanEcsLaunchSpec_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceanEcsLaunchSpec_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanEcsLaunchSpec.OceanEcsLaunchSpec",
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
func OceanEcsLaunchSpec_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanEcsLaunchSpec_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanEcsLaunchSpec.OceanEcsLaunchSpec",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanEcsLaunchSpec_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanEcsLaunchSpec_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanEcsLaunchSpec.OceanEcsLaunchSpec",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanEcsLaunchSpec_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanEcsLaunchSpec_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanEcsLaunchSpec.OceanEcsLaunchSpec",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanEcsLaunchSpec_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanEcsLaunchSpec.OceanEcsLaunchSpec",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpec) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcsLaunchSpec) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) PutAttributes(value interface{}) {
	if err := o.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAttributes",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) PutAutoscaleHeadrooms(value interface{}) {
	if err := o.validatePutAutoscaleHeadroomsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleHeadrooms",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) PutBlockDeviceMappings(value interface{}) {
	if err := o.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) PutImages(value interface{}) {
	if err := o.validatePutImagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putImages",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) PutInstanceMetadataOptions(value *OceanEcsLaunchSpecInstanceMetadataOptions) {
	if err := o.validatePutInstanceMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInstanceMetadataOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) PutSchedulingTask(value interface{}) {
	if err := o.validatePutSchedulingTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSchedulingTask",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) PutStrategy(value interface{}) {
	if err := o.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStrategy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) PutTags(value interface{}) {
	if err := o.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTags",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetAttributes() {
	_jsii_.InvokeVoid(
		o,
		"resetAttributes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetAutoscaleHeadrooms() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleHeadrooms",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		o,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		o,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetImageId() {
	_jsii_.InvokeVoid(
		o,
		"resetImageId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetImages() {
	_jsii_.InvokeVoid(
		o,
		"resetImages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetInstanceMetadataOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceMetadataOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetInstanceTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetPreferredSpotTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetPreferredSpotTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetRestrictScaleDown() {
	_jsii_.InvokeVoid(
		o,
		"resetRestrictScaleDown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetSchedulingTask() {
	_jsii_.InvokeVoid(
		o,
		"resetSchedulingTask",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetStrategy() {
	_jsii_.InvokeVoid(
		o,
		"resetStrategy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		o,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ResetUserData() {
	_jsii_.InvokeVoid(
		o,
		"resetUserData",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpec) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpec) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpec) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

