package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		// 解析请求到对应的处理器中
		var request WSRequest
		var response WSResponse
		err = json.Unmarshal(message, &request)
		if err != nil {
			fmt.Printf("ws request json fail: %s", err.Error())
		}

		// 请求处理器
		switch request.Action {
		case "classServer":
			response = ClassServer(message)
		}

		response.RequestID = request.RequestID
		jsonResBody, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("ws response json fail: %s", err.Error())
		}

		err = conn.WriteMessage(mt, jsonResBody)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func aliveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/alive", aliveHandler)
	http.HandleFunc("/ws", wsHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
