package main

import (
	"testing"

	"github.com/bfontaine/stargazer/Godeps/_workspace/src/github.com/nlopes/slack"
	"github.com/bfontaine/stargazer/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestStarNotification(t *testing.T) {
	s := Star{
		Added:     true,
		User:      "U123",
		Timestamp: "1401",
		Message: Message{
			User:      "U321",
			Channel:   "C123",
			Text:      "yo",
			Timestamp: "1400",
		},
	}

	slackInfo = &slack.Info{
		Team: &slack.Team{
			Name:   "yolo",
			Domain: "yoloo",
		},
	}

	n, err := s.notification()
	assert.Nil(t, err)

	assert.Equal(t,
		"@usernameU123 just starred your message in #channelC123: https://yoloo.slack.com/archives/channelC123/p1400",
		n)
}
