package statefulnodeazure


type StatefulNodeAzureSecretVaultCertificates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs/resources/stateful_node_azure#certificate_store StatefulNodeAzure#certificate_store}.
	CertificateStore *string `field:"optional" json:"certificateStore" yaml:"certificateStore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs/resources/stateful_node_azure#certificate_url StatefulNodeAzure#certificate_url}.
	CertificateUrl *string `field:"optional" json:"certificateUrl" yaml:"certificateUrl"`
}

