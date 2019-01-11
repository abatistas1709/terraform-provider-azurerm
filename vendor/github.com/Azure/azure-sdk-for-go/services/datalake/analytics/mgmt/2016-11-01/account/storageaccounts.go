package account

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// StorageAccountsClient is the creates an Azure Data Lake Analytics account management client.
type StorageAccountsClient struct {
	BaseClient
}

// NewStorageAccountsClient creates an instance of the StorageAccountsClient client.
func NewStorageAccountsClient(subscriptionID string) StorageAccountsClient {
	return NewStorageAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStorageAccountsClientWithBaseURI creates an instance of the StorageAccountsClient client.
func NewStorageAccountsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountsClient {
	return StorageAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Add updates the specified Data Lake Analytics account to add an Azure Storage account.
// Parameters:
// resourceGroupName - the name of the Azure resource group.
// accountName - the name of the Data Lake Analytics account.
// storageAccountName - the name of the Azure Storage account to add
// parameters - the parameters containing the access key and optional suffix for the Azure Storage Account.
func (client StorageAccountsClient) Add(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters AddStorageAccountParameters) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.Add")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.AddStorageAccountProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.AddStorageAccountProperties.AccessKey", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("account.StorageAccountsClient", "Add", err.Error())
	}

	req, err := client.AddPreparer(ctx, resourceGroupName, accountName, storageAccountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Add", resp, "Failure responding to request")
	}

	return
}

// AddPreparer prepares the Add request.
func (client StorageAccountsClient) AddPreparer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters AddStorageAccountParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) AddSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) AddResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete updates the specified Data Lake Analytics account to remove an Azure Storage account.
// Parameters:
// resourceGroupName - the name of the Azure resource group.
// accountName - the name of the Data Lake Analytics account.
// storageAccountName - the name of the Azure Storage account to remove
func (client StorageAccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, storageAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StorageAccountsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified Azure Storage account linked to the given Data Lake Analytics account.
// Parameters:
// resourceGroupName - the name of the Azure resource group.
// accountName - the name of the Data Lake Analytics account.
// storageAccountName - the name of the Azure Storage account for which to retrieve the details.
func (client StorageAccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result StorageAccountInformation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, storageAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client StorageAccountsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) GetResponder(resp *http.Response) (result StorageAccountInformation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetStorageContainer gets the specified Azure Storage container associated with the given Data Lake Analytics and
// Azure Storage accounts.
// Parameters:
// resourceGroupName - the name of the Azure resource group.
// accountName - the name of the Data Lake Analytics account.
// storageAccountName - the name of the Azure storage account from which to retrieve the blob container.
// containerName - the name of the Azure storage container to retrieve
func (client StorageAccountsClient) GetStorageContainer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result StorageContainer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.GetStorageContainer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetStorageContainerPreparer(ctx, resourceGroupName, accountName, storageAccountName, containerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "GetStorageContainer", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStorageContainerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "GetStorageContainer", resp, "Failure sending request")
		return
	}

	result, err = client.GetStorageContainerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "GetStorageContainer", resp, "Failure responding to request")
	}

	return
}

// GetStorageContainerPreparer prepares the GetStorageContainer request.
func (client StorageAccountsClient) GetStorageContainerPreparer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"containerName":      autorest.Encode("path", containerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetStorageContainerSender sends the GetStorageContainer request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) GetStorageContainerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetStorageContainerResponder handles the response to the GetStorageContainer request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) GetStorageContainerResponder(resp *http.Response) (result StorageContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAccount gets the first page of Azure Storage accounts, if any, linked to the specified Data Lake Analytics
// account. The response includes a link to the next page, if any.
// Parameters:
// resourceGroupName - the name of the Azure resource group.
// accountName - the name of the Data Lake Analytics account.
// filter - the OData filter. Optional.
// top - the number of items to return. Optional.
// skip - the number of items to skip over before returning elements. Optional.
// selectParameter - oData Select statement. Limits the properties on each entry to just those requested, e.g.
// Categories?$select=CategoryName,Description. Optional.
// orderby - orderBy clause. One or more comma-separated expressions with an optional "asc" (the default) or
// "desc" depending on the order you'd like the values sorted, e.g. Categories?$orderby=CategoryName desc.
// Optional.
// count - the Boolean value of true or false to request a count of the matching resources included with the
// resources in the response, e.g. Categories?$count=true. Optional.
func (client StorageAccountsClient) ListByAccount(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result StorageAccountInformationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.ListByAccount")
		defer func() {
			sc := -1
			if result.sailr.Response.Response != nil {
				sc = result.sailr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("account.StorageAccountsClient", "ListByAccount", err.Error())
	}

	result.fn = client.listByAccountNextResults
	req, err := client.ListByAccountPreparer(ctx, resourceGroupName, accountName, filter, top, skip, selectParameter, orderby, count)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "ListByAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.sailr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "ListByAccount", resp, "Failure sending request")
		return
	}

	result.sailr, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "ListByAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAccountPreparer prepares the ListByAccount request.
func (client StorageAccountsClient) ListByAccountPreparer(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if count != nil {
		queryParameters["$count"] = autorest.Encode("query", *count)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAccountSender sends the ListByAccount request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) ListByAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByAccountResponder handles the response to the ListByAccount request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) ListByAccountResponder(resp *http.Response) (result StorageAccountInformationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAccountNextResults retrieves the next set of results, if any.
func (client StorageAccountsClient) listByAccountNextResults(ctx context.Context, lastResults StorageAccountInformationListResult) (result StorageAccountInformationListResult, err error) {
	req, err := lastResults.storageAccountInformationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.StorageAccountsClient", "listByAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.StorageAccountsClient", "listByAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "listByAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client StorageAccountsClient) ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result StorageAccountInformationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.ListByAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAccount(ctx, resourceGroupName, accountName, filter, top, skip, selectParameter, orderby, count)
	return
}

// ListSasTokens gets the SAS token associated with the specified Data Lake Analytics and Azure Storage account and
// container combination.
// Parameters:
// resourceGroupName - the name of the Azure resource group.
// accountName - the name of the Data Lake Analytics account.
// storageAccountName - the name of the Azure storage account for which the SAS token is being requested.
// containerName - the name of the Azure storage container for which the SAS token is being requested.
func (client StorageAccountsClient) ListSasTokens(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result SasTokenInformationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.ListSasTokens")
		defer func() {
			sc := -1
			if result.stilr.Response.Response != nil {
				sc = result.stilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listSasTokensNextResults
	req, err := client.ListSasTokensPreparer(ctx, resourceGroupName, accountName, storageAccountName, containerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "ListSasTokens", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSasTokensSender(req)
	if err != nil {
		result.stilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "ListSasTokens", resp, "Failure sending request")
		return
	}

	result.stilr, err = client.ListSasTokensResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "ListSasTokens", resp, "Failure responding to request")
	}

	return
}

// ListSasTokensPreparer prepares the ListSasTokens request.
func (client StorageAccountsClient) ListSasTokensPreparer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"containerName":      autorest.Encode("path", containerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}/listSasTokens", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSasTokensSender sends the ListSasTokens request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) ListSasTokensSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListSasTokensResponder handles the response to the ListSasTokens request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) ListSasTokensResponder(resp *http.Response) (result SasTokenInformationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listSasTokensNextResults retrieves the next set of results, if any.
func (client StorageAccountsClient) listSasTokensNextResults(ctx context.Context, lastResults SasTokenInformationListResult) (result SasTokenInformationListResult, err error) {
	req, err := lastResults.sasTokenInformationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.StorageAccountsClient", "listSasTokensNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSasTokensSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.StorageAccountsClient", "listSasTokensNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListSasTokensResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "listSasTokensNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListSasTokensComplete enumerates all values, automatically crossing page boundaries as required.
func (client StorageAccountsClient) ListSasTokensComplete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result SasTokenInformationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.ListSasTokens")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListSasTokens(ctx, resourceGroupName, accountName, storageAccountName, containerName)
	return
}

// ListStorageContainers lists the Azure Storage containers, if any, associated with the specified Data Lake Analytics
// and Azure Storage account combination. The response includes a link to the next page of results, if any.
// Parameters:
// resourceGroupName - the name of the Azure resource group.
// accountName - the name of the Data Lake Analytics account.
// storageAccountName - the name of the Azure storage account from which to list blob containers.
func (client StorageAccountsClient) ListStorageContainers(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result StorageContainerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.ListStorageContainers")
		defer func() {
			sc := -1
			if result.sclr.Response.Response != nil {
				sc = result.sclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listStorageContainersNextResults
	req, err := client.ListStorageContainersPreparer(ctx, resourceGroupName, accountName, storageAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "ListStorageContainers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListStorageContainersSender(req)
	if err != nil {
		result.sclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "ListStorageContainers", resp, "Failure sending request")
		return
	}

	result.sclr, err = client.ListStorageContainersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "ListStorageContainers", resp, "Failure responding to request")
	}

	return
}

// ListStorageContainersPreparer prepares the ListStorageContainers request.
func (client StorageAccountsClient) ListStorageContainersPreparer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListStorageContainersSender sends the ListStorageContainers request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) ListStorageContainersSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListStorageContainersResponder handles the response to the ListStorageContainers request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) ListStorageContainersResponder(resp *http.Response) (result StorageContainerListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listStorageContainersNextResults retrieves the next set of results, if any.
func (client StorageAccountsClient) listStorageContainersNextResults(ctx context.Context, lastResults StorageContainerListResult) (result StorageContainerListResult, err error) {
	req, err := lastResults.storageContainerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.StorageAccountsClient", "listStorageContainersNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListStorageContainersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.StorageAccountsClient", "listStorageContainersNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListStorageContainersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "listStorageContainersNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListStorageContainersComplete enumerates all values, automatically crossing page boundaries as required.
func (client StorageAccountsClient) ListStorageContainersComplete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result StorageContainerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.ListStorageContainers")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListStorageContainers(ctx, resourceGroupName, accountName, storageAccountName)
	return
}

// Update updates the Data Lake Analytics account to replace Azure Storage blob account details, such as the access key
// and/or suffix.
// Parameters:
// resourceGroupName - the name of the Azure resource group.
// accountName - the name of the Data Lake Analytics account.
// storageAccountName - the Azure Storage account to modify
// parameters - the parameters containing the access key and suffix to update the storage account with, if any.
// Passing nothing results in no change.
func (client StorageAccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters *UpdateStorageAccountParameters) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, storageAccountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.StorageAccountsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client StorageAccountsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters *UpdateStorageAccountParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
