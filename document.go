package gnosql_client

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Create(document Document) DocumentCreateResult {
	var result DocumentCreateResult

	if collection.IsgRPC {
		result = GRPC_Create_Document(collection, document)
	} else {
		result = REST_Create_Document(collection, document)
	}

	return result
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Read(id string) DocumentReadResult {
	var result DocumentReadResult

	if collection.IsgRPC {
		result = GRPC_Read_Document(collection, id)
	} else {
		result = REST_Read_Document(collection, id)
	}

	return result
}

// return { Data : []Document, Error: "Error message" }, error
func (collection *Collection) Filter(filter DocumentFilterQuery) DocumentFilterResult {
	var result DocumentFilterResult

	if collection.IsgRPC {
		result = GRPC_Filter_Document(collection, filter)
	} else {
		result = REST_Filter_Document(collection, filter)
	}
	return result
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Update(id string, document Document) DocumentUpdateResult {
	var result DocumentUpdateResult

	if collection.IsgRPC {
		result = GRPC_Update_Document(collection, id, document)
	} else {
		result = REST_Update_Document(collection, id, document)
	}

	return result
}

// return { Data : "Success Message", Error: "Error message" }, error
func (collection *Collection) Delete(id string) DocumentDeleteResult {
	var result DocumentDeleteResult

	if collection.IsgRPC {
		result = GRPC_Delete_Document(collection, id)
	} else {
		result = REST_Delete_Document(collection, id)
	}
	return result
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) GetAll() DocumentGetAllResult {
	var result DocumentGetAllResult

	if collection.IsgRPC {
		result = GRPC_GetAll_Document(collection)
	} else {
		result = REST_GetAll_Document(collection)
	}

	return result
}
