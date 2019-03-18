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
		if len(notice.User) != 0 {
			notice.Message = fmt.Sprintf("<@%s> %s", notice.User, notice.Message)
		}
		c.AddFunc(notice.Times, func() {
			log.Println("send")
			line.PushMessage(notice.Message)
			slack.PushMessage(notice.Message)
		})
	}

	return c
}
