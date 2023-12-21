package network

import (
	"MultiState-P2P/pkg/protocol"
	"net"
    "encoding/json"
    "fmt"
    "io/ioutil"
)

// NodeState represents the state of a node in the network.
type NodeState string

const (
	Idle    NodeState = "Idle"
	Share   NodeState = "Share"
	Request NodeState = "Request"
	Update  NodeState = "Update"
	Transmit NodeState = "Transmit"
	Dead    NodeState = "Dead" // Inactive state
)

// Node represents a peer in the network.
type Node struct {
	IP         string
	State      NodeState
	Buffer     []protocol.Request // Assuming Request is a struct that handles requests.
	TableH     *TableH // Map of file names to a list of IPs that store them.
	AccessToken string
	Connections map[string]net.Conn
}

// NewNode creates a new node with the given IP and access token.
func NewNode(IP, accessToken string) *Node {
	return &Node{
		IP:         IP,
		State:      Idle,
		Buffer:     make([]protocol.Request, 0),
		TableH:     NewTableH(),
		AccessToken: accessToken,
		Connections: make(map[string]net.Conn),
	}
}

func (n *Node) HandleRequest(conn net.Conn) error {
    // Get the sender's address
    senderAddr := conn.RemoteAddr().String()

    // Split the address into IP and port
    senderIP, senderPort, err := net.SplitHostPort(senderAddr)
    if err != nil {
        return fmt.Errorf("failed to parse remote address: %w", err)
    }

    fmt.Printf("Received connection from %s:%s\n", senderIP, senderPort)
    // Add logic here to handle the request

	jsonData, err := ioutil.ReadAll(conn)
    if err != nil {
        return fmt.Errorf("error reading request data: %w", err)
    }

    // Unmarshal JSON data into a Request struct
    var req protocol.Request
    err = json.Unmarshal(jsonData, &req)
    if err != nil {
        return fmt.Errorf("error unmarshalling request JSON: %w", err)
    }

    // Process the request based on the content of the req object
    // Example:
    fmt.Printf("Received request of type: %s\n", req.Type)
    // Add more handling logic based on your application's requirements

    return nil
}