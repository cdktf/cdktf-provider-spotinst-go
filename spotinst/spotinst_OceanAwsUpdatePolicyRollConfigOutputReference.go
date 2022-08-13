// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAwsUpdatePolicyRollConfigOutputReference interface {
	cdktf.ComplexObject
	BatchMinHealthyPercentage() *float64
	SetBatchMinHealthyPercentage(val *float64)
	BatchMinHealthyPercentageInput() *float64
	BatchSizePercentage() *float64
	SetBatchSizePercentage(val *float64)
	BatchSizePercentageInput() *float64
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
	InternalValue() *OceanAwsUpdatePolicyRollConfig
	SetInternalValue(val *OceanAwsUpdatePolicyRollConfig)
	LaunchSpecIds() *[]*string
	SetLaunchSpecIds(val *[]*string)
	LaunchSpecIdsInput() *[]*string
	RespectPdb() interface{}
	SetRespectPdb(val interface{})
	RespectPdbInput() interface{}
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
	ResetBatchMinHealthyPercentage()
	ResetLaunchSpecIds()
	ResetRespectPdb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAwsUpdatePolicyRollConfigOutputReference
type jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) BatchMinHealthyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchMinHealthyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) BatchMinHealthyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchMinHealthyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) BatchSizePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) BatchSizePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) InternalValue() *OceanAwsUpdatePolicyRollConfig {
	var returns *OceanAwsUpdatePolicyRollConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) LaunchSpecIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"launchSpecIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) LaunchSpecIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"launchSpecIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) RespectPdb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectPdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) RespectPdbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectPdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanAwsUpdatePolicyRollConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAwsUpdatePolicyRollConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanAwsUpdatePolicyRollConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAwsUpdatePolicyRollConfigOutputReference_Override(o OceanAwsUpdatePolicyRollConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanAwsUpdatePolicyRollConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) SetBatchMinHealthyPercentage(val *float64) {
	_jsii_.Set(
		j,
		"batchMinHealthyPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) SetBatchSizePercentage(val *float64) {
	_jsii_.Set(
		j,
		"batchSizePercentage",
		val,
	)
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) SetInternalValue(val *OceanAwsUpdatePolicyRollConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) SetLaunchSpecIds(val *[]*string) {
	_jsii_.Set(
		j,
		"launchSpecIds",
		val,
	)
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) SetRespectPdb(val interface{}) {
	_jsii_.Set(
		j,
		"respectPdb",
		val,
	)
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) ResetBatchMinHealthyPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchMinHealthyPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) ResetLaunchSpecIds() {
	_jsii_.InvokeVoid(
		o,
		"resetLaunchSpecIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) ResetRespectPdb() {
	_jsii_.InvokeVoid(
		o,
		"resetRespectPdb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsUpdatePolicyRollConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

