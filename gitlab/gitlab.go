package gitlab

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/webhooks.v5/gitlab"
)

const (
	path = "/webhooks"
)

func GetInstance() *GitlabHook {
	var instance *GitlabHook
	var once sync.Once
	hook, _ := gitlab.New()
	once.Do(func() {
		instance = &GitlabHook{
			hook: hook,
		}
	})
	return instance
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
			log.Println("cant find event")
		}
	}

	switch payload.(type) {
	case gitlab.TagEventPayload:
		data := payload.(gitlab.TagEventPayload)
		log.Printf("%+v", data)
	default:
		log.Println("no support event")
	}
}
