package gitlab

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/webhooks.v5/gitlab"
)

const (
	path = "/webhooks"
)

func New() *GitlabHook {
	hook, _ := gitlab.New()
	return &GitlabHook{
		hook: hook,
	}
}

// GitlabHook handle gitlab hook
type GitlabHook struct {
	hook *gitlab.Webhook
}

// Event handle gitlab tag event only
func (g *GitlabHook) Event(c *gin.Context) {

	payload, err := g.hook.Parse(c.Request, gitlab.TagEvents)
	if err != nil {
		if err == gitlab.ErrEventNotFound {
			fmt.Println("cant find event")
		}
	}

	switch payload.(type) {
	case gitlab.TagEventPayload:
		data := payload.(gitlab.TagEventPayload)
		fmt.Printf("%+v", data)
	default:
		fmt.Println("no support event")
	}
}
