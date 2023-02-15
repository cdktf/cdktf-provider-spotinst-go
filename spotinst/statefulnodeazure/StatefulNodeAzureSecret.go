package statefulnodeazure


type StatefulNodeAzureSecret struct {
	// source_vault block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#source_vault StatefulNodeAzure#source_vault}
	SourceVault interface{} `field:"required" json:"sourceVault" yaml:"sourceVault"`
	// vault_certificates block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#vault_certificates StatefulNodeAzure#vault_certificates}
	VaultCertificates interface{} `field:"required" json:"vaultCertificates" yaml:"vaultCertificates"`
}

