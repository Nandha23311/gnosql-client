package gnosql_client

// return { Data : "Sucess message", Error: "Error message" }, error
func (client *Client) Connect(databaseName string, collections []CollectionInput) DatabaseConnectResult {
	var result = DatabaseConnectResult{}

	requestBody := DatabaseCreateRequest{
		DatabaseName: databaseName,
		Collections:  collections,
	}

	result = GRPC_Connect_DB(client, requestBody)

	if result.Error == nil {
		db := CreateDatabaseInstance(client, databaseName)
		CreateCollectionsInstance(db, result.Data.Collections)
	}

	return result
}

func CreateDatabaseInstance(client *Client, DatabaseName string) *Database {
	db := &Database{
		DBName:      DatabaseName,
		URI:         client.URI,
		GrpcClient:  client.GrpcClient,
		Collections: make(map[string]*Collection),
	}

	client.DB[DatabaseName] = db

	return db
}

// return { Data : "Sucess message", Error: "Error message" }, error
func (database *Database) DeleteDatabase(databaseName string) DatabaseDeleteResult {
	var result DatabaseDeleteResult

	requestBody := DatabaseDeleteRequest{
		DatabaseName: databaseName,
	}

	result = GRPC_Delete_DB(database, requestBody)

	return result
}
