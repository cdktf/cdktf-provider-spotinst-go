package provider


type SpotinstProviderConfig struct {
	// Spotinst Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs#account SpotinstProvider#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs#alias SpotinstProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Spotinst SDK Feature Flags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs#feature_flags SpotinstProvider#feature_flags}
	FeatureFlags *string `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// Spotinst Personal API Access Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs#token SpotinstProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

