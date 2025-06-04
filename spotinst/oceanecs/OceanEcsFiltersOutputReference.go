// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v14/oceanecs/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanEcsFiltersOutputReference interface {
	cdktf.ComplexObject
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	ArchitecturesInput() *[]*string
	Categories() *[]*string
	SetCategories(val *[]*string)
	CategoriesInput() *[]*string
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
	DiskTypes() *[]*string
	SetDiskTypes(val *[]*string)
	DiskTypesInput() *[]*string
	ExcludeFamilies() *[]*string
	SetExcludeFamilies(val *[]*string)
	ExcludeFamiliesInput() *[]*string
	ExcludeMetal() interface{}
	SetExcludeMetal(val interface{})
	ExcludeMetalInput() interface{}
	// Experimental.
	Fqn() *string
	Hypervisor() *[]*string
	SetHypervisor(val *[]*string)
	HypervisorInput() *[]*string
	IncludeFamilies() *[]*string
	SetIncludeFamilies(val *[]*string)
	IncludeFamiliesInput() *[]*string
	InternalValue() *OceanEcsFilters
	SetInternalValue(val *OceanEcsFilters)
	IsEnaSupported() *string
	SetIsEnaSupported(val *string)
	IsEnaSupportedInput() *string
	MaxGpu() *float64
	SetMaxGpu(val *float64)
	MaxGpuInput() *float64
	MaxMemoryGib() *float64
	SetMaxMemoryGib(val *float64)
	MaxMemoryGibInput() *float64
	MaxNetworkPerformance() *float64
	SetMaxNetworkPerformance(val *float64)
	MaxNetworkPerformanceInput() *float64
	MaxVcpu() *float64
	SetMaxVcpu(val *float64)
	MaxVcpuInput() *float64
	MinEnis() *float64
	SetMinEnis(val *float64)
	MinEnisInput() *float64
	MinGpu() *float64
	SetMinGpu(val *float64)
	MinGpuInput() *float64
	MinMemoryGib() *float64
	SetMinMemoryGib(val *float64)
	MinMemoryGibInput() *float64
	MinNetworkPerformance() *float64
	SetMinNetworkPerformance(val *float64)
	MinNetworkPerformanceInput() *float64
	MinVcpu() *float64
	SetMinVcpu(val *float64)
	MinVcpuInput() *float64
	RootDeviceTypes() *[]*string
	SetRootDeviceTypes(val *[]*string)
	RootDeviceTypesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualizationTypes() *[]*string
	SetVirtualizationTypes(val *[]*string)
	VirtualizationTypesInput() *[]*string
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
	ResetArchitectures()
	ResetCategories()
	ResetDiskTypes()
	ResetExcludeFamilies()
	ResetExcludeMetal()
	ResetHypervisor()
	ResetIncludeFamilies()
	ResetIsEnaSupported()
	ResetMaxGpu()
	ResetMaxMemoryGib()
	ResetMaxNetworkPerformance()
	ResetMaxVcpu()
	ResetMinEnis()
	ResetMinGpu()
	ResetMinMemoryGib()
	ResetMinNetworkPerformance()
	ResetMinVcpu()
	ResetRootDeviceTypes()
	ResetVirtualizationTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanEcsFiltersOutputReference
type jsiiProxy_OceanEcsFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) Categories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) CategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) DiskTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"diskTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) DiskTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"diskTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) ExcludeFamilies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) ExcludeFamiliesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) ExcludeMetal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) ExcludeMetalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeMetalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) Hypervisor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hypervisor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) HypervisorInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hypervisorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) IncludeFamilies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) IncludeFamiliesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) InternalValue() *OceanEcsFilters {
	var returns *OceanEcsFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) IsEnaSupported() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isEnaSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) IsEnaSupportedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isEnaSupportedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MaxGpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MaxGpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MaxMemoryGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MaxMemoryGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MaxNetworkPerformance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNetworkPerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MaxNetworkPerformanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNetworkPerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MaxVcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MaxVcpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinEnis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEnis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinEnisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEnisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinGpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinGpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minGpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinMemoryGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinMemoryGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinNetworkPerformance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNetworkPerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinNetworkPerformanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNetworkPerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinVcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) MinVcpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) RootDeviceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootDeviceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) RootDeviceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootDeviceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) VirtualizationTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"virtualizationTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference) VirtualizationTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"virtualizationTypesInput",
		&returns,
	)
	return returns
}


func NewOceanEcsFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanEcsFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewOceanEcsFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanEcsFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcsFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanEcsFiltersOutputReference_Override(o OceanEcsFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanEcs.OceanEcsFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetCategories(val *[]*string) {
	if err := j.validateSetCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"categories",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetDiskTypes(val *[]*string) {
	if err := j.validateSetDiskTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskTypes",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetExcludeFamilies(val *[]*string) {
	if err := j.validateSetExcludeFamiliesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeFamilies",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetExcludeMetal(val interface{}) {
	if err := j.validateSetExcludeMetalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeMetal",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetHypervisor(val *[]*string) {
	if err := j.validateSetHypervisorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hypervisor",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetIncludeFamilies(val *[]*string) {
	if err := j.validateSetIncludeFamiliesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeFamilies",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetInternalValue(val *OceanEcsFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetIsEnaSupported(val *string) {
	if err := j.validateSetIsEnaSupportedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnaSupported",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetMaxGpu(val *float64) {
	if err := j.validateSetMaxGpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxGpu",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetMaxMemoryGib(val *float64) {
	if err := j.validateSetMaxMemoryGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMemoryGib",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetMaxNetworkPerformance(val *float64) {
	if err := j.validateSetMaxNetworkPerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNetworkPerformance",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetMaxVcpu(val *float64) {
	if err := j.validateSetMaxVcpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVcpu",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetMinEnis(val *float64) {
	if err := j.validateSetMinEnisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minEnis",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetMinGpu(val *float64) {
	if err := j.validateSetMinGpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minGpu",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetMinMemoryGib(val *float64) {
	if err := j.validateSetMinMemoryGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMemoryGib",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetMinNetworkPerformance(val *float64) {
	if err := j.validateSetMinNetworkPerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNetworkPerformance",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetMinVcpu(val *float64) {
	if err := j.validateSetMinVcpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minVcpu",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetRootDeviceTypes(val *[]*string) {
	if err := j.validateSetRootDeviceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootDeviceTypes",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceanEcsFiltersOutputReference)SetVirtualizationTypes(val *[]*string) {
	if err := j.validateSetVirtualizationTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualizationTypes",
		val,
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetArchitectures() {
	_jsii_.InvokeVoid(
		o,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetCategories() {
	_jsii_.InvokeVoid(
		o,
		"resetCategories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetDiskTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetDiskTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetExcludeFamilies() {
	_jsii_.InvokeVoid(
		o,
		"resetExcludeFamilies",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetExcludeMetal() {
	_jsii_.InvokeVoid(
		o,
		"resetExcludeMetal",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetHypervisor() {
	_jsii_.InvokeVoid(
		o,
		"resetHypervisor",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetIncludeFamilies() {
	_jsii_.InvokeVoid(
		o,
		"resetIncludeFamilies",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetIsEnaSupported() {
	_jsii_.InvokeVoid(
		o,
		"resetIsEnaSupported",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetMaxGpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxGpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetMaxMemoryGib() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxMemoryGib",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetMaxNetworkPerformance() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxNetworkPerformance",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetMaxVcpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxVcpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetMinEnis() {
	_jsii_.InvokeVoid(
		o,
		"resetMinEnis",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetMinGpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMinGpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetMinMemoryGib() {
	_jsii_.InvokeVoid(
		o,
		"resetMinMemoryGib",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetMinNetworkPerformance() {
	_jsii_.InvokeVoid(
		o,
		"resetMinNetworkPerformance",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetMinVcpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMinVcpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetRootDeviceTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetRootDeviceTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ResetVirtualizationTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetVirtualizationTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanEcsFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanEcsFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

