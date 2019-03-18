package slack

import (
	"fmt"
	"releasebot/config"
	"sync"

	"github.com/nlopes/slack"
)

func GetInstance(config config.SlackConfig) *API {
	var instance *API
	var once sync.Once
	client := slack.New(config.Token)
	rtm := client.NewRTM()
	once.Do(func() {
		instance = &API{
			channelID: config.Channel,
			Client:    client,
			RTM:       rtm,
		}
	})
	return instance
}

type API struct {
	channelID string
	Client    *slack.Client
	RTM       *slack.RTM
}

func (t *API) PushMessage(message string) {
	// t.RTM.SendMessage(t.RTM.NewOutgoingMessage(message, t.channelID))
	t.Client.PostMessage(t.channelID, slack.MsgOptionText(message, false))
}

func (t *API) HandleConnection() {
	t.RTM.ManageConnection()
}

func (t *API) HandleEvent() {

	for msg := range t.RTM.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			// t.RTM.SendMessage(t.RTM.NewOutgoingMessage("Connected", t.channelID))

		case *slack.MessageEvent:
			fmt.Printf("Message: %+v\n", ev)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:
			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
