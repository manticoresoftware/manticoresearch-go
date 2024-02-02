# \SearchAPI

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Percolate**](SearchAPI.md#Percolate) | **Post** /pq/{index}/search | Perform reverse search on a percolate index
[**Search**](SearchAPI.md#Search) | **Post** /search | Performs a search on an index



## Percolate

> SearchResponse Percolate(ctx, index).PercolateRequest(percolateRequest).Execute()

Perform reverse search on a percolate index



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/manticoresoftware/manticoresearch-go"
)

func main() {
	index := "index_example" // string | Name of the percolate index
	percolateRequest := *openapiclient.NewPercolateRequest(*openapiclient.NewPercolateRequestQuery(map[string]interface{}(123))) // PercolateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.Percolate(context.Background(), index).PercolateRequest(percolateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Percolate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Percolate`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Percolate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Name of the percolate index | 

### Other Parameters

Other parameters are passed through a pointer to a apiPercolateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **percolateRequest** | [**PercolateRequest**](PercolateRequest.md) |  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SearchResponse Search(ctx).SearchRequest(searchRequest).Execute()

Performs a search on an index



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/manticoresoftware/manticoresearch-go"
)

func main() {
	searchRequest := *openapiclient.NewSearchRequest("test") // SearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.Search(context.Background()).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md) |  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

