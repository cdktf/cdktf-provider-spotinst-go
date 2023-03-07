package oceanaks


type OceanAksLoadBalancer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#backend_pool_names OceanAks#backend_pool_names}.
	BackendPoolNames *[]*string `field:"optional" json:"backendPoolNames" yaml:"backendPoolNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#load_balancer_sku OceanAks#load_balancer_sku}.
	LoadBalancerSku *string `field:"optional" json:"loadBalancerSku" yaml:"loadBalancerSku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#name OceanAks#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#resource_group_name OceanAks#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#type OceanAks#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}
