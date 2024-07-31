package gnosql_client

// return { Data : "Sucess message", Error: "Error message" }, error
func (client *Client) Connect(databaseName string, collections []CollectionInput) DatabaseConnectResult {
	var result = DatabaseConnectResult{}

	requestBody := DatabaseCreateRequest{
		DatabaseName: databaseName,
		Collections:  collections,
	}

	if client.IsgRPC {
		result = GRPC_Connect_DB(client, requestBody)
	} else {
		result = REST_Connect_DB(client, requestBody)
	}

	if result.Error == "" {
		db := CreateDatabaseInstance(client, databaseName)
		CreateCollectionsInstance(db, result.Data.Collections)
	}

	return result
}

func CreateDatabaseInstance(client *Client, DatabaseName string) *Database {
	db := &Database{
		DBName:      DatabaseName,
		URI:         client.URI,
		IsgRPC:      client.IsgRPC,
		GrpcClient:  client.GrpcClient,
		RestClient:  client.RestClient,
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

	if database.IsgRPC {
		result = GRPC_Delete_DB(database, requestBody)
	} else {
		result = REST_Delete_DB(database, requestBody)
	}

	// if result.Error == "" {
	// 	if client.DB[databaseName] != nil {
	// 		delete(client.DB, databaseName)
	// 	}
	// }

	return result
}

// // return { Data : [DatabaseName1, DatabaseName2...], Error: "Error message" }, error
// func (client *Client) GetAllDatabase() DatabaseGetAllResult {
// 	var result = DatabaseGetAllResult{}

// 	if client.IsgRPC {
// 		result = GRPC_GetAll_DB(client)
// 	} else {
// 		result = REST_GetAll_DB(client)
// 	}

// 	return result

// }

// // return { Data : "Sucess message", Error: "Error message" }, error
// func (client *Client) Create(databaseName string, collections []CollectionInput) DatabaseCreateResult {
// 	var result = DatabaseCreateResult{}

// 	requestBody := DatabaseCreateRequest{
// 		DatabaseName: databaseName,
// 		Collections:  collections,
// 	}

// 	if client.IsgRPC {
// 		result = GRPC_Create_DB(client, requestBody)
// 	} else {
// 		result = REST_Create_DB(client, requestBody)
// 	}

// 	if result.Error == "" {
// 		db := CreateDatabaseInstance(client, databaseName)

// 		var collectionNames []string
// 		for _, collection := range collections {
// 			collectionNames = append(collectionNames, collection.CollectionName)
// 		}

// 		CreateCollectionsInstance(db, collectionNames)
// 	}

// 	return result
// }

// // return { Data : [DatabaseName1, DatabaseName2...], Error: "Error message" }, error
// func (client *Client) GetAllDatabase() DatabaseGetAllResult {
// 	var result = DatabaseGetAllResult{}

// 	if client.IsgRPC {
// 		result = GRPC_GetAll_DB(client)
// 	} else {
// 		result = REST_GetAll_DB(client)
// 	}

// 	return result

// }
