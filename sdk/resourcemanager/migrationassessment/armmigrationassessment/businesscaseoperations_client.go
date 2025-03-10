// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationassessment

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BusinessCaseOperationsClient contains the methods for the BusinessCaseOperations group.
// Don't use this type directly, use NewBusinessCaseOperationsClient() instead.
type BusinessCaseOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBusinessCaseOperationsClient creates a new instance of BusinessCaseOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBusinessCaseOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BusinessCaseOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BusinessCaseOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCompareSummary - A long-running resource action.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - businessCaseName - Business case ARM name
//   - body - The content of the action request
//   - options - BusinessCaseOperationsClientBeginCompareSummaryOptions contains the optional parameters for the BusinessCaseOperationsClient.BeginCompareSummary
//     method.
func (client *BusinessCaseOperationsClient) BeginCompareSummary(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, body any, options *BusinessCaseOperationsClientBeginCompareSummaryOptions) (*runtime.Poller[BusinessCaseOperationsClientCompareSummaryResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.compareSummary(ctx, resourceGroupName, projectName, businessCaseName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BusinessCaseOperationsClientCompareSummaryResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BusinessCaseOperationsClientCompareSummaryResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CompareSummary - A long-running resource action.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *BusinessCaseOperationsClient) compareSummary(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, body any, options *BusinessCaseOperationsClientBeginCompareSummaryOptions) (*http.Response, error) {
	var err error
	const operationName = "BusinessCaseOperationsClient.BeginCompareSummary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.compareSummaryCreateRequest(ctx, resourceGroupName, projectName, businessCaseName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// compareSummaryCreateRequest creates the CompareSummary request.
func (client *BusinessCaseOperationsClient) compareSummaryCreateRequest(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, body any, _ *BusinessCaseOperationsClientBeginCompareSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases/{businessCaseName}/compareSummary"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if businessCaseName == "" {
		return nil, errors.New("parameter businessCaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessCaseName}", url.PathEscape(businessCaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreate - Create a BusinessCase
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - businessCaseName - Business case ARM name
//   - resource - Resource create parameters.
//   - options - BusinessCaseOperationsClientBeginCreateOptions contains the optional parameters for the BusinessCaseOperationsClient.BeginCreate
//     method.
func (client *BusinessCaseOperationsClient) BeginCreate(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, resource BusinessCase, options *BusinessCaseOperationsClientBeginCreateOptions) (*runtime.Poller[BusinessCaseOperationsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, projectName, businessCaseName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BusinessCaseOperationsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BusinessCaseOperationsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a BusinessCase
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *BusinessCaseOperationsClient) create(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, resource BusinessCase, options *BusinessCaseOperationsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "BusinessCaseOperationsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, projectName, businessCaseName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *BusinessCaseOperationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, resource BusinessCase, _ *BusinessCaseOperationsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases/{businessCaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if businessCaseName == "" {
		return nil, errors.New("parameter businessCaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessCaseName}", url.PathEscape(businessCaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a BusinessCase
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - businessCaseName - Business case ARM name
//   - options - BusinessCaseOperationsClientDeleteOptions contains the optional parameters for the BusinessCaseOperationsClient.Delete
//     method.
func (client *BusinessCaseOperationsClient) Delete(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, options *BusinessCaseOperationsClientDeleteOptions) (BusinessCaseOperationsClientDeleteResponse, error) {
	var err error
	const operationName = "BusinessCaseOperationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, businessCaseName, options)
	if err != nil {
		return BusinessCaseOperationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BusinessCaseOperationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BusinessCaseOperationsClientDeleteResponse{}, err
	}
	return BusinessCaseOperationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BusinessCaseOperationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, _ *BusinessCaseOperationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases/{businessCaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if businessCaseName == "" {
		return nil, errors.New("parameter businessCaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessCaseName}", url.PathEscape(businessCaseName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a BusinessCase
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - businessCaseName - Business case ARM name
//   - options - BusinessCaseOperationsClientGetOptions contains the optional parameters for the BusinessCaseOperationsClient.Get
//     method.
func (client *BusinessCaseOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, options *BusinessCaseOperationsClientGetOptions) (BusinessCaseOperationsClientGetResponse, error) {
	var err error
	const operationName = "BusinessCaseOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, businessCaseName, options)
	if err != nil {
		return BusinessCaseOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BusinessCaseOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BusinessCaseOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BusinessCaseOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, _ *BusinessCaseOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases/{businessCaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if businessCaseName == "" {
		return nil, errors.New("parameter businessCaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessCaseName}", url.PathEscape(businessCaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BusinessCaseOperationsClient) getHandleResponse(resp *http.Response) (BusinessCaseOperationsClientGetResponse, error) {
	result := BusinessCaseOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BusinessCase); err != nil {
		return BusinessCaseOperationsClientGetResponse{}, err
	}
	return result, nil
}

// BeginGetReportDownloadURL - Get the URL for downloading the business case in a report format.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - businessCaseName - Business case ARM name
//   - body - The content of the action request
//   - options - BusinessCaseOperationsClientBeginGetReportDownloadURLOptions contains the optional parameters for the BusinessCaseOperationsClient.BeginGetReportDownloadURL
//     method.
func (client *BusinessCaseOperationsClient) BeginGetReportDownloadURL(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, body any, options *BusinessCaseOperationsClientBeginGetReportDownloadURLOptions) (*runtime.Poller[BusinessCaseOperationsClientGetReportDownloadURLResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.getReportDownloadURL(ctx, resourceGroupName, projectName, businessCaseName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BusinessCaseOperationsClientGetReportDownloadURLResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BusinessCaseOperationsClientGetReportDownloadURLResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// GetReportDownloadURL - Get the URL for downloading the business case in a report format.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *BusinessCaseOperationsClient) getReportDownloadURL(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, body any, options *BusinessCaseOperationsClientBeginGetReportDownloadURLOptions) (*http.Response, error) {
	var err error
	const operationName = "BusinessCaseOperationsClient.BeginGetReportDownloadURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getReportDownloadURLCreateRequest(ctx, resourceGroupName, projectName, businessCaseName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// getReportDownloadURLCreateRequest creates the GetReportDownloadURL request.
func (client *BusinessCaseOperationsClient) getReportDownloadURLCreateRequest(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, body any, _ *BusinessCaseOperationsClientBeginGetReportDownloadURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases/{businessCaseName}/getReportDownloadUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if businessCaseName == "" {
		return nil, errors.New("parameter businessCaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessCaseName}", url.PathEscape(businessCaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// NewListByAssessmentProjectPager - List BusinessCase resources by AssessmentProject
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - options - BusinessCaseOperationsClientListByAssessmentProjectOptions contains the optional parameters for the BusinessCaseOperationsClient.NewListByAssessmentProjectPager
//     method.
func (client *BusinessCaseOperationsClient) NewListByAssessmentProjectPager(resourceGroupName string, projectName string, options *BusinessCaseOperationsClientListByAssessmentProjectOptions) *runtime.Pager[BusinessCaseOperationsClientListByAssessmentProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[BusinessCaseOperationsClientListByAssessmentProjectResponse]{
		More: func(page BusinessCaseOperationsClientListByAssessmentProjectResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BusinessCaseOperationsClientListByAssessmentProjectResponse) (BusinessCaseOperationsClientListByAssessmentProjectResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BusinessCaseOperationsClient.NewListByAssessmentProjectPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAssessmentProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			}, nil)
			if err != nil {
				return BusinessCaseOperationsClientListByAssessmentProjectResponse{}, err
			}
			return client.listByAssessmentProjectHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAssessmentProjectCreateRequest creates the ListByAssessmentProject request.
func (client *BusinessCaseOperationsClient) listByAssessmentProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, _ *BusinessCaseOperationsClientListByAssessmentProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAssessmentProjectHandleResponse handles the ListByAssessmentProject response.
func (client *BusinessCaseOperationsClient) listByAssessmentProjectHandleResponse(resp *http.Response) (BusinessCaseOperationsClientListByAssessmentProjectResponse, error) {
	result := BusinessCaseOperationsClientListByAssessmentProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BusinessCaseListResult); err != nil {
		return BusinessCaseOperationsClientListByAssessmentProjectResponse{}, err
	}
	return result, nil
}
