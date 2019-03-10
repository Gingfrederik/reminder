package main

import (
	"releasebot/config"
	"releasebot/gitlab"
	"releasebot/line"
	"releasebot/schedule"
	"releasebot/slack"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {

}

func main() {

	configuration := config.New()

	slackAPI := slack.New(configuration.Slack.Token, configuration.Slack.Channel)
	lineAPI := line.New(configuration.Line.Secret, configuration.Line.Token, configuration.Line.Group)
	gitlabAPI := gitlab.New()
	cron := schedule.New(configuration, slackAPI, lineAPI)

	router := gin.Default()
	router.POST("/gitlab/callback", gitlabAPI.Event)

	go slackAPI.HandleConnection()
	cron.Start()
	router.Run(":8080")
}
