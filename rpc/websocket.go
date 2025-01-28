package rpc

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	CallTimeout   time.Duration
	id            int64
	conn          *websocket.Conn
	channels      map[int64]chan RPCResponse
	events        map[uint64]int64
	mutex         sync.Mutex
	ConnectionErr chan error
}

func NewWebSocket(endpoint string, header http.Header) (*WebSocket, error) {
	socketUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	conn, _, err := websocket.DefaultDialer.Dial(socketUrl.String(), header)
	if err != nil {
		return nil, err
	}

	ws := &WebSocket{
		CallTimeout:   3 * time.Second,
		conn:          conn,
		channels:      make(map[int64]chan RPCResponse),
		events:        make(map[uint64]int64),
		ConnectionErr: make(chan error),
	}

	go ws.listen()
	return ws, nil
}

func (w *WebSocket) listen() {
	go func() {
		for {
			_, msg, err := w.conn.ReadMessage()
			if err != nil {
				w.ConnectionErr <- err
				return
			}

			w.mutex.Lock()
			var rpcResponse RPCResponse
			json.Unmarshal(msg, &rpcResponse)
			id := rpcResponse.ID
			ch, ok := w.channels[rpcResponse.ID]
			if ok {
				ch <- rpcResponse

				// Close channel if it's not an event.
				// We will never receive data from that channel ever again, because the id is incremented each call.
				// I'm not sure if it's OK to leave a channel open. Maybe it's picked up by GC, but I prefer to close manually and avoid leak.
				isEvent := false
				for _, eventId := range w.events {
					if eventId == id {
						isEvent = true
						break
					}
				}

				if !isEvent {
					close(ch)
					delete(w.channels, id)
				}
			}
			w.mutex.Unlock()
		}
	}()
}

func (w *WebSocket) subscribeEvent(event interface{}) (RPCResponse, error) {
	return w.Call("subscribe", map[string]interface{}{
		"notify": event,
	})
}

func (w *WebSocket) unsubscribeEvent(event interface{}) (RPCResponse, error) {
	return w.Call("unsubscribe", map[string]interface{}{
		"notify": event,
	})
}

func (w *WebSocket) Close() error {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	// Remove channels and events.
	// We don't need to send unsubscribe event if we just close the connection.
	for id := range w.channels {
		ch := w.channels[id]
		close(ch)
		delete(w.channels, id)
	}

	for event := range w.events {
		delete(w.events, event)
	}

	return w.conn.Close()
}

func (w *WebSocket) CloseEvent(event interface{}) error {
	eventHash, err := HashEvent(event)
	if err != nil {
		return err
	}

	id, ok := w.events[eventHash]
	if ok {
		res, err := w.unsubscribeEvent(event)
		if err != nil {
			return err
		}

		if res.Error != nil {
			return fmt.Errorf(res.Error.Message)
		}

		w.mutex.Lock()
		ch := w.channels[id]
		close(ch)
		delete(w.channels, id)
		delete(w.events, eventHash)
		w.mutex.Unlock()
	}

	return nil
}

func (w *WebSocket) ListenEvent(event interface{}) (ch chan RPCResponse, err error) {
	eventHash, err := HashEvent(event)
	if err != nil {
		return
	}

	id, ok := w.events[eventHash]
	if !ok {
		var res RPCResponse
		res, err = w.subscribeEvent(event)
		if err != nil {
			return
		}

		if res.Error != nil {
			err = fmt.Errorf(res.Error.Message)
			return
		}

		id = res.ID
		w.events[eventHash] = id
	}

	ch, ok = w.channels[id]
	if !ok {
		ch = make(chan RPCResponse)
		w.channels[id] = ch
	}

	return
}

func (w *WebSocket) ListenEventFunc(event interface{}, onData func(RPCResponse)) (err error) {
	ch, err := w.ListenEvent(event)
	if err != nil {
		return
	}

	go func() {
		for res := range ch {
			onData(res)
		}
	}()

	return
}

func (w *WebSocket) Call(method string, params interface{}) (res RPCResponse, err error) {
	w.id++
	rpcRequest := RPCRequest{ID: w.id, JSONRPC: "2.0", Method: method, Params: params}
	data, err := json.Marshal(rpcRequest)
	if err != nil {
		return
	}

	return w.RawCall(w.id, data)
}

func (w *WebSocket) RawCall(id int64, data []byte) (res RPCResponse, err error) {
	resChan := make(chan RPCResponse)
	w.channels[id] = resChan

	err = w.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return
	}

	var timerChan <-chan time.Time
	if w.CallTimeout > 0 {
		timerChan = time.After(w.CallTimeout)
	}

	select {
	case res = <-resChan:
		break
	case <-timerChan:
		defer w.mutex.Unlock()
		w.mutex.Lock()

		ch, ok := w.channels[w.id]
		if ok {
			close(ch)
			delete(w.channels, w.id)
			err = fmt.Errorf("timeout waiting for response")
		}
		break
	}

	return
}

func JsonFormatResponse(res RPCResponse, resErr error, result any) (err error) {
	if resErr != nil {
		err = resErr
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &result)
	return
}

func EventParamsWrap(event string, params interface{}) interface{} {
	eventMap := make(map[string]interface{})
	eventMap[event] = params
	return eventMap
}

func HashEvent(event interface{}) (data uint64, err error) {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return
	}

	h := fnv.New64a()
	_, err = h.Write(eventJson)
	if err != nil {
		return
	}

	data = h.Sum64()
	return
}
