package gnosql_client

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func REST_GetAll_DB(client *Client) DatabaseGetAllResult {
	var result DatabaseGetAllResult

	path := fmt.Sprintf("%s/%s", client.URI, EndpointsMap.DatabaseGetAll)

	restyResp, restyErr := resty.New().R().Get(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result

}

func REST_Create_DB(client *Client, requestBody interface{}) DatabaseCreateResult {
	var result DatabaseCreateResult

	path := fmt.Sprintf("%s/%s", client.URI, EndpointsMap.DatabaseAdd)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Delete_DB(client *Client, requestBody interface{}) DatabaseDeleteResult {
	var result DatabaseDeleteResult

	path := fmt.Sprintf("%s/%s", client.URI, EndpointsMap.DatabaseDelete)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Create_Collections(database *Database, requestBody interface{}) CollectionCreateResult {
	var result CollectionCreateResult

	path := fmt.Sprintf("%s/%s/%s/add", database.URI, EndpointsMap.Collection, database.DBName)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Delete_Collections(database *Database, requestBody interface{}) CollectionDeleteResult {
	var result CollectionDeleteResult

	path := fmt.Sprintf("%s/%s/%s/delete", database.URI, EndpointsMap.Collection, database.DBName)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Delete(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_GetAll_Collections(database *Database) CollectionGetAllResult {
	var result CollectionGetAllResult

	path := fmt.Sprintf("%s/%s/%s/get-all", database.URI, EndpointsMap.Collection, database.DBName)

	restyResp, restyErr := resty.New().R().Get(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_GetAll_Collection_Stats(database *Database, collectionName string) CollectionStatsResult {
	var result CollectionStatsResult

	path := fmt.Sprintf("%s/%s/%s/%s/stats", database.URI, EndpointsMap.Collection, database.DBName, collectionName)

	restyResp, restyErr := resty.New().R().Get(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Create_Document(collection *Collection, document Document) DocumentCreateResult {

	var result DocumentCreateResult

	path := fmt.Sprintf("%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName)

	restyResp, restyErr := resty.New().
		R().
		SetBody(document).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Read_Document(collection *Collection, id string) DocumentReadResult {
	var result DocumentReadResult

	path := fmt.Sprintf("%s/%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName, id)

	restyResp, restyErr := resty.New().
		R().
		Get(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Filter_Document(collection *Collection, filter DocumentFilterQuery) DocumentFilterResult {
	var result DocumentFilterResult

	path := fmt.Sprintf("%s/%s/%s/%s/filter", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName)

	restyResp, restyErr := resty.New().
		R().
		SetBody(filter).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Update_Document(collection *Collection, id string, document Document) DocumentUpdateResult {

	var result DocumentUpdateResult

	path := fmt.Sprintf("%s/%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName, id)

	restyResp, restyErr := resty.New().
		R().
		SetBody(document).
		Put(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Delete_Document(collection *Collection, id string) DocumentDeleteResult {
	var result DocumentDeleteResult

	path := fmt.Sprintf("%s/%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName, id)

	restyResp, restyErr := resty.New().
		R().
		Delete(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_GetAll_Document(collection *Collection) DocumentGetAllResult {
	var result DocumentGetAllResult

	path := fmt.Sprintf("%s/%s/%s/%s/get-all", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName)

	restyResp, restyErr := resty.New().
		R().
		Get(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}
