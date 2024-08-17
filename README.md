# GnoSQL Client for GnoSQL In-Memory Database

GnoSQL Client is a Go library that provides a client for interacting with the GnoSQL in-memory database. It simplifies the process of managing connections, executing queries, and working with data in a GnoSQL database.

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Features

-   **Connection Management:** Easily connect to a Gnosql in-memory database.
-   **Query Execution:** Execute queries and retrieve results.
-   **Data Manipulation:** Perform CRUD operations on your Gnosql database.
-   **Documentation:** Well-documented code and usage examples.

## Installation

```bash
go get -u github.com/nanda03dev/gnosql_client
```

## Usage

### 1. Connect to the Database

First, connect to the GnoSQL database using the GRPC URI and the desired database name.

```go
var DatabaseName = "test-g-c"

// Connect to the GnoSQL database
db := Connect(GRPC_URI, DatabaseName, true)
```

### 2. Create a Collection

Define the collection and its index keys, and create it in the database.

```go
var userCollectionName = "users"

// Define the collection input with index keys
UserCollectionInput := CollectionInput{
	CollectionName: userCollectionName,
	IndexKeys:      []string{"city", "pincode"},
}

// Create the collection
db.CreateCollections([]CollectionInput{UserCollectionInput})
```

### 3. Perform CRUD Operations on Documents

Get the collection, and then perform Create, Read, Update, and Delete operations on documents within the collection.

#### a. Create a Document

```go
// Get the collection
var userCollection *Collection = db.Collections[userCollectionName]

// Create a document
user1 := make(Document)
user1["name"] = "Nandakumar"
user1["city"] = "Chennai"
user1["pincode"] = "600100"

// Insert the document into the collection
userCollection.Create(user1)
```

#### b. Read a Document

```go
// Read the document by its ID
var docId = "your-document-id" // Replace with your actual document ID
userCollection.Read(docId)
```

#### c. Update a Document

```go
// Update the document
user1["designation"] = "developer"
userCollection.Update(docId, user1)
```

#### d. Filter Documents

```go
// Define a filter to find documents by a specific field
var filter MapInterface = MapInterface{
	"name": "Nandakumar",
}

// Filter documents in the collection
userCollection.Filter(filter)
```

#### e. Delete a Document

```go
// Delete the document by its ID
userCollection.Delete(docId)
```

### 4. Delete a Collection

Delete a specific collection from the database.

```go
// Define the collection to be deleted
var collectionDeleteInput = CollectionDeleteInput{
	Collections: []string{userCollectionName},
}

// Delete the collection
db.DeleteCollections(collectionDeleteInput)
```