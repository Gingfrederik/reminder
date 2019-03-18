package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"releasebot/config"

	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// NewHandler handle slack callback
func NewHandler(config config.SlackConfig, api *API) *Handler {
	return &Handler{
		config: config,
		api:    api,
	}
}

type Handler struct {
	config config.SlackConfig
	api    *API
}

func (h *Handler) Command(c *gin.Context) {
	s, err := slack.SlashCommandParse(c.Request)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	log.Printf("%+v\n", s)

	switch s.Command {
	case "/add":
		params := &slack.Msg{Text: s.Text}
		c.JSON(http.StatusOK, params)
	default:
		c.Status(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Event(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	body := buf.String()
	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: h.config.Verification}))
	if e != nil {
		c.Status(http.StatusInternalServerError)
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		r := slackevents.ChallengeResponse{}
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			c.Status(http.StatusInternalServerError)
		}
		c.String(http.StatusOK, r.Challenge)
	}
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			message := handleCommand(ev)
			att := slack.Attachment{
				CallbackID: "test",
				Actions: []slack.AttachmentAction{
					{
						Name:  "actionStart",
						Text:  "Yes",
						Type:  "button",
						Value: "start",
						Style: "primary",
					},
					{
						Name:  "actionCancel",
						Text:  "No",
						Type:  "button",
						Style: "danger",
					},
				},
			}
			h.api.Client.PostMessage(ev.Channel, slack.MsgOptionText(message, false), slack.MsgOptionAttachments(att))
		}
	}
}

func (h *Handler) Action(c *gin.Context) {

	r := c.Request
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read request body: %s", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		log.Printf("[ERROR] Failed to unespace request body: %s", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	message := slack.AttachmentActionCallback{}
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		log.Printf("[ERROR] Failed to decode json message from slack: %s", jsonStr)
		c.Status(http.StatusInternalServerError)
		return
	}
	log.Printf("%+v\n", message)
}
