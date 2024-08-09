package gnosql_client

import (
	"log"

	pb "github.com/nanda03dev/gnosql_client/proto"
	"google.golang.org/grpc"
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

type Client struct {
	URI        string // Ex: http://localhost:5454
	GrpcClient pb.GnoSQLServiceClient
	DB         map[string]*Database
}

// Create new GnoSQL client,
// URI string  Ex: http://localhost:5454
func Connect(URI string, databaseName string, isgRPC bool) *Database {
	var client = &Client{
		URI: URI,
		DB:  make(map[string]*Database),
	}

	conn, err := grpc.Dial(URI, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	} else {
		log.Println("conected to GNOSQL - gRPC Server")
	}

	client.GrpcClient = pb.NewGnoSQLServiceClient(conn)

	client.Connect(databaseName, []CollectionInput{})

	return client.DB[databaseName]
}
