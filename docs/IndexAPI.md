# \IndexAPI

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Bulk**](IndexAPI.md#Bulk) | **Post** /bulk | Bulk index operations
[**Delete**](IndexAPI.md#Delete) | **Post** /delete | Delete a document in an index
[**Insert**](IndexAPI.md#Insert) | **Post** /insert | Create a new document in an index
[**PartialReplace**](IndexAPI.md#PartialReplace) | **Post** /{index}/_update/{id} | Partially replaces a document in an index
[**Replace**](IndexAPI.md#Replace) | **Post** /replace | Replace new document in an index
[**Update**](IndexAPI.md#Update) | **Post** /update | Update a document in an index



## Bulk

> BulkResponse Bulk(ctx).Body(body).Execute()

Bulk index operations



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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.Bulk(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.Bulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Bulk`: BulkResponse
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.Bulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-ndjson
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> DeleteResponse Delete(ctx).DeleteDocumentRequest(deleteDocumentRequest).Execute()

Delete a document in an index



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
	deleteDocumentRequest := *openapiclient.NewDeleteDocumentRequest("Index_example") // DeleteDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.Delete(context.Background()).DeleteDocumentRequest(deleteDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Delete`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.Delete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDocumentRequest** | [**DeleteDocumentRequest**](DeleteDocumentRequest.md) |  | 

### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Insert

> SuccessResponse Insert(ctx).InsertDocumentRequest(insertDocumentRequest).Execute()

Create a new document in an index



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
	insertDocumentRequest := *openapiclient.NewInsertDocumentRequest("Index_example", map[string]interface{}(123)) // InsertDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(insertDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.Insert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Insert`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.Insert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertDocumentRequest** | [**InsertDocumentRequest**](InsertDocumentRequest.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialReplace

> UpdateResponse PartialReplace(ctx, index, id).ReplaceDocumentRequest(replaceDocumentRequest).Execute()

Partially replaces a document in an index



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
	id := float32(8.14) // float32 | Id of the document to replace
	replaceDocumentRequest := *openapiclient.NewReplaceDocumentRequest(map[string]interface{}(123)) // ReplaceDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.PartialReplace(context.Background(), index, id).ReplaceDocumentRequest(replaceDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PartialReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialReplace`: UpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.PartialReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Name of the percolate index | 
**id** | **float32** | Id of the document to replace | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceDocumentRequest** | [**ReplaceDocumentRequest**](ReplaceDocumentRequest.md) |  | 

### Return type

[**UpdateResponse**](UpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Replace

> SuccessResponse Replace(ctx).InsertDocumentRequest(insertDocumentRequest).Execute()

Replace new document in an index



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
	insertDocumentRequest := *openapiclient.NewInsertDocumentRequest("Index_example", map[string]interface{}(123)) // InsertDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.Replace(context.Background()).InsertDocumentRequest(insertDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.Replace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Replace`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.Replace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertDocumentRequest** | [**InsertDocumentRequest**](InsertDocumentRequest.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateResponse Update(ctx).UpdateDocumentRequest(updateDocumentRequest).Execute()

Update a document in an index



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
	updateDocumentRequest := *openapiclient.NewUpdateDocumentRequest("Index_example", map[string]interface{}({gid=10})) // UpdateDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.Update(context.Background()).UpdateDocumentRequest(updateDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDocumentRequest** | [**UpdateDocumentRequest**](UpdateDocumentRequest.md) |  | 

### Return type

[**UpdateResponse**](UpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

