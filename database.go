package gnosql_client

import (
	pb "github.com/nanda03dev/gnosql_client/proto"
)

// return { Data : [DatabaseName1, DatabaseName2...], Error: "Error message" }, error
func (client *Client) GetAll() DatabaseGetAllResult {
	var result = DatabaseGetAllResult{}

	if client.IsgRPC {
		result = GRPC_GetAll_DB(client)
	} else {
		result = REST_GetAll_DB(client)
	}

	return result

}

// return { Data : "Sucess message", Error: "Error message" }, error
func (client *Client) Create(databaseName string, collections []CollectionInput) DatabaseCreateResult {
	var result = DatabaseCreateResult{}

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
		result = GRPC_Create_DB(client, requestBody)
	} else {
		result = REST_Create_DB(client, requestBody)
	}

	if result.Error == "" {
		db := CreateDatabaseInstance(client, databaseName)
		CreateCollectionsInstance(db, collections)
	}

	return result
}

// return { Data : "Sucess message", Error: "Error message" }, error
func (client *Client) Delete(databaseName string) DatabaseDeleteResult {
	var result DatabaseDeleteResult

	requestBody := &pb.DatabaseDeleteRequest{
		DatabaseName: databaseName,
	}

	if client.IsgRPC {
		result = GRPC_Delete_DB(client, requestBody)
	} else {
		result = REST_Delete_DB(client, requestBody)
	}

	if result.Error == "" {
		if client.DB[databaseName] != nil {
			delete(client.DB, databaseName)
		}
	}

	return result
}

func CreateDatabaseInstance(client *Client, DatabaseName string) *Database {
	db := &Database{
		DBName:      DatabaseName,
		URI:         client.URI,
		IsgRPC:      client.IsgRPC,
		ClientgRPC:  client.ClientgRPC,
		Collections: make(map[string]*Collection),
	}

	client.DB[DatabaseName] = db

	return db
}
