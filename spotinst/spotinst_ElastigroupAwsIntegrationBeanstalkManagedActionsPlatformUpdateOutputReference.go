// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdate
	SetInternalValue(val *ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdate)
	PerformAt() *string
	SetPerformAt(val *string)
	PerformAtInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeWindow() *string
	SetTimeWindow(val *string)
	TimeWindowInput() *string
	UpdateLevel() *string
	SetUpdateLevel(val *string)
	UpdateLevelInput() *string
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
	ResetPerformAt()
	ResetTimeWindow()
	ResetUpdateLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference
type jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) InternalValue() *ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdate {
	var returns *ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) PerformAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) PerformAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) TimeWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) TimeWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) UpdateLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) UpdateLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateLevelInput",
		&returns,
	)
	return returns
}


func NewElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference_Override(e ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference)SetInternalValue(val *ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference)SetPerformAt(val *string) {
	if err := j.validateSetPerformAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performAt",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference)SetTimeWindow(val *string) {
	if err := j.validateSetTimeWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeWindow",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference)SetUpdateLevel(val *string) {
	if err := j.validateSetUpdateLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateLevel",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) ResetPerformAt() {
	_jsii_.InvokeVoid(
		e,
		"resetPerformAt",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) ResetTimeWindow() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeWindow",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) ResetUpdateLevel() {
	_jsii_.InvokeVoid(
		e,
		"resetUpdateLevel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

