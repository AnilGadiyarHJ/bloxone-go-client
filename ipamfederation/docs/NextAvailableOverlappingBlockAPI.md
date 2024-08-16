# NextAvailableOverlappingBlockAPI

All URIs are relative to */api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNextAvailableOverlappingBlocks**](NextAvailableOverlappingBlockAPI.md#ListNextAvailableOverlappingBlocks) | **Get** /federation/federated_block/{id}/next_available_overlapping_block | List the next available overlapping block.



## ListNextAvailableOverlappingBlocks

> FederationListNextAvailableOverlappingBlockResponse ListNextAvailableOverlappingBlocks(ctx, id).Cidr(cidr).Count(count).Name(name).Comment(comment).Execute()

List the next available overlapping block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableOverlappingBlockAPI.ListNextAvailableOverlappingBlocks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableOverlappingBlockAPI.ListNextAvailableOverlappingBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNextAvailableOverlappingBlocks`: FederationListNextAvailableOverlappingBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableOverlappingBlockAPI.ListNextAvailableOverlappingBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableOverlappingBlockAPIListNextAvailableOverlappingBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**cidr** | **int64** | The CIDR of the federated block. This is required, if _address_ does not specify it in its input. | 
**count** | **int64** | The count of __Block__ required. If not provided, it will default to 1. | 
**name** | **string** | The name to be provided. | 
**comment** | **string** | The description for the _federation/federated_block_. May contain 0 to 1024 characters. Can include UTF-8. | 

### Return type

[**FederationListNextAvailableOverlappingBlockResponse**](FederationListNextAvailableOverlappingBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
