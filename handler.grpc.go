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

func GRPC_Create_DB(client *Client, requestBody *pb.DatabaseCreateRequest) DatabaseCreateResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = client.ClientgRPC

	var result DatabaseCreateResult

	res, gRPCError := gRPC.CreateNewDatabase(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Delete_DB(client *Client, requestBody *pb.DatabaseDeleteRequest) DatabaseDeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = client.ClientgRPC

	var result DatabaseDeleteResult

	res, gRPCError := gRPC.DeleteDatabase(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Create_Collections(database *Database, requestBody *pb.CollectionCreateRequest) CollectionCreateResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.ClientgRPC

	var result CollectionCreateResult

	res, gRPCError := gRPC.CreateNewCollection(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Delete_Collections(database *Database, requestBody *pb.CollectionDeleteRequest) CollectionDeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.ClientgRPC

	var result CollectionDeleteResult

	res, gRPCError := gRPC.DeleteCollections(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_GetAll_Collections(database *Database, requestBody *pb.CollectionGetAllRequest) CollectionGetAllResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.ClientgRPC

	var result CollectionGetAllResult

	res, gRPCError := gRPC.GetAllCollections(ctx, requestBody)

	result.Data = res.GetData()
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Get_Collection_Stats(database *Database, requestBody *pb.CollectionStatsRequest) CollectionStatsResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = database.ClientgRPC

	var result CollectionStatsResult

	res, gRPCError := gRPC.GetCollectionStats(ctx, requestBody)

	result.Data = CollectionStats{
		CollectionName: res.GetData().GetCollectionName(),
		IndexKeys:      res.GetData().GetIndexKeys(),
		Documents:      res.GetData().GetDocuments(),
	}

	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_Create_Document(collection *Collection, document Document) DocumentCreateResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC

	var result DocumentCreateResult

	documentData, MarshallErr := json.Marshal(document)

	if MarshallErr != nil {
		result.Error = ERROR_WHILE_MARSHAL_JSON
		return result
	}

	requestBody := &pb.DocumentCreateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
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

func GRPC_Read_Document(collection *Collection, id string) DocumentReadResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC
	var result DocumentReadResult

	requestBody := &pb.DocumentReadRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Id:             id,
	}

	res, gRPCError := gRPC.ReadDocument(ctx, requestBody)

	var newDocument Document

	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

	result.Error = ValidateResponse(nil, UnMarshallErr, gRPCError, res.GetError())
	result.Data = newDocument

	return result
}

func GRPC_Filter_Document(collection *Collection, filter DocumentFilterQuery) DocumentFilterResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC

	var result DocumentFilterResult

	filterQuery, MarshallErr := json.Marshal(filter)

	if MarshallErr != nil {
		result.Error = ERROR_WHILE_MARSHAL_JSON
		return result
	}

	requestBody := &pb.DocumentFilterRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Filter:         string(filterQuery),
	}

	res, gRPCError := gRPC.FilterDocument(ctx, requestBody)

	var documents []Document
	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &documents)

	result.Error = ValidateResponse(nil, UnMarshallErr, gRPCError, res.GetError())
	result.Data = documents

	return result
}

func GRPC_Update_Document(collection *Collection, id string, document Document) DocumentUpdateResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC
	var result DocumentUpdateResult

	documentData, MarshallErr := json.Marshal(document)

	if MarshallErr != nil {
		result.Error = ERROR_WHILE_MARSHAL_JSON
		return result
	}

	requestBody := &pb.DocumentUpdateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Id:             id,
		Document:       string(documentData),
	}

	res, gRPCError := gRPC.UpdateDocument(ctx, requestBody)

	var newDocument Document

	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &newDocument)

	result.Error = ValidateResponse(nil, UnMarshallErr, gRPCError, res.GetError())

	result.Data = newDocument

	return result
}

func GRPC_Delete_Document(collection *Collection, id string) DocumentDeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC

	var result DocumentDeleteResult

	requestBody := &pb.DocumentDeleteRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Id:             id,
	}

	res, gRPCError := gRPC.DeleteDocument(ctx, requestBody)

	result.Data = res.Data
	result.Error = ValidateResponse(nil, nil, gRPCError, res.GetError())

	return result
}

func GRPC_GetAll_Document(collection *Collection) DocumentGetAllResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = collection.ClientgRPC

	var result DocumentGetAllResult

	requestBody := &pb.DocumentGetAllRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
	}

	res, gRPCError := gRPC.GetAllDocuments(ctx, requestBody)

	var documents []Document
	var UnMarshallErr = json.Unmarshal([]byte(res.Data), &documents)

	result.Data = documents
	result.Error = ValidateResponse(nil, UnMarshallErr, gRPCError, res.GetError())

	return result
}
