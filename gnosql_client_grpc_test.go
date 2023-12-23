package gnosql_client

import (
	"fmt"
	"testing"
)

const (
	ADDRESS = "localhost:5455"
)

func TestGnoSQLGRPC(t *testing.T) {
	newClient := Connect(ADDRESS, true)

	var GetAllDatabaseResult = newClient.GetAll()
	fmt.Printf("\n GetAllDatabaseResult %v \n", GetAllDatabaseResult)

	var DatabaseName = "test-g-c"
	var userCollectionName = "users"
	var orderCollectionName = "orders"

	UserCollectionInput := CollectionInput{
		CollectionName: userCollectionName,
		IndexKeys:      []string{"city", "pincode"},
	}
	collectionsInput1 := []CollectionInput{UserCollectionInput}

	var CreateDatabaseResult = newClient.Create(DatabaseName, collectionsInput1)

	fmt.Printf("\n CreateDatabaseResult %v \n", CreateDatabaseResult)

	var CreateDatabaseResult2 = newClient.Create(DatabaseName, collectionsInput1)

	fmt.Printf("\n CreateDatabaseResult2 %v \n", CreateDatabaseResult2)

	var GetAllDatabaseResult2 = newClient.GetAll()
	fmt.Printf("\n GetAllDatabaseResult %v \n", GetAllDatabaseResult2)

	// // ------------------------------------------------------------------------------// ------------------------------------------------------------------------------

	var db *Database = newClient.DB[DatabaseName]

	if db != nil {

		OrderCollectionInput := CollectionInput{
			CollectionName: orderCollectionName,
			IndexKeys:      []string{"userId", "category"},
		}
		collectionsInput2 := []CollectionInput{OrderCollectionInput}

		var CreateCollectionResult = db.CreateCollections(collectionsInput2)
		fmt.Printf("\n CreateCollectionResult %v \n", CreateCollectionResult)

		var GetCollectionsResult = db.GetAll()
		fmt.Printf("\n GetCollectionsResult %v \n", GetCollectionsResult)

		var collectionDeleteInput = CollectionDeleteInput{
			Collections: []string{orderCollectionName},
		}

		var DeleteCollectionResult = db.DeleteCollections(collectionDeleteInput)
		fmt.Printf("\n DeleteCollectionResult %v \n", DeleteCollectionResult)

		var GetCollectionsResult2 = db.GetAll()
		fmt.Printf("\n GetCollectionsResult %v \n", GetCollectionsResult2)

		var GetCollectionStatsResult2 = db.GetCollectionStats(userCollectionName)
		fmt.Printf("\n GetCollectionStatsResult2 %v \n", GetCollectionStatsResult2)

		var userCollection *Collection = db.Collections[userCollectionName]

		if userCollection != nil {

			user1 := make(Document)
			user1["name"] = "Nandakumar"
			user1["city"] = "Chennai"
			user1["pincode"] = "600100"

			var DocumentCreateResult = userCollection.Create(user1)
			fmt.Printf("\n DocumentCreateResult %v \n", DocumentCreateResult)

			var GetCollectionStatsResult3 = db.GetCollectionStats(userCollectionName)
			fmt.Printf("\n GetCollectionStatsResult3 %v \n", GetCollectionStatsResult3)

			if _, exists := DocumentCreateResult.Data["id"]; exists {

				var id = DocumentCreateResult.Data["id"].(string)

				var DocumentReadResult = userCollection.Read(id)
				fmt.Printf("\n DocumentReadResult %v \n", DocumentReadResult)

				user2 := make(Document)
				user2["name"] = "kumar"
				user2["city"] = "Chennai"
				user2["pincode"] = "600101"

				var DocumentCreateResult2 = userCollection.Create(user2)
				fmt.Printf("\n DocumentCreateResult2 %v \n", DocumentCreateResult2)

				var filter DocumentFilterQuery = DocumentFilterQuery{
					"name": "Nandakumar",
				}

				var DocumentFilterResult = userCollection.Filter(filter)
				fmt.Printf("\n DocumentFilterResult %v \n", DocumentFilterResult)

				user1["designation"] = "developer"

				var DocumentUpdateResult = userCollection.Update(id, user1)
				fmt.Printf("\n DocumentUpdateResult %v \n", DocumentUpdateResult)

				var DocumentReadResult3 = userCollection.Read(id)
				fmt.Printf("\n DocumentReadResult3 %v \n", DocumentReadResult3)

				var DocumentDeleteResult = userCollection.Delete(id)
				fmt.Printf("\n DocumentDeleteResult %v \n", DocumentDeleteResult)

				var DocumentReadResult4 = userCollection.Read(id)
				fmt.Printf("\n DocumentReadResult4 %v \n", DocumentReadResult4)

				var DocumentGetAllResult = userCollection.GetAll()
				fmt.Printf("\n DocumentGetAllResult %v \n", DocumentGetAllResult)
			}

		} else {
			println("User collection nil")
		}

	}

	// ------------------------------------------------------------------------------
	var DeleteDatabaseResult = newClient.Delete(DatabaseName)
	fmt.Printf("\n DeleteDatabaseResult %v ", DeleteDatabaseResult)
}
