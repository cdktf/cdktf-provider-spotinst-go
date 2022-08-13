// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAzureV3Strategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure_v3#draining_timeout ElastigroupAzureV3#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure_v3#fallback_to_on_demand ElastigroupAzureV3#fallback_to_on_demand}.
	FallbackToOnDemand interface{} `field:"optional" json:"fallbackToOnDemand" yaml:"fallbackToOnDemand"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure_v3#od_count ElastigroupAzureV3#od_count}.
	OdCount *float64 `field:"optional" json:"odCount" yaml:"odCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure_v3#spot_percentage ElastigroupAzureV3#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
}

