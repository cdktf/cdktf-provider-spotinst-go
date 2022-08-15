// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference interface {
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
	InternalValue() *OceanAksVirtualNodeGroupLaunchSpecificationOsDisk
	SetInternalValue(val *OceanAksVirtualNodeGroupLaunchSpecificationOsDisk)
	SizeGb() *float64
	SetSizeGb(val *float64)
	SizeGbInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UtilizeEphemeralStorage() interface{}
	SetUtilizeEphemeralStorage(val interface{})
	UtilizeEphemeralStorageInput() interface{}
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
	ResetType()
	ResetUtilizeEphemeralStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference
type jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) InternalValue() *OceanAksVirtualNodeGroupLaunchSpecificationOsDisk {
	var returns *OceanAksVirtualNodeGroupLaunchSpecificationOsDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) UtilizeEphemeralStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeEphemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) UtilizeEphemeralStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizeEphemeralStorageInput",
		&returns,
	)
	return returns
}


func NewOceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference_Override(o OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SetInternalValue(val *OceanAksVirtualNodeGroupLaunchSpecificationOsDisk) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SetSizeGb(val *float64) {
	_jsii_.Set(
		j,
		"sizeGb",
		val,
	)
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) SetUtilizeEphemeralStorage(val interface{}) {
	_jsii_.Set(
		j,
		"utilizeEphemeralStorage",
		val,
	)
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		o,
		"resetType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) ResetUtilizeEphemeralStorage() {
	_jsii_.InvokeVoid(
		o,
		"resetUtilizeEphemeralStorage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksVirtualNodeGroupLaunchSpecificationOsDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
