package gnosql_client

import (
	"fmt"
	"testing"
)

const (
	HTTP_URI = "http://localhost:5454"
)

func TestGnoSQLREST(t *testing.T) {

	var DatabaseName = "test-g-c"

	db := Connect(HTTP_URI, DatabaseName, false)

	var userCollectionName = "users"

	UserCollectionInput := CollectionInput{
		CollectionName: userCollectionName,
		IndexKeys:      []string{"city", "pincode"},
	}
	collectionsInput1 := []CollectionInput{UserCollectionInput}

	var CreateCollectionResult1 = db.CreateCollections(collectionsInput1)
	fmt.Printf("\n CreateCollectionResult1 %v \n", CreateCollectionResult1)

	var orderCollectionName = "orders"

	OrderCollectionInput := CollectionInput{
		CollectionName: orderCollectionName,
		IndexKeys:      []string{"userId", "category"},
	}

	collectionsInput2 := []CollectionInput{OrderCollectionInput}

	var CreateCollectionResult2 = db.CreateCollections(collectionsInput2)
	fmt.Printf("\n CreateCollectionResult %v \n", CreateCollectionResult2)

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

		var docId = DocumentCreateResult.Data["docId"].(string)

		var DocumentReadResult = userCollection.Read(docId)
		fmt.Printf("\n DocumentReadResult %v \n", DocumentReadResult)

		var filter MapInterface = MapInterface{
			"name": "Nandakumar",
		}

		var DocumentFilterResult = userCollection.Filter(filter)
		fmt.Printf("\n DocumentFilterResult %v \n", DocumentFilterResult)

		user1["designation"] = "developer"

		var DocumentUpdateResult = userCollection.Update(docId, user1)
		fmt.Printf("\n DocumentUpdateResult %v \n", DocumentUpdateResult)

		var DocumentDeleteResult = userCollection.Delete(docId)
		fmt.Printf("\n DocumentDeleteResult %v \n", DocumentDeleteResult)

	}

	var DeleteDatabaseResult = db.DeleteDatabase(DatabaseName)
	fmt.Printf("\n DeleteDatabaseResult %v ", DeleteDatabaseResult)

}
