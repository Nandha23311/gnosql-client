package gnosql_client

import (
	"github.com/go-resty/resty/v2"
	pb "github.com/nanda03dev/gnosql_client/proto"
)

type ReqBody map[string]interface{}
type MapInterface map[string]interface{}

// Collection Types
type Collection struct {
	URI            string
	CollectionName string
	DBName         string
	IsgRPC         bool
	GrpcClient     pb.GnoSQLServiceClient
	RestClient     *resty.Client
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

// Database Types

type Database struct {
	URI         string
	DBName      string
	IsgRPC      bool
	GrpcClient  pb.GnoSQLServiceClient
	Collections map[string]*Collection
	RestClient  *resty.Client
}

type DatabaseCreateInput struct {
	DatabaseName string
}

type Document map[string]interface{}

type DocumentFilterQuery map[string]interface{}

type Result struct {
	Data interface{} `json:"data"`
}

type DatabaseCreateRequest struct {
	DatabaseName string            `json:"databaseName"`
	Collections  []CollectionInput `json:"collections"`
}

type DatabaseCreateResult struct {
	Data string `json:"data"`
}

type DatabaseResult struct {
	DatabaseName string   `json:"databaseName"`
	Collections  []string `json:"collections"`
}

type DatabaseConnectResult struct {
	Data DatabaseResult `json:"data"`
}

type DatabaseDeleteRequest struct {
	DatabaseName string `json:"databaseName"`
}
type DatabaseDeleteResult struct {
	Data string `json:"data"`
}

type DatabaseGetAllResult struct {
	Data []string `json:"data"`
}

type DatabaseLoadToDiskResult struct {
	Data string `json:"data"`
}

type CollectionCreateRequest struct {
	DatabaseName string            `json:"databaseName"`
	Collections  []CollectionInput `json:"collections"`
}

type CollectionCreateResult struct {
	Data string `json:"data"`
}

type CollectionDeleteRequest struct {
	DatabaseName string   `json:"databaseName"`
	Collections  []string `json:"collections"`
}

type CollectionDeleteResult struct {
	Data string `json:"data"`
}

type CollectionGetAllRequest struct {
	DatabaseName string `json:"databaseName"`
}

type CollectionGetAllResult struct {
	Data []string `json:"data"`
}

type CollectionStatsRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
}

type CollectionStatsResult struct {
	Data CollectionStats
}

type DocumentCreateRequest struct {
	DatabaseName   string   `json:"databaseName"`
	CollectionName string   `json:"collectionName"`
	Document       Document `json:"document"`
}

type DocumentCreateResult struct {
	Data Document `json:"data"`
}

type DocumentReadRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
	DocId          string `json:"docId"`
}

type DocumentReadResult struct {
	Data Document `json:"data"`
}

type DocumentFilterRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
	Filter         MapInterface
}

type DocumentFilterResult struct {
	Data []Document `json:"data"`
}

type DocumentUpdateRequest struct {
	DatabaseName   string   `json:"databaseName"`
	CollectionName string   `json:"collectionName"`
	DocId          string   `json:"docId"`
	Document       Document `json:"document"`
}

type DocumentUpdateResult struct {
	Data Document `json:"data"`
}

type DocumentDeleteRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
	DocId          string `json:"docId"`
}

type DocumentDeleteResult struct {
	Data string `json:"data"`
}

type DocumentGetAllRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
}

type DocumentGetAllResult struct {
	Data []Document `json:"data"`
}
