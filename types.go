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
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

type DatabaseCreateRequest struct {
	DatabaseName string            `json:"databaseName"`
	Collections  []CollectionInput `json:"collections"`
}

type DatabaseCreateResult struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

type DatabaseResult struct {
	DatabaseName string   `json:"databaseName"`
	Collections  []string `json:"collections"`
}

type DatabaseConnectResult struct {
	Data  DatabaseResult `json:"data"`
	Error string         `json:"error"`
}

type DatabaseDeleteRequest struct {
	DatabaseName string `json:"databaseName"`
}
type DatabaseDeleteResult struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

type DatabaseGetAllResult struct {
	Data  []string `json:"data"`
	Error string   `json:"error"`
}

type DatabaseLoadToDiskResult struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

type CollectionCreateRequest struct {
	DatabaseName string            `json:"databaseName"`
	Collections  []CollectionInput `json:"collections"`
}

type CollectionCreateResult struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

type CollectionDeleteRequest struct {
	DatabaseName string   `json:"databaseName"`
	Collections  []string `json:"collections"`
}

type CollectionDeleteResult struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

type CollectionGetAllRequest struct {
	DatabaseName string `json:"databaseName"`
}

type CollectionGetAllResult struct {
	Data  []string `json:"data"`
	Error string   `json:"error"`
}

type CollectionStatsRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
}

type CollectionStatsResult struct {
	Data  CollectionStats
	Error string `json:"error"`
}

type DocumentCreateRequest struct {
	DatabaseName   string   `json:"databaseName"`
	CollectionName string   `json:"collectionName"`
	Document       Document `json:"document"`
}

type DocumentCreateResult struct {
	Data  Document `json:"data"`
	Error string   `json:"error"`
}

type DocumentReadRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
	Id             string `json:"docId"`
}

type DocumentReadResult struct {
	Data  Document `json:"data"`
	Error string   `json:"error"`
}

type DocumentFilterRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
	Filter         MapInterface
}

type DocumentFilterResult struct {
	Data  []Document `json:"data"`
	Error string     `json:"error"`
}

type DocumentUpdateRequest struct {
	DatabaseName   string   `json:"databaseName"`
	CollectionName string   `json:"collectionName"`
	Id             string   `json:"docId"`
	Document       Document `json:"document"`
}

type DocumentUpdateResult struct {
	Data  Document `json:"data"`
	Error string   `json:"error"`
}

type DocumentDeleteRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
	Id             string `json:"docId"`
}

type DocumentDeleteResult struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

type DocumentGetAllRequest struct {
	DatabaseName   string `json:"databaseName"`
	CollectionName string `json:"collectionName"`
}

type DocumentGetAllResult struct {
	Data  []Document `json:"data"`
	Error string     `json:"error"`
}
