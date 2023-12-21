package gnosql_client

type Endpoints struct {
	DatabaseGetAll string
	DatabaseAdd    string
	DatabaseDelete string
	Collection     string
	Document       string
}

var EndpointsMap = Endpoints{
	DatabaseGetAll: "database/get-all",
	DatabaseAdd:    "database/add",
	DatabaseDelete: "database/delete",
	Collection:     "collection",
	Document:       "document",
}

type Client struct {
	URI string // Ex: http://localhost:5454
	DB  map[string]*Database
}

// Create new GnoSQL client, 
// URI string  Ex: http://localhost:5454
func Connect(URI string) *Client {
	return &Client{
		URI: URI,
		DB:  make(map[string]*Database),
	}
}
