package gnosql_client

// return { Data : "Success Message", Error: "Error message" }, error
func (database *Database) CreateCollections(collections []CollectionInput) (CollectionCreateResult, error) {

	requestBody := CollectionCreateRequest{
		DatabaseName: database.DBName,
		Collections:  collections,
	}

	result, err := GRPC_Create_Collections(database, requestBody)
	if err == nil {
		var collectionNames []string
		collectionsResult, collectionsResultErr := database.GetAll()

		if collectionsResultErr == nil {
			collectionNames = append(collectionNames, collectionsResult.Data...)
			CreateCollectionsInstance(database, collectionNames)
		}

	}

	return result, nil
}

// return { Data : "Success Message", Error: "Error message" }, error
func (database *Database) DeleteCollections(collectionDeleteInput CollectionDeleteInput) (CollectionDeleteResult, error) {

	requestBody := CollectionDeleteRequest{
		DatabaseName: database.DBName,
		Collections:  collectionDeleteInput.Collections,
	}

	result, err := GRPC_Delete_Collections(database, requestBody)

	if err == nil && result.Data == COLLECTION_DELETE_SUCCESS_MSG {
		DeleteCollectionInstances(database, collectionDeleteInput.Collections)
	}

	return result, nil
}

// return { Data : [collection1, collection2...], Error: "Error message" }, error
func (database *Database) GetAll() (CollectionGetAllResult, error) {

	requestBody := CollectionGetAllRequest{
		DatabaseName: database.DBName,
	}

	result, err := GRPC_GetAll_Collections(database, requestBody)
	return result, err
}

// return { Data : { CollectionName string, IndexKeys []string, Documents int} , Error: "Error message" }, error
func (database *Database) GetCollectionStats(collectionName string) (CollectionStatsResult, error) {
	var result CollectionStatsResult

	requestBody := CollectionStatsRequest{
		DatabaseName:   database.DBName,
		CollectionName: collectionName,
	}

	result, err := GRPC_Get_Collection_Stats(database, requestBody)

	return result, err
}

func CreateCollectionsInstance(database *Database, collectionNames []string) {

	for _, collectionName := range collectionNames {

		if database.Collections[collectionName] == nil {
			collectionInstance := &Collection{
				CollectionName: collectionName,
				URI:            database.URI,
				DBName:         database.DBName,
				IsgRPC:         database.IsgRPC,
				GrpcClient:     database.GrpcClient,
				RestClient:     database.RestClient,
			}

			database.Collections[collectionName] = collectionInstance
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
