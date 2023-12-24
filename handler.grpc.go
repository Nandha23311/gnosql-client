package gnosql_client

import (
	"context"
	"encoding/json"
	"time"

	pb "github.com/nanda03dev/gnosql_client/proto"
)

func GRPC_GetAll_DB(client *Client) DatabaseGetAllResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = client.ClientgRPC

	var result DatabaseGetAllResult

	res, gRPCError := gRPC.GetAllDatabases(ctx, &pb.NoRequestBody{})

	result.Data = res.GetData()

	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())
	return result
}

func GRPC_Create_DB(client *Client, request DatabaseCreateRequest) DatabaseCreateResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = client.ClientgRPC

	var result DatabaseCreateResult

	requestBody := &pb.DatabaseCreateRequest{
		DatabaseName: request.DatabaseName,
		Collections:  ConvertToPBCollectionInput(request.Collections),
	}

	res, gRPCError := gRPC.CreateNewDatabase(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Delete_DB(client *Client, request DatabaseDeleteRequest) DatabaseDeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = client.ClientgRPC

	var result DatabaseDeleteResult

	requestBody := &pb.DatabaseDeleteRequest{
		DatabaseName: request.DatabaseName,
	}

	res, gRPCError := gRPC.DeleteDatabase(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Create_Collections(database *Database, request CollectionCreateRequest) CollectionCreateResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.ClientgRPC
	var result CollectionCreateResult

	requestBody := &pb.CollectionCreateRequest{
		DatabaseName: request.DatabaseName,
		Collections:  ConvertToPBCollectionInput(request.Collections),
	}

	res, gRPCError := gRPC.CreateNewCollection(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Delete_Collections(database *Database, request CollectionDeleteRequest) CollectionDeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.ClientgRPC

	var result CollectionDeleteResult

	requestBody := &pb.CollectionDeleteRequest{
		DatabaseName: request.DatabaseName,
		Collections:  request.Collections,
	}

	res, gRPCError := gRPC.DeleteCollections(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_GetAll_Collections(database *Database, request CollectionGetAllRequest) CollectionGetAllResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.ClientgRPC

	var result CollectionGetAllResult

	requestBody := &pb.CollectionGetAllRequest{
		DatabaseName: request.DatabaseName,
	}

	res, gRPCError := gRPC.GetAllCollections(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Get_Collection_Stats(database *Database, request CollectionStatsRequest) CollectionStatsResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.ClientgRPC

	var result CollectionStatsResult

	requestBody := &pb.CollectionStatsRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
	}

	res, gRPCError := gRPC.GetCollectionStats(ctx, requestBody)

	result.Data = CollectionStats{
		CollectionName: res.GetData().GetCollectionName(),
		IndexKeys:      res.GetData().GetIndexKeys(),
		Documents:      res.GetData().GetDocuments(),
	}

	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Create_Document(collection *Collection, request DocumentCreateRequest) DocumentCreateResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC

	var result DocumentCreateResult

	documentData, MarshallErr := json.Marshal(request.Document)

	if MarshallErr != nil {
		result.Error = ERROR_WHILE_MARSHAL_JSON
		return result
	}

	requestBody := &pb.DocumentCreateRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		Document:       string(documentData),
	}

	res, gRPCError := gRPC.CreateDocument(ctx, requestBody)

	var newDocument Document
	result.Error = res.GetError()

	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

	result.Data = newDocument
	result.Error = ValidateResponse(nil, UnMarshallErr, gRPCError, res.GetError())

	return result
}

func GRPC_Read_Document(collection *Collection, request DocumentReadRequest) DocumentReadResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC
	var result DocumentReadResult

	requestBody := &pb.DocumentReadRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		Id:             request.Id,
	}

	res, gRPCError := gRPC.ReadDocument(ctx, requestBody)

	var newDocument Document

	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

	result.Error = ValidateResponse(nil, UnMarshallErr, gRPCError, res.GetError())
	result.Data = newDocument

	return result
}

func GRPC_Filter_Document(collection *Collection, request DocumentFilterRequest) DocumentFilterResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC

	var result DocumentFilterResult

	filterQuery, MarshallErr := json.Marshal(request.Filter)

	if MarshallErr != nil {
		result.Error = ERROR_WHILE_MARSHAL_JSON
		return result
	}

	requestBody := &pb.DocumentFilterRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		Filter:         string(filterQuery),
	}

	res, gRPCError := gRPC.FilterDocument(ctx, requestBody)

	var documents []Document
	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &documents)

	result.Error = ValidateResponse(nil, UnMarshallErr, gRPCError, res.GetError())
	result.Data = documents

	return result
}

func GRPC_Update_Document(collection *Collection, request DocumentUpdateRequest) DocumentUpdateResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC
	var result DocumentUpdateResult

	documentData, MarshallErr := json.Marshal(request.Document)

	if MarshallErr != nil {
		result.Error = ERROR_WHILE_MARSHAL_JSON
		return result
	}

	requestBody := &pb.DocumentUpdateRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		Id:             request.Id,
		Document:       string(documentData),
	}

	res, gRPCError := gRPC.UpdateDocument(ctx, requestBody)

	var newDocument Document

	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

	result.Error = ValidateResponse(nil, UnMarshallErr, gRPCError, res.GetError())

	result.Data = newDocument

	return result
}

func GRPC_Delete_Document(collection *Collection, request DocumentDeleteRequest) DocumentDeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC

	var result DocumentDeleteResult

	requestBody := &pb.DocumentDeleteRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		Id:             request.Id,
	}

	res, gRPCError := gRPC.DeleteDocument(ctx, requestBody)

	result.Data = res.Data
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_GetAll_Document(collection *Collection, request DocumentGetAllRequest) DocumentGetAllResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC

	var result DocumentGetAllResult

	requestBody := &pb.DocumentGetAllRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
	}

	res, gRPCError := gRPC.GetAllDocuments(ctx, requestBody)

	var documents []Document
	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &documents)

	result.Data = documents
	result.Error = ValidateResponse(nil, UnMarshallErr, gRPCError, res.GetError())

	return result
}

func ConvertToPBCollectionInput(collections []CollectionInput) []*pb.CollectionInput {
	PBCollectionsInput := make([]*pb.CollectionInput, 0)

	for _, coll := range collections {
		coll1 := &pb.CollectionInput{
			CollectionName: coll.CollectionName,
			IndexKeys:      coll.IndexKeys,
		}
		PBCollectionsInput = append(PBCollectionsInput, coll1)
	}

	return PBCollectionsInput
}
