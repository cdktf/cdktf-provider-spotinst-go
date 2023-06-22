package oceanawslaunchspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v9/oceanawslaunchspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.123.0/docs/resources/ocean_aws_launch_spec spotinst_ocean_aws_launch_spec}.
type OceanAwsLaunchSpec interface {
	cdktf.TerraformResource
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	AssociatePublicIpAddressInput() interface{}
	AutoscaleDown() OceanAwsLaunchSpecAutoscaleDownList
	AutoscaleDownInput() interface{}
	AutoscaleHeadrooms() OceanAwsLaunchSpecAutoscaleHeadroomsList
	AutoscaleHeadroomsAutomatic() OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList
	AutoscaleHeadroomsAutomaticInput() interface{}
	AutoscaleHeadroomsInput() interface{}
	BlockDeviceMappings() OceanAwsLaunchSpecBlockDeviceMappingsList
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
	CreateOptions() OceanAwsLaunchSpecCreateOptionsOutputReference
	CreateOptionsInput() *OceanAwsLaunchSpecCreateOptions
	DeleteOptions() OceanAwsLaunchSpecDeleteOptionsOutputReference
	DeleteOptionsInput() *OceanAwsLaunchSpecDeleteOptions
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ElasticIpPool() OceanAwsLaunchSpecElasticIpPoolList
	ElasticIpPoolInput() interface{}
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
	Images() OceanAwsLaunchSpecImagesList
	ImagesInput() interface{}
	InstanceMetadataOptions() OceanAwsLaunchSpecInstanceMetadataOptionsOutputReference
	InstanceMetadataOptionsInput() *OceanAwsLaunchSpecInstanceMetadataOptions
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	InstanceTypesInput() *[]*string
	Labels() OceanAwsLaunchSpecLabelsList
	LabelsInput() interface{}
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
	ResourceLimits() OceanAwsLaunchSpecResourceLimitsList
	ResourceLimitsInput() interface{}
	RestrictScaleDown() interface{}
	SetRestrictScaleDown(val interface{})
	RestrictScaleDownInput() interface{}
	RootVolumeSize() *float64
	SetRootVolumeSize(val *float64)
	RootVolumeSizeInput() *float64
	SchedulingShutdownHours() OceanAwsLaunchSpecSchedulingShutdownHoursOutputReference
	SchedulingShutdownHoursInput() *OceanAwsLaunchSpecSchedulingShutdownHours
	SchedulingTask() OceanAwsLaunchSpecSchedulingTaskList
	SchedulingTaskInput() interface{}
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	Strategy() OceanAwsLaunchSpecStrategyList
	StrategyInput() interface{}
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() OceanAwsLaunchSpecTagsList
	TagsInput() interface{}
	Taints() OceanAwsLaunchSpecTaintsList
	TaintsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatePolicy() OceanAwsLaunchSpecUpdatePolicyOutputReference
	UpdatePolicyInput() *OceanAwsLaunchSpecUpdatePolicy
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
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
	PutAutoscaleDown(value interface{})
	PutAutoscaleHeadrooms(value interface{})
	PutAutoscaleHeadroomsAutomatic(value interface{})
	PutBlockDeviceMappings(value interface{})
	PutCreateOptions(value *OceanAwsLaunchSpecCreateOptions)
	PutDeleteOptions(value *OceanAwsLaunchSpecDeleteOptions)
	PutElasticIpPool(value interface{})
	PutImages(value interface{})
	PutInstanceMetadataOptions(value *OceanAwsLaunchSpecInstanceMetadataOptions)
	PutLabels(value interface{})
	PutResourceLimits(value interface{})
	PutSchedulingShutdownHours(value *OceanAwsLaunchSpecSchedulingShutdownHours)
	PutSchedulingTask(value interface{})
	PutStrategy(value interface{})
	PutTags(value interface{})
	PutTaints(value interface{})
	PutUpdatePolicy(value *OceanAwsLaunchSpecUpdatePolicy)
	ResetAssociatePublicIpAddress()
	ResetAutoscaleDown()
	ResetAutoscaleHeadrooms()
	ResetAutoscaleHeadroomsAutomatic()
	ResetBlockDeviceMappings()
	ResetCreateOptions()
	ResetDeleteOptions()
	ResetElasticIpPool()
	ResetIamInstanceProfile()
	ResetId()
	ResetImageId()
	ResetImages()
	ResetInstanceMetadataOptions()
	ResetInstanceTypes()
	ResetLabels()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreferredSpotTypes()
	ResetResourceLimits()
	ResetRestrictScaleDown()
	ResetRootVolumeSize()
	ResetSchedulingShutdownHours()
	ResetSchedulingTask()
	ResetSecurityGroups()
	ResetStrategy()
	ResetSubnetIds()
	ResetTags()
	ResetTaints()
	ResetUpdatePolicy()
	ResetUserData()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OceanAwsLaunchSpec
type jsiiProxy_OceanAwsLaunchSpec struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanAwsLaunchSpec) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) AutoscaleDown() OceanAwsLaunchSpecAutoscaleDownList {
	var returns OceanAwsLaunchSpecAutoscaleDownList
	_jsii_.Get(
		j,
		"autoscaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) AutoscaleDownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) AutoscaleHeadrooms() OceanAwsLaunchSpecAutoscaleHeadroomsList {
	var returns OceanAwsLaunchSpecAutoscaleHeadroomsList
	_jsii_.Get(
		j,
		"autoscaleHeadrooms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) AutoscaleHeadroomsAutomatic() OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList {
	var returns OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList
	_jsii_.Get(
		j,
		"autoscaleHeadroomsAutomatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) AutoscaleHeadroomsAutomaticInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleHeadroomsAutomaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) AutoscaleHeadroomsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleHeadroomsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) BlockDeviceMappings() OceanAwsLaunchSpecBlockDeviceMappingsList {
	var returns OceanAwsLaunchSpecBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) CreateOptions() OceanAwsLaunchSpecCreateOptionsOutputReference {
	var returns OceanAwsLaunchSpecCreateOptionsOutputReference
	_jsii_.Get(
		j,
		"createOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) CreateOptionsInput() *OceanAwsLaunchSpecCreateOptions {
	var returns *OceanAwsLaunchSpecCreateOptions
	_jsii_.Get(
		j,
		"createOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) DeleteOptions() OceanAwsLaunchSpecDeleteOptionsOutputReference {
	var returns OceanAwsLaunchSpecDeleteOptionsOutputReference
	_jsii_.Get(
		j,
		"deleteOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) DeleteOptionsInput() *OceanAwsLaunchSpecDeleteOptions {
	var returns *OceanAwsLaunchSpecDeleteOptions
	_jsii_.Get(
		j,
		"deleteOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) ElasticIpPool() OceanAwsLaunchSpecElasticIpPoolList {
	var returns OceanAwsLaunchSpecElasticIpPoolList
	_jsii_.Get(
		j,
		"elasticIpPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) ElasticIpPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticIpPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Images() OceanAwsLaunchSpecImagesList {
	var returns OceanAwsLaunchSpecImagesList
	_jsii_.Get(
		j,
		"images",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) ImagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) InstanceMetadataOptions() OceanAwsLaunchSpecInstanceMetadataOptionsOutputReference {
	var returns OceanAwsLaunchSpecInstanceMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"instanceMetadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) InstanceMetadataOptionsInput() *OceanAwsLaunchSpecInstanceMetadataOptions {
	var returns *OceanAwsLaunchSpecInstanceMetadataOptions
	_jsii_.Get(
		j,
		"instanceMetadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Labels() OceanAwsLaunchSpecLabelsList {
	var returns OceanAwsLaunchSpecLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) OceanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) OceanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) PreferredSpotTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredSpotTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) PreferredSpotTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredSpotTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) ResourceLimits() OceanAwsLaunchSpecResourceLimitsList {
	var returns OceanAwsLaunchSpecResourceLimitsList
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) ResourceLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) RestrictScaleDown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictScaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) RestrictScaleDownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictScaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) RootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) RootVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) SchedulingShutdownHours() OceanAwsLaunchSpecSchedulingShutdownHoursOutputReference {
	var returns OceanAwsLaunchSpecSchedulingShutdownHoursOutputReference
	_jsii_.Get(
		j,
		"schedulingShutdownHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) SchedulingShutdownHoursInput() *OceanAwsLaunchSpecSchedulingShutdownHours {
	var returns *OceanAwsLaunchSpecSchedulingShutdownHours
	_jsii_.Get(
		j,
		"schedulingShutdownHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) SchedulingTask() OceanAwsLaunchSpecSchedulingTaskList {
	var returns OceanAwsLaunchSpecSchedulingTaskList
	_jsii_.Get(
		j,
		"schedulingTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) SchedulingTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedulingTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Strategy() OceanAwsLaunchSpecStrategyList {
	var returns OceanAwsLaunchSpecStrategyList
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) StrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Tags() OceanAwsLaunchSpecTagsList {
	var returns OceanAwsLaunchSpecTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) Taints() OceanAwsLaunchSpecTaintsList {
	var returns OceanAwsLaunchSpecTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) UpdatePolicy() OceanAwsLaunchSpecUpdatePolicyOutputReference {
	var returns OceanAwsLaunchSpecUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) UpdatePolicyInput() *OceanAwsLaunchSpecUpdatePolicy {
	var returns *OceanAwsLaunchSpecUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpec) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.123.0/docs/resources/ocean_aws_launch_spec spotinst_ocean_aws_launch_spec} Resource.
func NewOceanAwsLaunchSpec(scope constructs.Construct, id *string, config *OceanAwsLaunchSpecConfig) OceanAwsLaunchSpec {
	_init_.Initialize()

	if err := validateNewOceanAwsLaunchSpecParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAwsLaunchSpec{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAwsLaunchSpec.OceanAwsLaunchSpec",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.123.0/docs/resources/ocean_aws_launch_spec spotinst_ocean_aws_launch_spec} Resource.
func NewOceanAwsLaunchSpec_Override(o OceanAwsLaunchSpec, scope constructs.Construct, id *string, config *OceanAwsLaunchSpecConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAwsLaunchSpec.OceanAwsLaunchSpec",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetOceanId(val *string) {
	if err := j.validateSetOceanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oceanId",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetPreferredSpotTypes(val *[]*string) {
	if err := j.validateSetPreferredSpotTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredSpotTypes",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetRestrictScaleDown(val interface{}) {
	if err := j.validateSetRestrictScaleDownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictScaleDown",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetRootVolumeSize(val *float64) {
	if err := j.validateSetRootVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpec)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
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
func OceanAwsLaunchSpec_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAwsLaunchSpec_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAwsLaunchSpec.OceanAwsLaunchSpec",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAwsLaunchSpec_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAwsLaunchSpec_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAwsLaunchSpec.OceanAwsLaunchSpec",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAwsLaunchSpec_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAwsLaunchSpec_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAwsLaunchSpec.OceanAwsLaunchSpec",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanAwsLaunchSpec_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanAwsLaunchSpec.OceanAwsLaunchSpec",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanAwsLaunchSpec) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAwsLaunchSpec) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutAutoscaleDown(value interface{}) {
	if err := o.validatePutAutoscaleDownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleDown",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutAutoscaleHeadrooms(value interface{}) {
	if err := o.validatePutAutoscaleHeadroomsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleHeadrooms",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutAutoscaleHeadroomsAutomatic(value interface{}) {
	if err := o.validatePutAutoscaleHeadroomsAutomaticParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaleHeadroomsAutomatic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutBlockDeviceMappings(value interface{}) {
	if err := o.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutCreateOptions(value *OceanAwsLaunchSpecCreateOptions) {
	if err := o.validatePutCreateOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCreateOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutDeleteOptions(value *OceanAwsLaunchSpecDeleteOptions) {
	if err := o.validatePutDeleteOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDeleteOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutElasticIpPool(value interface{}) {
	if err := o.validatePutElasticIpPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putElasticIpPool",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutImages(value interface{}) {
	if err := o.validatePutImagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putImages",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutInstanceMetadataOptions(value *OceanAwsLaunchSpecInstanceMetadataOptions) {
	if err := o.validatePutInstanceMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInstanceMetadataOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutLabels(value interface{}) {
	if err := o.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLabels",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutResourceLimits(value interface{}) {
	if err := o.validatePutResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putResourceLimits",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutSchedulingShutdownHours(value *OceanAwsLaunchSpecSchedulingShutdownHours) {
	if err := o.validatePutSchedulingShutdownHoursParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSchedulingShutdownHours",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutSchedulingTask(value interface{}) {
	if err := o.validatePutSchedulingTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSchedulingTask",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutStrategy(value interface{}) {
	if err := o.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStrategy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutTags(value interface{}) {
	if err := o.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTags",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutTaints(value interface{}) {
	if err := o.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTaints",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) PutUpdatePolicy(value *OceanAwsLaunchSpecUpdatePolicy) {
	if err := o.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		o,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetAutoscaleDown() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleDown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetAutoscaleHeadrooms() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleHeadrooms",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetAutoscaleHeadroomsAutomatic() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaleHeadroomsAutomatic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		o,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetCreateOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetCreateOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetDeleteOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetDeleteOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetElasticIpPool() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticIpPool",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		o,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetImageId() {
	_jsii_.InvokeVoid(
		o,
		"resetImageId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetImages() {
	_jsii_.InvokeVoid(
		o,
		"resetImages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetInstanceMetadataOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceMetadataOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetInstanceTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetLabels() {
	_jsii_.InvokeVoid(
		o,
		"resetLabels",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetPreferredSpotTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetPreferredSpotTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		o,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetRestrictScaleDown() {
	_jsii_.InvokeVoid(
		o,
		"resetRestrictScaleDown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetRootVolumeSize() {
	_jsii_.InvokeVoid(
		o,
		"resetRootVolumeSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetSchedulingShutdownHours() {
	_jsii_.InvokeVoid(
		o,
		"resetSchedulingShutdownHours",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetSchedulingTask() {
	_jsii_.InvokeVoid(
		o,
		"resetSchedulingTask",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		o,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetStrategy() {
	_jsii_.InvokeVoid(
		o,
		"resetStrategy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		o,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetTaints() {
	_jsii_.InvokeVoid(
		o,
		"resetTaints",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ResetUserData() {
	_jsii_.InvokeVoid(
		o,
		"resetUserData",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpec) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsLaunchSpec) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

