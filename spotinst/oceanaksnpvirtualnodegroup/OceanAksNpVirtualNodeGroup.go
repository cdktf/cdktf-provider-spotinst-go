package oceanaksnpvirtualnodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v9/oceanaksnpvirtualnodegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aks_np_virtual_node_group spotinst_ocean_aks_np_virtual_node_group}.
type OceanAksNpVirtualNodeGroup interface {
	cdktf.TerraformResource
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
	Filters() OceanAksNpVirtualNodeGroupFiltersOutputReference
	FiltersInput() *OceanAksNpVirtualNodeGroupFilters
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Headrooms() OceanAksNpVirtualNodeGroupHeadroomsList
	HeadroomsInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	OceanId() *string
	SetOceanId(val *string)
	OceanIdInput() *string
	OsDiskSizeGb() *float64
	SetOsDiskSizeGb(val *float64)
	OsDiskSizeGbInput() *float64
	OsDiskType() *string
	SetOsDiskType(val *string)
	OsDiskTypeInput() *string
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
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
	SpotPercentage() *float64
	SetSpotPercentage(val *float64)
	SpotPercentageInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Taints() OceanAksNpVirtualNodeGroupTaintsList
	TaintsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutFilters(value *OceanAksNpVirtualNodeGroupFilters)
	PutHeadrooms(value interface{})
	PutTaints(value interface{})
	ResetAvailabilityZones()
	ResetEnableNodePublicIp()
	ResetFallbackToOndemand()
	ResetFilters()
	ResetHeadrooms()
	ResetId()
	ResetLabels()
	ResetMaxCount()
	ResetMaxPodsPerNode()
	ResetMinCount()
	ResetOsDiskSizeGb()
	ResetOsDiskType()
	ResetOsType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSpotPercentage()
	ResetTags()
	ResetTaints()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OceanAksNpVirtualNodeGroup
type jsiiProxy_OceanAksNpVirtualNodeGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) EnableNodePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) EnableNodePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) FallbackToOndemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) FallbackToOndemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOndemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Filters() OceanAksNpVirtualNodeGroupFiltersOutputReference {
	var returns OceanAksNpVirtualNodeGroupFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) FiltersInput() *OceanAksNpVirtualNodeGroupFilters {
	var returns *OceanAksNpVirtualNodeGroupFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Headrooms() OceanAksNpVirtualNodeGroupHeadroomsList {
	var returns OceanAksNpVirtualNodeGroupHeadroomsList
	_jsii_.Get(
		j,
		"headrooms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) HeadroomsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headroomsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) MaxCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) MaxCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) MaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) MaxPodsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) MinCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) MinCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) OceanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) OceanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oceanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) OsDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) OsDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) OsDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) OsDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) SpotPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) SpotPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) Taints() OceanAksNpVirtualNodeGroupTaintsList {
	var returns OceanAksNpVirtualNodeGroupTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aks_np_virtual_node_group spotinst_ocean_aks_np_virtual_node_group} Resource.
func NewOceanAksNpVirtualNodeGroup(scope constructs.Construct, id *string, config *OceanAksNpVirtualNodeGroupConfig) OceanAksNpVirtualNodeGroup {
	_init_.Initialize()

	if err := validateNewOceanAksNpVirtualNodeGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAksNpVirtualNodeGroup{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aks_np_virtual_node_group spotinst_ocean_aks_np_virtual_node_group} Resource.
func NewOceanAksNpVirtualNodeGroup_Override(o OceanAksNpVirtualNodeGroup, scope constructs.Construct, id *string, config *OceanAksNpVirtualNodeGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroup",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetEnableNodePublicIp(val interface{}) {
	if err := j.validateSetEnableNodePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNodePublicIp",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetFallbackToOndemand(val interface{}) {
	if err := j.validateSetFallbackToOndemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackToOndemand",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetMaxCount(val *float64) {
	if err := j.validateSetMaxCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCount",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetMaxPodsPerNode(val *float64) {
	if err := j.validateSetMaxPodsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPodsPerNode",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetMinCount(val *float64) {
	if err := j.validateSetMinCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCount",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetOceanId(val *string) {
	if err := j.validateSetOceanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oceanId",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetOsDiskSizeGb(val *float64) {
	if err := j.validateSetOsDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetOsDiskType(val *string) {
	if err := j.validateSetOsDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskType",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetSpotPercentage(val *float64) {
	if err := j.validateSetSpotPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func OceanAksNpVirtualNodeGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAksNpVirtualNodeGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAksNpVirtualNodeGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAksNpVirtualNodeGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAksNpVirtualNodeGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAksNpVirtualNodeGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanAksNpVirtualNodeGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) PutFilters(value *OceanAksNpVirtualNodeGroupFilters) {
	if err := o.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFilters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) PutHeadrooms(value interface{}) {
	if err := o.validatePutHeadroomsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHeadrooms",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) PutTaints(value interface{}) {
	if err := o.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTaints",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		o,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetEnableNodePublicIp() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableNodePublicIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetFallbackToOndemand() {
	_jsii_.InvokeVoid(
		o,
		"resetFallbackToOndemand",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetFilters() {
	_jsii_.InvokeVoid(
		o,
		"resetFilters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetHeadrooms() {
	_jsii_.InvokeVoid(
		o,
		"resetHeadrooms",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetLabels() {
	_jsii_.InvokeVoid(
		o,
		"resetLabels",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetMaxCount() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetMaxPodsPerNode() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxPodsPerNode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetMinCount() {
	_jsii_.InvokeVoid(
		o,
		"resetMinCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetOsDiskSizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetOsDiskSizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetOsDiskType() {
	_jsii_.InvokeVoid(
		o,
		"resetOsDiskType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetOsType() {
	_jsii_.InvokeVoid(
		o,
		"resetOsType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetSpotPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetSpotPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ResetTaints() {
	_jsii_.InvokeVoid(
		o,
		"resetTaints",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

