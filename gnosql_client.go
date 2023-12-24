package gnosql_client

import (
	pb "github.com/nanda03dev/gnosql_client/proto"
	"log"

	"google.golang.org/grpc"
)

type Endpoints struct {
	DatabaseGetAll   string
	DatabaseAdd      string
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
	ClientgRPC pb.GnoSQLServiceClient
	DB         map[string]*Database
}

// Create new GnoSQL client,
// URI string  Ex: http://localhost:5454
func Connect(URI string, isgRPC bool) *Client {
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

		client.ClientgRPC = pb.NewGnoSQLServiceClient(conn)

	}
	return client
}
