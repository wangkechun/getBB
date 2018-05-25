package v2ex

import (
	"regexp"
	"strings"

	rpc "qiniupkg.com/x/rpc.v7"
)

// Topic 主题
type Topic struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	URL             string `json:"url"`
	Content         string `json:"content"`
	ContentRendered string `json:"content_rendered"`
	Replies         int    `json:"replies"`
	Member          struct {
		ID           int    `json:"id"`
		Username     string `json:"username"`
		Tagline      string `json:"tagline"`
		AvatarMini   string `json:"avatar_mini"`
		AvatarNormal string `json:"avatar_normal"`
		AvatarLarge  string `json:"avatar_large"`
	} `json:"member"`
	Node struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		Title            string `json:"title"`
		TitleAlternative string `json:"title_alternative"`
		URL              string `json:"url"`
		Topics           int    `json:"topics"`
		AvatarMini       string `json:"avatar_mini"`
		AvatarNormal     string `json:"avatar_normal"`
		AvatarLarge      string `json:"avatar_large"`
	} `json:"node"`
	Created      int `json:"created"`
	LastModified int `json:"last_modified"`
	LastTouched  int `json:"last_touched"`
}

// GetLatestTipics 获取最新主题
func GetLatestTipics(keys []string) (topics []*Topic, err error) {
	var latest []*Topic
	err = rpc.DefaultClient.CallWithJson(nil, &latest, "GET", "https://www.v2ex.com/api/topics/latest.json", nil)
	if err != nil {
		return
	}

	regStr := strings.Join(keys, "|")
	reg := regexp.MustCompile(regStr)

	for _, topic := range latest {
		if len(reg.FindAllString(topic.Title, -1)) > 0 || len(reg.FindAllString(topic.Content, -1)) > 0 {
			topics = append(topics, topic)
		}
	}
	return
}
