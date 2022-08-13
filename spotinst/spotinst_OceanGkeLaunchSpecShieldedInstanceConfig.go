// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeLaunchSpecShieldedInstanceConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#enable_integrity_monitoring OceanGkeLaunchSpec#enable_integrity_monitoring}.
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#enable_secure_boot OceanGkeLaunchSpec#enable_secure_boot}.
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
}

