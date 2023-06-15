package statefulnodeazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v8/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v8/statefulnodeazure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulNodeAzureStrategyOutputReference interface {
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
	DrainingTimeout() *float64
	SetDrainingTimeout(val *float64)
	DrainingTimeoutInput() *float64
	FallbackToOnDemand() interface{}
	SetFallbackToOnDemand(val interface{})
	FallbackToOnDemandInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *StatefulNodeAzureStrategy
	SetInternalValue(val *StatefulNodeAzureStrategy)
	OptimizationWindows() *[]*string
	SetOptimizationWindows(val *[]*string)
	OptimizationWindowsInput() *[]*string
	PreferredLifeCycle() *string
	SetPreferredLifeCycle(val *string)
	PreferredLifeCycleInput() *string
	RevertToSpot() StatefulNodeAzureStrategyRevertToSpotOutputReference
	RevertToSpotInput() *StatefulNodeAzureStrategyRevertToSpot
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
	PutRevertToSpot(value *StatefulNodeAzureStrategyRevertToSpot)
	ResetDrainingTimeout()
	ResetOptimizationWindows()
	ResetPreferredLifeCycle()
	ResetRevertToSpot()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulNodeAzureStrategyOutputReference
type jsiiProxy_StatefulNodeAzureStrategyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) DrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) DrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) FallbackToOnDemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOnDemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) FallbackToOnDemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackToOnDemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) InternalValue() *StatefulNodeAzureStrategy {
	var returns *StatefulNodeAzureStrategy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) OptimizationWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optimizationWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) OptimizationWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optimizationWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) PreferredLifeCycle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLifeCycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) PreferredLifeCycleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLifeCycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) RevertToSpot() StatefulNodeAzureStrategyRevertToSpotOutputReference {
	var returns StatefulNodeAzureStrategyRevertToSpotOutputReference
	_jsii_.Get(
		j,
		"revertToSpot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) RevertToSpotInput() *StatefulNodeAzureStrategyRevertToSpot {
	var returns *StatefulNodeAzureStrategyRevertToSpot
	_jsii_.Get(
		j,
		"revertToSpotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStatefulNodeAzureStrategyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StatefulNodeAzureStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulNodeAzureStrategyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulNodeAzureStrategyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStatefulNodeAzureStrategyOutputReference_Override(s StatefulNodeAzureStrategyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.statefulNodeAzure.StatefulNodeAzureStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference)SetDrainingTimeout(val *float64) {
	if err := j.validateSetDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainingTimeout",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference)SetFallbackToOnDemand(val interface{}) {
	if err := j.validateSetFallbackToOnDemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackToOnDemand",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference)SetInternalValue(val *StatefulNodeAzureStrategy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference)SetOptimizationWindows(val *[]*string) {
	if err := j.validateSetOptimizationWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optimizationWindows",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference)SetPreferredLifeCycle(val *string) {
	if err := j.validateSetPreferredLifeCycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredLifeCycle",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulNodeAzureStrategyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) PutRevertToSpot(value *StatefulNodeAzureStrategyRevertToSpot) {
	if err := s.validatePutRevertToSpotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRevertToSpot",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) ResetDrainingTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetDrainingTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) ResetOptimizationWindows() {
	_jsii_.InvokeVoid(
		s,
		"resetOptimizationWindows",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) ResetPreferredLifeCycle() {
	_jsii_.InvokeVoid(
		s,
		"resetPreferredLifeCycle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) ResetRevertToSpot() {
	_jsii_.InvokeVoid(
		s,
		"resetRevertToSpot",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulNodeAzureStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

