package alarm

import (
	rpc "qiniupkg.com/x/rpc.v7"
)

// Payload slack msg
type Payload struct {
	Parse       string `json:"parse,omitempty"`
	Username    string `json:"username,omitempty"`
	Channel     string `json:"channel,omitempty"`
	Text        string `json:"text,omitempty"`
	LinkNames   string `json:"link_names,omitempty"`
	UnfurlLinks bool   `json:"unfurl_links,omitempty"`
	UnfurlMedia bool   `json:"unfurl_media,omitempty"`
	Markdown    bool   `json:"mrkdwn,omitempty"`
}

// SlackAlarm slack alarm
type SlackAlarm struct {
	WebhookURL string
}

// Send alarm
func (alarm *SlackAlarm) Send(payload *Payload) (err error) {
	_, err = rpc.DefaultClient.DoRequestWithJson(nil, "POST", alarm.WebhookURL, payload)
	return
}
