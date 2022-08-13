// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsIntegrationEcsOutputReference interface {
	cdktf.ComplexObject
	AutoscaleAttributes() ElastigroupAwsIntegrationEcsAutoscaleAttributesList
	AutoscaleAttributesInput() interface{}
	AutoscaleCooldown() *float64
	SetAutoscaleCooldown(val *float64)
	AutoscaleCooldownInput() *float64
	AutoscaleDown() ElastigroupAwsIntegrationEcsAutoscaleDownOutputReference
	AutoscaleDownInput() *ElastigroupAwsIntegrationEcsAutoscaleDown
	AutoscaleHeadroom() ElastigroupAwsIntegrationEcsAutoscaleHeadroomOutputReference
	AutoscaleHeadroomInput() *ElastigroupAwsIntegrationEcsAutoscaleHeadroom
	AutoscaleIsAutoConfig() interface{}
	SetAutoscaleIsAutoConfig(val interface{})
	AutoscaleIsAutoConfigInput() interface{}
	AutoscaleIsEnabled() interface{}
	SetAutoscaleIsEnabled(val interface{})
	AutoscaleIsEnabledInput() interface{}
	AutoscaleScaleDownNonServiceTasks() interface{}
	SetAutoscaleScaleDownNonServiceTasks(val interface{})
	AutoscaleScaleDownNonServiceTasksInput() interface{}
	Batch() ElastigroupAwsIntegrationEcsBatchOutputReference
	BatchInput() *ElastigroupAwsIntegrationEcsBatch
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	InternalValue() *ElastigroupAwsIntegrationEcs
	SetInternalValue(val *ElastigroupAwsIntegrationEcs)
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
	PutAutoscaleAttributes(value interface{})
	PutAutoscaleDown(value *ElastigroupAwsIntegrationEcsAutoscaleDown)
	PutAutoscaleHeadroom(value *ElastigroupAwsIntegrationEcsAutoscaleHeadroom)
	PutBatch(value *ElastigroupAwsIntegrationEcsBatch)
	ResetAutoscaleAttributes()
	ResetAutoscaleCooldown()
	ResetAutoscaleDown()
	ResetAutoscaleHeadroom()
	ResetAutoscaleIsAutoConfig()
	ResetAutoscaleIsEnabled()
	ResetAutoscaleScaleDownNonServiceTasks()
	ResetBatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsIntegrationEcsOutputReference
type jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleAttributes() ElastigroupAwsIntegrationEcsAutoscaleAttributesList {
	var returns ElastigroupAwsIntegrationEcsAutoscaleAttributesList
	_jsii_.Get(
		j,
		"autoscaleAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleDown() ElastigroupAwsIntegrationEcsAutoscaleDownOutputReference {
	var returns ElastigroupAwsIntegrationEcsAutoscaleDownOutputReference
	_jsii_.Get(
		j,
		"autoscaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleDownInput() *ElastigroupAwsIntegrationEcsAutoscaleDown {
	var returns *ElastigroupAwsIntegrationEcsAutoscaleDown
	_jsii_.Get(
		j,
		"autoscaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleHeadroom() ElastigroupAwsIntegrationEcsAutoscaleHeadroomOutputReference {
	var returns ElastigroupAwsIntegrationEcsAutoscaleHeadroomOutputReference
	_jsii_.Get(
		j,
		"autoscaleHeadroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleHeadroomInput() *ElastigroupAwsIntegrationEcsAutoscaleHeadroom {
	var returns *ElastigroupAwsIntegrationEcsAutoscaleHeadroom
	_jsii_.Get(
		j,
		"autoscaleHeadroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleIsAutoConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleIsAutoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsAutoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleIsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleIsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleScaleDownNonServiceTasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleScaleDownNonServiceTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) AutoscaleScaleDownNonServiceTasksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleScaleDownNonServiceTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) Batch() ElastigroupAwsIntegrationEcsBatchOutputReference {
	var returns ElastigroupAwsIntegrationEcsBatchOutputReference
	_jsii_.Get(
		j,
		"batch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) BatchInput() *ElastigroupAwsIntegrationEcsBatch {
	var returns *ElastigroupAwsIntegrationEcsBatch
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) InternalValue() *ElastigroupAwsIntegrationEcs {
	var returns *ElastigroupAwsIntegrationEcs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsIntegrationEcsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsIntegrationEcsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsIntegrationEcsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsIntegrationEcsOutputReference_Override(e ElastigroupAwsIntegrationEcsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsIntegrationEcsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetAutoscaleCooldown(val *float64) {
	_jsii_.Set(
		j,
		"autoscaleCooldown",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetAutoscaleIsAutoConfig(val interface{}) {
	_jsii_.Set(
		j,
		"autoscaleIsAutoConfig",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetAutoscaleIsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"autoscaleIsEnabled",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetAutoscaleScaleDownNonServiceTasks(val interface{}) {
	_jsii_.Set(
		j,
		"autoscaleScaleDownNonServiceTasks",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetInternalValue(val *ElastigroupAwsIntegrationEcs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) PutAutoscaleAttributes(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleAttributes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) PutAutoscaleDown(value *ElastigroupAwsIntegrationEcsAutoscaleDown) {
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleDown",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) PutAutoscaleHeadroom(value *ElastigroupAwsIntegrationEcsAutoscaleHeadroom) {
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleHeadroom",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) PutBatch(value *ElastigroupAwsIntegrationEcsBatch) {
	_jsii_.InvokeVoid(
		e,
		"putBatch",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ResetAutoscaleAttributes() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleAttributes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ResetAutoscaleCooldown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleCooldown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ResetAutoscaleDown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleDown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ResetAutoscaleHeadroom() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleHeadroom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ResetAutoscaleIsAutoConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleIsAutoConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ResetAutoscaleIsEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleIsEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ResetAutoscaleScaleDownNonServiceTasks() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleScaleDownNonServiceTasks",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ResetBatch() {
	_jsii_.InvokeVoid(
		e,
		"resetBatch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationEcsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

