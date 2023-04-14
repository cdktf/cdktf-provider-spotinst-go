package multailistener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v6/multailistener/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MultaiListenerTlsConfigOutputReference interface {
	cdktf.ComplexObject
	CertificateIds() *[]*string
	SetCertificateIds(val *[]*string)
	CertificateIdsInput() *[]*string
	CipherSuites() *[]*string
	SetCipherSuites(val *[]*string)
	CipherSuitesInput() *[]*string
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
	InternalValue() *MultaiListenerTlsConfig
	SetInternalValue(val *MultaiListenerTlsConfig)
	MaxVersion() *string
	SetMaxVersion(val *string)
	MaxVersionInput() *string
	MinVersion() *string
	SetMinVersion(val *string)
	MinVersionInput() *string
	PreferServerCipherSuites() interface{}
	SetPreferServerCipherSuites(val interface{})
	PreferServerCipherSuitesInput() interface{}
	SessionTicketsDisabled() interface{}
	SetSessionTicketsDisabled(val interface{})
	SessionTicketsDisabledInput() interface{}
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MultaiListenerTlsConfigOutputReference
type jsiiProxy_MultaiListenerTlsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) CertificateIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) CertificateIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) CipherSuites() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cipherSuites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) CipherSuitesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cipherSuitesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) InternalValue() *MultaiListenerTlsConfig {
	var returns *MultaiListenerTlsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) MaxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) MaxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) MinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) MinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) PreferServerCipherSuites() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferServerCipherSuites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) PreferServerCipherSuitesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferServerCipherSuitesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) SessionTicketsDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionTicketsDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) SessionTicketsDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionTicketsDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMultaiListenerTlsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MultaiListenerTlsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewMultaiListenerTlsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MultaiListenerTlsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.multaiListener.MultaiListenerTlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMultaiListenerTlsConfigOutputReference_Override(m MultaiListenerTlsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.multaiListener.MultaiListenerTlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetCertificateIds(val *[]*string) {
	if err := j.validateSetCertificateIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateIds",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetCipherSuites(val *[]*string) {
	if err := j.validateSetCipherSuitesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cipherSuites",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetInternalValue(val *MultaiListenerTlsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetMaxVersion(val *string) {
	if err := j.validateSetMaxVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVersion",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetMinVersion(val *string) {
	if err := j.validateSetMinVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minVersion",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetPreferServerCipherSuites(val interface{}) {
	if err := j.validateSetPreferServerCipherSuitesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferServerCipherSuites",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetSessionTicketsDisabled(val interface{}) {
	if err := j.validateSetSessionTicketsDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTicketsDisabled",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MultaiListenerTlsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MultaiListenerTlsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

