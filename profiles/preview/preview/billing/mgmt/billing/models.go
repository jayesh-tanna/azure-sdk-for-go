// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package billing

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/billing/mgmt/2018-03-01-preview/billing"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DiscoverTenant = original.DiscoverTenant
type DiscoverTenantsClient = original.DiscoverTenantsClient
type DownloadURL = original.DownloadURL
type EnrollmentAccount = original.EnrollmentAccount
type EnrollmentAccountListResult = original.EnrollmentAccountListResult
type EnrollmentAccountListResultIterator = original.EnrollmentAccountListResultIterator
type EnrollmentAccountListResultPage = original.EnrollmentAccountListResultPage
type EnrollmentAccountProperties = original.EnrollmentAccountProperties
type EnrollmentAccountsClient = original.EnrollmentAccountsClient
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Invoice = original.Invoice
type InvoiceProperties = original.InvoiceProperties
type InvoicesClient = original.InvoicesClient
type InvoicesListResult = original.InvoicesListResult
type InvoicesListResultIterator = original.InvoicesListResultIterator
type InvoicesListResultPage = original.InvoicesListResultPage
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Period = original.Period
type PeriodProperties = original.PeriodProperties
type PeriodsClient = original.PeriodsClient
type PeriodsListResult = original.PeriodsListResult
type PeriodsListResultIterator = original.PeriodsListResultIterator
type PeriodsListResultPage = original.PeriodsListResultPage
type Resource = original.Resource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDiscoverTenantsClient(subscriptionID string) DiscoverTenantsClient {
	return original.NewDiscoverTenantsClient(subscriptionID)
}
func NewDiscoverTenantsClientWithBaseURI(baseURI string, subscriptionID string) DiscoverTenantsClient {
	return original.NewDiscoverTenantsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEnrollmentAccountListResultIterator(page EnrollmentAccountListResultPage) EnrollmentAccountListResultIterator {
	return original.NewEnrollmentAccountListResultIterator(page)
}
func NewEnrollmentAccountListResultPage(getNextPage func(context.Context, EnrollmentAccountListResult) (EnrollmentAccountListResult, error)) EnrollmentAccountListResultPage {
	return original.NewEnrollmentAccountListResultPage(getNextPage)
}
func NewEnrollmentAccountsClient(subscriptionID string) EnrollmentAccountsClient {
	return original.NewEnrollmentAccountsClient(subscriptionID)
}
func NewEnrollmentAccountsClientWithBaseURI(baseURI string, subscriptionID string) EnrollmentAccountsClient {
	return original.NewEnrollmentAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoicesClient(subscriptionID string) InvoicesClient {
	return original.NewInvoicesClient(subscriptionID)
}
func NewInvoicesClientWithBaseURI(baseURI string, subscriptionID string) InvoicesClient {
	return original.NewInvoicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoicesListResultIterator(page InvoicesListResultPage) InvoicesListResultIterator {
	return original.NewInvoicesListResultIterator(page)
}
func NewInvoicesListResultPage(getNextPage func(context.Context, InvoicesListResult) (InvoicesListResult, error)) InvoicesListResultPage {
	return original.NewInvoicesListResultPage(getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPeriodsClient(subscriptionID string) PeriodsClient {
	return original.NewPeriodsClient(subscriptionID)
}
func NewPeriodsClientWithBaseURI(baseURI string, subscriptionID string) PeriodsClient {
	return original.NewPeriodsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPeriodsListResultIterator(page PeriodsListResultPage) PeriodsListResultIterator {
	return original.NewPeriodsListResultIterator(page)
}
func NewPeriodsListResultPage(getNextPage func(context.Context, PeriodsListResult) (PeriodsListResult, error)) PeriodsListResultPage {
	return original.NewPeriodsListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
