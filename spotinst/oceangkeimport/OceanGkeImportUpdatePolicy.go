package oceangkeimport


type OceanGkeImportUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_gke_import#should_roll OceanGkeImport#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_gke_import#conditioned_roll OceanGkeImport#conditioned_roll}.
	ConditionedRoll interface{} `field:"optional" json:"conditionedRoll" yaml:"conditionedRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_gke_import#roll_config OceanGkeImport#roll_config}
	RollConfig *OceanGkeImportUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

