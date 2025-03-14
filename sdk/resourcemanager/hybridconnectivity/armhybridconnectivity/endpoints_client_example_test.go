// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhybridconnectivity_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
	"log"
)

// Generated from example definition: 2024-12-01/EndpointsPutCustom.json
func ExampleEndpointsClient_CreateOrUpdate_hybridConnectivityEndpointsPutCustom() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().CreateOrUpdate(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "custom", armhybridconnectivity.EndpointResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientCreateOrUpdateResponse{
	// 	EndpointResource: &armhybridconnectivity.EndpointResource{
	// 		Name: to.Ptr("custom"),
	// 		Type: to.Ptr("Microsoft.HybridConnectivity/endpoints"),
	// 		ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/custom"),
	// 		Properties: &armhybridconnectivity.EndpointProperties{
	// 			Type: to.Ptr(armhybridconnectivity.TypeCustom),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ResourceID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.Relay/namespaces/custom-relay-namespace"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/EndpointsPutDefault.json
func ExampleEndpointsClient_CreateOrUpdate_hybridConnectivityEndpointsPutDefault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().CreateOrUpdate(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "default", armhybridconnectivity.EndpointResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientCreateOrUpdateResponse{
	// 	EndpointResource: &armhybridconnectivity.EndpointResource{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.HybridConnectivity/endpoints"),
	// 		ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default"),
	// 		Properties: &armhybridconnectivity.EndpointProperties{
	// 			Type: to.Ptr(armhybridconnectivity.TypeDefault),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/EndpointsDeleteDefault.json
func ExampleEndpointsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().Delete(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientDeleteResponse{
	// }
}

// Generated from example definition: 2024-12-01/EndpointsGetCustom.json
func ExampleEndpointsClient_Get_hybridConnectivityEndpointsGetCustom() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().Get(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "custom", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientGetResponse{
	// 	EndpointResource: &armhybridconnectivity.EndpointResource{
	// 		Name: to.Ptr("custom"),
	// 		Type: to.Ptr("Microsoft.HybridConnectivity/endpoints"),
	// 		ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/custom"),
	// 		Properties: &armhybridconnectivity.EndpointProperties{
	// 			Type: to.Ptr(armhybridconnectivity.TypeCustom),
	// 			ResourceID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.Relay/namespaces/custom-relay-namespace"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/EndpointsGetDefault.json
func ExampleEndpointsClient_Get_hybridConnectivityEndpointsGetDefault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().Get(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientGetResponse{
	// 	EndpointResource: &armhybridconnectivity.EndpointResource{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.HybridConnectivity/endpoints"),
	// 		ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default"),
	// 		Properties: &armhybridconnectivity.EndpointProperties{
	// 			Type: to.Ptr(armhybridconnectivity.TypeDefault),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/EndpointsList.json
func ExampleEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEndpointsClient().NewListPager("subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armhybridconnectivity.EndpointsClientListResponse{
		// 	EndpointsList: armhybridconnectivity.EndpointsList{
		// 		Value: []*armhybridconnectivity.EndpointResource{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.HybridConnectivity/endpoints"),
		// 				ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default"),
		// 				Properties: &armhybridconnectivity.EndpointProperties{
		// 					Type: to.Ptr(armhybridconnectivity.TypeDefault),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("custom"),
		// 				Type: to.Ptr("Microsoft.HybridConnectivity/endpoints"),
		// 				ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/custom"),
		// 				Properties: &armhybridconnectivity.EndpointProperties{
		// 					Type: to.Ptr(armhybridconnectivity.TypeCustom),
		// 					ResourceID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.Relay/namespaces/custom-relay-namespace"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2024-12-01/EndpointsPostListCredentials.json
func ExampleEndpointsClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().ListCredentials(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "default", &armhybridconnectivity.EndpointsClientListCredentialsOptions{
		ListCredentialsRequest: &armhybridconnectivity.ListCredentialsRequest{
			ServiceName: to.Ptr(armhybridconnectivity.ServiceNameSSH),
		},
		Expiresin: to.Ptr[int64](10800)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientListCredentialsResponse{
	// 	EndpointAccessResource: &armhybridconnectivity.EndpointAccessResource{
	// 		Relay: &armhybridconnectivity.RelayNamespaceAccessProperties{
	// 			AccessKey: to.Ptr("SharedAccessSignature sr=http%3A%2F%2Fazgnrelay-eastus-l1.servicebus.windows.net%2Fmicrosoft.kubernetes%2Fconnectedclusters%2Fa0e1fd7d1d974ddf6b11a952d67679c9f12c006eee16861857a8268da4eb1498%2F1619989456957411072%2F&sig=WxDwPF6AmmODaMHNnBGDSm773UG%2B%2Be"),
	// 			ExpiresOn: to.Ptr[int64](1620000256),
	// 			HybridConnectionName: to.Ptr("microsoft.kubernetes/connectedclusters/a0e1fd7d1d974ddf6b11a952d67679c9f12c006eee16861857a8268da4eb1498/1619989456957411072"),
	// 			NamespaceName: to.Ptr("azgnrelay-eastus-l1"),
	// 			NamespaceNameSuffix: to.Ptr("servicebus.windows.net"),
	// 			ServiceConfigurationToken: to.Ptr("SSHvjqH=pTlKql=RtMGw/-k5VFBxSYHIiq5ZgbGFcLkNrDNz5fDsinCN2zkG"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/EndpointsPostListIngressGatewayCredentials.json
func ExampleEndpointsClient_ListIngressGatewayCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().ListIngressGatewayCredentials(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/arcGroup/providers/Microsoft.ArcPlaceHolder/ProvisionedClusters/cluster0", "default", &armhybridconnectivity.EndpointsClientListIngressGatewayCredentialsOptions{
		Expiresin: to.Ptr[int64](10800)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientListIngressGatewayCredentialsResponse{
	// 	IngressGatewayResource: &armhybridconnectivity.IngressGatewayResource{
	// 		Ingress: &armhybridconnectivity.IngressProfileProperties{
	// 			AADProfile: &armhybridconnectivity.AADProfileProperties{
	// 				ServerID: to.Ptr("6256c85f-0aad-4d50-b960-e6e9b21efe35"),
	// 				TenantID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
	// 			},
	// 			Hostname: to.Ptr("clusterhostname"),
	// 		},
	// 		Relay: &armhybridconnectivity.RelayNamespaceAccessProperties{
	// 			AccessKey: to.Ptr("SharedAccessSignature sr=http%3A%2F%2Fazgnrelay-eastus-l1.servicebus.windows.net%2Fmicrosoft.provisionedcluster%hci"),
	// 			ExpiresOn: to.Ptr[int64](1620000256),
	// 			HybridConnectionName: to.Ptr("microsoft.arcplaceholder/provisionedclusters/000/1619989456957411072"),
	// 			NamespaceName: to.Ptr("relaynamespace"),
	// 			NamespaceNameSuffix: to.Ptr("servicebus.windows.net"),
	// 			ServiceConfigurationToken: to.Ptr("SSHvjqH=pTlKql=RtMGw/-k5VFBxSYHIiq5ZgbGFcLkNrDNz5fDsinCN2zkG"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/EndpointsPostListManagedProxyDetails.json
func ExampleEndpointsClient_ListManagedProxyDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().ListManagedProxyDetails(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/arcGroup/providers/Microsoft.Compute/virtualMachines/vm00006", "default", armhybridconnectivity.ManagedProxyRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientListManagedProxyDetailsResponse{
	// 	ManagedProxyResource: &armhybridconnectivity.ManagedProxyResource{
	// 		ExpiresOn: to.Ptr[int64](1620000256),
	// 		Proxy: to.Ptr("uid.r.proxy.arc.com"),
	// 	},
	// }
}

// Generated from example definition: 2024-12-01/EndpointsPatchDefault.json
func ExampleEndpointsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().Update(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "default", armhybridconnectivity.EndpointResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientUpdateResponse{
	// 	EndpointResource: &armhybridconnectivity.EndpointResource{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.HybridConnectivity/endpoints"),
	// 		ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default"),
	// 		Properties: &armhybridconnectivity.EndpointProperties{
	// 			Type: to.Ptr(armhybridconnectivity.TypeDefault),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
