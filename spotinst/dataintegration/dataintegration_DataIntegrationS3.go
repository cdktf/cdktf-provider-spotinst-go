package dataintegration


type DataIntegrationS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/data_integration#bucket_name DataIntegration#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/data_integration#subdir DataIntegration#subdir}.
	Subdir *string `field:"optional" json:"subdir" yaml:"subdir"`
}

