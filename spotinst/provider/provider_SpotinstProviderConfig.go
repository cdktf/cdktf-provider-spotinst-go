package provider


type SpotinstProviderConfig struct {
	// Spotinst Account ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst#account SpotinstProvider#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst#alias SpotinstProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Spotinst SDK Feature Flags.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst#feature_flags SpotinstProvider#feature_flags}
	FeatureFlags *string `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// Spotinst Personal API Access Token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst#token SpotinstProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

