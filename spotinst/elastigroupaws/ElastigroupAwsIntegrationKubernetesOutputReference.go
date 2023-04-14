package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsIntegrationKubernetesOutputReference interface {
	cdktf.ComplexObject
	ApiServer() *string
	SetApiServer(val *string)
	ApiServerInput() *string
	AutoscaleCooldown() *float64
	SetAutoscaleCooldown(val *float64)
	AutoscaleCooldownInput() *float64
	AutoscaleDown() ElastigroupAwsIntegrationKubernetesAutoscaleDownOutputReference
	AutoscaleDownInput() *ElastigroupAwsIntegrationKubernetesAutoscaleDown
	AutoscaleHeadroom() ElastigroupAwsIntegrationKubernetesAutoscaleHeadroomOutputReference
	AutoscaleHeadroomInput() *ElastigroupAwsIntegrationKubernetesAutoscaleHeadroom
	AutoscaleIsAutoConfig() interface{}
	SetAutoscaleIsAutoConfig(val interface{})
	AutoscaleIsAutoConfigInput() interface{}
	AutoscaleIsEnabled() interface{}
	SetAutoscaleIsEnabled(val interface{})
	AutoscaleIsEnabledInput() interface{}
	AutoscaleLabels() ElastigroupAwsIntegrationKubernetesAutoscaleLabelsList
	AutoscaleLabelsInput() interface{}
	ClusterIdentifier() *string
	SetClusterIdentifier(val *string)
	ClusterIdentifierInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IntegrationMode() *string
	SetIntegrationMode(val *string)
	IntegrationModeInput() *string
	InternalValue() *ElastigroupAwsIntegrationKubernetes
	SetInternalValue(val *ElastigroupAwsIntegrationKubernetes)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAutoscaleDown(value *ElastigroupAwsIntegrationKubernetesAutoscaleDown)
	PutAutoscaleHeadroom(value *ElastigroupAwsIntegrationKubernetesAutoscaleHeadroom)
	PutAutoscaleLabels(value interface{})
	ResetApiServer()
	ResetAutoscaleCooldown()
	ResetAutoscaleDown()
	ResetAutoscaleHeadroom()
	ResetAutoscaleIsAutoConfig()
	ResetAutoscaleIsEnabled()
	ResetAutoscaleLabels()
	ResetClusterIdentifier()
	ResetIntegrationMode()
	ResetToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsIntegrationKubernetesOutputReference
type jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ApiServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ApiServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleDown() ElastigroupAwsIntegrationKubernetesAutoscaleDownOutputReference {
	var returns ElastigroupAwsIntegrationKubernetesAutoscaleDownOutputReference
	_jsii_.Get(
		j,
		"autoscaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleDownInput() *ElastigroupAwsIntegrationKubernetesAutoscaleDown {
	var returns *ElastigroupAwsIntegrationKubernetesAutoscaleDown
	_jsii_.Get(
		j,
		"autoscaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleHeadroom() ElastigroupAwsIntegrationKubernetesAutoscaleHeadroomOutputReference {
	var returns ElastigroupAwsIntegrationKubernetesAutoscaleHeadroomOutputReference
	_jsii_.Get(
		j,
		"autoscaleHeadroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleHeadroomInput() *ElastigroupAwsIntegrationKubernetesAutoscaleHeadroom {
	var returns *ElastigroupAwsIntegrationKubernetesAutoscaleHeadroom
	_jsii_.Get(
		j,
		"autoscaleHeadroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleIsAutoConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleIsAutoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleIsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleIsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleLabels() ElastigroupAwsIntegrationKubernetesAutoscaleLabelsList {
	var returns ElastigroupAwsIntegrationKubernetesAutoscaleLabelsList
	_jsii_.Get(
		j,
		"autoscaleLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) AutoscaleLabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) IntegrationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) IntegrationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) InternalValue() *ElastigroupAwsIntegrationKubernetes {
	var returns *ElastigroupAwsIntegrationKubernetes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


func NewElastigroupAwsIntegrationKubernetesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsIntegrationKubernetesOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsIntegrationKubernetesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsIntegrationKubernetesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsIntegrationKubernetesOutputReference_Override(e ElastigroupAwsIntegrationKubernetesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsIntegrationKubernetesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetApiServer(val *string) {
	if err := j.validateSetApiServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiServer",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetAutoscaleCooldown(val *float64) {
	if err := j.validateSetAutoscaleCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleCooldown",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetAutoscaleIsAutoConfig(val interface{}) {
	if err := j.validateSetAutoscaleIsAutoConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleIsAutoConfig",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetAutoscaleIsEnabled(val interface{}) {
	if err := j.validateSetAutoscaleIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleIsEnabled",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetClusterIdentifier(val *string) {
	if err := j.validateSetClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetIntegrationMode(val *string) {
	if err := j.validateSetIntegrationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationMode",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetInternalValue(val *ElastigroupAwsIntegrationKubernetes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) PutAutoscaleDown(value *ElastigroupAwsIntegrationKubernetesAutoscaleDown) {
	if err := e.validatePutAutoscaleDownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleDown",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) PutAutoscaleHeadroom(value *ElastigroupAwsIntegrationKubernetesAutoscaleHeadroom) {
	if err := e.validatePutAutoscaleHeadroomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleHeadroom",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) PutAutoscaleLabels(value interface{}) {
	if err := e.validatePutAutoscaleLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleLabels",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetApiServer() {
	_jsii_.InvokeVoid(
		e,
		"resetApiServer",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetAutoscaleCooldown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleCooldown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetAutoscaleDown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleDown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetAutoscaleHeadroom() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleHeadroom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetAutoscaleIsAutoConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleIsAutoConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetAutoscaleIsEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleIsEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetAutoscaleLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetClusterIdentifier() {
	_jsii_.InvokeVoid(
		e,
		"resetClusterIdentifier",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetIntegrationMode() {
	_jsii_.InvokeVoid(
		e,
		"resetIntegrationMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ResetToken() {
	_jsii_.InvokeVoid(
		e,
		"resetToken",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationKubernetesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

