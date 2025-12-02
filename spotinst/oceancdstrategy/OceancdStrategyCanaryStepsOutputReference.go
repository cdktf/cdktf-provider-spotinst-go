// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceancdstrategy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceancdStrategyCanaryStepsOutputReference interface {
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
	Pause() OceancdStrategyCanaryStepsPauseOutputReference
	PauseInput() *OceancdStrategyCanaryStepsPause
	SetCanaryScale() OceancdStrategyCanaryStepsSetCanaryScaleOutputReference
	SetCanaryScaleInput() *OceancdStrategyCanaryStepsSetCanaryScale
	SetHeaderRoute() OceancdStrategyCanaryStepsSetHeaderRouteOutputReference
	SetHeaderRouteInput() *OceancdStrategyCanaryStepsSetHeaderRoute
	SetWeight() *float64
	SetSetWeight(val *float64)
	SetWeightInput() *float64
	StepName() *string
	SetStepName(val *string)
	StepNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Verification() OceancdStrategyCanaryStepsVerificationOutputReference
	VerificationInput() *OceancdStrategyCanaryStepsVerification
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutPause(value *OceancdStrategyCanaryStepsPause)
	PutSetCanaryScale(value *OceancdStrategyCanaryStepsSetCanaryScale)
	PutSetHeaderRoute(value *OceancdStrategyCanaryStepsSetHeaderRoute)
	PutVerification(value *OceancdStrategyCanaryStepsVerification)
	ResetPause()
	ResetSetCanaryScale()
	ResetSetHeaderRoute()
	ResetSetWeight()
	ResetStepName()
	ResetVerification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceancdStrategyCanaryStepsOutputReference
type jsiiProxy_OceancdStrategyCanaryStepsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) Pause() OceancdStrategyCanaryStepsPauseOutputReference {
	var returns OceancdStrategyCanaryStepsPauseOutputReference
	_jsii_.Get(
		j,
		"pause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) PauseInput() *OceancdStrategyCanaryStepsPause {
	var returns *OceancdStrategyCanaryStepsPause
	_jsii_.Get(
		j,
		"pauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) SetCanaryScale() OceancdStrategyCanaryStepsSetCanaryScaleOutputReference {
	var returns OceancdStrategyCanaryStepsSetCanaryScaleOutputReference
	_jsii_.Get(
		j,
		"setCanaryScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) SetCanaryScaleInput() *OceancdStrategyCanaryStepsSetCanaryScale {
	var returns *OceancdStrategyCanaryStepsSetCanaryScale
	_jsii_.Get(
		j,
		"setCanaryScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) SetHeaderRoute() OceancdStrategyCanaryStepsSetHeaderRouteOutputReference {
	var returns OceancdStrategyCanaryStepsSetHeaderRouteOutputReference
	_jsii_.Get(
		j,
		"setHeaderRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) SetHeaderRouteInput() *OceancdStrategyCanaryStepsSetHeaderRoute {
	var returns *OceancdStrategyCanaryStepsSetHeaderRoute
	_jsii_.Get(
		j,
		"setHeaderRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) SetWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"setWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) SetWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"setWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) StepName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) StepNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) Verification() OceancdStrategyCanaryStepsVerificationOutputReference {
	var returns OceancdStrategyCanaryStepsVerificationOutputReference
	_jsii_.Get(
		j,
		"verification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) VerificationInput() *OceancdStrategyCanaryStepsVerification {
	var returns *OceancdStrategyCanaryStepsVerification
	_jsii_.Get(
		j,
		"verificationInput",
		&returns,
	)
	return returns
}


func NewOceancdStrategyCanaryStepsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OceancdStrategyCanaryStepsOutputReference {
	_init_.Initialize()

	if err := validateNewOceancdStrategyCanaryStepsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceancdStrategyCanaryStepsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdStrategy.OceancdStrategyCanaryStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOceancdStrategyCanaryStepsOutputReference_Override(o OceancdStrategyCanaryStepsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceancdStrategy.OceancdStrategyCanaryStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference)SetSetWeight(val *float64) {
	if err := j.validateSetSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setWeight",
		val,
	)
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference)SetStepName(val *string) {
	if err := j.validateSetStepNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepName",
		val,
	)
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceancdStrategyCanaryStepsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) PutPause(value *OceancdStrategyCanaryStepsPause) {
	if err := o.validatePutPauseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPause",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) PutSetCanaryScale(value *OceancdStrategyCanaryStepsSetCanaryScale) {
	if err := o.validatePutSetCanaryScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSetCanaryScale",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) PutSetHeaderRoute(value *OceancdStrategyCanaryStepsSetHeaderRoute) {
	if err := o.validatePutSetHeaderRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSetHeaderRoute",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) PutVerification(value *OceancdStrategyCanaryStepsVerification) {
	if err := o.validatePutVerificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVerification",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ResetPause() {
	_jsii_.InvokeVoid(
		o,
		"resetPause",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ResetSetCanaryScale() {
	_jsii_.InvokeVoid(
		o,
		"resetSetCanaryScale",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ResetSetHeaderRoute() {
	_jsii_.InvokeVoid(
		o,
		"resetSetHeaderRoute",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ResetSetWeight() {
	_jsii_.InvokeVoid(
		o,
		"resetSetWeight",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ResetStepName() {
	_jsii_.InvokeVoid(
		o,
		"resetStepName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ResetVerification() {
	_jsii_.InvokeVoid(
		o,
		"resetVerification",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceancdStrategyCanaryStepsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

