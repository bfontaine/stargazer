package starsbot

import "github.com/nlopes/slack"

type Message struct {
	User, Channel, Text, Timestamp string
}

type Star struct {
	User      string
	Timestamp string
	Message   Message
	added     bool
}

func newStar(user string, added bool, item slack.StarredItem) Star {
	return Star{
		User:      user,
		Timestamp: item.Timestamp,
		Message: Message{
			User:      item.Message.User,
			Channel:   item.Channel,
			Text:      item.Message.Text,
			Timestamp: item.Message.Timestamp,
		},
		added: added,
	}
}

func (s Star) Added() bool   { return s.added }
func (s Star) Removed() bool { return !s.added }
