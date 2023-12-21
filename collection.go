package gnosql_client

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// return { Data : "Success Message", Error: "Error message" }, error
func (database *Database) CreateCollections(collections []CollectionInput) (CollectionCreateResult, error) {
	path := fmt.Sprintf("%s/%s/%s/add", database.URI, EndpointsMap.Collection, database.DBName)

	restyResp, restyErr := resty.New().
		R().
		SetBody(collections).
		Post(path)

	var result CollectionCreateResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	error := ValidateResponse(restyErr, UnMarshallErr)

	if error == nil {
		if result.Data == "collection created successfully" {
			for _, collection := range collections {
				collName := collection.CollectionName
				if database.Collections[collName] == nil {
					collectionInstance := Collection{
						CollectionName: collName,
						URI:            database.URI,
						DBName:         database.DBName,
					}

					database.Collections[collName] = &collectionInstance
				}

			}
		}

	}

	return result, error
}

// return { Data : "Success Message", Error: "Error message" }, error
func (database *Database) DeleteCollections(collections []string) (CollectionDeleteResult, error) {
	path := fmt.Sprintf("%s/%s/%s/delete", database.URI, EndpointsMap.Collection, database.DBName)

	requestBody := MapInterface{
		"collections": collections,
	}

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Delete(path)

	var result CollectionDeleteResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	error := ValidateResponse(restyErr, UnMarshallErr)

	if error == nil {
		if result.Data == "collection deleted successfully" {
			for _, collection := range collections {

				if database.Collections[collection] == nil {
					delete(database.Collections, collection)
				}

			}
		}

	}

	return result, error
}

// return { Data : [collection1, collection2...], Error: "Error message" }, error
func (database *Database) GetAll() (CollectionGetAllResult, error) {
	path := fmt.Sprintf("%s/%s/%s/get-all", database.URI, EndpointsMap.Collection, database.DBName)

	restyResp, restyErr := resty.New().R().Get(path)

	var result CollectionGetAllResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	return result, ValidateResponse(restyErr, UnMarshallErr)
}

// return { Data : { CollectionName string, IndexKeys []string, Documents int} , Error: "Error message" }, error
func (database *Database) GetCollectionStats(collectionName string) (CollectionStatsResult, error) {
	path := fmt.Sprintf("%s/%s/%s/%s/stats", database.URI, EndpointsMap.Collection, database.DBName, collectionName)

	restyResp, restyErr := resty.New().R().Get(path)

	var result CollectionStatsResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	return result, ValidateResponse(restyErr, UnMarshallErr)
}
