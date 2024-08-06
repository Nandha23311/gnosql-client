package gnosql_client

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	pb "github.com/nanda03dev/gnosql_client/proto"
)

func GRPC_GetAll_DB(client *Client) (DatabaseGetAllResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = client.GrpcClient

	var result DatabaseGetAllResult

	res, gRPCError := gRPC.GetAllDatabases(ctx, &pb.NoRequestBody{})

	result.Data = res.GetData()

	err := ValidateResponse(nil, gRPCError)

	return result, err
}

func GRPC_Create_DB(client *Client, request DatabaseCreateRequest) (DatabaseCreateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = client.GrpcClient

	var result DatabaseCreateResult

	requestBody := &pb.DatabaseCreateRequest{
		DatabaseName: request.DatabaseName,
		Collections:  ConvertToPBCollectionInput(request.Collections),
	}

	res, gRPCError := gRPC.CreateNewDatabase(ctx, requestBody)

	result.Data = res.GetData()

	err := ValidateResponse(nil, gRPCError)

	return result, err
}

func GRPC_Connect_DB(client *Client, request DatabaseCreateRequest) (DatabaseConnectResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = client.GrpcClient

	var result DatabaseConnectResult

	requestBody := &pb.DatabaseCreateRequest{
		DatabaseName: request.DatabaseName,
		Collections:  ConvertToPBCollectionInput(request.Collections),
	}

	res, gRPCError := gRPC.ConnectDatabase(ctx, requestBody)
	resultData := res.GetData()

	result.Data = DatabaseResult{
		DatabaseName: resultData.DatabaseName,
		Collections:  resultData.Collections,
	}

	err := ValidateResponse(nil, gRPCError)

	return result, err
}

func GRPC_Delete_DB(database *Database, request DatabaseDeleteRequest) (DatabaseDeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.GrpcClient

	var result DatabaseDeleteResult

	requestBody := &pb.DatabaseDeleteRequest{
		DatabaseName: request.DatabaseName,
	}

	res, gRPCError := gRPC.DeleteDatabase(ctx, requestBody)

	result.Data = res.GetData()
	err := ValidateResponse(nil, gRPCError)

	return result, err
}

func GRPC_Create_Collections(database *Database, request CollectionCreateRequest) (CollectionCreateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.GrpcClient
	var result CollectionCreateResult

	requestBody := &pb.CollectionCreateRequest{
		DatabaseName: request.DatabaseName,
		Collections:  ConvertToPBCollectionInput(request.Collections),
	}

	res, gRPCError := gRPC.CreateNewCollection(ctx, requestBody)

	result.Data = res.GetData()
	err := ValidateResponse(nil, gRPCError)

	return result, err
}

func GRPC_Delete_Collections(database *Database, request CollectionDeleteRequest) (CollectionDeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.GrpcClient

	var result CollectionDeleteResult

	requestBody := &pb.CollectionDeleteRequest{
		DatabaseName: request.DatabaseName,
		Collections:  request.Collections,
	}

	res, gRPCError := gRPC.DeleteCollections(ctx, requestBody)

	result.Data = res.GetData()
	err := ValidateResponse(nil, gRPCError)

	return result, err
}

func GRPC_GetAll_Collections(database *Database, request CollectionGetAllRequest) (CollectionGetAllResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.GrpcClient

	var result CollectionGetAllResult

	requestBody := &pb.CollectionGetAllRequest{
		DatabaseName: request.DatabaseName,
	}

	res, gRPCError := gRPC.GetAllCollections(ctx, requestBody)

	result.Data = res.GetData()
	err := ValidateResponse(nil, gRPCError)

	return result, err
}

func GRPC_Get_Collection_Stats(database *Database, request CollectionStatsRequest) (CollectionStatsResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.GrpcClient

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

	err := ValidateResponse(nil, gRPCError)

	return result, err
}

func GRPC_Create_Document(collection *Collection, request DocumentCreateRequest) (DocumentCreateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.GrpcClient

	var result DocumentCreateResult

	documentData, MarshallErr := json.Marshal(request.Document)

	if MarshallErr != nil {
		return result, errors.New(ERROR_WHILE_MARSHAL_JSON)
	}

	requestBody := &pb.DocumentCreateRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		Document:       string(documentData),
	}

	res, gRPCError := gRPC.CreateDocument(ctx, requestBody)

	var newDocument Document

	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

	result.Data = newDocument
	err := ValidateResponse(UnMarshallErr, gRPCError)

	return result, err
}

func GRPC_Read_Document(collection *Collection, request DocumentReadRequest) (DocumentReadResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.GrpcClient
	var result DocumentReadResult

	requestBody := &pb.DocumentReadRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		DocId:          request.DocId,
	}

	res, gRPCError := gRPC.ReadDocument(ctx, requestBody)

	var newDocument Document

	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

	err := ValidateResponse(UnMarshallErr, gRPCError)
	result.Data = newDocument

	return result, err
}

func GRPC_Filter_Document(collection *Collection, request DocumentFilterRequest) (DocumentFilterResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.GrpcClient

	var result DocumentFilterResult

	filterQuery, MarshallErr := json.Marshal(request.Filter)

	if MarshallErr != nil {
		return result, errors.New(ERROR_WHILE_MARSHAL_JSON)
	}

	requestBody := &pb.DocumentFilterRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		Filter:         string(filterQuery),
	}

	res, gRPCError := gRPC.FilterDocument(ctx, requestBody)

	var documents []Document
	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &documents)

	err := ValidateResponse(UnMarshallErr, gRPCError)
	result.Data = documents

	return result, err
}

func GRPC_Update_Document(collection *Collection, request DocumentUpdateRequest) (DocumentUpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.GrpcClient
	var result DocumentUpdateResult

	documentData, MarshallErr := json.Marshal(request.Document)

	if MarshallErr != nil {
		return result, errors.New(ERROR_WHILE_MARSHAL_JSON)
	}

	requestBody := &pb.DocumentUpdateRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		DocId:          request.DocId,
		Document:       string(documentData),
	}

	res, gRPCError := gRPC.UpdateDocument(ctx, requestBody)

	var newDocument Document

	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

	err := ValidateResponse(UnMarshallErr, gRPCError)

	result.Data = newDocument

	return result, err
}

func GRPC_Delete_Document(collection *Collection, request DocumentDeleteRequest) (DocumentDeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.GrpcClient

	var result DocumentDeleteResult

	requestBody := &pb.DocumentDeleteRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
		DocId:          request.DocId,
	}

	res, gRPCError := gRPC.DeleteDocument(ctx, requestBody)

	result.Data = res.Data
	err := ValidateResponse(nil, gRPCError)

	return result, err
}

func GRPC_GetAll_Document(collection *Collection, request DocumentGetAllRequest) (DocumentGetAllResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.GrpcClient

	var result DocumentGetAllResult

	requestBody := &pb.DocumentGetAllRequest{
		DatabaseName:   request.DatabaseName,
		CollectionName: request.CollectionName,
	}

	res, gRPCError := gRPC.GetAllDocuments(ctx, requestBody)

	var documents []Document
	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &documents)

	result.Data = documents
	err := ValidateResponse(UnMarshallErr, gRPCError)

	return result, err
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
