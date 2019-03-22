package schedule

import (
	"fmt"
	"log"

	"releasebot/config"
	"releasebot/line"
	"releasebot/slack"

	"github.com/robfig/cron"
)

func New(config *config.Config) *cron.Cron {

	slack := slack.GetInstance(config.Slack)
	line := line.GetInstance(config.Line)

	c := cron.New()
	for _, notice := range config.Notice {
		n := notice
		if len(n.User) != 0 {
			n.Message = fmt.Sprintf("<@%s> %s", n.User, n.Message)
		}
		c.AddFunc(n.Times, func() {
			log.Printf("send %+v\n", n)
			line.PushMessage(n.Message)
			slack.PushMessage(n.Message)
		})
		log.Printf("add %+v\n", n)
	}

	return c
}
