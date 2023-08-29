package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Ws(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("Failed to set websocket upgrade: %+v\n", err)
		return
	}
	defer conn.Close()
	type Message struct {
		Type string      `json:"type"`
		Data interface{} `json:"data"`
	}
	ticker := time.NewTicker(1 * time.Second) // 每秒发送一次消息
	defer ticker.Stop()
	if err != nil {
		log.Println("Failed to marshal message to JSON:", err)
		return
	}
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("Received message: %s\n", message)
		}
	}()
	for {
		select {
		case <-ticker.C:
			randNum := rand.Intn(100)
			msg := Message{
				Type: "random",
				Data: randNum,
			}
			data, err := json.Marshal(msg)
			if err != nil {
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return
			}
		default:

			fmt.Println("No message received")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func Wstest(c *gin.Context) {
	type DataAddr struct {
		Address string
	}
	req := new(DataAddr)
	c.BindJSON(req)
	fmt.Println(req)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	//utils.XlsxToJson()
	if err != nil {
		log.Println("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	type Message struct {
		Type string      `json:"type"`
		Data interface{} `json:"data"`
	}
	ticker := time.NewTicker(1 * time.Second) // 每秒发送一次消息
	defer ticker.Stop()
	if err != nil {
		log.Println("Failed to marshal message to JSON:", err)
		return
	}
	i := 0
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("Received message: %s\n", message)
		select {
		case <-ticker.C:
			randNum := rand.Intn(100)
			msg := Message{
				Type: "random",
				Data: randNum,
			}
			data, err := json.Marshal(msg)
			if err != nil {
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return
			}
			i++
		}
		if i > 100 {
			ticker.Stop()
		}

	}

}
