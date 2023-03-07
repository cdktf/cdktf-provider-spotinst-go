package elastigroupgke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v5/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v5/elastigroupgke/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupGkeNetworkInterfaceAliasIpRangesList interface {
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
	Get(index *float64) ElastigroupGkeNetworkInterfaceAliasIpRangesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastigroupGkeNetworkInterfaceAliasIpRangesList
type jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewElastigroupGkeNetworkInterfaceAliasIpRangesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ElastigroupGkeNetworkInterfaceAliasIpRangesList {
	_init_.Initialize()

	if err := validateNewElastigroupGkeNetworkInterfaceAliasIpRangesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGkeNetworkInterfaceAliasIpRangesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewElastigroupGkeNetworkInterfaceAliasIpRangesList_Override(e ElastigroupGkeNetworkInterfaceAliasIpRangesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.elastigroupGke.ElastigroupGkeNetworkInterfaceAliasIpRangesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) Get(index *float64) ElastigroupGkeNetworkInterfaceAliasIpRangesOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ElastigroupGkeNetworkInterfaceAliasIpRangesOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElastigroupGkeNetworkInterfaceAliasIpRangesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
