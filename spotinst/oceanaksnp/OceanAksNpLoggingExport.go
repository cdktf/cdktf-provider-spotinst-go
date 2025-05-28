// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpLoggingExport struct {
	// azure_blob block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/ocean_aks_np#azure_blob OceanAksNp#azure_blob}
	AzureBlob interface{} `field:"optional" json:"azureBlob" yaml:"azureBlob"`
}

