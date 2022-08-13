// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/jsii"

	"github.com/hashicorp/cdktf-provider-spotinst-go/spotinst/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList
type jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList {
	_init_.Initialize()

	j := jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList_Override(o OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) Get(index *float64) OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticOutputReference {
	var returns OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAwsLaunchSpecAutoscaleHeadroomsAutomaticList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

