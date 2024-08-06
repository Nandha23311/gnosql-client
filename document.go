package gnosql_client

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Create(document Document) (DocumentCreateResult, error) {
	requestBody := DocumentCreateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Document:       document,
	}

	result, err := GRPC_Create_Document(collection, requestBody)
	return result, err
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Read(docId string) (DocumentReadResult, error) {
	requestBody := DocumentReadRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		DocId:          docId,
	}

	result, err := GRPC_Read_Document(collection, requestBody)

	return result, err
}

// return { Data : []Document, Error: "Error message" }, error
func (collection *Collection) Filter(filter MapInterface) (DocumentFilterResult, error) {
	requestBody := DocumentFilterRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Filter:         filter,
	}

	result, err := GRPC_Filter_Document(collection, requestBody)

	return result, err
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) Update(docId string, document Document) (DocumentUpdateResult, error) {
	requestBody := DocumentUpdateRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		Document:       document,
		DocId:          docId,
	}

	result, err := GRPC_Update_Document(collection, requestBody)

	return result, err
}

// return { Data : "Success Message", Error: "Error message" }, error
func (collection *Collection) Delete(docId string) (DocumentDeleteResult, error) {
	requestBody := DocumentDeleteRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
		DocId:          docId,
	}

	result, err := GRPC_Delete_Document(collection, requestBody)

	return result, err
}

// return { Data : Document, Error: "Error message" }, error
func (collection *Collection) GetAll() (DocumentGetAllResult, error) {
	requestBody := DocumentGetAllRequest{
		DatabaseName:   collection.DBName,
		CollectionName: collection.CollectionName,
	}

	result, err := GRPC_GetAll_Document(collection, requestBody)

	return result, err
}
