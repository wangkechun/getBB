package main

import (
	"fmt"

	"github.com/spf13/viper"
	cron "gopkg.in/robfig/cron.v2"
	log "qiniupkg.com/x/log.v7"

	"github.com/miclle/getBB/alarm"
	"github.com/miclle/getBB/store"
	"github.com/miclle/getBB/v2ex"
)

var cloudStore *store.Store
var slackAlarm *alarm.SlackAlarm

func init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func main() {
	viper.SetDefault("slack_channel", "#user-voice")

	ak := viper.GetString("access_key")
	sk := viper.GetString("secret_key")
	bucket := viper.GetString("bucket")
	slackWebhook := viper.GetString("slack_webhook")

	cloudStore = store.Init(ak, sk, bucket)
	slackAlarm = &alarm.SlackAlarm{WebhookURL: slackWebhook}

	cron := cron.New()
	cron.Start()

	spec := "@every 30s"
	if _, err := cron.AddFunc(spec, WatchV2ex); err != nil {
		log.Error("add watch v2ex cron task error:", err.Error())
	}

	select {}
}

// WatchV2ex watch v2ex
func WatchV2ex() {
	log.Info("watch v2ex topics")
	keys := viper.GetStringSlice("keys")
	topics, err := v2ex.GetLatestTipics(keys)
	if err != nil {
		log.Error("get v2ex latest topics error:", err.Error())
	}
	for _, topic := range topics {

		if cloudStore.IfExists(topic.URL) {
			break
		}

		if err := cloudStore.Save(topic.URL); err != nil {
			log.Error("send slack alarm error:", err.Error())
			break
		}

		payload := &alarm.Payload{
			Channel:   viper.GetString("slack_channel"),
			Username:  fmt.Sprintf("%s - v2ex", topic.Title),
			Text:      fmt.Sprintf("%s\n%s", topic.URL, topic.PlainText),
			Markdown:  true,
			LinkNames: topic.URL,
		}
		if err := slackAlarm.Send(payload); err != nil {
			log.Error("send slack alarm error:", err.Error())
		}
	}
}
