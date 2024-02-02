# \UtilsAPI

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Sql**](UtilsAPI.md#Sql) | **Post** /sql | Perform SQL requests



## Sql

> []map[string]interface{} Sql(ctx).Body(body).RawResponse(rawResponse).Execute()

Perform SQL requests



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
	body := "SHOW TABLES" // string | A query parameter string. 
	rawResponse := true // bool | Optional parameter, defines a format of response. Can be set to `False` for Select only queries and set to `True` or omitted for any type of queries:  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.Sql(context.Background()).Body(body).RawResponse(rawResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.Sql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Sql`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.Sql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | A query parameter string.  | 
 **rawResponse** | **bool** | Optional parameter, defines a format of response. Can be set to &#x60;False&#x60; for Select only queries and set to &#x60;True&#x60; or omitted for any type of queries:  | [default to true]

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

