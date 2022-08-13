// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupGcpIntegrationGkeOutputReference interface {
	cdktf.ComplexObject
	AutoscaleCooldown() *float64
	SetAutoscaleCooldown(val *float64)
	AutoscaleCooldownInput() *float64
	AutoscaleDown() ElastigroupGcpIntegrationGkeAutoscaleDownOutputReference
	AutoscaleDownInput() *ElastigroupGcpIntegrationGkeAutoscaleDown
	AutoscaleHeadroom() ElastigroupGcpIntegrationGkeAutoscaleHeadroomOutputReference
	AutoscaleHeadroomInput() *ElastigroupGcpIntegrationGkeAutoscaleHeadroom
	AutoscaleIsAutoConfig() interface{}
	SetAutoscaleIsAutoConfig(val interface{})
	AutoscaleIsAutoConfigInput() interface{}
	AutoscaleIsEnabled() interface{}
	SetAutoscaleIsEnabled(val interface{})
	AutoscaleIsEnabledInput() interface{}
	AutoscaleLabels() ElastigroupGcpIntegrationGkeAutoscaleLabelsList
	AutoscaleLabelsInput() interface{}
	AutoUpdate() interface{}
	SetAutoUpdate(val interface{})
	AutoUpdateInput() interface{}
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	InternalValue() *ElastigroupGcpIntegrationGke
	SetInternalValue(val *ElastigroupGcpIntegrationGke)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAutoscaleDown(value *ElastigroupGcpIntegrationGkeAutoscaleDown)
	PutAutoscaleHeadroom(value *ElastigroupGcpIntegrationGkeAutoscaleHeadroom)
	PutAutoscaleLabels(value interface{})
	ResetAutoscaleCooldown()
	ResetAutoscaleDown()
	ResetAutoscaleHeadroom()
	ResetAutoscaleIsAutoConfig()
	ResetAutoscaleIsEnabled()
	ResetAutoscaleLabels()
	ResetAutoUpdate()
	ResetClusterId()
	ResetLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupGcpIntegrationGkeOutputReference
type jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleDown() ElastigroupGcpIntegrationGkeAutoscaleDownOutputReference {
	var returns ElastigroupGcpIntegrationGkeAutoscaleDownOutputReference
	_jsii_.Get(
		j,
		"autoscaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleDownInput() *ElastigroupGcpIntegrationGkeAutoscaleDown {
	var returns *ElastigroupGcpIntegrationGkeAutoscaleDown
	_jsii_.Get(
		j,
		"autoscaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleHeadroom() ElastigroupGcpIntegrationGkeAutoscaleHeadroomOutputReference {
	var returns ElastigroupGcpIntegrationGkeAutoscaleHeadroomOutputReference
	_jsii_.Get(
		j,
		"autoscaleHeadroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleHeadroomInput() *ElastigroupGcpIntegrationGkeAutoscaleHeadroom {
	var returns *ElastigroupGcpIntegrationGkeAutoscaleHeadroom
	_jsii_.Get(
		j,
		"autoscaleHeadroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleIsAutoConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleIsAutoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleIsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleIsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleLabels() ElastigroupGcpIntegrationGkeAutoscaleLabelsList {
	var returns ElastigroupGcpIntegrationGkeAutoscaleLabelsList
	_jsii_.Get(
		j,
		"autoscaleLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoscaleLabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) AutoUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) InternalValue() *ElastigroupGcpIntegrationGke {
	var returns *ElastigroupGcpIntegrationGke
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupGcpIntegrationGkeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupGcpIntegrationGkeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupGcpIntegrationGkeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupGcpIntegrationGkeOutputReference_Override(e ElastigroupGcpIntegrationGkeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupGcpIntegrationGkeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetAutoscaleCooldown(val *float64) {
	_jsii_.Set(
		j,
		"autoscaleCooldown",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetAutoscaleIsAutoConfig(val interface{}) {
	_jsii_.Set(
		j,
		"autoscaleIsAutoConfig",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetAutoscaleIsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"autoscaleIsEnabled",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetAutoUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"autoUpdate",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetInternalValue(val *ElastigroupGcpIntegrationGke) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) PutAutoscaleDown(value *ElastigroupGcpIntegrationGkeAutoscaleDown) {
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleDown",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) PutAutoscaleHeadroom(value *ElastigroupGcpIntegrationGkeAutoscaleHeadroom) {
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleHeadroom",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) PutAutoscaleLabels(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleLabels",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ResetAutoscaleCooldown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleCooldown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ResetAutoscaleDown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleDown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ResetAutoscaleHeadroom() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleHeadroom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ResetAutoscaleIsAutoConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleIsAutoConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ResetAutoscaleIsEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleIsEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ResetAutoscaleLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ResetAutoUpdate() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoUpdate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		e,
		"resetClusterId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		e,
		"resetLocation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGcpIntegrationGkeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

