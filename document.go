package gnosql_client

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Create(document Document) DocumentCreateResult {
	var result DocumentCreateResult

	requestBody := DocumentCreateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Document:       document,
	}

	result = GRPC_Create_Document(collection, requestBody)

	return result
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Read(docId string) DocumentReadResult {
	var result DocumentReadResult

	requestBody := DocumentReadRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		DocId:          docId,
	}

	result = GRPC_Read_Document(collection, requestBody)

	return result
}

// return { Data : []Document, Error: "Error message" }, error
func (collection *Collection) Filter(filter MapInterface) DocumentFilterResult {
	var result DocumentFilterResult

	requestBody := DocumentFilterRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Filter:         filter,
	}

	result = GRPC_Filter_Document(collection, requestBody)

	return result
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Update(docId string, document Document) DocumentUpdateResult {
	var result DocumentUpdateResult

	requestBody := DocumentUpdateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Document:       document,
		DocId:          docId,
	}

	result = GRPC_Update_Document(collection, requestBody)

	return result
}

// return { Data : "Success Message", Error: "Error message" }, error
func (collection *Collection) Delete(docId string) DocumentDeleteResult {
	var result DocumentDeleteResult

	requestBody := DocumentDeleteRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		DocId:          docId,
	}

	result = GRPC_Delete_Document(collection, requestBody)

	return result
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) GetAll() DocumentGetAllResult {
	var result DocumentGetAllResult

	requestBody := DocumentGetAllRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
	}

	result = GRPC_GetAll_Document(collection, requestBody)

	return result
}
