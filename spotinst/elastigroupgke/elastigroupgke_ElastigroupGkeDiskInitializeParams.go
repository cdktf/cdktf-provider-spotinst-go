package elastigroupgke


type ElastigroupGkeDiskInitializeParams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#source_image ElastigroupGke#source_image}.
	SourceImage *string `field:"required" json:"sourceImage" yaml:"sourceImage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#disk_size_gb ElastigroupGke#disk_size_gb}.
	DiskSizeGb *string `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#disk_type ElastigroupGke#disk_type}.
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
}

