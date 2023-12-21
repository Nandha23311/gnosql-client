package gnosql_client

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Create(document Document) (DocumentCreateResult, error) {

	path := fmt.Sprintf("%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName)

	restyResp, restyErr := resty.New().
		R().
		SetBody(document).
		Post(path)

	var result DocumentCreateResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	error := ValidateResponse(restyErr, UnMarshallErr)

	return result, error
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Read(id string) (DocumentReadResult, error) {

	path := fmt.Sprintf("%s/%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName, id)

	restyResp, restyErr := resty.New().
		R().
		Get(path)

	var result DocumentReadResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	error := ValidateResponse(restyErr, UnMarshallErr)

	return result, error
}

// return { Data : []Document, Error: "Error message" }, error
func (collection *Collection) Filter(filter DocumentFilterQuery) (DocumentFilterResult, error) {

	path := fmt.Sprintf("%s/%s/%s/%s/filter", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName)

	restyResp, restyErr := resty.New().
		R().
		SetBody(filter).
		Post(path)

	var result DocumentFilterResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	error := ValidateResponse(restyErr, UnMarshallErr)

	return result, error
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Update(id string, document Document) (DocumentUpdateResult, error) {

	path := fmt.Sprintf("%s/%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName, id)

	restyResp, restyErr := resty.New().
		R().
		SetBody(document).
		Put(path)

	var result DocumentUpdateResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	error := ValidateResponse(restyErr, UnMarshallErr)

	return result, error
}

// return { Data : "Success Message", Error: "Error message" }, error
func (collection *Collection) Delete(id string) (DocumentDeleteResult, error) {

	path := fmt.Sprintf("%s/%s/%s/%s/%s", collection.URI, EndpointsMap.Document, collection.DBName, collection.CollectionName, id)

	restyResp, restyErr := resty.New().
		R().
		Delete(path)

	var result DocumentDeleteResult

	var UnMarshallErr = json.Unmarshal(restyResp.Body(), &result)

	error := ValidateResponse(restyErr, UnMarshallErr)

	return result, error
}
