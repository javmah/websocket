package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

//  In This Package I Will Create A  Project Base On WEB SOCKET
//  This Project Badically a Chat Applicatoion Base On WEB SOket Technology
//  Help From YouTube : https://www.youtube.com/watch?v=CIh8qN7LO8M

type myStruct struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var upgrader = websocket.Upgrader{}

func main() {
	// fmt.Println("Hello, playground")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../template/index.html")
	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				mType, msg, _ := conn.ReadMessage()

				conn.WriteMessage(mType, msg)
			}
		}(conn)
	})

	http.HandleFunc("/v2/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				_, msg, _ := conn.ReadMessage()
				println(string(msg))
			}
		}(conn)
	})

	http.HandleFunc("/v3/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			ch := time.Tick(10 * time.Second)

			for range ch {
				conn.WriteJSON(myStruct{
					Username:  "mvansickle",
					FirstName: "Michael",
					LastName:  "Van Sickle",
				})
			}
		}(conn)
	})

	http.HandleFunc("/v4/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					conn.Close()
				}
			}
		}(conn)

		go func(conn *websocket.Conn) {
			ch := time.Tick(5 * time.Second)

			for range ch {
				conn.WriteJSON(myStruct{
					Username:  "mvansickle",
					FirstName: "Michael",
					LastName:  "Van Sickle",
				})
			}
		}(conn)
	})

	http.ListenAndServe(":3000", nil)
}
