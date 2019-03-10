package slack

import (
	"fmt"

	"github.com/nlopes/slack"
)

func New(token string, channelID string) *API {
	rtm := slack.New(token).NewRTM()
	return &API{
		channelID: channelID,
		Client:    rtm,
	}
}

type API struct {
	channelID string
	Client    *slack.RTM
}

func (t *API) PushMessage(message string) {
	t.Client.SendMessage(t.Client.NewOutgoingMessage("<@UFKUT5NJD> "+message, t.channelID))
}

func (t *API) HandleConnection() {
	t.Client.ManageConnection()
}

func (t *API) HandleEvent() {

	for msg := range t.Client.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			t.Client.SendMessage(t.Client.NewOutgoingMessage("Connected", t.channelID))

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)

		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

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
