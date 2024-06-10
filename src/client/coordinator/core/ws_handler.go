package core

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

const (
	WS_URL = "wss://api.witnesschain.com/tracer/v1/watchtower/websocket"
)

type WebsocketClient struct {
	dialer     websocket.Dialer
	Connection *websocket.Conn
}

func (wc WebsocketClient) GetWriteBufferSize() int {
	return wc.dialer.WriteBufferSize
}

func (wc *WebsocketClient) ConnectToCoordinator(headers http.Header) error {
	url, _ := url.Parse(WS_URL)
	wc.dialer = websocket.Dialer{
		WriteBufferSize: 65536,
	}

	connection, response, err := wc.dialer.Dial(url.String(), headers)
	connection.EnableWriteCompression(true)
	if err != nil {
		wtCommon.Error(response) //Note: more info in case ws connection fails
		return err
	}
	wc.Connection = connection
	return nil
}

func (wc *WebsocketClient) CloseConnection() error {
	wtCommon.Info("Disconnecting")
	err := wc.Connection.Close()
	return err
}

func (wc *WebsocketClient) ListenForMessages(dataChannel chan string) error {
	defer func() {
		wtCommon.Warning("Closing listener routine")
		wc.Connection.Close()
	}()

	var lastRequest string = ""
	for {
		msgType, payload, err := wc.Connection.ReadMessage()
		if err != nil {
			return err
		}
		content := string(payload)
		cutoff := len(content)
		if cutoff > 100 {
			cutoff = 100
		}
		if len(content) > 0 && content[0] == '{' {
			if lastRequest == content {
				// wtCommon.Warning("Ignoring repeated request")
			} else {
				wtCommon.Info(fmt.Sprintf("WS:REC Type[%d] Content[%s]", msgType, content))

				lastRequest = content
				dataChannel <- content
			}
		}
	}
}
