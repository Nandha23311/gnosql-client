package gnosql_client

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Create(document Document) DocumentCreateResult {
	var result DocumentCreateResult

	requestBody := DocumentCreateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Document:       document,
	}

	if collection.IsgRPC {
		result = GRPC_Create_Document(collection, requestBody)
	} else {
		result = REST_Create_Document(collection, requestBody)
	}

	return result
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Read(id string) DocumentReadResult {
	var result DocumentReadResult

	requestBody := DocumentReadRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Id:             id,
	}

	if collection.IsgRPC {
		result = GRPC_Read_Document(collection, requestBody)
	} else {
		result = REST_Read_Document(collection, requestBody)
	}

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

	if collection.IsgRPC {
		result = GRPC_Filter_Document(collection, requestBody)
	} else {
		result = REST_Filter_Document(collection, requestBody)
	}
	return result
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Update(id string, document Document) DocumentUpdateResult {
	var result DocumentUpdateResult

	requestBody := DocumentUpdateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Document:       document,
		Id:             id,
	}

	if collection.IsgRPC {
		result = GRPC_Update_Document(collection, requestBody)
	} else {
		result = REST_Update_Document(collection, requestBody)
	}

	return result
}

// return { Data : "Success Message", Error: "Error message" }, error
func (collection *Collection) Delete(id string) DocumentDeleteResult {
	var result DocumentDeleteResult

	requestBody := DocumentDeleteRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Id:             id,
	}

	if collection.IsgRPC {
		result = GRPC_Delete_Document(collection, requestBody)
	} else {
		result = REST_Delete_Document(collection, requestBody)
	}
	return result
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) GetAll() DocumentGetAllResult {
	var result DocumentGetAllResult

	requestBody := DocumentGetAllRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
	}

	if collection.IsgRPC {
		result = GRPC_GetAll_Document(collection, requestBody)
	} else {
		result = REST_GetAll_Document(collection, requestBody)
	}

	return result
}
