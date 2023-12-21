package gnosql_client

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
