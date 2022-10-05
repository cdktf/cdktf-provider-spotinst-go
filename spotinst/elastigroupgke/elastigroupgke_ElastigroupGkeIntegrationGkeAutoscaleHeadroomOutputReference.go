package elastigroupgke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v3/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v3/elastigroupgke/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference interface {
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
	CpuPerUnit() *float64
	SetCpuPerUnit(val *float64)
	CpuPerUnitInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ElastigroupGkeIntegrationGkeAutoscaleHeadroom
	SetInternalValue(val *ElastigroupGkeIntegrationGkeAutoscaleHeadroom)
	MemoryPerUnit() *float64
	SetMemoryPerUnit(val *float64)
	MemoryPerUnitInput() *float64
	NumOfUnits() *float64
	SetNumOfUnits(val *float64)
	NumOfUnitsInput() *float64
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
	ResetCpuPerUnit()
	ResetMemoryPerUnit()
	ResetNumOfUnits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference
type jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) CpuPerUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPerUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) CpuPerUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPerUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) InternalValue() *ElastigroupGkeIntegrationGkeAutoscaleHeadroom {
	var returns *ElastigroupGkeIntegrationGkeAutoscaleHeadroom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) MemoryPerUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) MemoryPerUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) NumOfUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numOfUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) NumOfUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numOfUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference_Override(e ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference)SetCpuPerUnit(val *float64) {
	if err := j.validateSetCpuPerUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuPerUnit",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference)SetInternalValue(val *ElastigroupGkeIntegrationGkeAutoscaleHeadroom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference)SetMemoryPerUnit(val *float64) {
	if err := j.validateSetMemoryPerUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryPerUnit",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference)SetNumOfUnits(val *float64) {
	if err := j.validateSetNumOfUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numOfUnits",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) ResetCpuPerUnit() {
	_jsii_.InvokeVoid(
		e,
		"resetCpuPerUnit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) ResetMemoryPerUnit() {
	_jsii_.InvokeVoid(
		e,
		"resetMemoryPerUnit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) ResetNumOfUnits() {
	_jsii_.InvokeVoid(
		e,
		"resetNumOfUnits",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupGkeIntegrationGkeAutoscaleHeadroomOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

