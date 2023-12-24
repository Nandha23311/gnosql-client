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

func REST_Create_DB(client *Client, requestBody DatabaseCreateRequest) DatabaseCreateResult {
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

func REST_Delete_DB(client *Client, requestBody DatabaseDeleteRequest) DatabaseDeleteResult {
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

func REST_Create_Collections(database *Database, requestBody CollectionCreateRequest) CollectionCreateResult {
	var result CollectionCreateResult

	path := fmt.Sprintf("%s/%s", database.URI, EndpointsMap.CollectionAdd)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Delete_Collections(database *Database, requestBody CollectionDeleteRequest) CollectionDeleteResult {
	var result CollectionDeleteResult

	path := fmt.Sprintf("%s/%s", database.URI, EndpointsMap.CollectionDelete)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_GetAll_Collections(database *Database, requestBody CollectionGetAllRequest) CollectionGetAllResult {
	var result CollectionGetAllResult

	path := fmt.Sprintf("%s/%s", database.URI, EndpointsMap.CollectionGetAll)

	restyResp, restyErr := resty.New().R().SetBody(requestBody).Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Get_Collection_Stats(database *Database, requestBody CollectionStatsRequest) CollectionStatsResult {
	var result CollectionStatsResult

	path := fmt.Sprintf("%s/%s", database.URI, EndpointsMap.CollectionStats)

	restyResp, restyErr := resty.New().R().SetBody(requestBody).Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Create_Document(collection *Collection, requestBody DocumentCreateRequest) DocumentCreateResult {

	var result DocumentCreateResult

	path := fmt.Sprintf("%s/%s", collection.URI, EndpointsMap.DocumentAdd)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Read_Document(collection *Collection, requestBody DocumentReadRequest) DocumentReadResult {
	var result DocumentReadResult

	path := fmt.Sprintf("%s/%s", collection.URI, EndpointsMap.DocumentRead)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Filter_Document(collection *Collection, requestBody DocumentFilterRequest) DocumentFilterResult {
	var result DocumentFilterResult

	path := fmt.Sprintf("%s/%s", collection.URI, EndpointsMap.DocumentFilter)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Update_Document(collection *Collection, requestBody DocumentUpdateRequest) DocumentUpdateResult {

	var result DocumentUpdateResult

	path := fmt.Sprintf("%s/%s", collection.URI, EndpointsMap.DocumentUpdate)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_Delete_Document(collection *Collection, requestBody DocumentDeleteRequest) DocumentDeleteResult {
	var result DocumentDeleteResult

	path := fmt.Sprintf("%s/%s", collection.URI, EndpointsMap.DocumentDelete)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}

func REST_GetAll_Document(collection *Collection, requestBody DocumentGetAllRequest) DocumentGetAllResult {
	var result DocumentGetAllResult

	path := fmt.Sprintf("%s/%s", collection.URI, EndpointsMap.DocumentGetAll)

	restyResp, restyErr := resty.New().
		R().
		SetBody(requestBody).
		Post(path)

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	result.Error = ValidateResponse(restyErr, UnMarshallErr, nil, "")

	return result
}
