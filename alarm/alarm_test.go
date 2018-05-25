package alarm

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSlack(t *testing.T) {
	assert := assert.New(t)
	webhookURL := os.Getenv("SLACK_WEBHOOK")

	assert.NotEmpty(webhookURL)

	slackAlarm := &SlackAlarm{
		WebhookURL: webhookURL,
	}

	payload := &Payload{
		Channel:  "#mars-alarm",
		Username: "Mars Bot",
		Text:     time.Now().String(),
		Markdown: true,
	}

	err := slackAlarm.Send(payload)
	assert.Nil(err)
}
