package managedinstanceaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/managedinstanceaws/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedInstanceAwsResourceTagSpecificationOutputReference interface {
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

// The jsii proxy struct for ManagedInstanceAwsResourceTagSpecificationOutputReference
type jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ShouldTagAmis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagAmis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ShouldTagAmisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagAmisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ShouldTagEnis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagEnis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ShouldTagEnisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagEnisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ShouldTagSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ShouldTagSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ShouldTagVolumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ShouldTagVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTagVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedInstanceAwsResourceTagSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ManagedInstanceAwsResourceTagSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewManagedInstanceAwsResourceTagSpecificationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsResourceTagSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewManagedInstanceAwsResourceTagSpecificationOutputReference_Override(m ManagedInstanceAwsResourceTagSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsResourceTagSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference)SetShouldTagAmis(val interface{}) {
	if err := j.validateSetShouldTagAmisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTagAmis",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference)SetShouldTagEnis(val interface{}) {
	if err := j.validateSetShouldTagEnisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTagEnis",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference)SetShouldTagSnapshots(val interface{}) {
	if err := j.validateSetShouldTagSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTagSnapshots",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference)SetShouldTagVolumes(val interface{}) {
	if err := j.validateSetShouldTagVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTagVolumes",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ResetShouldTagAmis() {
	_jsii_.InvokeVoid(
		m,
		"resetShouldTagAmis",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ResetShouldTagEnis() {
	_jsii_.InvokeVoid(
		m,
		"resetShouldTagEnis",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ResetShouldTagSnapshots() {
	_jsii_.InvokeVoid(
		m,
		"resetShouldTagSnapshots",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ResetShouldTagVolumes() {
	_jsii_.InvokeVoid(
		m,
		"resetShouldTagVolumes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

