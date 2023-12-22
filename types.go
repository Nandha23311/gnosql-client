package gnosql_client

import (
	pb "github.com/nanda03dev/gnosql_client/proto"
)

type ReqBody map[string]interface{}

// Collection Types
type Collection struct {
	URI            string
	CollectionName string
	DBName         string
	IsgRPC         bool
	ClientgRPC     pb.GnoSQLServiceClient
}

type CollectionInput struct {
	CollectionName string
	IndexKeys      []string
}

type CollectionDeleteInput struct {
	Collections []string
}

type CollectionStats struct {
	CollectionName string
	IndexKeys      []string
	Documents      int32
}

type CollectionStatsResult struct {
	Data  CollectionStats
	Error string
}

type CollectionCreateResult struct {
	Data  string // Data : collection created succesfully
	Error string
}

type CollectionGetAllResult struct {
	Data  []string // List of collection names [collection1, collection2...]
	Error string
}

type CollectionDeleteResult struct {
	Data  string // Data : collection deleted succesfully
	Error string
}

// Database Types

type Database struct {
	URI         string
	DBName      string
	IsgRPC      bool
	ClientgRPC  pb.GnoSQLServiceClient
	Collections map[string]*Collection
}

type DatabaseCreateInput struct {
	DatabaseName string
}

type DatabaseCreateResult struct {
	Data  string // Data : database created succesfully
	Error string
}

type DatabaseGetAllResult struct {
	Data  []string // List of database names [DatabaseName1, DatabaseName2...]
	Error string
}

type DatabaseDeleteResult struct {
	Data  string // Data : database deleted succesfully
	Error string
}

// Document types
type Document map[string]interface{}

type DocumentCreateResult struct {
	Data  Document // Data : Document
	Error string
}

type DocumentUpdateResult struct {
	Data  Document // Data : Document
	Error string
}

type DocumentReadResult struct {
	Data  Document // Data : Document
	Error string
}

type DocumentFilterResult struct {
	Data  []Document // List of documents [Document, Document...]
	Error string
}
type DocumentFilterQuery map[string]interface{}

type DocumentDeleteResult struct {
	Data  string // Data : document deleted succesfully
	Error string
}
