package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsIntegrationDockerSwarmOutputReference interface {
	cdktf.ComplexObject
	AutoscaleCooldown() *float64
	SetAutoscaleCooldown(val *float64)
	AutoscaleCooldownInput() *float64
	AutoscaleDown() ElastigroupAwsIntegrationDockerSwarmAutoscaleDownOutputReference
	AutoscaleDownInput() *ElastigroupAwsIntegrationDockerSwarmAutoscaleDown
	AutoscaleHeadroom() ElastigroupAwsIntegrationDockerSwarmAutoscaleHeadroomOutputReference
	AutoscaleHeadroomInput() *ElastigroupAwsIntegrationDockerSwarmAutoscaleHeadroom
	AutoscaleIsEnabled() interface{}
	SetAutoscaleIsEnabled(val interface{})
	AutoscaleIsEnabledInput() interface{}
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
	InternalValue() *ElastigroupAwsIntegrationDockerSwarm
	SetInternalValue(val *ElastigroupAwsIntegrationDockerSwarm)
	MasterHost() *string
	SetMasterHost(val *string)
	MasterHostInput() *string
	MasterPort() *float64
	SetMasterPort(val *float64)
	MasterPortInput() *float64
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
	PutAutoscaleDown(value *ElastigroupAwsIntegrationDockerSwarmAutoscaleDown)
	PutAutoscaleHeadroom(value *ElastigroupAwsIntegrationDockerSwarmAutoscaleHeadroom)
	ResetAutoscaleCooldown()
	ResetAutoscaleDown()
	ResetAutoscaleHeadroom()
	ResetAutoscaleIsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsIntegrationDockerSwarmOutputReference
type jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) AutoscaleCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) AutoscaleCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscaleCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) AutoscaleDown() ElastigroupAwsIntegrationDockerSwarmAutoscaleDownOutputReference {
	var returns ElastigroupAwsIntegrationDockerSwarmAutoscaleDownOutputReference
	_jsii_.Get(
		j,
		"autoscaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) AutoscaleDownInput() *ElastigroupAwsIntegrationDockerSwarmAutoscaleDown {
	var returns *ElastigroupAwsIntegrationDockerSwarmAutoscaleDown
	_jsii_.Get(
		j,
		"autoscaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) AutoscaleHeadroom() ElastigroupAwsIntegrationDockerSwarmAutoscaleHeadroomOutputReference {
	var returns ElastigroupAwsIntegrationDockerSwarmAutoscaleHeadroomOutputReference
	_jsii_.Get(
		j,
		"autoscaleHeadroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) AutoscaleHeadroomInput() *ElastigroupAwsIntegrationDockerSwarmAutoscaleHeadroom {
	var returns *ElastigroupAwsIntegrationDockerSwarmAutoscaleHeadroom
	_jsii_.Get(
		j,
		"autoscaleHeadroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) AutoscaleIsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) AutoscaleIsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleIsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) InternalValue() *ElastigroupAwsIntegrationDockerSwarm {
	var returns *ElastigroupAwsIntegrationDockerSwarm
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) MasterHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) MasterHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) MasterPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) MasterPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsIntegrationDockerSwarmOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsIntegrationDockerSwarmOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsIntegrationDockerSwarmOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsIntegrationDockerSwarmOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsIntegrationDockerSwarmOutputReference_Override(e ElastigroupAwsIntegrationDockerSwarmOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsIntegrationDockerSwarmOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference)SetAutoscaleCooldown(val *float64) {
	if err := j.validateSetAutoscaleCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleCooldown",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference)SetAutoscaleIsEnabled(val interface{}) {
	if err := j.validateSetAutoscaleIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscaleIsEnabled",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference)SetInternalValue(val *ElastigroupAwsIntegrationDockerSwarm) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference)SetMasterHost(val *string) {
	if err := j.validateSetMasterHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterHost",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference)SetMasterPort(val *float64) {
	if err := j.validateSetMasterPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPort",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) PutAutoscaleDown(value *ElastigroupAwsIntegrationDockerSwarmAutoscaleDown) {
	if err := e.validatePutAutoscaleDownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleDown",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) PutAutoscaleHeadroom(value *ElastigroupAwsIntegrationDockerSwarmAutoscaleHeadroom) {
	if err := e.validatePutAutoscaleHeadroomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoscaleHeadroom",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) ResetAutoscaleCooldown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleCooldown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) ResetAutoscaleDown() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleDown",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) ResetAutoscaleHeadroom() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleHeadroom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) ResetAutoscaleIsEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscaleIsEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationDockerSwarmOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

