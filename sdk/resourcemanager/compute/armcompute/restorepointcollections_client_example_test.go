//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_CreateOrUpdate_ForCrossRegionCopy.json
func ExampleRestorePointCollectionsClient_CreateOrUpdate_createOrUpdateARestorePointCollectionForCrossRegionCopy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().CreateOrUpdate(ctx, "myResourceGroup", "myRpc", armcompute.RestorePointCollection{
		Location: to.Ptr("norwayeast"),
		Tags: map[string]*string{
			"myTag1": to.Ptr("tagValue1"),
		},
		Properties: &armcompute.RestorePointCollectionProperties{
			Source: &armcompute.RestorePointCollectionSourceProperties{
				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePointCollection = armcompute.RestorePointCollection{
	// 	Name: to.Ptr("myRpc"),
	// 	Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/myRpc"),
	// 	Location: to.Ptr("norwayeast"),
	// 	Tags: map[string]*string{
	// 		"myTag1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armcompute.RestorePointCollectionProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RestorePointCollectionID: to.Ptr("638f052b-a7c2-450c-89e7-6a3b8f1d6a7c"),
	// 		Source: &armcompute.RestorePointCollectionSourceProperties{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 			Location: to.Ptr("eastus"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_CreateOrUpdate.json
func ExampleRestorePointCollectionsClient_CreateOrUpdate_createOrUpdateARestorePointCollection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().CreateOrUpdate(ctx, "myResourceGroup", "myRpc", armcompute.RestorePointCollection{
		Location: to.Ptr("norwayeast"),
		Tags: map[string]*string{
			"myTag1": to.Ptr("tagValue1"),
		},
		Properties: &armcompute.RestorePointCollectionProperties{
			Source: &armcompute.RestorePointCollectionSourceProperties{
				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePointCollection = armcompute.RestorePointCollection{
	// 	Name: to.Ptr("myRpc"),
	// 	Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/myRpc"),
	// 	Location: to.Ptr("norwayeast"),
	// 	Tags: map[string]*string{
	// 		"myTag1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armcompute.RestorePointCollectionProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RestorePointCollectionID: to.Ptr("638f052b-a7c2-450c-89e7-6a3b8f1d6a7c"),
	// 		Source: &armcompute.RestorePointCollectionSourceProperties{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 			Location: to.Ptr("eastus"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_Update_MaximumSet_Gen.json
func ExampleRestorePointCollectionsClient_Update_restorePointCollectionUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaa", armcompute.RestorePointCollectionUpdate{
		Tags: map[string]*string{
			"key8536": to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		},
		Properties: &armcompute.RestorePointCollectionProperties{
			Source: &armcompute.RestorePointCollectionSourceProperties{
				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePointCollection = armcompute.RestorePointCollection{
	// 	Name: to.Ptr("myRpc"),
	// 	Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/myRpc"),
	// 	Location: to.Ptr("norwayeast"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcompute.RestorePointCollectionProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RestorePointCollectionID: to.Ptr("638f052b-a7c2-450c-89e7-6a3b8f1d6a7c"),
	// 		RestorePoints: []*armcompute.RestorePoint{
	// 			{
	// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 				Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				ID: to.Ptr("aaaaaaaaaaa"),
	// 				Properties: &armcompute.RestorePointProperties{
	// 					ConsistencyMode: to.Ptr(armcompute.ConsistencyModeTypesCrashConsistent),
	// 					ExcludeDisks: []*armcompute.APIEntityReference{
	// 						{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
	// 					}},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					SourceMetadata: &armcompute.RestorePointSourceMetadata{
	// 						DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 							BootDiagnostics: &armcompute.BootDiagnostics{
	// 								Enabled: to.Ptr(true),
	// 								StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 							},
	// 						},
	// 						HardwareProfile: &armcompute.HardwareProfile{
	// 							VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardB1S),
	// 							VMSizeProperties: &armcompute.VMSizeProperties{
	// 								VCPUsAvailable: to.Ptr[int32](9),
	// 								VCPUsPerCore: to.Ptr[int32](12),
	// 							},
	// 						},
	// 						LicenseType: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 						Location: to.Ptr("westus"),
	// 						OSProfile: &armcompute.OSProfile{
	// 							AdminUsername: to.Ptr("admin"),
	// 							AllowExtensionOperations: to.Ptr(true),
	// 							ComputerName: to.Ptr("computerName"),
	// 							CustomData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 							LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								PatchSettings: &armcompute.LinuxPatchSettings{
	// 									AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
	// 									PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
	// 								},
	// 								ProvisionVMAgent: to.Ptr(true),
	// 								SSH: &armcompute.SSHConfiguration{
	// 									PublicKeys: []*armcompute.SSHPublicKey{
	// 										{
	// 											Path: to.Ptr("aaa"),
	// 											KeyData: to.Ptr("aaaaaa"),
	// 									}},
	// 								},
	// 							},
	// 							RequireGuestProvisionSignal: to.Ptr(true),
	// 							Secrets: []*armcompute.VaultSecretGroup{
	// 								{
	// 									SourceVault: &armcompute.SubResource{
	// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 									},
	// 									VaultCertificates: []*armcompute.VaultCertificate{
	// 										{
	// 											CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 											CertificateURL: to.Ptr("aaaaaaa"),
	// 									}},
	// 							}},
	// 							WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 								AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
	// 									{
	// 										ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
	// 										Content: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 										PassName: to.Ptr("OobeSystem"),
	// 										SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
	// 								}},
	// 								EnableAutomaticUpdates: to.Ptr(true),
	// 								PatchSettings: &armcompute.PatchSettings{
	// 									AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
	// 									EnableHotpatching: to.Ptr(true),
	// 									PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
	// 								},
	// 								ProvisionVMAgent: to.Ptr(true),
	// 								TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 								WinRM: &armcompute.WinRMConfiguration{
	// 									Listeners: []*armcompute.WinRMListener{
	// 										{
	// 											CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 											Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
	// 									}},
	// 								},
	// 							},
	// 						},
	// 						SecurityProfile: &armcompute.SecurityProfile{
	// 							EncryptionAtHost: to.Ptr(true),
	// 							SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
	// 							UefiSettings: &armcompute.UefiSettings{
	// 								SecureBootEnabled: to.Ptr(true),
	// 								VTpmEnabled: to.Ptr(true),
	// 							},
	// 						},
	// 						StorageProfile: &armcompute.RestorePointSourceVMStorageProfile{
	// 							DataDisks: []*armcompute.RestorePointSourceVMDataDisk{
	// 								{
	// 									Name: to.Ptr("testingexcludedisk_DataDisk_1"),
	// 									Caching: to.Ptr(armcompute.CachingTypesNone),
	// 									DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_DataDisk_1_68785190-1acb-4d5e-a8ae-705b45f3dca5"),
	// 									},
	// 									DiskSizeGB: to.Ptr[int32](24),
	// 									Lun: to.Ptr[int32](1),
	// 									ManagedDisk: &armcompute.ManagedDiskParameters{
	// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_DataDisk_1"),
	// 										DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 											ID: to.Ptr("aaaaaaaaaaaa"),
	// 										},
	// 										StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 									},
	// 							}},
	// 							DiskControllerType: to.Ptr(armcompute.DiskControllerTypesNVMe),
	// 							OSDisk: &armcompute.RestorePointSourceVMOSDisk{
	// 								Name: to.Ptr("testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 								Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 								DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
	// 								},
	// 								DiskSizeGB: to.Ptr[int32](3),
	// 								EncryptionSettings: &armcompute.DiskEncryptionSettings{
	// 									DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
	// 										SecretURL: to.Ptr("aaaaaaaa"),
	// 										SourceVault: &armcompute.SubResource{
	// 											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 										},
	// 									},
	// 									Enabled: to.Ptr(true),
	// 									KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
	// 										KeyURL: to.Ptr("aaaaaaaaaaaaaa"),
	// 										SourceVault: &armcompute.SubResource{
	// 											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 										},
	// 									},
	// 								},
	// 								ManagedDisk: &armcompute.ManagedDiskParameters{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 										ID: to.Ptr("aaaaaaaaaaaa"),
	// 									},
	// 									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 								},
	// 								OSType: to.Ptr(armcompute.OperatingSystemTypeWindows),
	// 							},
	// 						},
	// 						VMID: to.Ptr("76d6541e-80bd-4dc1-932b-3cae4cfb80e7"),
	// 					},
	// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.593Z"); return t}()),
	// 				},
	// 		}},
	// 		Source: &armcompute.RestorePointCollectionSourceProperties{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 			Location: to.Ptr("eastus"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_Update_MinimumSet_Gen.json
func ExampleRestorePointCollectionsClient_Update_restorePointCollectionUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaa", armcompute.RestorePointCollectionUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePointCollection = armcompute.RestorePointCollection{
	// 	Location: to.Ptr("norwayeast"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_Delete_MaximumSet_Gen.json
func ExampleRestorePointCollectionsClient_BeginDelete_restorePointCollectionDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestorePointCollectionsClient().BeginDelete(ctx, "rgcompute", "aaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_Delete_MinimumSet_Gen.json
func ExampleRestorePointCollectionsClient_BeginDelete_restorePointCollectionDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestorePointCollectionsClient().BeginDelete(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_Get.json
func ExampleRestorePointCollectionsClient_Get_getARestorePointCollectionButNotTheRestorePointsContainedInTheRestorePointCollection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().Get(ctx, "myResourceGroup", "myRpc", &armcompute.RestorePointCollectionsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePointCollection = armcompute.RestorePointCollection{
	// 	Name: to.Ptr("myRpc"),
	// 	Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/myRpc"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"myTag1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armcompute.RestorePointCollectionProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RestorePointCollectionID: to.Ptr("59f04a5d-f783-4200-a1bd-d3f464e8c4b4"),
	// 		Source: &armcompute.RestorePointCollectionSourceProperties{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/myRpc"),
	// 			Location: to.Ptr("eastus"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_Get_WithContainedRestorePoints.json
func ExampleRestorePointCollectionsClient_Get_getARestorePointCollectionIncludingTheRestorePointsContainedInTheRestorePointCollection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().Get(ctx, "myResourceGroup", "rpcName", &armcompute.RestorePointCollectionsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePointCollection = armcompute.RestorePointCollection{
	// 	Name: to.Ptr("rpcName"),
	// 	Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"myTag1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armcompute.RestorePointCollectionProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RestorePointCollectionID: to.Ptr("59f04a5d-f783-4200-a1bd-d3f464e8c4b4"),
	// 		RestorePoints: []*armcompute.RestorePoint{
	// 			{
	// 				Name: to.Ptr("restorePointName"),
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName"),
	// 				Properties: &armcompute.RestorePointProperties{
	// 					ConsistencyMode: to.Ptr(armcompute.ConsistencyModeTypesApplicationConsistent),
	// 					ExcludeDisks: []*armcompute.APIEntityReference{
	// 						{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm8768_disk2_fe6ffde4f69b491ca33fb984d5bcd89f"),
	// 					}},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					SourceMetadata: &armcompute.RestorePointSourceMetadata{
	// 						DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 							BootDiagnostics: &armcompute.BootDiagnostics{
	// 								Enabled: to.Ptr(true),
	// 							},
	// 						},
	// 						HardwareProfile: &armcompute.HardwareProfile{
	// 							VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardB1S),
	// 						},
	// 						Location: to.Ptr("westus"),
	// 						OSProfile: &armcompute.OSProfile{
	// 							AdminUsername: to.Ptr("admin"),
	// 							AllowExtensionOperations: to.Ptr(true),
	// 							ComputerName: to.Ptr("computerName"),
	// 							RequireGuestProvisionSignal: to.Ptr(true),
	// 							Secrets: []*armcompute.VaultSecretGroup{
	// 							},
	// 							WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 								EnableAutomaticUpdates: to.Ptr(true),
	// 								ProvisionVMAgent: to.Ptr(true),
	// 							},
	// 						},
	// 						StorageProfile: &armcompute.RestorePointSourceVMStorageProfile{
	// 							DataDisks: []*armcompute.RestorePointSourceVMDataDisk{
	// 								{
	// 									Name: to.Ptr("testingexcludedisk_DataDisk_1"),
	// 									Caching: to.Ptr(armcompute.CachingTypesNone),
	// 									DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_DataDisk_1_68785190-1acb-4d5e-a8ae-705b45f3dca5"),
	// 									},
	// 									Lun: to.Ptr[int32](1),
	// 									ManagedDisk: &armcompute.ManagedDiskParameters{
	// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_DataDisk_1"),
	// 										StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 									},
	// 							}},
	// 							DiskControllerType: to.Ptr(armcompute.DiskControllerTypesNVMe),
	// 							OSDisk: &armcompute.RestorePointSourceVMOSDisk{
	// 								Name: to.Ptr("testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 								Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 								DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
	// 								},
	// 								ManagedDisk: &armcompute.ManagedDiskParameters{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 								},
	// 								OSType: to.Ptr(armcompute.OperatingSystemTypeWindows),
	// 							},
	// 						},
	// 						VMID: to.Ptr("76d6541e-80bd-4dc1-932b-3cae4cfb80e7"),
	// 					},
	// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-27T20:35:05.840Z"); return t}()),
	// 				},
	// 		}},
	// 		Source: &armcompute.RestorePointCollectionSourceProperties{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 			Location: to.Ptr("eastus"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_ListByResourceGroup.json
func ExampleRestorePointCollectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorePointCollectionsClient().NewListPager("myResourceGroup", nil)
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
		// page.RestorePointCollectionListResult = armcompute.RestorePointCollectionListResult{
		// 	Value: []*armcompute.RestorePointCollection{
		// 		{
		// 			Name: to.Ptr("restorePointCollection1"),
		// 			Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/restorePointCollection1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armcompute.RestorePointCollectionProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RestorePointCollectionID: to.Ptr("59f04a5d-f783-4200-a1bd-d3f464e8c4b4"),
		// 				Source: &armcompute.RestorePointCollectionSourceProperties{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/restorePointCollection1"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("restorePointCollection2"),
		// 			Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/restorePointCollection2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armcompute.RestorePointCollectionProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RestorePointCollectionID: to.Ptr("2875c590-e337-4102-9668-4f5b7e3f98a4"),
		// 				Source: &armcompute.RestorePointCollectionSourceProperties{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/restorePointCollection2"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePointCollection_ListBySubscription.json
func ExampleRestorePointCollectionsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorePointCollectionsClient().NewListAllPager(nil)
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
		// page.RestorePointCollectionListResult = armcompute.RestorePointCollectionListResult{
		// 	Value: []*armcompute.RestorePointCollection{
		// 		{
		// 			Name: to.Ptr("restorePointCollection1"),
		// 			Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resourceGroup1/providers/Microsoft.Compute/restorePointCollections/restorePointCollection1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armcompute.RestorePointCollectionProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RestorePointCollectionID: to.Ptr("59f04a5d-f783-4200-a1bd-d3f464e8c4b4"),
		// 				Source: &armcompute.RestorePointCollectionSourceProperties{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/VM_Test"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("restorePointCollection2"),
		// 			Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resourceGroup2/providers/Microsoft.Compute/restorePointCollections/restorePointCollection2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armcompute.RestorePointCollectionProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RestorePointCollectionID: to.Ptr("2875c590-e337-4102-9668-4f5b7e3f98a4"),
		// 				Source: &armcompute.RestorePointCollectionSourceProperties{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/VM_Prod"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
