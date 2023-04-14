package elastigroupgke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/elastigroupgke/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupGkeIntegrationGkeOutputReference interface {
	cdktf.ComplexObject
	AutoscaleCooldown() *float64
	SetAutoscaleCooldown(val *float64)
	AutoscaleCooldownInput() *float64
	AutoscaleDown() ElastigroupGkeIntegrationGkeAutoscaleDownOutputReference
	AutoscaleDownInput() *ElastigroupGkeIntegrationGkeAutoscaleDown
	AutoscaleHeadroom() ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference
	AutoscaleHeadroomInput() *ElastigroupGkeIntegrationGkeAutoscaleHeadroom
	AutoscaleIsAutoConfig() interface{}
	SetAutoscaleIsAutoConfig(val interface{})
	AutoscaleIsAutoConfigInput() interface{}
	AutoscaleIsEnabled() interface{}
	SetAutoscaleIsEnabled(val interface{})
	AutoscaleIsEnabledInput() interface{}
	AutoscaleLabels() ElastigroupGkeIntegrationGkeAutoscaleLabelsList
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
	InternalValue() *ElastigroupGkeIntegrationGke
	SetInternalValue(val *ElastigroupGkeIntegrationGke)
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
	PutAutoscaleDown(value *ElastigroupGkeIntegrationGkeAutoscaleDown)
	PutAutoscaleHeadroom(value *ElastigroupGkeIntegrationGkeAutoscaleHeadroom)
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

// The jsii proxy struct for ElastigroupGkeIntegrationGkeOutputReference
type jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleDown() ElastigroupGkeIntegrationGkeAutoscaleDownOutputReference {
	var returns ElastigroupGkeIntegrationGkeAutoscaleDownOutputReference
	_jsii_.Get(
		j,
		"autoscaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleDownInput() *ElastigroupGkeIntegrationGkeAutoscaleDown {
	var returns *ElastigroupGkeIntegrationGkeAutoscaleDown
	_jsii_.Get(
		j,
		"autoscaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleHeadroom() ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference {
	var returns ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference
	_jsii_.Get(
		j,
		"autoscaleHeadroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleHeadroomInput() *ElastigroupGkeIntegrationGkeAutoscaleHeadroom {
	var returns *ElastigroupGkeIntegrationGkeAutoscaleHeadroom
	_jsii_.Get(
		j,
		"autoscaleHeadroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleIsAutoConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleIsAutoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleIsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleIsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleLabels() ElastigroupGkeIntegrationGkeAutoscaleLabelsList {
	var returns ElastigroupGkeIntegrationGkeAutoscaleLabelsList
	_jsii_.Get(
		j,
		"autoscaleLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoscaleLabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) AutoUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) InternalValue() *ElastigroupGkeIntegrationGke {
	var returns *ElastigroupGkeIntegrationGke
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupGkeIntegrationGkeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupGkeIntegrationGkeOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupGkeIntegrationGkeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGkeIntegrationGkeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupGkeIntegrationGkeOutputReference_Override(e ElastigroupGkeIntegrationGkeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGkeIntegrationGkeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetAutoscaleCooldown(val *float64) {
	if err := j.validateSetAutoscaleCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleCooldown",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetAutoscaleIsAutoConfig(val interface{}) {
	if err := j.validateSetAutoscaleIsAutoConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleIsAutoConfig",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetAutoscaleIsEnabled(val interface{}) {
	if err := j.validateSetAutoscaleIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleIsEnabled",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetAutoUpdate(val interface{}) {
	if err := j.validateSetAutoUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoUpdate",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetInternalValue(val *ElastigroupGkeIntegrationGke) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) PutAutoscaleDown(value *ElastigroupGkeIntegrationGkeAutoscaleDown) {
	if err := e.validatePutAutoscaleDownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleDown",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) PutAutoscaleHeadroom(value *ElastigroupGkeIntegrationGkeAutoscaleHeadroom) {
	if err := e.validatePutAutoscaleHeadroomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleHeadroom",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) PutAutoscaleLabels(value interface{}) {
	if err := e.validatePutAutoscaleLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleLabels",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ResetAutoscaleCooldown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleCooldown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ResetAutoscaleDown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleDown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ResetAutoscaleHeadroom() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleHeadroom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ResetAutoscaleIsAutoConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleIsAutoConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ResetAutoscaleIsEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleIsEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ResetAutoscaleLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ResetAutoUpdate() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoUpdate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		e,
		"resetClusterId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		e,
		"resetLocation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

