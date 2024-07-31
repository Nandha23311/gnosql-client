package gnosql_client

import (
	"github.com/go-resty/resty/v2"
	pb "github.com/nanda03dev/gnosql_client/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type Endpoints struct {
	DatabaseGetAll   string
	DatabaseAdd      string
	DatabaseConnect  string
	DatabaseDelete   string
	CollectionAdd    string
	CollectionDelete string
	CollectionGetAll string
	CollectionStats  string
	DocumentAdd      string
	DocumentDelete   string
	DocumentRead     string
	DocumentFilter   string
	DocumentUpdate   string
	DocumentGetAll   string
}

var EndpointsMap = Endpoints{
	DatabaseGetAll:   "database/get-all",
	DatabaseAdd:      "database/add",
	DatabaseConnect:  "database/connect",
	DatabaseDelete:   "database/delete",
	CollectionAdd:    "collection/add",
	CollectionDelete: "collection/delete",
	CollectionGetAll: "collection/get-all",
	CollectionStats:  "collection/stats",
	DocumentAdd:      "document/add",
	DocumentDelete:   "document/delete",
	DocumentRead:     "document/find",
	DocumentFilter:   "document/filter",
	DocumentUpdate:   "document/update",
	DocumentGetAll:   "document/get-all",
}

type Client struct {
	URI        string // Ex: http://localhost:5454
	IsgRPC     bool
	GrpcClient pb.GnoSQLServiceClient
	DB         map[string]*Database
	RestClient *resty.Client
}

// Create new GnoSQL client,
// URI string  Ex: http://localhost:5454
func Connect(URI string, databaseName string, isgRPC bool) *Database {
	var client = &Client{
		URI:    URI,
		IsgRPC: isgRPC,
		DB:     make(map[string]*Database),
	}

	if isgRPC {
		conn, err := grpc.Dial(URI, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect : %v", err)
		} else {
			log.Println("conected to gRPC Server")
		}

		client.GrpcClient = pb.NewGnoSQLServiceClient(conn)

	} else {
		restClient := resty.New().SetTransport(&http.Transport{
			MaxIdleConns:        5,
			MaxIdleConnsPerHost: 5,
			MaxConnsPerHost:     10,
		})
		client.RestClient = restClient
		log.Println("conected to Http Server")
	}

	client.Connect(databaseName, []CollectionInput{})

	return client.DB[databaseName]
}
