package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Accepting all requests
	},
}

type Server struct {
	server      *http.Server
	connections map[*websocket.Conn]bool
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	s.connections[connection] = true
	defer func() {
		logrus.Debug("Closing websocket connection")
		connection.Close()
		delete(s.connections, connection)
	}()

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			logrus.Debug("Websocket connection closed by client")
			break // Exit the loop if the client tries to close the connection or the connection is interrupted
		}
		if "__ping__" == string(message) {
			logrus.Debug("Received ping message from client")
			s.WriteMessage(connection, []byte("__pong__"))
			continue
		}
		go s.handleMessage(connection, message)
	}
}

func (s *Server) WriteMessage(conn *websocket.Conn, message []byte) {
	logrus.Debug("Sending message to client")
	conn.WriteMessage(websocket.TextMessage, message)
}

func (s Server) BroadcastMessage(message []byte) {
	for conn := range s.connections {
		s.WriteMessage(conn, message)
	}
}

func StartServer() *Server {
	server := Server{
		connections: make(map[*websocket.Conn]bool),
		server: &http.Server{
			Addr: ":8081",
		},
	}
	server.server.RegisterOnShutdown(func() {
		logrus.Info("Websocket server shutdown")
	})

	http.HandleFunc("/ws", server.handler)

	go func() {
		if err := server.server.ListenAndServe(); err != nil {
			logrus.Error(err)
		}
	}()
	logrus.Info("Websocket server started on port 8081")
	return &server
}

func (s *Server) Close(ctx context.Context) {
	s.server.Shutdown(ctx)
}
