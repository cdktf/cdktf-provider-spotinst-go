package multailistener


type MultaiListenerTlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/multai_listener#certificate_ids MultaiListener#certificate_ids}.
	CertificateIds *[]*string `field:"required" json:"certificateIds" yaml:"certificateIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/multai_listener#cipher_suites MultaiListener#cipher_suites}.
	CipherSuites *[]*string `field:"required" json:"cipherSuites" yaml:"cipherSuites"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/multai_listener#max_version MultaiListener#max_version}.
	MaxVersion *string `field:"required" json:"maxVersion" yaml:"maxVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/multai_listener#min_version MultaiListener#min_version}.
	MinVersion *string `field:"required" json:"minVersion" yaml:"minVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/multai_listener#prefer_server_cipher_suites MultaiListener#prefer_server_cipher_suites}.
	PreferServerCipherSuites interface{} `field:"required" json:"preferServerCipherSuites" yaml:"preferServerCipherSuites"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/multai_listener#session_tickets_disabled MultaiListener#session_tickets_disabled}.
	SessionTicketsDisabled interface{} `field:"required" json:"sessionTicketsDisabled" yaml:"sessionTicketsDisabled"`
}

