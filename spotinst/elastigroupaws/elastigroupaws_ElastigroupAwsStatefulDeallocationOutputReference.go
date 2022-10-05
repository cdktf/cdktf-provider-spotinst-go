package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v3/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v3/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsStatefulDeallocationOutputReference interface {
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
	InternalValue() *ElastigroupAwsStatefulDeallocation
	SetInternalValue(val *ElastigroupAwsStatefulDeallocation)
	ShouldDeleteImages() interface{}
	SetShouldDeleteImages(val interface{})
	ShouldDeleteImagesInput() interface{}
	ShouldDeleteNetworkInterfaces() interface{}
	SetShouldDeleteNetworkInterfaces(val interface{})
	ShouldDeleteNetworkInterfacesInput() interface{}
	ShouldDeleteSnapshots() interface{}
	SetShouldDeleteSnapshots(val interface{})
	ShouldDeleteSnapshotsInput() interface{}
	ShouldDeleteVolumes() interface{}
	SetShouldDeleteVolumes(val interface{})
	ShouldDeleteVolumesInput() interface{}
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
	ResetShouldDeleteImages()
	ResetShouldDeleteNetworkInterfaces()
	ResetShouldDeleteSnapshots()
	ResetShouldDeleteVolumes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsStatefulDeallocationOutputReference
type jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) InternalValue() *ElastigroupAwsStatefulDeallocation {
	var returns *ElastigroupAwsStatefulDeallocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ShouldDeleteImages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ShouldDeleteImagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ShouldDeleteNetworkInterfaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteNetworkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ShouldDeleteNetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteNetworkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ShouldDeleteSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ShouldDeleteSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ShouldDeleteVolumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ShouldDeleteVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldDeleteVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsStatefulDeallocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastigroupAwsStatefulDeallocationOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsStatefulDeallocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsStatefulDeallocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastigroupAwsStatefulDeallocationOutputReference_Override(e ElastigroupAwsStatefulDeallocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsStatefulDeallocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference)SetInternalValue(val *ElastigroupAwsStatefulDeallocation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference)SetShouldDeleteImages(val interface{}) {
	if err := j.validateSetShouldDeleteImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDeleteImages",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference)SetShouldDeleteNetworkInterfaces(val interface{}) {
	if err := j.validateSetShouldDeleteNetworkInterfacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDeleteNetworkInterfaces",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference)SetShouldDeleteSnapshots(val interface{}) {
	if err := j.validateSetShouldDeleteSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDeleteSnapshots",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference)SetShouldDeleteVolumes(val interface{}) {
	if err := j.validateSetShouldDeleteVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldDeleteVolumes",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ResetShouldDeleteImages() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldDeleteImages",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ResetShouldDeleteNetworkInterfaces() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldDeleteNetworkInterfaces",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ResetShouldDeleteSnapshots() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldDeleteSnapshots",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ResetShouldDeleteVolumes() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldDeleteVolumes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsStatefulDeallocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

