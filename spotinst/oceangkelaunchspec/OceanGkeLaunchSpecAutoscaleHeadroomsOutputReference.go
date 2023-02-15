package oceangkelaunchspec

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v5/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v5/oceangkelaunchspec/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference interface {
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
	GpuPerUnit() *float64
	SetGpuPerUnit(val *float64)
	GpuPerUnitInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetGpuPerUnit()
	ResetMemoryPerUnit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference
type jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) CpuPerUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPerUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) CpuPerUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPerUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GpuPerUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuPerUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GpuPerUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuPerUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) MemoryPerUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) MemoryPerUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) NumOfUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numOfUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) NumOfUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numOfUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanGkeLaunchSpecAutoscaleHeadroomsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference {
	_init_.Initialize()

	if err := validateNewOceanGkeLaunchSpecAutoscaleHeadroomsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOceanGkeLaunchSpecAutoscaleHeadroomsOutputReference_Override(o OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeLaunchSpec.OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference)SetCpuPerUnit(val *float64) {
	if err := j.validateSetCpuPerUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuPerUnit",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference)SetGpuPerUnit(val *float64) {
	if err := j.validateSetGpuPerUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuPerUnit",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference)SetMemoryPerUnit(val *float64) {
	if err := j.validateSetMemoryPerUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryPerUnit",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference)SetNumOfUnits(val *float64) {
	if err := j.validateSetNumOfUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numOfUnits",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) ResetCpuPerUnit() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuPerUnit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) ResetGpuPerUnit() {
	_jsii_.InvokeVoid(
		o,
		"resetGpuPerUnit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) ResetMemoryPerUnit() {
	_jsii_.InvokeVoid(
		o,
		"resetMemoryPerUnit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeLaunchSpecAutoscaleHeadroomsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

