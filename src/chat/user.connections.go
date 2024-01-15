package chat

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

var connectionMutex = &sync.Mutex{}
var UserConnections = make(map[string]*websocket.Conn)

// Register a new connection in the Map
func AddConnection(userID string, conn *websocket.Conn) {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()
	UserConnections[userID] = conn
}

// Remove a connection from the Map
func RemoveConnection(userID string) {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()
	delete(UserConnections, userID)
}

// Get a connection from the Map
func GetConnection(userID string) (*websocket.Conn, bool) {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()
	conn, ok := UserConnections[userID]
	return conn, ok
}
