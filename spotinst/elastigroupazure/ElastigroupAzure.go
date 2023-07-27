package elastigroupazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v9/elastigroupazure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/spotinst/spotinst/1.129.0/docs/resources/elastigroup_azure spotinst_elastigroup_azure}.
type ElastigroupAzure interface {
	cdktf.TerraformResource
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
	CustomData() *string
	SetCustomData(val *string)
	CustomDataInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredCapacity() *float64
	SetDesiredCapacity(val *float64)
	DesiredCapacityInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheck() ElastigroupAzureHealthCheckOutputReference
	HealthCheckInput() *ElastigroupAzureHealthCheck
	Id() *string
	SetId(val *string)
	IdInput() *string
	Image() ElastigroupAzureImageList
	ImageInput() interface{}
	IntegrationKubernetes() ElastigroupAzureIntegrationKubernetesOutputReference
	IntegrationKubernetesInput() *ElastigroupAzureIntegrationKubernetes
	IntegrationMultaiRuntime() ElastigroupAzureIntegrationMultaiRuntimeOutputReference
	IntegrationMultaiRuntimeInput() *ElastigroupAzureIntegrationMultaiRuntime
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancers() ElastigroupAzureLoadBalancersList
	LoadBalancersInput() interface{}
	Login() ElastigroupAzureLoginOutputReference
	LoginInput() *ElastigroupAzureLogin
	LowPrioritySizes() *[]*string
	SetLowPrioritySizes(val *[]*string)
	LowPrioritySizesInput() *[]*string
	ManagedServiceIdentities() ElastigroupAzureManagedServiceIdentitiesList
	ManagedServiceIdentitiesInput() interface{}
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() ElastigroupAzureNetworkOutputReference
	NetworkInput() *ElastigroupAzureNetwork
	// The tree node.
	Node() constructs.Node
	OdSizes() *[]*string
	SetOdSizes(val *[]*string)
	OdSizesInput() *[]*string
	Product() *string
	SetProduct(val *string)
	ProductInput() *string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ScalingDownPolicy() ElastigroupAzureScalingDownPolicyList
	ScalingDownPolicyInput() interface{}
	ScalingUpPolicy() ElastigroupAzureScalingUpPolicyList
	ScalingUpPolicyInput() interface{}
	ScheduledTask() ElastigroupAzureScheduledTaskList
	ScheduledTaskInput() interface{}
	ShutdownScript() *string
	SetShutdownScript(val *string)
	ShutdownScriptInput() *string
	Strategy() ElastigroupAzureStrategyOutputReference
	StrategyInput() *ElastigroupAzureStrategy
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatePolicy() ElastigroupAzureUpdatePolicyOutputReference
	UpdatePolicyInput() *ElastigroupAzureUpdatePolicy
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
	PutHealthCheck(value *ElastigroupAzureHealthCheck)
	PutImage(value interface{})
	PutIntegrationKubernetes(value *ElastigroupAzureIntegrationKubernetes)
	PutIntegrationMultaiRuntime(value *ElastigroupAzureIntegrationMultaiRuntime)
	PutLoadBalancers(value interface{})
	PutLogin(value *ElastigroupAzureLogin)
	PutManagedServiceIdentities(value interface{})
	PutNetwork(value *ElastigroupAzureNetwork)
	PutScalingDownPolicy(value interface{})
	PutScalingUpPolicy(value interface{})
	PutScheduledTask(value interface{})
	PutStrategy(value *ElastigroupAzureStrategy)
	PutUpdatePolicy(value *ElastigroupAzureUpdatePolicy)
	ResetCustomData()
	ResetDesiredCapacity()
	ResetHealthCheck()
	ResetId()
	ResetImage()
	ResetIntegrationKubernetes()
	ResetIntegrationMultaiRuntime()
	ResetLoadBalancers()
	ResetLogin()
	ResetManagedServiceIdentities()
	ResetMaxSize()
	ResetMinSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScalingDownPolicy()
	ResetScalingUpPolicy()
	ResetScheduledTask()
	ResetShutdownScript()
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

// The jsii proxy struct for ElastigroupAzure
type jsiiProxy_ElastigroupAzure struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElastigroupAzure) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) CustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) CustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) HealthCheck() ElastigroupAzureHealthCheckOutputReference {
	var returns ElastigroupAzureHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) HealthCheckInput() *ElastigroupAzureHealthCheck {
	var returns *ElastigroupAzureHealthCheck
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Image() ElastigroupAzureImageList {
	var returns ElastigroupAzureImageList
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) IntegrationKubernetes() ElastigroupAzureIntegrationKubernetesOutputReference {
	var returns ElastigroupAzureIntegrationKubernetesOutputReference
	_jsii_.Get(
		j,
		"integrationKubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) IntegrationKubernetesInput() *ElastigroupAzureIntegrationKubernetes {
	var returns *ElastigroupAzureIntegrationKubernetes
	_jsii_.Get(
		j,
		"integrationKubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) IntegrationMultaiRuntime() ElastigroupAzureIntegrationMultaiRuntimeOutputReference {
	var returns ElastigroupAzureIntegrationMultaiRuntimeOutputReference
	_jsii_.Get(
		j,
		"integrationMultaiRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) IntegrationMultaiRuntimeInput() *ElastigroupAzureIntegrationMultaiRuntime {
	var returns *ElastigroupAzureIntegrationMultaiRuntime
	_jsii_.Get(
		j,
		"integrationMultaiRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) LoadBalancers() ElastigroupAzureLoadBalancersList {
	var returns ElastigroupAzureLoadBalancersList
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) LoadBalancersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Login() ElastigroupAzureLoginOutputReference {
	var returns ElastigroupAzureLoginOutputReference
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) LoginInput() *ElastigroupAzureLogin {
	var returns *ElastigroupAzureLogin
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) LowPrioritySizes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lowPrioritySizes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) LowPrioritySizesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lowPrioritySizesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ManagedServiceIdentities() ElastigroupAzureManagedServiceIdentitiesList {
	var returns ElastigroupAzureManagedServiceIdentitiesList
	_jsii_.Get(
		j,
		"managedServiceIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ManagedServiceIdentitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedServiceIdentitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Network() ElastigroupAzureNetworkOutputReference {
	var returns ElastigroupAzureNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) NetworkInput() *ElastigroupAzureNetwork {
	var returns *ElastigroupAzureNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) OdSizes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"odSizes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) OdSizesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"odSizesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Product() *string {
	var returns *string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ProductInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ScalingDownPolicy() ElastigroupAzureScalingDownPolicyList {
	var returns ElastigroupAzureScalingDownPolicyList
	_jsii_.Get(
		j,
		"scalingDownPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ScalingDownPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingDownPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ScalingUpPolicy() ElastigroupAzureScalingUpPolicyList {
	var returns ElastigroupAzureScalingUpPolicyList
	_jsii_.Get(
		j,
		"scalingUpPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ScalingUpPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingUpPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ScheduledTask() ElastigroupAzureScheduledTaskList {
	var returns ElastigroupAzureScheduledTaskList
	_jsii_.Get(
		j,
		"scheduledTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ScheduledTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ShutdownScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) ShutdownScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) Strategy() ElastigroupAzureStrategyOutputReference {
	var returns ElastigroupAzureStrategyOutputReference
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) StrategyInput() *ElastigroupAzureStrategy {
	var returns *ElastigroupAzureStrategy
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) UpdatePolicy() ElastigroupAzureUpdatePolicyOutputReference {
	var returns ElastigroupAzureUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) UpdatePolicyInput() *ElastigroupAzureUpdatePolicy {
	var returns *ElastigroupAzureUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAzure) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.129.0/docs/resources/elastigroup_azure spotinst_elastigroup_azure} Resource.
func NewElastigroupAzure(scope constructs.Construct, id *string, config *ElastigroupAzureConfig) ElastigroupAzure {
	_init_.Initialize()

	if err := validateNewElastigroupAzureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAzure{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzure.ElastigroupAzure",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/spotinst/spotinst/1.129.0/docs/resources/elastigroup_azure spotinst_elastigroup_azure} Resource.
func NewElastigroupAzure_Override(e ElastigroupAzure, scope constructs.Construct, id *string, config *ElastigroupAzureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAzure.ElastigroupAzure",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetCustomData(val *string) {
	if err := j.validateSetCustomDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customData",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetLowPrioritySizes(val *[]*string) {
	if err := j.validateSetLowPrioritySizesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lowPrioritySizes",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetOdSizes(val *[]*string) {
	if err := j.validateSetOdSizesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odSizes",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetProduct(val *string) {
	if err := j.validateSetProductParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"product",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetShutdownScript(val *string) {
	if err := j.validateSetShutdownScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownScript",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAzure)SetUserData(val *string) {
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
func ElastigroupAzure_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAzure_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAzure.ElastigroupAzure",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupAzure_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAzure_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAzure.ElastigroupAzure",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastigroupAzure_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastigroupAzure_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-spotinst.elastigroupAzure.ElastigroupAzure",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElastigroupAzure_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-spotinst.elastigroupAzure.ElastigroupAzure",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElastigroupAzure) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAzure) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzure) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAzure) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAzure) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAzure) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAzure) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAzure) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAzure) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAzure) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAzure) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutHealthCheck(value *ElastigroupAzureHealthCheck) {
	if err := e.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutImage(value interface{}) {
	if err := e.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putImage",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutIntegrationKubernetes(value *ElastigroupAzureIntegrationKubernetes) {
	if err := e.validatePutIntegrationKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationKubernetes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutIntegrationMultaiRuntime(value *ElastigroupAzureIntegrationMultaiRuntime) {
	if err := e.validatePutIntegrationMultaiRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntegrationMultaiRuntime",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutLoadBalancers(value interface{}) {
	if err := e.validatePutLoadBalancersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLoadBalancers",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutLogin(value *ElastigroupAzureLogin) {
	if err := e.validatePutLoginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogin",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutManagedServiceIdentities(value interface{}) {
	if err := e.validatePutManagedServiceIdentitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedServiceIdentities",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutNetwork(value *ElastigroupAzureNetwork) {
	if err := e.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetwork",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutScalingDownPolicy(value interface{}) {
	if err := e.validatePutScalingDownPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingDownPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutScalingUpPolicy(value interface{}) {
	if err := e.validatePutScalingUpPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingUpPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutScheduledTask(value interface{}) {
	if err := e.validatePutScheduledTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScheduledTask",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutStrategy(value *ElastigroupAzureStrategy) {
	if err := e.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStrategy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) PutUpdatePolicy(value *ElastigroupAzureUpdatePolicy) {
	if err := e.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetCustomData() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomData",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetImage() {
	_jsii_.InvokeVoid(
		e,
		"resetImage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetIntegrationKubernetes() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationKubernetes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetIntegrationMultaiRuntime() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationMultaiRuntime",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetLoadBalancers() {
	_jsii_.InvokeVoid(
		e,
		"resetLoadBalancers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetLogin() {
	_jsii_.InvokeVoid(
		e,
		"resetLogin",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetManagedServiceIdentities() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedServiceIdentities",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetMaxSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetMinSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMinSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetScalingDownPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingDownPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetScalingUpPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingUpPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetScheduledTask() {
	_jsii_.InvokeVoid(
		e,
		"resetScheduledTask",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetShutdownScript() {
	_jsii_.InvokeVoid(
		e,
		"resetShutdownScript",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) ResetUserData() {
	_jsii_.InvokeVoid(
		e,
		"resetUserData",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAzure) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzure) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzure) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAzure) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

