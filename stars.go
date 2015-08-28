package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

type Message struct {
	User, Channel, Text, Timestamp string
}

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

	return fmt.Sprintf("%s %s your message in #%s", author, verb, s.Message.Channel), nil
}
