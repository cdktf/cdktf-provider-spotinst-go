// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAws",
		reflect.TypeOf((*ManagedInstanceAws)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoHealing", GoGetter: "AutoHealing"},
			_jsii_.MemberProperty{JsiiProperty: "autoHealingInput", GoGetter: "AutoHealingInput"},
			_jsii_.MemberProperty{JsiiProperty: "blockDeviceMappings", GoGetter: "BlockDeviceMappings"},
			_jsii_.MemberProperty{JsiiProperty: "blockDeviceMappingsInput", GoGetter: "BlockDeviceMappingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "blockDevicesMode", GoGetter: "BlockDevicesMode"},
			_jsii_.MemberProperty{JsiiProperty: "blockDevicesModeInput", GoGetter: "BlockDevicesModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCredits", GoGetter: "CpuCredits"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCreditsInput", GoGetter: "CpuCreditsInput"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "drainingTimeout", GoGetter: "DrainingTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "drainingTimeoutInput", GoGetter: "DrainingTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimized", GoGetter: "EbsOptimized"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimizedInput", GoGetter: "EbsOptimizedInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticIp", GoGetter: "ElasticIp"},
			_jsii_.MemberProperty{JsiiProperty: "elasticIpInput", GoGetter: "ElasticIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableMonitoring", GoGetter: "EnableMonitoring"},
			_jsii_.MemberProperty{JsiiProperty: "enableMonitoringInput", GoGetter: "EnableMonitoringInput"},
			_jsii_.MemberProperty{JsiiProperty: "fallbackToOndemand", GoGetter: "FallbackToOndemand"},
			_jsii_.MemberProperty{JsiiProperty: "fallbackToOndemandInput", GoGetter: "FallbackToOndemandInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gracePeriod", GoGetter: "GracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "gracePeriodInput", GoGetter: "GracePeriodInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckType", GoGetter: "HealthCheckType"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckTypeInput", GoGetter: "HealthCheckTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "iamInstanceProfile", GoGetter: "IamInstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "iamInstanceProfileInput", GoGetter: "IamInstanceProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageId", GoGetter: "ImageId"},
			_jsii_.MemberProperty{JsiiProperty: "imageIdInput", GoGetter: "ImageIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypes", GoGetter: "InstanceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypesInput", GoGetter: "InstanceTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "integrationRoute53", GoGetter: "IntegrationRoute53"},
			_jsii_.MemberProperty{JsiiProperty: "integrationRoute53Input", GoGetter: "IntegrationRoute53Input"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyPair", GoGetter: "KeyPair"},
			_jsii_.MemberProperty{JsiiProperty: "keyPairInput", GoGetter: "KeyPairInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "lifeCycle", GoGetter: "LifeCycle"},
			_jsii_.MemberProperty{JsiiProperty: "lifeCycleInput", GoGetter: "LifeCycleInput"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancers", GoGetter: "LoadBalancers"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancersInput", GoGetter: "LoadBalancersInput"},
			_jsii_.MemberProperty{JsiiProperty: "managedInstanceAction", GoGetter: "ManagedInstanceAction"},
			_jsii_.MemberProperty{JsiiProperty: "managedInstanceActionInput", GoGetter: "ManagedInstanceActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "minimumInstanceLifetime", GoGetter: "MinimumInstanceLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "minimumInstanceLifetimeInput", GoGetter: "MinimumInstanceLifetimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterface", GoGetter: "NetworkInterface"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceInput", GoGetter: "NetworkInterfaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "optimizationWindows", GoGetter: "OptimizationWindows"},
			_jsii_.MemberProperty{JsiiProperty: "optimizationWindowsInput", GoGetter: "OptimizationWindowsInput"},
			_jsii_.MemberProperty{JsiiProperty: "orientation", GoGetter: "Orientation"},
			_jsii_.MemberProperty{JsiiProperty: "orientationInput", GoGetter: "OrientationInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "persistBlockDevices", GoGetter: "PersistBlockDevices"},
			_jsii_.MemberProperty{JsiiProperty: "persistBlockDevicesInput", GoGetter: "PersistBlockDevicesInput"},
			_jsii_.MemberProperty{JsiiProperty: "persistPrivateIp", GoGetter: "PersistPrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "persistPrivateIpInput", GoGetter: "PersistPrivateIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "persistRootDevice", GoGetter: "PersistRootDevice"},
			_jsii_.MemberProperty{JsiiProperty: "persistRootDeviceInput", GoGetter: "PersistRootDeviceInput"},
			_jsii_.MemberProperty{JsiiProperty: "placementTenancy", GoGetter: "PlacementTenancy"},
			_jsii_.MemberProperty{JsiiProperty: "placementTenancyInput", GoGetter: "PlacementTenancyInput"},
			_jsii_.MemberProperty{JsiiProperty: "preferredType", GoGetter: "PreferredType"},
			_jsii_.MemberProperty{JsiiProperty: "preferredTypeInput", GoGetter: "PreferredTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateIp", GoGetter: "PrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpInput", GoGetter: "PrivateIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "product", GoGetter: "Product"},
			_jsii_.MemberProperty{JsiiProperty: "productInput", GoGetter: "ProductInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putBlockDeviceMappings", GoMethod: "PutBlockDeviceMappings"},
			_jsii_.MemberMethod{JsiiMethod: "putDelete", GoMethod: "PutDelete"},
			_jsii_.MemberMethod{JsiiMethod: "putIntegrationRoute53", GoMethod: "PutIntegrationRoute53"},
			_jsii_.MemberMethod{JsiiMethod: "putLoadBalancers", GoMethod: "PutLoadBalancers"},
			_jsii_.MemberMethod{JsiiMethod: "putManagedInstanceAction", GoMethod: "PutManagedInstanceAction"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkInterface", GoMethod: "PutNetworkInterface"},
			_jsii_.MemberMethod{JsiiMethod: "putResourceTagSpecification", GoMethod: "PutResourceTagSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "putRevertToSpot", GoMethod: "PutRevertToSpot"},
			_jsii_.MemberMethod{JsiiMethod: "putScheduledTask", GoMethod: "PutScheduledTask"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoHealing", GoMethod: "ResetAutoHealing"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockDeviceMappings", GoMethod: "ResetBlockDeviceMappings"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockDevicesMode", GoMethod: "ResetBlockDevicesMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuCredits", GoMethod: "ResetCpuCredits"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDrainingTimeout", GoMethod: "ResetDrainingTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsOptimized", GoMethod: "ResetEbsOptimized"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticIp", GoMethod: "ResetElasticIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableMonitoring", GoMethod: "ResetEnableMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "resetFallbackToOndemand", GoMethod: "ResetFallbackToOndemand"},
			_jsii_.MemberMethod{JsiiMethod: "resetGracePeriod", GoMethod: "ResetGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckType", GoMethod: "ResetHealthCheckType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIamInstanceProfile", GoMethod: "ResetIamInstanceProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegrationRoute53", GoMethod: "ResetIntegrationRoute53"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyPair", GoMethod: "ResetKeyPair"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifeCycle", GoMethod: "ResetLifeCycle"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoadBalancers", GoMethod: "ResetLoadBalancers"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagedInstanceAction", GoMethod: "ResetManagedInstanceAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumInstanceLifetime", GoMethod: "ResetMinimumInstanceLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkInterface", GoMethod: "ResetNetworkInterface"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptimizationWindows", GoMethod: "ResetOptimizationWindows"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrientation", GoMethod: "ResetOrientation"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPersistPrivateIp", GoMethod: "ResetPersistPrivateIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetPersistRootDevice", GoMethod: "ResetPersistRootDevice"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlacementTenancy", GoMethod: "ResetPlacementTenancy"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreferredType", GoMethod: "ResetPreferredType"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIp", GoMethod: "ResetPrivateIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceTagSpecification", GoMethod: "ResetResourceTagSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetRevertToSpot", GoMethod: "ResetRevertToSpot"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduledTask", GoMethod: "ResetScheduledTask"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroupIds", GoMethod: "ResetSecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetShutdownScript", GoMethod: "ResetShutdownScript"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnhealthyDuration", GoMethod: "ResetUnhealthyDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserData", GoMethod: "ResetUserData"},
			_jsii_.MemberMethod{JsiiMethod: "resetUtilizeReservedInstances", GoMethod: "ResetUtilizeReservedInstances"},
			_jsii_.MemberProperty{JsiiProperty: "resourceTagSpecification", GoGetter: "ResourceTagSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "resourceTagSpecificationInput", GoGetter: "ResourceTagSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "revertToSpot", GoGetter: "RevertToSpot"},
			_jsii_.MemberProperty{JsiiProperty: "revertToSpotInput", GoGetter: "RevertToSpotInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledTask", GoGetter: "ScheduledTask"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledTaskInput", GoGetter: "ScheduledTaskInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "shutdownScript", GoGetter: "ShutdownScript"},
			_jsii_.MemberProperty{JsiiProperty: "shutdownScriptInput", GoGetter: "ShutdownScriptInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyDuration", GoGetter: "UnhealthyDuration"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyDurationInput", GoGetter: "UnhealthyDurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "userData", GoGetter: "UserData"},
			_jsii_.MemberProperty{JsiiProperty: "userDataInput", GoGetter: "UserDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "utilizeReservedInstances", GoGetter: "UtilizeReservedInstances"},
			_jsii_.MemberProperty{JsiiProperty: "utilizeReservedInstancesInput", GoGetter: "UtilizeReservedInstancesInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIdInput", GoGetter: "VpcIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAws{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsBlockDeviceMappings",
		reflect.TypeOf((*ManagedInstanceAwsBlockDeviceMappings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsBlockDeviceMappingsEbs",
		reflect.TypeOf((*ManagedInstanceAwsBlockDeviceMappingsEbs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsBlockDeviceMappingsEbsOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsBlockDeviceMappingsEbsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnTermination", GoGetter: "DeleteOnTermination"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnTerminationInput", GoGetter: "DeleteOnTerminationInput"},
			_jsii_.MemberProperty{JsiiProperty: "encrypted", GoGetter: "Encrypted"},
			_jsii_.MemberProperty{JsiiProperty: "encryptedInput", GoGetter: "EncryptedInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iops", GoGetter: "Iops"},
			_jsii_.MemberProperty{JsiiProperty: "iopsInput", GoGetter: "IopsInput"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteOnTermination", GoMethod: "ResetDeleteOnTermination"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncrypted", GoMethod: "ResetEncrypted"},
			_jsii_.MemberMethod{JsiiMethod: "resetIops", GoMethod: "ResetIops"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnapshotId", GoMethod: "ResetSnapshotId"},
			_jsii_.MemberMethod{JsiiMethod: "resetThroughput", GoMethod: "ResetThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeSize", GoMethod: "ResetVolumeSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeType", GoMethod: "ResetVolumeType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotId", GoGetter: "SnapshotId"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotIdInput", GoGetter: "SnapshotIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "throughput", GoGetter: "Throughput"},
			_jsii_.MemberProperty{JsiiProperty: "throughputInput", GoGetter: "ThroughputInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSize", GoGetter: "VolumeSize"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSizeInput", GoGetter: "VolumeSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "volumeType", GoGetter: "VolumeType"},
			_jsii_.MemberProperty{JsiiProperty: "volumeTypeInput", GoGetter: "VolumeTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsBlockDeviceMappingsEbsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsBlockDeviceMappingsList",
		reflect.TypeOf((*ManagedInstanceAwsBlockDeviceMappingsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsBlockDeviceMappingsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsBlockDeviceMappingsOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsBlockDeviceMappingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deviceName", GoGetter: "DeviceName"},
			_jsii_.MemberProperty{JsiiProperty: "deviceNameInput", GoGetter: "DeviceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebs", GoGetter: "Ebs"},
			_jsii_.MemberProperty{JsiiProperty: "ebsInput", GoGetter: "EbsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putEbs", GoMethod: "PutEbs"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbs", GoMethod: "ResetEbs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsBlockDeviceMappingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsConfig",
		reflect.TypeOf((*ManagedInstanceAwsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsDelete",
		reflect.TypeOf((*ManagedInstanceAwsDelete)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsDeleteList",
		reflect.TypeOf((*ManagedInstanceAwsDeleteList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsDeleteList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsDeleteOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsDeleteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amiBackupShouldDeleteImages", GoGetter: "AmiBackupShouldDeleteImages"},
			_jsii_.MemberProperty{JsiiProperty: "amiBackupShouldDeleteImagesInput", GoGetter: "AmiBackupShouldDeleteImagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deallocationConfigShouldDeleteImages", GoGetter: "DeallocationConfigShouldDeleteImages"},
			_jsii_.MemberProperty{JsiiProperty: "deallocationConfigShouldDeleteImagesInput", GoGetter: "DeallocationConfigShouldDeleteImagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmiBackupShouldDeleteImages", GoMethod: "ResetAmiBackupShouldDeleteImages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeallocationConfigShouldDeleteImages", GoMethod: "ResetDeallocationConfigShouldDeleteImages"},
			_jsii_.MemberMethod{JsiiMethod: "resetShouldDeleteNetworkInterfaces", GoMethod: "ResetShouldDeleteNetworkInterfaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetShouldDeleteSnapshots", GoMethod: "ResetShouldDeleteSnapshots"},
			_jsii_.MemberMethod{JsiiMethod: "resetShouldDeleteVolumes", GoMethod: "ResetShouldDeleteVolumes"},
			_jsii_.MemberMethod{JsiiMethod: "resetShouldTerminateInstance", GoMethod: "ResetShouldTerminateInstance"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "shouldDeleteNetworkInterfaces", GoGetter: "ShouldDeleteNetworkInterfaces"},
			_jsii_.MemberProperty{JsiiProperty: "shouldDeleteNetworkInterfacesInput", GoGetter: "ShouldDeleteNetworkInterfacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "shouldDeleteSnapshots", GoGetter: "ShouldDeleteSnapshots"},
			_jsii_.MemberProperty{JsiiProperty: "shouldDeleteSnapshotsInput", GoGetter: "ShouldDeleteSnapshotsInput"},
			_jsii_.MemberProperty{JsiiProperty: "shouldDeleteVolumes", GoGetter: "ShouldDeleteVolumes"},
			_jsii_.MemberProperty{JsiiProperty: "shouldDeleteVolumesInput", GoGetter: "ShouldDeleteVolumesInput"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTerminateInstance", GoGetter: "ShouldTerminateInstance"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTerminateInstanceInput", GoGetter: "ShouldTerminateInstanceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsDeleteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53",
		reflect.TypeOf((*ManagedInstanceAwsIntegrationRoute53)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53Domains",
		reflect.TypeOf((*ManagedInstanceAwsIntegrationRoute53Domains)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53DomainsList",
		reflect.TypeOf((*ManagedInstanceAwsIntegrationRoute53DomainsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53DomainsOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsIntegrationRoute53DomainsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hostedZoneId", GoGetter: "HostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "hostedZoneIdInput", GoGetter: "HostedZoneIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putRecordSets", GoMethod: "PutRecordSets"},
			_jsii_.MemberProperty{JsiiProperty: "recordSets", GoGetter: "RecordSets"},
			_jsii_.MemberProperty{JsiiProperty: "recordSetsInput", GoGetter: "RecordSetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "recordSetType", GoGetter: "RecordSetType"},
			_jsii_.MemberProperty{JsiiProperty: "recordSetTypeInput", GoGetter: "RecordSetTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecordSetType", GoMethod: "ResetRecordSetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpotinstAcctId", GoMethod: "ResetSpotinstAcctId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "spotinstAcctId", GoGetter: "SpotinstAcctId"},
			_jsii_.MemberProperty{JsiiProperty: "spotinstAcctIdInput", GoGetter: "SpotinstAcctIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53DomainsRecordSets",
		reflect.TypeOf((*ManagedInstanceAwsIntegrationRoute53DomainsRecordSets)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsList",
		reflect.TypeOf((*ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsePublicDns", GoMethod: "ResetUsePublicDns"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsePublicIp", GoMethod: "ResetUsePublicIp"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "usePublicDns", GoGetter: "UsePublicDns"},
			_jsii_.MemberProperty{JsiiProperty: "usePublicDnsInput", GoGetter: "UsePublicDnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "usePublicIp", GoGetter: "UsePublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "usePublicIpInput", GoGetter: "UsePublicIpInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsIntegrationRoute53DomainsRecordSetsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsIntegrationRoute53OutputReference",
		reflect.TypeOf((*ManagedInstanceAwsIntegrationRoute53OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domains", GoGetter: "Domains"},
			_jsii_.MemberProperty{JsiiProperty: "domainsInput", GoGetter: "DomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putDomains", GoMethod: "PutDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsIntegrationRoute53OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsLoadBalancers",
		reflect.TypeOf((*ManagedInstanceAwsLoadBalancers)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsLoadBalancersList",
		reflect.TypeOf((*ManagedInstanceAwsLoadBalancersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsLoadBalancersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsLoadBalancersOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsLoadBalancersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "arnInput", GoGetter: "ArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoWeight", GoGetter: "AutoWeight"},
			_jsii_.MemberProperty{JsiiProperty: "autoWeightInput", GoGetter: "AutoWeightInput"},
			_jsii_.MemberProperty{JsiiProperty: "azAwareness", GoGetter: "AzAwareness"},
			_jsii_.MemberProperty{JsiiProperty: "azAwarenessInput", GoGetter: "AzAwarenessInput"},
			_jsii_.MemberProperty{JsiiProperty: "balancerId", GoGetter: "BalancerId"},
			_jsii_.MemberProperty{JsiiProperty: "balancerIdInput", GoGetter: "BalancerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetArn", GoMethod: "ResetArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoWeight", GoMethod: "ResetAutoWeight"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzAwareness", GoMethod: "ResetAzAwareness"},
			_jsii_.MemberMethod{JsiiMethod: "resetBalancerId", GoMethod: "ResetBalancerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetSetId", GoMethod: "ResetTargetSetId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetSetId", GoGetter: "TargetSetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetSetIdInput", GoGetter: "TargetSetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsLoadBalancersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsManagedInstanceAction",
		reflect.TypeOf((*ManagedInstanceAwsManagedInstanceAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsManagedInstanceActionOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsManagedInstanceActionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsManagedInstanceActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsNetworkInterface",
		reflect.TypeOf((*ManagedInstanceAwsNetworkInterface)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsNetworkInterfaceList",
		reflect.TypeOf((*ManagedInstanceAwsNetworkInterfaceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsNetworkInterfaceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsNetworkInterfaceOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsNetworkInterfaceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "associateIpv6Address", GoGetter: "AssociateIpv6Address"},
			_jsii_.MemberProperty{JsiiProperty: "associateIpv6AddressInput", GoGetter: "AssociateIpv6AddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "associatePublicIpAddress", GoGetter: "AssociatePublicIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "associatePublicIpAddressInput", GoGetter: "AssociatePublicIpAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deviceIndex", GoGetter: "DeviceIndex"},
			_jsii_.MemberProperty{JsiiProperty: "deviceIndexInput", GoGetter: "DeviceIndexInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssociateIpv6Address", GoMethod: "ResetAssociateIpv6Address"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssociatePublicIpAddress", GoMethod: "ResetAssociatePublicIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsNetworkInterfaceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsResourceTagSpecification",
		reflect.TypeOf((*ManagedInstanceAwsResourceTagSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsResourceTagSpecificationList",
		reflect.TypeOf((*ManagedInstanceAwsResourceTagSpecificationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsResourceTagSpecificationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsResourceTagSpecificationOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsResourceTagSpecificationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetShouldTagAmis", GoMethod: "ResetShouldTagAmis"},
			_jsii_.MemberMethod{JsiiMethod: "resetShouldTagEnis", GoMethod: "ResetShouldTagEnis"},
			_jsii_.MemberMethod{JsiiMethod: "resetShouldTagSnapshots", GoMethod: "ResetShouldTagSnapshots"},
			_jsii_.MemberMethod{JsiiMethod: "resetShouldTagVolumes", GoMethod: "ResetShouldTagVolumes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTagAmis", GoGetter: "ShouldTagAmis"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTagAmisInput", GoGetter: "ShouldTagAmisInput"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTagEnis", GoGetter: "ShouldTagEnis"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTagEnisInput", GoGetter: "ShouldTagEnisInput"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTagSnapshots", GoGetter: "ShouldTagSnapshots"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTagSnapshotsInput", GoGetter: "ShouldTagSnapshotsInput"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTagVolumes", GoGetter: "ShouldTagVolumes"},
			_jsii_.MemberProperty{JsiiProperty: "shouldTagVolumesInput", GoGetter: "ShouldTagVolumesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsResourceTagSpecificationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsRevertToSpot",
		reflect.TypeOf((*ManagedInstanceAwsRevertToSpot)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsRevertToSpotOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsRevertToSpotOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "performAt", GoGetter: "PerformAt"},
			_jsii_.MemberProperty{JsiiProperty: "performAtInput", GoGetter: "PerformAtInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsRevertToSpotOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsScheduledTask",
		reflect.TypeOf((*ManagedInstanceAwsScheduledTask)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsScheduledTaskList",
		reflect.TypeOf((*ManagedInstanceAwsScheduledTaskList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsScheduledTaskList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsScheduledTaskOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsScheduledTaskOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "cronExpression", GoGetter: "CronExpression"},
			_jsii_.MemberProperty{JsiiProperty: "cronExpressionInput", GoGetter: "CronExpressionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frequency", GoGetter: "Frequency"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyInput", GoGetter: "FrequencyInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isEnabled", GoGetter: "IsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isEnabledInput", GoGetter: "IsEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCronExpression", GoMethod: "ResetCronExpression"},
			_jsii_.MemberMethod{JsiiMethod: "resetFrequency", GoMethod: "ResetFrequency"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsEnabled", GoMethod: "ResetIsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartTime", GoMethod: "ResetStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startTime", GoGetter: "StartTime"},
			_jsii_.MemberProperty{JsiiProperty: "startTimeInput", GoGetter: "StartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskType", GoGetter: "TaskType"},
			_jsii_.MemberProperty{JsiiProperty: "taskTypeInput", GoGetter: "TaskTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsScheduledTaskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsTags",
		reflect.TypeOf((*ManagedInstanceAwsTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsTagsList",
		reflect.TypeOf((*ManagedInstanceAwsTagsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-spotinst.managedInstanceAws.ManagedInstanceAwsTagsOutputReference",
		reflect.TypeOf((*ManagedInstanceAwsTagsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKey", GoMethod: "ResetKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedInstanceAwsTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
