package oceanecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v4/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v4/oceanecs/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanEcsAutoscalerOutputReference interface {
	cdktf.ComplexObject
	AutoHeadroomPercentage() *float64
	SetAutoHeadroomPercentage(val *float64)
	AutoHeadroomPercentageInput() *float64
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
	Cooldown() *float64
	SetCooldown(val *float64)
	CooldownInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Down() OceanEcsAutoscalerDownOutputReference
	DownInput() *OceanEcsAutoscalerDown
	// Experimental.
	Fqn() *string
	Headroom() OceanEcsAutoscalerHeadroomOutputReference
	HeadroomInput() *OceanEcsAutoscalerHeadroom
	InternalValue() *OceanEcsAutoscaler
	SetInternalValue(val *OceanEcsAutoscaler)
	IsAutoConfig() interface{}
	SetIsAutoConfig(val interface{})
	IsAutoConfigInput() interface{}
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	ResourceLimits() OceanEcsAutoscalerResourceLimitsOutputReference
	ResourceLimitsInput() *OceanEcsAutoscalerResourceLimits
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
	PutDown(value *OceanEcsAutoscalerDown)
	PutHeadroom(value *OceanEcsAutoscalerHeadroom)
	PutResourceLimits(value *OceanEcsAutoscalerResourceLimits)
	ResetAutoHeadroomPercentage()
	ResetCooldown()
	ResetDown()
	ResetHeadroom()
	ResetIsAutoConfig()
	ResetIsEnabled()
	ResetResourceLimits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanEcsAutoscalerOutputReference
type jsiiProxy_OceanEcsAutoscalerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) AutoHeadroomPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoHeadroomPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) AutoHeadroomPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoHeadroomPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) Cooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) CooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) Down() OceanEcsAutoscalerDownOutputReference {
	var returns OceanEcsAutoscalerDownOutputReference
	_jsii_.Get(
		j,
		"down",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) DownInput() *OceanEcsAutoscalerDown {
	var returns *OceanEcsAutoscalerDown
	_jsii_.Get(
		j,
		"downInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) Headroom() OceanEcsAutoscalerHeadroomOutputReference {
	var returns OceanEcsAutoscalerHeadroomOutputReference
	_jsii_.Get(
		j,
		"headroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) HeadroomInput() *OceanEcsAutoscalerHeadroom {
	var returns *OceanEcsAutoscalerHeadroom
	_jsii_.Get(
		j,
		"headroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) InternalValue() *OceanEcsAutoscaler {
	var returns *OceanEcsAutoscaler
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) IsAutoConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) IsAutoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) ResourceLimits() OceanEcsAutoscalerResourceLimitsOutputReference {
	var returns OceanEcsAutoscalerResourceLimitsOutputReference
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) ResourceLimitsInput() *OceanEcsAutoscalerResourceLimits {
	var returns *OceanEcsAutoscalerResourceLimits
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanEcsAutoscalerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanEcsAutoscalerOutputReference {
	_init_.Initialize()

	if err := validateNewOceanEcsAutoscalerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanEcsAutoscalerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcsAutoscalerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanEcsAutoscalerOutputReference_Override(o OceanEcsAutoscalerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcsAutoscalerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference)SetAutoHeadroomPercentage(val *float64) {
	if err := j.validateSetAutoHeadroomPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHeadroomPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference)SetCooldown(val *float64) {
	if err := j.validateSetCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference)SetInternalValue(val *OceanEcsAutoscaler) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference)SetIsAutoConfig(val interface{}) {
	if err := j.validateSetIsAutoConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoConfig",
		val,
	)
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanEcsAutoscalerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) PutDown(value *OceanEcsAutoscalerDown) {
	if err := o.validatePutDownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDown",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) PutHeadroom(value *OceanEcsAutoscalerHeadroom) {
	if err := o.validatePutHeadroomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHeadroom",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) PutResourceLimits(value *OceanEcsAutoscalerResourceLimits) {
	if err := o.validatePutResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putResourceLimits",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) ResetAutoHeadroomPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHeadroomPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) ResetCooldown() {
	_jsii_.InvokeVoid(
		o,
		"resetCooldown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) ResetDown() {
	_jsii_.InvokeVoid(
		o,
		"resetDown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) ResetHeadroom() {
	_jsii_.InvokeVoid(
		o,
		"resetHeadroom",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) ResetIsAutoConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetIsAutoConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		o,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanEcsAutoscalerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

