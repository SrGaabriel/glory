package gateway

import (
	"context"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Gateway struct {
	url                 url.URL
	connectionMutex     sync.RWMutex
	websocketConnection *websocket.Conn

	ByteChannel chan []byte

	context       context.Context
	contextCancel context.CancelFunc
}

func NewGateway() Gateway {
	gateway := Gateway{
		url:         url.URL{Scheme: "ws", Host: "api.guilded.gg", Path: "/v1/websocket"},
		ByteChannel: make(chan []byte),
	}
	gateway.context, gateway.contextCancel = context.WithCancel(context.Background())
	return gateway
}

func (gateway *Gateway) Start(token string) {
	gateway.Connect(token)
	go gateway.Listen()
}

func (gateway *Gateway) Connect(token string) {
	gateway.connectionMutex.Lock()
	defer gateway.connectionMutex.Unlock()

	connectionHeader := http.Header{}
	connectionHeader.Add("Authorization", "Bearer "+token)

	ws, _, err := websocket.DefaultDialer.Dial(gateway.url.String(), connectionHeader)
	if err != nil {
		return
	}
	gateway.websocketConnection = ws
}

func (gateway *Gateway) Listen() {
	ticker := time.NewTimer(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-gateway.context.Done():
			return
		case <-ticker.C:
			for {
				if gateway.websocketConnection == nil {
					return
				}
				_, byteMessage, err := gateway.websocketConnection.ReadMessage()
				if err != nil {
					break
				}
				gateway.ByteChannel <- byteMessage
			}
		}
	}
}
