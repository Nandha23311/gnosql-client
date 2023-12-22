package gnosql_client

import (
	"fmt"
	"testing"
)

const (
	ADDRESS = "localhost:50051"
)

func TestGnoSQLGRPC(t *testing.T) {
	newClient := Connect(ADDRESS, true)

	var GetAllDatabaseResult, _ = newClient.GetAll()
	fmt.Printf("\n GetAllDatabaseResult %v \n", GetAllDatabaseResult.Data)

	var DatabaseName = "test-g-c"
	var userCollectionName = "users"
	var orderCollectionName = "orders"

	UserCollectionInput := CollectionInput{
		CollectionName: userCollectionName,
		IndexKeys:      []string{"city", "pincode"},
	}
	collectionsInput1 := []CollectionInput{UserCollectionInput}

	var CreateDatabaseResult, _ = newClient.Create(DatabaseName, collectionsInput1)

	fmt.Printf("\n CreateDatabaseResult %v \n", CreateDatabaseResult)

	// // ------------------------------------------------------------------------------// ------------------------------------------------------------------------------

	var db *Database = newClient.DB[DatabaseName]

	if db != nil {

		OrderCollectionInput := CollectionInput{
			CollectionName: orderCollectionName,
			IndexKeys:      []string{"userId", "category"},
		}
		collectionsInput2 := []CollectionInput{OrderCollectionInput}

		var CreateCollectionResult, _ = db.CreateCollections(collectionsInput2)
		fmt.Printf("\n CreateCollectionResult %v \n", CreateCollectionResult)

		var GetCollectionsResult, _ = db.GetAll()
		fmt.Printf("\n GetCollectionsResult %v \n", GetCollectionsResult)

		var collectionDeleteInput = CollectionDeleteInput{
			Collections: []string{orderCollectionName},
		}

		var DeleteCollectionResult, _ = db.DeleteCollections(collectionDeleteInput)
		fmt.Printf("\n DeleteCollectionResult %v \n", DeleteCollectionResult)

		var GetCollectionsResult2, _ = db.GetAll()
		fmt.Printf("\n GetCollectionsResult %v \n", GetCollectionsResult2)

		var GetCollectionStatsResult2, _ = db.GetCollectionStats(userCollectionName)
		fmt.Printf("\n GetCollectionStatsResult2 %v \n", GetCollectionStatsResult2)

	}

	// ------------------------------------------------------------------------------
	var DeleteDatabaseResult, _ = newClient.Delete(DatabaseName)
	fmt.Printf("\n DeleteDatabaseResult %v ", DeleteDatabaseResult)
}
