package net

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/types"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/util"
	"github.com/gorilla/websocket"
)

type Matchmaking struct {
	InputEvents             []InputEvent
	GetPlayerEntityCallback func() EntityInterface
	DisconnectCallback      func(entity EntityInterface)
	MatchmakingConnections  []MatchmakingConnection
	Upgrader                websocket.Upgrader
}

func (matchmaking *Matchmaking) SetInputEvents(inputEvents []InputEvent) {
	matchmaking.InputEvents = inputEvents
}

func (matchmaking *Matchmaking) GetPlayerEntity(callback func() EntityInterface) {
	matchmaking.GetPlayerEntityCallback = callback
}

func (matchmaking *Matchmaking) Disconnect(callback func(entity EntityInterface)) {
	matchmaking.DisconnectCallback = callback
}

func (matchmaking *Matchmaking) Listen(port types.Port) {
	matchmaking.Upgrader = websocket.Upgrader{}

	cwsc := make(chan *websocket.Conn)
	cmmm := make(chan *MatchmakingMessage)
	ccwsc := make(chan *websocket.Conn)

	go func() {
		for {
			select {
			case connection := <-cwsc:
				matchmaking.addClient(connection)
			case message := <-cmmm:
				matchmaking.receive(message)
			case closedConnection := <-ccwsc:
				matchmaking.removeClient(closedConnection)
			}
		}
	}()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		connection, err := matchmaking.Upgrader.Upgrade(writer, request, nil)

		if err != nil {
			log.Panic(err)
		}

		defer connection.Close()

		defer func() {
			ccwsc <- connection
		}()

		cwsc <- connection

		for {
			_, message, err := connection.ReadMessage()

			if err != nil {
				break
			}

			cmmm <- &MatchmakingMessage{connection, message}
		}
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (matchmaking *Matchmaking) Sync(gameStateStr string) {

}

func (matchmaking *Matchmaking) addClient(connection *websocket.Conn) {
	var entity EntityInterface

	if matchmaking.GetPlayerEntityCallback != nil {
		entity = matchmaking.GetPlayerEntityCallback()
	} else {
		entity = nil
	}

	matchmaking.MatchmakingConnections = append(
		matchmaking.MatchmakingConnections,
		MatchmakingConnection{connection, entity},
	)
}

func (matchmaking *Matchmaking) receive(message *MatchmakingMessage) {
	for _, connection := range matchmaking.MatchmakingConnections {
		if message.WebSocketConn == connection.WebSocketConn {
			var digitalInputEvent DigitalInputEvent
			var analogicInputEvent AnalogicInputEvent

			eventType := INPUT_EVENT_TYPE_DIGITAL
			err := json.Unmarshal(message.Message, &digitalInputEvent)

			if err != nil {
				eventType = INPUT_EVENT_TYPE_ANALOGIC
				err = json.Unmarshal(message.Message, &analogicInputEvent)

				if err != nil {
					break
				}
			}

			for _, inputEvent := range matchmaking.InputEvents {
				if eventType == INPUT_EVENT_TYPE_DIGITAL &&
					inputEvent.Type == INPUT_EVENT_TYPE_DIGITAL &&
					digitalInputEvent.Event == inputEvent.Event {

					if connection.Entity != nil && connection.Entity.GetInput() != nil {
						input := connection.Entity.GetInput()
						input.SetAction(digitalInputEvent.Event, digitalInputEvent.Value)
					}

					break
				}

				if eventType == INPUT_EVENT_TYPE_ANALOGIC &&
					inputEvent.Type == INPUT_EVENT_TYPE_ANALOGIC &&
					analogicInputEvent.Event == inputEvent.Event {

					if connection.Entity != nil && connection.Entity.GetInput() != nil {
						input := connection.Entity.GetInput()

						input.SetValue(
							analogicInputEvent.Event,
							analogicInputEvent.Value.X,
							analogicInputEvent.Value.Y,
						)
					}

					break
				}
			}

			break
		}
	}
}

func (matchmaking *Matchmaking) removeClient(closedConnection *websocket.Conn) {
	indexToRemove := -1

	for i, connection := range matchmaking.MatchmakingConnections {
		if closedConnection == connection.WebSocketConn {
			indexToRemove = i

			break
		}
	}

	matchmaking.MatchmakingConnections = util.Splice(matchmaking.MatchmakingConnections, indexToRemove)
}
