// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsIntegrationCodedeployOutputReference interface {
	cdktf.ComplexObject
	CleanupOnFailure() interface{}
	SetCleanupOnFailure(val interface{})
	CleanupOnFailureInput() interface{}
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
	DeploymentGroups() ElastigroupAwsIntegrationCodedeployDeploymentGroupsList
	DeploymentGroupsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ElastigroupAwsIntegrationCodedeploy
	SetInternalValue(val *ElastigroupAwsIntegrationCodedeploy)
	TerminateInstanceOnFailure() interface{}
	SetTerminateInstanceOnFailure(val interface{})
	TerminateInstanceOnFailureInput() interface{}
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
	PutDeploymentGroups(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsIntegrationCodedeployOutputReference
type jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) CleanupOnFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) CleanupOnFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupOnFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) DeploymentGroups() ElastigroupAwsIntegrationCodedeployDeploymentGroupsList {
	var returns ElastigroupAwsIntegrationCodedeployDeploymentGroupsList
	_jsii_.Get(
		j,
		"deploymentGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) DeploymentGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) InternalValue() *ElastigroupAwsIntegrationCodedeploy {
	var returns *ElastigroupAwsIntegrationCodedeploy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) TerminateInstanceOnFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstanceOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) TerminateInstanceOnFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstanceOnFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsIntegrationCodedeployOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsIntegrationCodedeployOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsIntegrationCodedeployOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsIntegrationCodedeployOutputReference_Override(e ElastigroupAwsIntegrationCodedeployOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsIntegrationCodedeployOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) SetCleanupOnFailure(val interface{}) {
	_jsii_.Set(
		j,
		"cleanupOnFailure",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) SetInternalValue(val *ElastigroupAwsIntegrationCodedeploy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) SetTerminateInstanceOnFailure(val interface{}) {
	_jsii_.Set(
		j,
		"terminateInstanceOnFailure",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) PutDeploymentGroups(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putDeploymentGroups",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationCodedeployOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

