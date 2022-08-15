// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGcpDiskInitializeParams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#source_image ElastigroupGcp#source_image}.
	SourceImage *string `field:"required" json:"sourceImage" yaml:"sourceImage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#disk_size_gb ElastigroupGcp#disk_size_gb}.
	DiskSizeGb *string `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#disk_type ElastigroupGcp#disk_type}.
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
}
