package gnosql_client

import (
	"fmt"
	"testing"
)

func TestGnoSQLREST(t *testing.T) {
	uri := "http://localhost:5454"
	newClient := Connect(uri, false)

	var DatabaseName = "test-g-c"
	var userCollectionName = "users"

	UserCollectionInput := CollectionInput{
		CollectionName: userCollectionName,
		IndexKeys:      []string{"city", "pincode"},
	}
	collectionsInput1 := []CollectionInput{UserCollectionInput}

	var CreateDatabaseResult = newClient.Create(DatabaseName, collectionsInput1)
	fmt.Printf("\n CreateDatabaseResult %v \n", CreateDatabaseResult)

	var GetAllDatabaseResult = newClient.GetAll()
	fmt.Printf("\n GetAllDatabaseResult %v \n", GetAllDatabaseResult.Data)

	var db *Database = newClient.DB[DatabaseName]

	if db != nil {
		var orderCollectionName = "orders"

		OrderCollectionInput := CollectionInput{
			CollectionName: orderCollectionName,
			IndexKeys:      []string{"userId", "category"},
		}

		collectionsInput2 := []CollectionInput{OrderCollectionInput}

		var CreateCollectionResult = db.CreateCollections(collectionsInput2)
		fmt.Printf("\n CreateCollectionResult %v \n", CreateCollectionResult)

		var GetCollectionsResult = db.GetAll()
		fmt.Printf("\n GetCollectionsResult %v \n", GetCollectionsResult)

		var GetCollectionStatsResult = db.GetCollectionStats(userCollectionName)
		fmt.Printf("\n GetCollectionStatsResult %v \n", GetCollectionStatsResult)

		var collectionDeleteInput = CollectionDeleteInput{
			Collections: []string{orderCollectionName},
		}

		var DeleteCollectionResult = db.DeleteCollections(collectionDeleteInput)
		fmt.Printf("\n DeleteCollectionResult %v \n", DeleteCollectionResult)

		var GetCollectionsResult2 = db.GetAll()
		fmt.Printf("\n GetCollectionsResult2 %v \n", GetCollectionsResult2)

		var userCollection *Collection = db.Collections[userCollectionName]

		if userCollection != nil {

			user1 := make(Document)
			user1["name"] = "Nandakumar"
			user1["city"] = "Chennai"
			user1["pincode"] = "600100"

			var DocumentCreateResult = userCollection.Create(user1)
			fmt.Printf("\n DocumentCreateResult %v \n", DocumentCreateResult)

			user2 := make(Document)
			user2["name"] = "kumar"
			user2["city"] = "Chennai"
			user2["pincode"] = "600101"

			var DocumentCreateResult2 = userCollection.Create(user2)
			fmt.Printf("\n DocumentCreateResult2 %v \n", DocumentCreateResult2)

			var id = DocumentCreateResult.Data["id"].(string)

			var DocumentReadResult = userCollection.Read(id)
			fmt.Printf("\n DocumentReadResult %v \n", DocumentReadResult)

			var filter DocumentFilterQuery = DocumentFilterQuery{
				"name": "Nandakumar",
			}

			var DocumentFilterResult = userCollection.Filter(filter)
			fmt.Printf("\n DocumentFilterResult %v \n", DocumentFilterResult)

			user1["designation"] = "developer"

			var DocumentUpdateResult = userCollection.Update(id, user1)
			fmt.Printf("\n DocumentUpdateResult %v \n", DocumentUpdateResult)

			var DocumentDeleteResult = userCollection.Delete(id)
			fmt.Printf("\n DocumentDeleteResult %v \n", DocumentDeleteResult)

		}

	}

	var DeleteDatabaseResult = newClient.Delete(DatabaseName)
	fmt.Printf("\n DeleteDatabaseResult %v ", DeleteDatabaseResult)

}
