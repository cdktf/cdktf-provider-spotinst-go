// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference interface {
	cdktf.ComplexObject
	BaseSize() *float64
	SetBaseSize(val *float64)
	BaseSizeInput() *float64
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
	InternalValue() *OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSize
	SetInternalValue(val *OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSize)
	Resource() *string
	SetResource(val *string)
	ResourceInput() *string
	SizePerResourceUnit() *float64
	SetSizePerResourceUnit(val *float64)
	SizePerResourceUnitInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference
type jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) BaseSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) BaseSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) InternalValue() *OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSize {
	var returns *OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSize
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SizePerResourceUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizePerResourceUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SizePerResourceUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizePerResourceUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference_Override(o OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SetBaseSize(val *float64) {
	_jsii_.Set(
		j,
		"baseSize",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SetInternalValue(val *OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSize) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SetResource(val *string) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SetSizePerResourceUnit(val *float64) {
	_jsii_.Set(
		j,
		"sizePerResourceUnit",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSizeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

