package gnosql_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"

	pb "github.com/nanda03dev/gnosql_client/proto"
)

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Create(document Document) (DocumentCreateResult, error) {

	var result DocumentCreateResult
	var error error = nil

	documentData, MarshallErr := json.Marshal(document)

	if MarshallErr != nil {
		result.Error = "Error in document json"
		return result, error
	}

	requestBody := &pb.DocumentCreateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Document:       string(documentData),
	}

	if collection.IsgRPC {
		gRPC := collection.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.CreateDocument(ctx, requestBody)

		var newDocument Document

		var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

		error = ValidateResponse(nil, UnMarshallErr)

		result.Data = newDocument
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName)

		restyResp, restyErr := resty.New().
			R().
			SetBody(document).
			Post(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)
	}

	return result, error
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Read(id string) (DocumentReadResult, error) {
	var result DocumentReadResult
	var error error = nil

	requestBody := &pb.DocumentReadRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Id:             id,
	}

	if collection.IsgRPC {
		gRPC := collection.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.ReadDocument(ctx, requestBody)

		var newDocument Document
		var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

		error = ValidateResponse(nil, UnMarshallErr)

		result.Data = newDocument
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName, id)

		restyResp, restyErr := resty.New().
			R().
			Get(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)
	}

	return result, error
}

// return { Data : []Document, Error: "Error message" }, error
func (collection *Collection) Filter(filter DocumentFilterQuery) (DocumentFilterResult, error) {
	var result DocumentFilterResult

	var error error = nil

	filterQuery, MarshallErr := json.Marshal(filter)

	if MarshallErr != nil {
		result.Error = "Error in filter map"
		return result, error
	}

	requestBody := &pb.DocumentFilterRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Filter:         string(filterQuery),
	}

	if collection.IsgRPC {
		gRPC := collection.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.FilterDocument(ctx, requestBody)

		var documents []Document
		var UnMarshallErr = json.Unmarshal([]byte(res.Data), &documents)

		error = ValidateResponse(nil, UnMarshallErr)

		result.Data = documents
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s/%s/%s/filter", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName)

		restyResp, restyErr := resty.New().
			R().
			SetBody(filter).
			Post(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)
	}
	return result, error
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Update(id string, document Document) (DocumentUpdateResult, error) {

	var result DocumentUpdateResult
	var error error = nil

	documentData, MarshallErr := json.Marshal(document)

	if MarshallErr != nil {
		result.Error = "Error in document json"
		return result, error
	}

	requestBody := &pb.DocumentUpdateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Id:             id,
		Document:       string(documentData),
	}

	if collection.IsgRPC {
		gRPC := collection.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.UpdateDocument(ctx, requestBody)

		var newDocument Document
		var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

		error = ValidateResponse(nil, UnMarshallErr)

		result.Data = newDocument
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName, id)

		restyResp, restyErr := resty.New().
			R().
			SetBody(document).
			Put(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)
	}
	return result, error
}

// return { Data : "Success Message", Error: "Error message" }, error
func (collection *Collection) Delete(id string) (DocumentDeleteResult, error) {
	var result DocumentDeleteResult
	var error error = nil

	requestBody := &pb.DocumentDeleteRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Id:             id,
	}

	if collection.IsgRPC {
		gRPC := collection.ClientgRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, _ := gRPC.DeleteDocument(ctx, requestBody)

		result.Data = res.Data
		result.Error = res.GetError()

	} else {
		path := fmt.Sprintf("%s/%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName, id)

		restyResp, restyErr := resty.New().
			R().
			Delete(path)

		var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

		error = ValidateResponse(restyErr, UnMarshallErr)
	}
	return result, error
}
