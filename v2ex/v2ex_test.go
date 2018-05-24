package v2ex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestTipics(t *testing.T) {
	assert := assert.New(t)

	topics, err := GetLatestTipics([]string{"画师", "AppStore"})
	assert.NotNil(topics)
	assert.Nil(err)

	for _, topic := range topics {
		fmt.Println(topic.Title)
		fmt.Println(topic.URL)
		err := SendToSlack(topic)
		assert.Nil(err)
	}

}
