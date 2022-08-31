// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsItfOutputReference interface {
	cdktf.ComplexObject
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
	DefaultStaticTargetGroup() ElastigroupAwsItfDefaultStaticTargetGroupOutputReference
	DefaultStaticTargetGroupInput() *ElastigroupAwsItfDefaultStaticTargetGroup
	FixedTargetGroups() interface{}
	SetFixedTargetGroups(val interface{})
	FixedTargetGroupsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancer() ElastigroupAwsItfLoadBalancerList
	LoadBalancerInput() interface{}
	MigrationHealthinessThreshold() *float64
	SetMigrationHealthinessThreshold(val *float64)
	MigrationHealthinessThresholdInput() *float64
	TargetGroupConfig() ElastigroupAwsItfTargetGroupConfigList
	TargetGroupConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeightStrategy() *string
	SetWeightStrategy(val *string)
	WeightStrategyInput() *string
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
	PutDefaultStaticTargetGroup(value *ElastigroupAwsItfDefaultStaticTargetGroup)
	PutLoadBalancer(value interface{})
	PutTargetGroupConfig(value interface{})
	ResetDefaultStaticTargetGroup()
	ResetMigrationHealthinessThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsItfOutputReference
type jsiiProxy_ElastigroupAwsItfOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) DefaultStaticTargetGroup() ElastigroupAwsItfDefaultStaticTargetGroupOutputReference {
	var returns ElastigroupAwsItfDefaultStaticTargetGroupOutputReference
	_jsii_.Get(
		j,
		"defaultStaticTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) DefaultStaticTargetGroupInput() *ElastigroupAwsItfDefaultStaticTargetGroup {
	var returns *ElastigroupAwsItfDefaultStaticTargetGroup
	_jsii_.Get(
		j,
		"defaultStaticTargetGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) FixedTargetGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedTargetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) FixedTargetGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedTargetGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) LoadBalancer() ElastigroupAwsItfLoadBalancerList {
	var returns ElastigroupAwsItfLoadBalancerList
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) LoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) MigrationHealthinessThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"migrationHealthinessThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) MigrationHealthinessThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"migrationHealthinessThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) TargetGroupConfig() ElastigroupAwsItfTargetGroupConfigList {
	var returns ElastigroupAwsItfTargetGroupConfigList
	_jsii_.Get(
		j,
		"targetGroupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) TargetGroupConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) WeightStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weightStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference) WeightStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weightStrategyInput",
		&returns,
	)
	return returns
}


func NewElastigroupAwsItfOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElastigroupAwsItfOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsItfOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsItfOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsItfOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElastigroupAwsItfOutputReference_Override(e ElastigroupAwsItfOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsItfOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference)SetFixedTargetGroups(val interface{}) {
	if err := j.validateSetFixedTargetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedTargetGroups",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference)SetMigrationHealthinessThreshold(val *float64) {
	if err := j.validateSetMigrationHealthinessThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationHealthinessThreshold",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsItfOutputReference)SetWeightStrategy(val *string) {
	if err := j.validateSetWeightStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weightStrategy",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) PutDefaultStaticTargetGroup(value *ElastigroupAwsItfDefaultStaticTargetGroup) {
	if err := e.validatePutDefaultStaticTargetGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDefaultStaticTargetGroup",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) PutLoadBalancer(value interface{}) {
	if err := e.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) PutTargetGroupConfig(value interface{}) {
	if err := e.validatePutTargetGroupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargetGroupConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) ResetDefaultStaticTargetGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetDefaultStaticTargetGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) ResetMigrationHealthinessThreshold() {
	_jsii_.InvokeVoid(
		e,
		"resetMigrationHealthinessThreshold",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsItfOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

