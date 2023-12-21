package gnosql_client

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// return { Data : [DatabaseName1, DatabaseName2...], Error: "Error message" }, error
func (client *Client) GetAll() (DatabaseGetAllResult, error) {
	path := fmt.Sprintf("%s/%s", client.URI, EndpointsMap.DatabaseGetAll)

	restyResp, restyErr := resty.New().R().Get(path)

	var result DatabaseGetAllResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	return result, ValidateResponse(restyErr, UnMarshallErr)
}

// return { Data : "Sucess message", Error: "Error message" }, error
func (client *Client) Create(databaseName string) (DatabaseCreateResult, error) {

	path := fmt.Sprintf("%s/%s", client.URI, EndpointsMap.DatabaseAdd)

	requestBody := DatabaseCreateInput{
		DatabaseName: databaseName,
	}

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var result DatabaseCreateResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	error := ValidateResponse(restyErr, UnMarshallErr)

	if error == nil {
		database := Database{
			DBName:      databaseName,
			URI:         client.URI,
			Collections: make(map[string]*Collection),
		}
		client.DB[databaseName] = &database
	}

	return result, error
}

// return { Data : "Sucess message", Error: "Error message" }, error
func (client *Client) Delete(databaseName string) (DatabaseDeleteResult, error) {

	path := fmt.Sprintf("%s/%s", client.URI, EndpointsMap.DatabaseDelete)

	requestBody := ReqBody{
		"DatabaseName": databaseName,
	}

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var result DatabaseDeleteResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	error := ValidateResponse(restyErr, UnMarshallErr)

	if error == nil {
		if client.DB[databaseName] != nil {
			delete(client.DB, databaseName)
		}
	}

	return result, error
}
