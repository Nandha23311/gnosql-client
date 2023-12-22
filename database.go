package gnosql_client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	pb "github.com/nanda03dev/gnosql_client/proto"
)

// return { Data : [DatabaseName1, DatabaseName2...], Error: "Error message" }, error
func (client *Client) GetAll() (DatabaseGetAllResult, error) {
	var result DatabaseGetAllResult

	if client.IsgRPC {
		gRPC := client.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.GetAllDatabases(ctx, &pb.NoRequestBody{})

		result.Data = res.GetData()
		result.Error = res.GetError()
		return result, nil
	} else {
		path := fmt.Sprintf("%s/%s", client.URI, EndpointsMap.DatabaseGetAll)

		restyResp, restyErr := resty.New().R().Get(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)
		return result, ValidateResponse(restyErr, UnMarshallErr)

	}

}

// return { Data : "Sucess message", Error: "Error message" }, error
func (client *Client) Create(databaseName string, collections []CollectionInput) (DatabaseCreateResult, error) {
	var result DatabaseCreateResult
	var error error = nil

	colls := make([]*pb.CollectionInput, 0)

	for _, coll := range collections {
		coll1 := &pb.CollectionInput{
			CollectionName: coll.CollectionName,
			IndexKeys:      coll.IndexKeys,
		}
		colls = append(colls, coll1)
	}

	requestBody := &pb.DatabaseCreateRequest{
		DatabaseName: databaseName,
		Collections:  colls,
	}

	if client.IsgRPC {
		gRPC := client.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.CreateNewDatabase(ctx, requestBody)

		result.Data = res.GetData()
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s", client.URI, EndpointsMap.DatabaseAdd)

		restyResp, restyErr := resty.New().
			R().
			SetBody(requestBody).
			Post(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)

	}

	if error == nil {
		database := Database{
			DBName:      databaseName,
			URI:         client.URI,
			IsgRPC:      client.IsgRPC,
			ClientgRPC:  client.ClientgRPC,
			Collections: make(map[string]*Collection),
		}
		client.DB[databaseName] = &database
	}

	return result, nil
}

// return { Data : "Sucess message", Error: "Error message" }, error
func (client *Client) Delete(databaseName string) (DatabaseDeleteResult, error) {
	var result DatabaseDeleteResult
	var error error = nil

	requestBody := &pb.DatabaseDeleteRequest{
		DatabaseName: databaseName,
	}

	if client.IsgRPC {
		gRPC := client.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.DeleteDatabase(ctx, requestBody)

		result.Data = res.GetData()
		result.Error = res.GetError()
	} else {

		path := fmt.Sprintf("%s/%s", client.URI, EndpointsMap.DatabaseDelete)

		restyResp, restyErr := resty.New().
			R().
			SetBody(requestBody).
			Post(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)
	}

	if error == nil {
		if client.DB[databaseName] != nil {
			delete(client.DB, databaseName)
		}
	}

	return result, nil
}
