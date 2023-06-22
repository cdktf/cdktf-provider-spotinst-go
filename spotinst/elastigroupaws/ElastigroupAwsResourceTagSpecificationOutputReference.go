package elastigroupaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v9/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v9/elastigroupaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsResourceTagSpecificationOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ShouldTagAmis() interface{}
	SetShouldTagAmis(val interface{})
	ShouldTagAmisInput() interface{}
	ShouldTagEnis() interface{}
	SetShouldTagEnis(val interface{})
	ShouldTagEnisInput() interface{}
	ShouldTagSnapshots() interface{}
	SetShouldTagSnapshots(val interface{})
	ShouldTagSnapshotsInput() interface{}
	ShouldTagVolumes() interface{}
	SetShouldTagVolumes(val interface{})
	ShouldTagVolumesInput() interface{}
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
	ResetShouldTagAmis()
	ResetShouldTagEnis()
	ResetShouldTagSnapshots()
	ResetShouldTagVolumes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupAwsResourceTagSpecificationOutputReference
type jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ShouldTagAmis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagAmis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ShouldTagAmisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagAmisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ShouldTagEnis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagEnis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ShouldTagEnisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagEnisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ShouldTagSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ShouldTagSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ShouldTagVolumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ShouldTagVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastigroupAwsResourceTagSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElastigroupAwsResourceTagSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewElastigroupAwsResourceTagSpecificationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsResourceTagSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElastigroupAwsResourceTagSpecificationOutputReference_Override(e ElastigroupAwsResourceTagSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupAws.ElastigroupAwsResourceTagSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference)SetShouldTagAmis(val interface{}) {
	if err := j.validateSetShouldTagAmisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTagAmis",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference)SetShouldTagEnis(val interface{}) {
	if err := j.validateSetShouldTagEnisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTagEnis",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference)SetShouldTagSnapshots(val interface{}) {
	if err := j.validateSetShouldTagSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTagSnapshots",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference)SetShouldTagVolumes(val interface{}) {
	if err := j.validateSetShouldTagVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTagVolumes",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ResetShouldTagAmis() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldTagAmis",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ResetShouldTagEnis() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldTagEnis",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ResetShouldTagSnapshots() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldTagSnapshots",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ResetShouldTagVolumes() {
	_jsii_.InvokeVoid(
		e,
		"resetShouldTagVolumes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupAwsResourceTagSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

