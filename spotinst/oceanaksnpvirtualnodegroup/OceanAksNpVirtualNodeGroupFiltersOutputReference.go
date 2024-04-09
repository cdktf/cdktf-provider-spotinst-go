// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/jsii"

	"github.com/cdktf/cdktf-provider-spotinst-go/spotinst/v13/oceanaksnpvirtualnodegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksNpVirtualNodeGroupFiltersOutputReference interface {
	cdktf.ComplexObject
	AcceleratedNetworking() *string
	SetAcceleratedNetworking(val *string)
	AcceleratedNetworkingInput() *string
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	ArchitecturesInput() *[]*string
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
	DiskPerformance() *string
	SetDiskPerformance(val *string)
	DiskPerformanceInput() *string
	ExcludeSeries() *[]*string
	SetExcludeSeries(val *[]*string)
	ExcludeSeriesInput() *[]*string
	// Experimental.
	Fqn() *string
	GpuTypes() *[]*string
	SetGpuTypes(val *[]*string)
	GpuTypesInput() *[]*string
	InternalValue() *OceanAksNpVirtualNodeGroupFilters
	SetInternalValue(val *OceanAksNpVirtualNodeGroupFilters)
	MaxGpu() *float64
	SetMaxGpu(val *float64)
	MaxGpuInput() *float64
	MaxMemoryGib() *float64
	SetMaxMemoryGib(val *float64)
	MaxMemoryGibInput() *float64
	MaxVcpu() *float64
	SetMaxVcpu(val *float64)
	MaxVcpuInput() *float64
	MinDisk() *float64
	SetMinDisk(val *float64)
	MinDiskInput() *float64
	MinGpu() *float64
	SetMinGpu(val *float64)
	MinGpuInput() *float64
	MinMemoryGib() *float64
	SetMinMemoryGib(val *float64)
	MinMemoryGibInput() *float64
	MinNics() *float64
	SetMinNics(val *float64)
	MinNicsInput() *float64
	MinVcpu() *float64
	SetMinVcpu(val *float64)
	MinVcpuInput() *float64
	Series() *[]*string
	SetSeries(val *[]*string)
	SeriesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmTypes() *[]*string
	SetVmTypes(val *[]*string)
	VmTypesInput() *[]*string
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
	ResetAcceleratedNetworking()
	ResetArchitectures()
	ResetDiskPerformance()
	ResetExcludeSeries()
	ResetGpuTypes()
	ResetMaxGpu()
	ResetMaxMemoryGib()
	ResetMaxVcpu()
	ResetMinDisk()
	ResetMinGpu()
	ResetMinMemoryGib()
	ResetMinNics()
	ResetMinVcpu()
	ResetSeries()
	ResetVmTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OceanAksNpVirtualNodeGroupFiltersOutputReference
type jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) AcceleratedNetworking() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratedNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) AcceleratedNetworkingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratedNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) DiskPerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskPerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) DiskPerformanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskPerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ExcludeSeries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ExcludeSeriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GpuTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gpuTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GpuTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gpuTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) InternalValue() *OceanAksNpVirtualNodeGroupFilters {
	var returns *OceanAksNpVirtualNodeGroupFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MaxGpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MaxGpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MaxMemoryGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MaxMemoryGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MaxVcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MaxVcpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinDisk() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinDiskInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinGpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinGpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minGpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinMemoryGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinMemoryGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinNics() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinNicsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinVcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) MinVcpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) Series() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"series",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) SeriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"seriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) VmTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) VmTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmTypesInput",
		&returns,
	)
	return returns
}


func NewOceanAksNpVirtualNodeGroupFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OceanAksNpVirtualNodeGroupFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewOceanAksNpVirtualNodeGroupFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroupFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOceanAksNpVirtualNodeGroupFiltersOutputReference_Override(o OceanAksNpVirtualNodeGroupFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-spotinst.oceanAksNpVirtualNodeGroup.OceanAksNpVirtualNodeGroupFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetAcceleratedNetworking(val *string) {
	if err := j.validateSetAcceleratedNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratedNetworking",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetDiskPerformance(val *string) {
	if err := j.validateSetDiskPerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskPerformance",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetExcludeSeries(val *[]*string) {
	if err := j.validateSetExcludeSeriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeSeries",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetGpuTypes(val *[]*string) {
	if err := j.validateSetGpuTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuTypes",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetInternalValue(val *OceanAksNpVirtualNodeGroupFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetMaxGpu(val *float64) {
	if err := j.validateSetMaxGpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxGpu",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetMaxMemoryGib(val *float64) {
	if err := j.validateSetMaxMemoryGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMemoryGib",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetMaxVcpu(val *float64) {
	if err := j.validateSetMaxVcpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVcpu",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetMinDisk(val *float64) {
	if err := j.validateSetMinDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDisk",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetMinGpu(val *float64) {
	if err := j.validateSetMinGpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minGpu",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetMinMemoryGib(val *float64) {
	if err := j.validateSetMinMemoryGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMemoryGib",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetMinNics(val *float64) {
	if err := j.validateSetMinNicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNics",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetMinVcpu(val *float64) {
	if err := j.validateSetMinVcpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minVcpu",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetSeries(val *[]*string) {
	if err := j.validateSetSeriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"series",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference)SetVmTypes(val *[]*string) {
	if err := j.validateSetVmTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmTypes",
		val,
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetAcceleratedNetworking() {
	_jsii_.InvokeVoid(
		o,
		"resetAcceleratedNetworking",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetArchitectures() {
	_jsii_.InvokeVoid(
		o,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetDiskPerformance() {
	_jsii_.InvokeVoid(
		o,
		"resetDiskPerformance",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetExcludeSeries() {
	_jsii_.InvokeVoid(
		o,
		"resetExcludeSeries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetGpuTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetGpuTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetMaxGpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxGpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetMaxMemoryGib() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxMemoryGib",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetMaxVcpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxVcpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetMinDisk() {
	_jsii_.InvokeVoid(
		o,
		"resetMinDisk",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetMinGpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMinGpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetMinMemoryGib() {
	_jsii_.InvokeVoid(
		o,
		"resetMinMemoryGib",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetMinNics() {
	_jsii_.InvokeVoid(
		o,
		"resetMinNics",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetMinVcpu() {
	_jsii_.InvokeVoid(
		o,
		"resetMinVcpu",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetSeries() {
	_jsii_.InvokeVoid(
		o,
		"resetSeries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ResetVmTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetVmTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OceanAksNpVirtualNodeGroupFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

