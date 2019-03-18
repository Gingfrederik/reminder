package main

import (
	"releasebot/config"
	"releasebot/gitlab"
	"releasebot/line"
	"releasebot/schedule"
	"releasebot/slack"

	"github.com/gin-gonic/gin"
)

func main() {

	configuration := config.New()

	slackAPI := slack.GetInstance(configuration.Slack)
	gitlabAPI := gitlab.GetInstance()
	_ = line.GetInstance(configuration.Line)

	cron := schedule.New(configuration)

	slackHandler := slack.NewHandler(configuration.Slack, slackAPI)

	router := gin.Default()
	router.POST("/gitlab/callback", gitlabAPI.Event)
	router.POST("/slack/callback", slackHandler.Command)
	router.POST("/slack/event", slackHandler.Event)
	router.POST("/slack/action", slackHandler.Action)

	go slackAPI.HandleConnection()
	go slackAPI.HandleEvent()
	cron.Start()
	router.Run(":8080")
}
