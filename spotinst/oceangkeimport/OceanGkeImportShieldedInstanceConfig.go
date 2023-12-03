// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport


type OceanGkeImportShieldedInstanceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_gke_import#enable_integrity_monitoring OceanGkeImport#enable_integrity_monitoring}.
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_gke_import#enable_secure_boot OceanGkeImport#enable_secure_boot}.
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
}

