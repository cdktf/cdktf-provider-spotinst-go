// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAksExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#api_version OceanAks#api_version}.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#minor_version_auto_upgrade OceanAks#minor_version_auto_upgrade}.
	MinorVersionAutoUpgrade interface{} `field:"optional" json:"minorVersionAutoUpgrade" yaml:"minorVersionAutoUpgrade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#name OceanAks#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#publisher OceanAks#publisher}.
	Publisher *string `field:"optional" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#type OceanAks#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

