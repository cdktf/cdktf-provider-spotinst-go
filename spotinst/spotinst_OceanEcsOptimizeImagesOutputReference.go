// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanEcsOptimizeImagesOutputReference interface {
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
	InternalValue() *OceanEcsOptimizeImages
	SetInternalValue(val *OceanEcsOptimizeImages)
	PerformAt() *string
	SetPerformAt(val *string)
	PerformAtInput() *string
	ShouldOptimizeEcsAmi() interface{}
	SetShouldOptimizeEcsAmi(val interface{})
	ShouldOptimizeEcsAmiInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeWindows() *[]*string
	SetTimeWindows(val *[]*string)
	TimeWindowsInput() *[]*string
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
	ResetTimeWindows()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanEcsOptimizeImagesOutputReference
type jsiiProxy_OceanEcsOptimizeImagesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) InternalValue() *OceanEcsOptimizeImages {
	var returns *OceanEcsOptimizeImages
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) PerformAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) PerformAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) ShouldOptimizeEcsAmi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldOptimizeEcsAmi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) ShouldOptimizeEcsAmiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldOptimizeEcsAmiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) TimeWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"timeWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) TimeWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"timeWindowsInput",
		&returns,
	)
	return returns
}


func NewOceanEcsOptimizeImagesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanEcsOptimizeImagesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OceanEcsOptimizeImagesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanEcsOptimizeImagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanEcsOptimizeImagesOutputReference_Override(o OceanEcsOptimizeImagesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanEcsOptimizeImagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) SetInternalValue(val *OceanEcsOptimizeImages) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) SetPerformAt(val *string) {
	_jsii_.Set(
		j,
		"performAt",
		val,
	)
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) SetShouldOptimizeEcsAmi(val interface{}) {
	_jsii_.Set(
		j,
		"shouldOptimizeEcsAmi",
		val,
	)
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceanEcsOptimizeImagesOutputReference) SetTimeWindows(val *[]*string) {
	_jsii_.Set(
		j,
		"timeWindows",
		val,
	)
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) ResetTimeWindows() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeWindows",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsOptimizeImagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

