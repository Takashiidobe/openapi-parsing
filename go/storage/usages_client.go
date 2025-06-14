// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.7, generator: @autorest/go@4.0.0-preview.72)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package storage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use a constructor function instead.
type UsagesClient struct {
	internal *azcore.Client
	subscriptionID string
}

// NewListByLocationPager - Gets the current usage count and the limit for the resources of the location under the subscription.
//
// Generated from API version 2024-01-01
//   - location - The location of the Azure Storage resource.
//   - options - UsagesClientListByLocationOptions contains the optional parameters for the UsagesClient.NewListByLocationPager
//     method.
func (client *UsagesClient) NewListByLocationPager(location string, options *UsagesClientListByLocationOptions) (*runtime.Pager[UsagesClientListByLocationResponse]) {
	return runtime.NewPager(runtime.PagingHandler[UsagesClientListByLocationResponse]{
		More: func(page UsagesClientListByLocationResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *UsagesClientListByLocationResponse) (UsagesClientListByLocationResponse, error) {
			req, err := client.listByLocationCreateRequest(ctx, location, options)
			if err != nil {
				return UsagesClientListByLocationResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return UsagesClientListByLocationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UsagesClientListByLocationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLocationHandleResponse(resp)
		},
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *UsagesClient) listByLocationCreateRequest(ctx context.Context, location string, _ *UsagesClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/locations/{location}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *UsagesClient) listByLocationHandleResponse(resp *http.Response) (UsagesClientListByLocationResponse, error) {
	result := UsagesClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageListResult); err != nil {
		return UsagesClientListByLocationResponse{}, err
	}
	return result, nil
}

