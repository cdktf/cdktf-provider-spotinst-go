package oceangkeimport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v3/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v3/oceangkeimport/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanGkeImportAutoscalerOutputReference interface {
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
	Down() OceanGkeImportAutoscalerDownOutputReference
	DownInput() *OceanGkeImportAutoscalerDown
	EnableAutomaticAndManualHeadroom() interface{}
	SetEnableAutomaticAndManualHeadroom(val interface{})
	EnableAutomaticAndManualHeadroomInput() interface{}
	// Experimental.
	Fqn() *string
	Headroom() OceanGkeImportAutoscalerHeadroomOutputReference
	HeadroomInput() *OceanGkeImportAutoscalerHeadroom
	InternalValue() *OceanGkeImportAutoscaler
	SetInternalValue(val *OceanGkeImportAutoscaler)
	IsAutoConfig() interface{}
	SetIsAutoConfig(val interface{})
	IsAutoConfigInput() interface{}
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	ResourceLimits() OceanGkeImportAutoscalerResourceLimitsOutputReference
	ResourceLimitsInput() *OceanGkeImportAutoscalerResourceLimits
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
	PutDown(value *OceanGkeImportAutoscalerDown)
	PutHeadroom(value *OceanGkeImportAutoscalerHeadroom)
	PutResourceLimits(value *OceanGkeImportAutoscalerResourceLimits)
	ResetAutoHeadroomPercentage()
	ResetCooldown()
	ResetDown()
	ResetEnableAutomaticAndManualHeadroom()
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

// The jsii proxy struct for OceanGkeImportAutoscalerOutputReference
type jsiiProxy_OceanGkeImportAutoscalerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) AutoHeadroomPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoHeadroomPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) AutoHeadroomPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoHeadroomPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) Cooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) CooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) Down() OceanGkeImportAutoscalerDownOutputReference {
	var returns OceanGkeImportAutoscalerDownOutputReference
	_jsii_.Get(
		j,
		"down",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) DownInput() *OceanGkeImportAutoscalerDown {
	var returns *OceanGkeImportAutoscalerDown
	_jsii_.Get(
		j,
		"downInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) EnableAutomaticAndManualHeadroom() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticAndManualHeadroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) EnableAutomaticAndManualHeadroomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticAndManualHeadroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) Headroom() OceanGkeImportAutoscalerHeadroomOutputReference {
	var returns OceanGkeImportAutoscalerHeadroomOutputReference
	_jsii_.Get(
		j,
		"headroom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) HeadroomInput() *OceanGkeImportAutoscalerHeadroom {
	var returns *OceanGkeImportAutoscalerHeadroom
	_jsii_.Get(
		j,
		"headroomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) InternalValue() *OceanGkeImportAutoscaler {
	var returns *OceanGkeImportAutoscaler
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) IsAutoConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) IsAutoConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResourceLimits() OceanGkeImportAutoscalerResourceLimitsOutputReference {
	var returns OceanGkeImportAutoscalerResourceLimitsOutputReference
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResourceLimitsInput() *OceanGkeImportAutoscalerResourceLimits {
	var returns *OceanGkeImportAutoscalerResourceLimits
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOceanGkeImportAutoscalerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanGkeImportAutoscalerOutputReference {
	_init_.Initialize()

	if err := validateNewOceanGkeImportAutoscalerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanGkeImportAutoscalerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImportAutoscalerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanGkeImportAutoscalerOutputReference_Override(o OceanGkeImportAutoscalerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanGkeImport.OceanGkeImportAutoscalerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetAutoHeadroomPercentage(val *float64) {
	if err := j.validateSetAutoHeadroomPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHeadroomPercentage",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetCooldown(val *float64) {
	if err := j.validateSetCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetEnableAutomaticAndManualHeadroom(val interface{}) {
	if err := j.validateSetEnableAutomaticAndManualHeadroomParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticAndManualHeadroom",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetInternalValue(val *OceanGkeImportAutoscaler) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetIsAutoConfig(val interface{}) {
	if err := j.validateSetIsAutoConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoConfig",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanGkeImportAutoscalerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) PutDown(value *OceanGkeImportAutoscalerDown) {
	if err := o.validatePutDownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDown",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) PutHeadroom(value *OceanGkeImportAutoscalerHeadroom) {
	if err := o.validatePutHeadroomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHeadroom",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) PutResourceLimits(value *OceanGkeImportAutoscalerResourceLimits) {
	if err := o.validatePutResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putResourceLimits",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResetAutoHeadroomPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHeadroomPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResetCooldown() {
	_jsii_.InvokeVoid(
		o,
		"resetCooldown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResetDown() {
	_jsii_.InvokeVoid(
		o,
		"resetDown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResetEnableAutomaticAndManualHeadroom() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableAutomaticAndManualHeadroom",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResetHeadroom() {
	_jsii_.InvokeVoid(
		o,
		"resetHeadroom",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResetIsAutoConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetIsAutoConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		o,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanGkeImportAutoscalerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

