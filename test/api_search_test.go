/*
Manticore Search Client

Testing SearchAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"fmt"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"encoding/json"
	"strings"
	Manticoresearch "github.com/manticoresoftware/manticoresearch-go"
)

func Test_openapi_SearchAPIService(t *testing.T) {

	configuration := Manticoresearch.NewConfiguration()
	configuration.Servers[0].URL = "http://localhost:9408"
	apiClient := Manticoresearch.NewAPIClient(configuration)

	t.Run("Test SearchAPIService Search", func(t *testing.T) {

		docs := []string {
 			`{"insert": {"index" : "movies", "id" : 1, "doc" : {"title" : "Star Trek 2: Nemesis", "plot": "The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.", "year": 2002, "rating": 6.4, "code": [1,2,3]}}}`,
			`{"insert": {"index" : "movies", "id" : 2, "doc" : {"title" : "Star Trek 1: Nemesis", "plot": "The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.", "year": 2001, "rating": 6.5, "code": [1,12,3]}}}`,
			`{"insert": {"index" : "movies", "id" : 3, "doc" : {"title" : "Star Trek 3: Nemesis", "plot": "The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.", "year": 2003, "rating": 6.6, "code": [11,2,3]}}}`,
			`{"insert": {"index" : "movies", "id" : 4, "doc" : {"title" : "Star Trek 4: Nemesis", "plot": "The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.", "year": 100000000000, "rating": 6.5, "code": [1,2,4]}}}`,
		};
 		apiClient.UtilsAPI.Sql(context.Background()).Body("DROP TABLE IF EXISTS movies").Execute()
		apiClient.UtilsAPI.Sql(context.Background()).Body("CREATE TABLE IF NOT EXISTS movies (title text, plot text, year bigint, rating float, code multi)").Execute()
		body := strings.Join(docs, "\n")
		bulkResp, bulkHttpRes, bulkErr := apiClient.IndexAPI.Bulk(context.Background()).Body(body).Execute()
   		fmt.Printf("%+v\n\n", bulkHttpRes)
   		require.Nil(t, bulkErr)
		require.NotNil(t, bulkResp)
		outRes, outErr := json.Marshal(bulkResp)
		require.Nil(t, outErr)
		fmt.Printf("%+v\n\n", string(outRes[:]))
    	assert.Equal(t, 200, bulkHttpRes.StatusCode)

		queryHighlight := Manticoresearch.NewHighlight()
 		queryHighlight.Fields =  map[string]interface{} {"title": map[string]interface{} {}}

		searchQuery := Manticoresearch.NewSearchQuery()
		searchQuery.QueryString = "@title Trek 4"
		searchRequest := Manticoresearch.NewSearchRequest("movies")
		searchRequest.Highlight = queryHighlight
		searchRequest.Query = searchQuery
		resp, httpRes, err := apiClient.SearchAPI.Search(context.Background()).SearchRequest(*searchRequest).Execute()
		fmt.Printf("%+v\n\n", httpRes)
		require.Nil(t, err)
		require.NotNil(t, resp)
		outRes, outErr = json.Marshal(resp)
		require.Nil(t, outErr)
		fmt.Printf("%+v\n\n", string(outRes[:]))
		assert.Equal(t, 200, httpRes.StatusCode)

		fmt.Println("Search tests finished");
	})

}
