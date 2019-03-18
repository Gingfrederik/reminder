package slack

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack/slackevents"
)

const botID = "<@UGTFXKJFK> "
const helpMessage = `
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?`

func handleCommand(data *slackevents.AppMentionEvent) string {
	text := strings.TrimPrefix(data.Text, botID)
	switch text {
	case "help":
		return fmt.Sprintf("<@%s> %s", data.User, helpMessage)
	case "lol":
		return fmt.Sprintf("<@%s> %s", data.User, "LAMO")
	default:
		return fmt.Sprintf("<@%s> %s", data.User, "噓")
	}
}
