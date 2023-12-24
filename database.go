package gnosql_client

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

	requestBody := DatabaseCreateRequest{
		DatabaseName: databaseName,
		Collections:  collections,
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

	requestBody := DatabaseDeleteRequest{
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
