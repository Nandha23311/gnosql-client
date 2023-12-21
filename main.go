package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Endpoints struct {
	DatabaseGetAll string
	DatabaseAdd    string
	DatabaseDelete string
	Collection     string
	Document       string
}

var EndpointsMap = Endpoints{
	DatabaseGetAll: "database/get-all",
	DatabaseAdd:    "database/add",
	DatabaseDelete: "database/delete",
	Collection:     "collection",
	Document:       "document",
}

type Client struct {
	URI string
	DB  map[string]*Database
}

type Result struct {
	message string
	error   interface{}
	data    interface{}
}

func GetResult() Result {
	return Result{
		error:   nil,
		message: "",
		data:    nil,
	}
}

func Connect(URI string) *Client {
	return &Client{
		URI: URI,
		DB:  make(map[string]*Database),
	}
}

func (client *Client) Check() interface{} {
	var result Result = GetResult()
	path := client.URI + EndpointsMap.DatabaseGetAll

	restyClient := resty.New()

	resp, err := restyClient.R().Get(path)
	if err != nil {
		fmt.Println("Error:", err)
		result.error = "error"

		return result
	}

	var database []string

	err = json.Unmarshal(resp.Body(), &database)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		result.error = "error"
	} else {
		result.data = database
	}

	return result

}

func main() {
	uri := "http://localhost:5454"
	newClinet := Connect(uri)

	println("newClinet ", newClinet.URI)

	var DatabaseName = "test-g-c"
	var CreateDatabaseResult, _ = newClinet.Create(DatabaseName)
	fmt.Printf("\n CreateDatabaseResult %v \n", CreateDatabaseResult)

	var GetAllDatabaseResult, _ = newClinet.GetAll()
	fmt.Printf("\nGetAllDatabaseResult %v \n", GetAllDatabaseResult.Data)

	var db *Database = newClinet.DB[DatabaseName]
	if db != nil {
		var userCollectionName = "users"
		var orderCollectionName = "orders"

		UserCollectionInput := CollectionInput{
			CollectionName: userCollectionName,
			IndexKeys:      []string{"city", "pincode"},
		}

		OrderCollectionInput := CollectionInput{
			CollectionName: orderCollectionName,
			IndexKeys:      []string{"userId", "category"},
		}

		collectionsInput := []CollectionInput{UserCollectionInput, OrderCollectionInput}

		var CreateCollectionResult, _ = db.CreateCollections(collectionsInput)
		fmt.Printf("\n CreateCollectionResult %v \n", CreateCollectionResult)

		var GetCollectionsResult, _ = db.GetAll()
		fmt.Printf("\n GetCollectionsResult %v \n", GetCollectionsResult)

		var GetCollectionStatsResult, _ = db.GetCollectionStats(userCollectionName)
		fmt.Printf("\n GetCollectionStatsResult %v \n", GetCollectionStatsResult)

		var DeleteCollectionResult, _ = db.DeleteCollections([]string{orderCollectionName})
		fmt.Printf("\n DeleteCollectionResult %v \n", DeleteCollectionResult)

		var userCollection *Collection = db.Collections[userCollectionName]

		if userCollection != nil {

			user1 := make(Document)
			user1["name"] = "Nandakumar"
			user1["city"] = "Chennai"
			user1["pincode"] = "600100"

			var DocumentCreateResult, _ = userCollection.Create(user1)
			fmt.Printf("\n DocumentCreateResult %v \n", DocumentCreateResult)

			user2 := make(Document)
			user2["name"] = "kumar"
			user2["city"] = "Chennai"
			user2["pincode"] = "600101"

			var DocumentCreateResult2, _ = userCollection.Create(user2)
			fmt.Printf("\n DocumentCreateResult2 %v \n", DocumentCreateResult2)

			var id = DocumentCreateResult.Data["id"].(string)

			var DocumentReadResult, _ = userCollection.Read(id)
			fmt.Printf("\n DocumentReadResult %v \n", DocumentReadResult)

			var filter MapInterface = MapInterface{
				"name": "Nandakumar",
			}

			var DocumentFilterResult, _ = userCollection.Filter(filter)
			fmt.Printf("\n DocumentFilterResult %v \n", DocumentFilterResult)

			user1["designation"] = "developer"

			var DocumentUpdateResult, _ = userCollection.Update(id, user1)
			fmt.Printf("\n DocumentUpdateResult %v \n", DocumentUpdateResult)

			var DocumentDeleteResult, _ = userCollection.Delete(id)
			fmt.Printf("\n DocumentDeleteResult %v \n", DocumentDeleteResult)

		}

	}

	var DeleteDatabaseResult, _ = newClinet.Delete(DatabaseName)
	fmt.Printf("\n DeleteDatabaseResult %v ", DeleteDatabaseResult)

}
