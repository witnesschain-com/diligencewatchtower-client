package core

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"

	wtCommon "github.com/diligencewatchtower-client/common"
)

const (
	WS_URL = "wss://api.witnesschain.com/tracer/v1/watchtower/websocket"
)

type WebsockerClient struct {
	dialer     websocket.Dialer
	Connection *websocket.Conn
}

func (wc *WebsockerClient) ConnectToCoordinator(headers http.Header) error {
	url, _ := url.Parse(WS_URL)
	connection, response, err := wc.dialer.Dial(url.String(), headers)
	if err != nil {
		wtCommon.Error(response) //Note: more info in case ws connection fails
		return err
	}
	wc.Connection = connection
	return nil
}

func (wc *WebsockerClient) CloseConnection() error {
	wtCommon.Info("Disconnecting")
	err := wc.Connection.Close()
	return err
}

func (wc *WebsockerClient) ListenForMessages(dataChannel chan string) error {
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
