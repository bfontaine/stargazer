package starsbot

import "github.com/nlopes/slack"

type Bot struct {
	API *slack.Client
	RTM *slack.RTM

	Stars chan Star
}

func New(token string) *Bot {
	return &Bot{
		API:   slack.New(token),
		Stars: make(chan Star),
	}
}

func (b *Bot) Start() {
	b.RTM = b.API.NewRTM()
	go b.RTM.ManageConnection()

	for {
		select {
		case msg := <-b.RTM.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.StarAddedEvent:
				b.Stars <- newStar(ev.User, true, ev.Item)
			case *slack.StarRemovedEvent:
				b.Stars <- newStar(ev.User, false, ev.Item)
			case *slack.DisconnectedEvent:
				return
			default:
				// skip
			}
		}
	}
}
