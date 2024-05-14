package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	// Define an upgrader
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Define a WebSocket handler function
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading connection:", err)
			return
		}
		defer conn.Close()

		for {
			// Read message from the WebSocket client
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}

			// Print the received message
			fmt.Printf("Received message: %s\n", message)

			// Echo the message back to the WebSocket client
			err = conn.WriteMessage(messageType, message)
			if err != nil {
				log.Println("Error writing message:", err)
				break
			}
		}
	})

	// Start the WebSocket server
	log.Println("Starting WebSocket server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
