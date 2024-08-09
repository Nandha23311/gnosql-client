package gnosql_client

import (
	"fmt"
	"testing"
)

const (
	GRPC_URI = "localhost:5455"
)

func TestGnoSQLGRPC(t *testing.T) {
	var DatabaseName = "test-g-c"

	db := Connect(GRPC_URI, DatabaseName, true)

	var userCollectionName = "users"
	var orderCollectionName = "orders"

	UserCollectionInput := CollectionInput{
		CollectionName: userCollectionName,
		IndexKeys:      []string{"city", "pincode"},
	}

	collectionsInput1 := []CollectionInput{UserCollectionInput}

	var CreateCollectionResult1, CreateCollectionResult1Err = db.CreateCollections(collectionsInput1)
	fmt.Printf("\n CreateCollectionResult1 %v \n", CreateCollectionResult1)
	fmt.Printf("\n CreateCollectionResultErr %v \n", CreateCollectionResult1Err)

	OrderCollectionInput := CollectionInput{
		CollectionName: orderCollectionName,
		IndexKeys:      []string{"userId", "category"},
	}
	collectionsInput2 := []CollectionInput{OrderCollectionInput}

	var CreateCollectionResult, CreateCollectionResultErr = db.CreateCollections(collectionsInput2)
	fmt.Printf("\n CreateCollectionResult %v \n", CreateCollectionResult)
	fmt.Printf("\n CreateCollectionResultErr %v \n", CreateCollectionResultErr)

	var GetCollectionsResult, GetCollectionsResultErr = db.GetAll()
	fmt.Printf("\n GetCollectionsResult %v \n", GetCollectionsResult)
	fmt.Printf("\n GetCollectionsResultErr %v \n", GetCollectionsResultErr)

	var collectionDeleteInput = CollectionDeleteInput{
		Collections: []string{orderCollectionName},
	}

	var DeleteCollectionResult, DeleteCollectionResultErr = db.DeleteCollections(collectionDeleteInput)
	fmt.Printf("\n DeleteCollectionResult %v \n", DeleteCollectionResult)
	fmt.Printf("\n DeleteCollectionResultErr %v \n", DeleteCollectionResultErr)

	var GetCollectionsResult2, GetCollectionsResult2Err = db.GetAll()
	fmt.Printf("\n GetCollectionsResult %v \n", GetCollectionsResult2)
	fmt.Printf("\n GetCollectionsResult2Err %v \n", GetCollectionsResult2Err)

	var GetCollectionStatsResult2, GetCollectionStatsResult2Err = db.GetCollectionStats(userCollectionName)
	fmt.Printf("\n GetCollectionStatsResult2 %v \n", GetCollectionStatsResult2)
	fmt.Printf("\n GetCollectionStatsResult2Err %v \n", GetCollectionStatsResult2Err)

	fmt.Printf("\n db.Collections %v \n", db.Collections)
	var userCollection *Collection = db.Collections[userCollectionName]

	if userCollection != nil {

		user1 := make(Document)
		user1["name"] = "Nandakumar"
		user1["city"] = "Chennai"
		user1["pincode"] = "600100"

		var DocumentCreateResult, DocumentCreateResultErr = userCollection.Create(user1)
		fmt.Printf("\n DocumentCreateResult %v \n", DocumentCreateResult)
		fmt.Printf("\n DocumentCreateResultErr %v \n", DocumentCreateResultErr)

		var GetCollectionStatsResult3, GetCollectionStatsResult3Err = db.GetCollectionStats(userCollectionName)
		fmt.Printf("\n GetCollectionStatsResult3 %v \n", GetCollectionStatsResult3)
		fmt.Printf("\n GetCollectionStatsResult3Err %v \n", GetCollectionStatsResult3Err)

		if _, exists := DocumentCreateResult.Data["docId"]; exists {

			var docId = DocumentCreateResult.Data["docId"].(string)

			var DocumentReadResult, DocumentReadResultErr = userCollection.Read(docId)
			fmt.Printf("\n DocumentReadResult %v \n", DocumentReadResult)
			fmt.Printf("\n DocumentReadResultErr %v \n", DocumentReadResultErr)

			user2 := make(Document)
			user2["name"] = "kumar"
			user2["city"] = "Chennai"
			user2["pincode"] = "600101"

			var DocumentCreateResult2, DocumentCreateResult2Err = userCollection.Create(user2)
			fmt.Printf("\n DocumentCreateResult2 %v \n", DocumentCreateResult2)
			fmt.Printf("\n DocumentCreateResult2Err %v \n", DocumentCreateResult2Err)

			var filter MapInterface = MapInterface{
				"name": "Nandakumar",
			}

			var DocumentFilterResult, DocumentFilterResultErr = userCollection.Filter(filter)
			fmt.Printf("\n DocumentFilterResult %v \n", DocumentFilterResult)
			fmt.Printf("\n DocumentFilterResultErr %v \n", DocumentFilterResultErr)

			user1["designation"] = "developer"

			var DocumentUpdateResult, DocumentUpdateResultErr = userCollection.Update(docId, user1)
			fmt.Printf("\n DocumentUpdateResult %v \n", DocumentUpdateResult)
			fmt.Printf("\n DocumentUpdateResultErr %v \n", DocumentUpdateResultErr)

			var DocumentReadResult3, DocumentReadResult3Err = userCollection.Read(docId)
			fmt.Printf("\n DocumentReadResult3 %v \n", DocumentReadResult3)
			fmt.Printf("\n DocumentReadResult3Err %v \n", DocumentReadResult3Err)

			var DocumentDeleteResult, DocumentDeleteResultErr = userCollection.Delete(docId)
			fmt.Printf("\n DocumentDeleteResult %v \n", DocumentDeleteResult)
			fmt.Printf("\n DocumentDeleteResultErr %v \n", DocumentDeleteResultErr)

			var DocumentReadResult4, DocumentReadResult4Err = userCollection.Read(docId)
			fmt.Printf("\n DocumentReadResult4 %v \n", DocumentReadResult4)
			fmt.Printf("\n DocumentReadResult4Err %v \n", DocumentReadResult4Err)

			var DocumentGetAllResult, DocumentGetAllResultErr = userCollection.GetAll()
			fmt.Printf("\n DocumentGetAllResult %v \n", DocumentGetAllResult)
			fmt.Printf("\n DocumentGetAllResultErr %v \n", DocumentGetAllResultErr)
		}

	} else {
		println("User collection nil")
	}

	// ------------------------------------------------------------------------------
	var DeleteDatabaseResult, DeleteDatabaseResultErr = db.DeleteDatabase(DatabaseName)
	fmt.Printf("\n DeleteDatabaseResult %v ", DeleteDatabaseResult)
	fmt.Printf("\n DeleteDatabaseResultErr %v ", DeleteDatabaseResultErr)
}
