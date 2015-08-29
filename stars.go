package main

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

// A Message is a starred message
type Message struct {
	User, Channel, Text, Timestamp string
}

// A Star is a star on Slack.
type Star struct {
	Added     bool
	User      string
	Timestamp string
	Message   Message
}

func newStar(user string, added bool, item slack.StarredItem) Star {
	return Star{
		Added:     added,
		User:      user,
		Timestamp: item.Timestamp,
		Message: Message{
			User:      item.Message.User,
			Channel:   item.Channel,
			Text:      item.Message.Text,
			Timestamp: item.Message.Timestamp,
		},
	}
}

func (s Star) notification() (string, error) {
	verb := "starred"
	if !s.Added {
		verb = "unstarred"
	}

	author, err := getUsername(s.User)
	if err != nil {
		return "", err
	}

	channel, err := getChannel(s.Message.Channel)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("@%s just %s your message in #%s: https://%s.slack.com/archives/%s/p%s",
		author, verb, channel,
		slackInfo.Team.Domain, channel,
		strings.Replace(s.Message.Timestamp, ".", "", -1)), nil
}
