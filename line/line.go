package line

import (
	"log"
	"releasebot/config"
	"sync"

	"github.com/line/line-bot-sdk-go/linebot"
)

func GetInstance(config config.LineConfig) *API {
	var instance *API
	var once sync.Once
	bot, err := linebot.New(config.Secret, config.Token)
	if err != nil {
		log.Fatal(err)
	}
	once.Do(func() {
		instance = &API{
			groupID: config.Group,
			bot:     bot,
		}
	})
	return instance
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
