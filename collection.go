package gnosql_client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	pb "github.com/nanda03dev/gnosql_client/proto"
)

// return { Data : "Success Message", Error: "Error message" }, error
func (database *Database) CreateCollections(collections []CollectionInput) (CollectionCreateResult, error) {
	var result CollectionCreateResult
	var error error = nil

	colls := make([]*pb.CollectionInput, 0)

	for _, coll := range collections {
		coll1 := &pb.CollectionInput{
			CollectionName: coll.CollectionName,
			IndexKeys:      coll.IndexKeys,
		}
		colls = append(colls, coll1)
	}

	requestBody := &pb.CollectionCreateRequest{
		DatabaseName: database.DBName,
		Collections:  colls,
	}

	if database.IsgRPC {
		gRPC := database.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.CreateNewCollection(ctx, requestBody)

		result.Data = res.GetData()
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s/%s/add", database.URI, EndpointsMap.Collection, database.DBName)

		restyResp, restyErr := resty.New().
			R().
			SetBody(collections).
			Post(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)

	}

	if error == nil {
		for _, collection := range collections {
			collName := collection.CollectionName

			if database.Collections[collName] == nil {
				collectionInstance := &Collection{
					CollectionName: collName,
					URI:            database.URI,
					DBName:         database.DBName,
					IsgRPC:         database.IsgRPC,
					ClientgRPC:     database.ClientgRPC,
				}

				database.Collections[collName] = collectionInstance
			}
		}
	}

	return result, nil
}

// return { Data : "Success Message", Error: "Error message" }, error
func (database *Database) DeleteCollections(collectionDeleteInput CollectionDeleteInput) (CollectionDeleteResult, error) {
	var result CollectionDeleteResult
	var error error = nil

	requestBody := &pb.CollectionDeleteRequest{
		DatabaseName: database.DBName,
		Collections:  collectionDeleteInput.Collections,
	}

	if database.IsgRPC {
		gRPC := database.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.DeleteCollections(ctx, requestBody)

		result.Data = res.GetData()
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s/%s/delete", database.URI, EndpointsMap.Collection, database.DBName)

		restyResp, restyErr := resty.New().
			R().
			SetBody(collectionDeleteInput).
			Delete(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)
	}

	if error == nil {
		if result.Data == "collection deleted successfully" {
			for _, collection := range collectionDeleteInput.Collections {

				if database.Collections[collection] == nil {
					delete(database.Collections, collection)
				}

			}
		}

	}

	return result, nil
}

// return { Data : [collection1, collection2...], Error: "Error message" }, error
func (database *Database) GetAll() (CollectionGetAllResult, error) {
	var result CollectionGetAllResult
	var error error = nil

	requestBody := &pb.CollectionGetAllRequest{
		DatabaseName: database.DBName,
	}

	if database.IsgRPC {
		gRPC := database.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.GetAllCollections(ctx, requestBody)

		result.Data = res.GetData()
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s/%s/get-all", database.URI, EndpointsMap.Collection, database.DBName)

		restyResp, restyErr := resty.New().R().Get(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)

	}
	return result, error
}

// return { Data : { CollectionName string, IndexKeys []string, Documents int} , Error: "Error message" }, error
func (database *Database) GetCollectionStats(collectionName string) (CollectionStatsResult, error) {
	var result CollectionStatsResult
	var error error = nil

	requestBody := &pb.CollectionStatsRequest{
		DatabaseName:   database.DBName,
		CollectionName: collectionName,
	}

	if database.IsgRPC {
		gRPC := database.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.GetCollectionStats(ctx, requestBody)

		result.Data = CollectionStats{
			CollectionName: res.GetData().GetCollectionName(),
			IndexKeys:      res.GetData().GetIndexKeys(),
			Documents:      res.GetData().GetDocuments(),
		}
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s/%s/%s/stats", database.URI, EndpointsMap.Collection, database.DBName, collectionName)

		restyResp, restyErr := resty.New().R().Get(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)

	}
	return result, error
}
