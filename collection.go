package gnosql_client

// return { Data : "Success Message", Error: "Error message" }, error
func (database *Database) CreateCollections(collections []CollectionInput) CollectionCreateResult {
	var result CollectionCreateResult

	requestBody := CollectionCreateRequest{
		DatabaseName: database.DBName,
		Collections:  collections,
	}

	if database.IsgRPC {
		result = GRPC_Create_Collections(database, requestBody)
	} else {
		result = REST_Create_Collections(database, requestBody)
	}

	if result.Error == "" {
		CreateCollectionsInstance(database, collections)
	}

	return result
}

// return { Data : "Success Message", Error: "Error message" }, error
func (database *Database) DeleteCollections(collectionDeleteInput CollectionDeleteInput) CollectionDeleteResult {
	var result CollectionDeleteResult

	requestBody := CollectionDeleteRequest{
		DatabaseName: database.DBName,
		Collections:  collectionDeleteInput.Collections,
	}

	if database.IsgRPC {
		result = GRPC_Delete_Collections(database, requestBody)
	} else {
		result = REST_Delete_Collections(database, requestBody)
	}

	if result.Error == "" {
		if result.Data == COLLECTION_DELETE_SUCCESS_MSG {
			DeleteCollectionInstances(database, collectionDeleteInput.Collections)
		}
	}

	return result
}

// return { Data : [collection1, collection2...], Error: "Error message" }, error
func (database *Database) GetAll() CollectionGetAllResult {
	var result CollectionGetAllResult

	requestBody := CollectionGetAllRequest{
		DatabaseName: database.DBName,
	}

	if database.IsgRPC {
		result = GRPC_GetAll_Collections(database, requestBody)
	} else {
		result = REST_GetAll_Collections(database, requestBody)
	}

	return result
}

// return { Data : { CollectionName string, IndexKeys []string, Documents int} , Error: "Error message" }, error
func (database *Database) GetCollectionStats(collectionName string) CollectionStatsResult {
	var result CollectionStatsResult

	requestBody := CollectionStatsRequest{
		DatabaseName:   database.DBName,
		CollectionName: collectionName,
	}
	if database.IsgRPC {
		result = GRPC_Get_Collection_Stats(database, requestBody)
	} else {
		result = REST_Get_Collection_Stats(database, requestBody)
	}

	return result
}

func CreateCollectionsInstance(database *Database, collections []CollectionInput) {

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

func DeleteCollectionInstances(database *Database, collections []string) {
	for _, collection := range collections {
		if database.Collections[collection] == nil {
			delete(database.Collections, collection)
		}
	}
}
