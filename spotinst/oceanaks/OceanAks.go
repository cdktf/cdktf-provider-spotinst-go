// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanaks/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.166.0/docs/resources/ocean_aks spotinst_ocean_aks}.
type OceanAks interface {
	cdktf.TerraformResource
	AcdIdentifier() *string
	SetAcdIdentifier(val *string)
	AcdIdentifierInput() *string
	AksName() *string
	SetAksName(val *string)
	AksNameInput() *string
	AksResourceGroupName() *string
	SetAksResourceGroupName(val *string)
	AksResourceGroupNameInput() *string
	Autoscaler() OceanAksAutoscalerOutputReference
	AutoscalerInput() *OceanAksAutoscaler
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
	CustomData() *string
	SetCustomData(val *string)
	CustomDataInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Extension() OceanAksExtensionList
	ExtensionInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Health() OceanAksHealthOutputReference
	HealthInput() *OceanAksHealth
	Id() *string
	SetId(val *string)
	IdInput() *string
	Image() OceanAksImageList
	ImageInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() OceanAksLoadBalancerList
	LoadBalancerInput() interface{}
	ManagedServiceIdentity() OceanAksManagedServiceIdentityList
	ManagedServiceIdentityInput() interface{}
	MaxPods() *float64
	SetMaxPods(val *float64)
	MaxPodsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() OceanAksNetworkOutputReference
	NetworkInput() *OceanAksNetwork
	// The tree node.
	Node() constructs.Node
	OsDisk() OceanAksOsDiskOutputReference
	OsDiskInput() *OceanAksOsDisk
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SshPublicKey() *string
	SetSshPublicKey(val *string)
	SshPublicKeyInput() *string
	Strategy() OceanAksStrategyList
	StrategyInput() interface{}
	Tag() OceanAksTagList
	TagInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
	VmSizes() OceanAksVmSizesList
	VmSizesInput() interface{}
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	PutAutoscaler(value *OceanAksAutoscaler)
	PutExtension(value interface{})
	PutHealth(value *OceanAksHealth)
	PutImage(value interface{})
	PutLoadBalancer(value interface{})
	PutManagedServiceIdentity(value interface{})
	PutNetwork(value *OceanAksNetwork)
	PutOsDisk(value *OceanAksOsDisk)
	PutStrategy(value interface{})
	PutTag(value interface{})
	PutVmSizes(value interface{})
	ResetAutoscaler()
	ResetControllerClusterId()
	ResetCustomData()
	ResetExtension()
	ResetHealth()
	ResetId()
	ResetImage()
	ResetLoadBalancer()
	ResetManagedServiceIdentity()
	ResetMaxPods()
	ResetNetwork()
	ResetOsDisk()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceGroupName()
	ResetStrategy()
	ResetTag()
	ResetUserName()
	ResetVmSizes()
	ResetZones()
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

// The jsii proxy struct for OceanAks
type jsiiProxy_OceanAks struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OceanAks) AcdIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acdIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) AcdIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acdIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) AksName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) AksNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) AksResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) AksResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aksResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Autoscaler() OceanAksAutoscalerOutputReference {
	var returns OceanAksAutoscalerOutputReference
	_jsii_.Get(
		j,
		"autoscaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) AutoscalerInput() *OceanAksAutoscaler {
	var returns *OceanAksAutoscaler
	_jsii_.Get(
		j,
		"autoscalerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ControllerClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ControllerClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) CustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) CustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Extension() OceanAksExtensionList {
	var returns OceanAksExtensionList
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Health() OceanAksHealthOutputReference {
	var returns OceanAksHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) HealthInput() *OceanAksHealth {
	var returns *OceanAksHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Image() OceanAksImageList {
	var returns OceanAksImageList
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) LoadBalancer() OceanAksLoadBalancerList {
	var returns OceanAksLoadBalancerList
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) LoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ManagedServiceIdentity() OceanAksManagedServiceIdentityList {
	var returns OceanAksManagedServiceIdentityList
	_jsii_.Get(
		j,
		"managedServiceIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ManagedServiceIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedServiceIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) MaxPods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) MaxPodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Network() OceanAksNetworkOutputReference {
	var returns OceanAksNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) NetworkInput() *OceanAksNetwork {
	var returns *OceanAksNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) OsDisk() OceanAksOsDiskOutputReference {
	var returns OceanAksOsDiskOutputReference
	_jsii_.Get(
		j,
		"osDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) OsDiskInput() *OceanAksOsDisk {
	var returns *OceanAksOsDisk
	_jsii_.Get(
		j,
		"osDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) SshPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) SshPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Strategy() OceanAksStrategyList {
	var returns OceanAksStrategyList
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) StrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Tag() OceanAksTagList {
	var returns OceanAksTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) VmSizes() OceanAksVmSizesList {
	var returns OceanAksVmSizesList
	_jsii_.Get(
		j,
		"vmSizes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) VmSizesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmSizesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAks) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.166.0/docs/resources/ocean_aks spotinst_ocean_aks} Resource.
func NewOceanAks(scope constructs.Construct, id *string, config *OceanAksConfig) OceanAks {
	_init_.Initialize()

	if err := validateNewOceanAksParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAks{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAks.OceanAks",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.166.0/docs/resources/ocean_aks spotinst_ocean_aks} Resource.
func NewOceanAks_Override(o OceanAks, scope constructs.Construct, id *string, config *OceanAksConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAks.OceanAks",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OceanAks)SetAcdIdentifier(val *string) {
	if err := j.validateSetAcdIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acdIdentifier",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetAksName(val *string) {
	if err := j.validateSetAksNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aksName",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetAksResourceGroupName(val *string) {
	if err := j.validateSetAksResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aksResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetControllerClusterId(val *string) {
	if err := j.validateSetControllerClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controllerClusterId",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetCustomData(val *string) {
	if err := j.validateSetCustomDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customData",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetMaxPods(val *float64) {
	if err := j.validateSetMaxPodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPods",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetSshPublicKey(val *string) {
	if err := j.validateSetSshPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKey",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetUserName(val *string) {
	if err := j.validateSetUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (j *jsiiProxy_OceanAks)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Generates CDKTF code for importing a OceanAks resource upon running "cdktf plan <stack-name>".
func OceanAks_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOceanAks_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAks.OceanAks",
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
func OceanAks_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAks_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAks.OceanAks",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAks_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAks_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAks.OceanAks",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OceanAks_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOceanAks_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.oceanAks.OceanAks",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OceanAks_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.oceanAks.OceanAks",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OceanAks) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OceanAks) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OceanAks) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAks) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAks) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAks) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAks) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAks) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAks) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAks) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAks) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAks) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAks) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OceanAks) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAks) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanAks) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OceanAks) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OceanAks) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OceanAks) PutAutoscaler(value *OceanAksAutoscaler) {
	if err := o.validatePutAutoscalerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAutoscaler",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutExtension(value interface{}) {
	if err := o.validatePutExtensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putExtension",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutHealth(value *OceanAksHealth) {
	if err := o.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHealth",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutImage(value interface{}) {
	if err := o.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putImage",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutLoadBalancer(value interface{}) {
	if err := o.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutManagedServiceIdentity(value interface{}) {
	if err := o.validatePutManagedServiceIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putManagedServiceIdentity",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutNetwork(value *OceanAksNetwork) {
	if err := o.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNetwork",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutOsDisk(value *OceanAksOsDisk) {
	if err := o.validatePutOsDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOsDisk",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutStrategy(value interface{}) {
	if err := o.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStrategy",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutTag(value interface{}) {
	if err := o.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTag",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) PutVmSizes(value interface{}) {
	if err := o.validatePutVmSizesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVmSizes",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanAks) ResetAutoscaler() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoscaler",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetControllerClusterId() {
	_jsii_.InvokeVoid(
		o,
		"resetControllerClusterId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetCustomData() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomData",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetExtension() {
	_jsii_.InvokeVoid(
		o,
		"resetExtension",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetHealth() {
	_jsii_.InvokeVoid(
		o,
		"resetHealth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetImage() {
	_jsii_.InvokeVoid(
		o,
		"resetImage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetManagedServiceIdentity() {
	_jsii_.InvokeVoid(
		o,
		"resetManagedServiceIdentity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetMaxPods() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxPods",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetNetwork() {
	_jsii_.InvokeVoid(
		o,
		"resetNetwork",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetOsDisk() {
	_jsii_.InvokeVoid(
		o,
		"resetOsDisk",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetResourceGroupName() {
	_jsii_.InvokeVoid(
		o,
		"resetResourceGroupName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetStrategy() {
	_jsii_.InvokeVoid(
		o,
		"resetStrategy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetTag() {
	_jsii_.InvokeVoid(
		o,
		"resetTag",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetUserName() {
	_jsii_.InvokeVoid(
		o,
		"resetUserName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetVmSizes() {
	_jsii_.InvokeVoid(
		o,
		"resetVmSizes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) ResetZones() {
	_jsii_.InvokeVoid(
		o,
		"resetZones",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAks) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAks) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAks) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAks) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAks) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAks) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

