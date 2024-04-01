# Manticorer Go client

Сlient for Manticore Search.

❗ WARNING: this is a development version of the client. The latest release's readme is https://github.com/manticoresoftware/manticoresearch-go/tree/v1.0.0


## Requiments

| Manticore Search  | manticoresearch-go           |     Go        |
| ----------------- | ---------------------------- | ------------- |
| dev               | dev                          | >= 1.17.0     |
| >= 6.2.12         | v1.0.0                       | >= 1.17.0     |

## Installation

```shell

go get github.com/manticoresoftware/manticoresearch-go@dev 

```

## Getting Started

```go

package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	openapiclient "github.com/manticoresoftware/manticoresearch-go"
)

func main() {

  // Initialize ApiClient
  configuration := openapiclient.NewConfiguration()
  // Set Manticore's host/port if necessary (http://127.0.0.1:9308 by default) 
  // configuration.Servers[0].URL = "http://127.0.0.2:9408"
  apiClient := openapiclient.NewAPIClient(configuration)
  
  // Add documents to an index 
  docs := []string {
	"{\"insert\": {\"index\" : \"test\", \"id\" : 1, \"doc\" : {\"title\" : \"Title 1\"}}}",
	"{\"insert\": {\"index\" : \"test\", \"id\" : 2, \"doc\" : {\"title\" : \"Title 2\"}}}",
  }
  apiClient.IndexAPI.Bulk(context.Background()).Body(strings.Join(docs[:], "\n")).Execute()
  
  // response from `Search`: SearchRequest
  searchRequest := *openapiclient.NewSearchRequest("test")
  
  // Perform a search
  query := map[string]interface{} {"query_string": "Title"}
  searchRequest.SetQuery(query)
  resp, r, err := apiClient.SearchAPI.Search(context.Background()).SearchRequest(searchRequest).Execute()
  
  if err != nil {
	fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search``: %v\n", err)
	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
  }
  // response from `Search`: SearchResponse
  fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Search`: %v\n", resp)
}


```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

For example, if your Manticore server is available through the '/manticore' path, you should set `basePath` appropriately. This will rewrite all your requests: e.g., '/search' to '/manticore/search', etc.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "manticore",
})
```



Note, enum values are always validated and all unused variables are silently ignored.

### Using proxy

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

More details on the use of `HTTP_PROXY` can be found [here](https://www.cyberciti.biz/faq/linux-unix-set-proxy-environment-variable/)

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:9308*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndexAPI* | [**Bulk**](docs/IndexAPI.md#bulk) | **Post** /bulk | Bulk index operations
*IndexAPI* | [**Delete**](docs/IndexAPI.md#delete) | **Post** /delete | Delete a document in an index
*IndexAPI* | [**Insert**](docs/IndexAPI.md#insert) | **Post** /insert | Create a new document in an index
*IndexAPI* | [**Replace**](docs/IndexAPI.md#replace) | **Post** /replace | Replace new document in an index
*IndexAPI* | [**Update**](docs/IndexAPI.md#update) | **Post** /update | Update a document in an index
*IndexAPI* | [**Update_0**](docs/IndexAPI.md#update_0) | **Post** /{index}/_update/{id} | Partially replaces a document in an index
*SearchAPI* | [**Percolate**](docs/SearchAPI.md#percolate) | **Post** /pq/{index}/search | Perform reverse search on a percolate index
*SearchAPI* | [**Search**](docs/SearchAPI.md#search) | **Post** /search | Performs a search on an index
*UtilsAPI* | [**Sql**](docs/UtilsAPI.md#sql) | **Post** /sql | Perform SQL requests


## Documentation For Models

 - [Aggregation](docs/Aggregation.md)
 - [AggregationSortInnerValue](docs/AggregationSortInnerValue.md)
 - [AggregationTerms](docs/AggregationTerms.md)
 - [BoolFilter](docs/BoolFilter.md)
 - [BulkResponse](docs/BulkResponse.md)
 - [DeleteDocumentRequest](docs/DeleteDocumentRequest.md)
 - [DeleteResponse](docs/DeleteResponse.md)
 - [EqualsFilter](docs/EqualsFilter.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorResponseString](docs/ErrorResponseString.md)
 - [Facet](docs/Facet.md)
 - [FilterBoolean](docs/FilterBoolean.md)
 - [FilterNumber](docs/FilterNumber.md)
 - [FilterString](docs/FilterString.md)
 - [GeoDistanceFilter](docs/GeoDistanceFilter.md)
 - [GeoDistanceFilterLocationAnchor](docs/GeoDistanceFilterLocationAnchor.md)
 - [Highlight](docs/Highlight.md)
 - [HighlightField](docs/HighlightField.md)
 - [InFilter](docs/InFilter.md)
 - [InsertDocumentRequest](docs/InsertDocumentRequest.md)
 - [KnnSearchRequestByDocId](docs/KnnSearchRequestByDocId.md)
 - [KnnSearchRequestByVector](docs/KnnSearchRequestByVector.md)
 - [MatchFilter](docs/MatchFilter.md)
 - [MatchOp](docs/MatchOp.md)
 - [MatchOpFilter](docs/MatchOpFilter.md)
 - [MatchPhraseFilter](docs/MatchPhraseFilter.md)
 - [NotFilterBoolean](docs/NotFilterBoolean.md)
 - [NotFilterNumber](docs/NotFilterNumber.md)
 - [NotFilterString](docs/NotFilterString.md)
 - [PercolateRequest](docs/PercolateRequest.md)
 - [PercolateRequestQuery](docs/PercolateRequestQuery.md)
 - [QueryFilter](docs/QueryFilter.md)
 - [RangeFilter](docs/RangeFilter.md)
 - [ReplaceDocumentRequest](docs/ReplaceDocumentRequest.md)
 - [SearchRequest](docs/SearchRequest.md)
 - [SearchResponse](docs/SearchResponse.md)
 - [SearchResponseHits](docs/SearchResponseHits.md)
 - [SortMVA](docs/SortMVA.md)
 - [SortMultiple](docs/SortMultiple.md)
 - [SortOrder](docs/SortOrder.md)
 - [SourceByRules](docs/SourceByRules.md)
 - [SqlDefaultResponse](docs/SqlDefaultResponse.md)
 - [SuccessResponse](docs/SuccessResponse.md)
 - [UpdateDocumentRequest](docs/UpdateDocumentRequest.md)
 - [UpdateResponse](docs/UpdateResponse.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
