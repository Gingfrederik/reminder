package line

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	channelSecret      = "beab03ede72a37ecb571269535c80041"
	channelAccessToken = "xdLZO3vavV9M7ZC5/fV1Itcq3HH8ZmS8tX5Vs2sdSMYSxPJ+mM0yazbQFxauex8pYFzWlPz4WRaMvpHrOFxmWpRNmVWoTthBaS0D+7UVE19ngfdKXn1eVjTx97mc8QuJjASkBt/NWjMXfBVVMtKN5AdB04t89/1O/w1cDnyilFU="
)

func New(secret string, token string, groupID string) *API {
	bot, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		log.Fatal(err)
	}
	return &API{
		groupID: groupID,
		bot:     bot,
	}
}

type API struct {
	groupID string
	bot     *linebot.Client
}

func (t *API) PushMessage(message string) {
	if t.groupID != "" {
		if _, err := t.bot.PushMessage(t.groupID, linebot.NewTextMessage(message)).Do(); err != nil {
		}
	}
}
