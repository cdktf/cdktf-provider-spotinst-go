// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanrightsizingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference interface {
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
	IntervalMonths() *[]*float64
	SetIntervalMonths(val *[]*float64)
	IntervalMonthsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklyRepetitionBasis() OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisWeeklyRepetitionBasisList
	WeeklyRepetitionBasisInput() interface{}
	WeekOfTheMonth() *[]*string
	SetWeekOfTheMonth(val *[]*string)
	WeekOfTheMonthInput() *[]*string
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
	PutWeeklyRepetitionBasis(value interface{})
	ResetWeeklyRepetitionBasis()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference
type jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) IntervalMonths() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"intervalMonths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) IntervalMonthsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"intervalMonthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) WeeklyRepetitionBasis() OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisWeeklyRepetitionBasisList {
	var returns OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisWeeklyRepetitionBasisList
	_jsii_.Get(
		j,
		"weeklyRepetitionBasis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) WeeklyRepetitionBasisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklyRepetitionBasisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) WeekOfTheMonth() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekOfTheMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) WeekOfTheMonthInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekOfTheMonthInput",
		&returns,
	)
	return returns
}


func NewOceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference {
	_init_.Initialize()

	if err := validateNewOceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference_Override(o OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanRightSizingRule.OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference)SetIntervalMonths(val *[]*float64) {
	if err := j.validateSetIntervalMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalMonths",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference)SetWeekOfTheMonth(val *[]*string) {
	if err := j.validateSetWeekOfTheMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekOfTheMonth",
		val,
	)
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) PutWeeklyRepetitionBasis(value interface{}) {
	if err := o.validatePutWeeklyRepetitionBasisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWeeklyRepetitionBasis",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) ResetWeeklyRepetitionBasis() {
	_jsii_.InvokeVoid(
		o,
		"resetWeeklyRepetitionBasis",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

