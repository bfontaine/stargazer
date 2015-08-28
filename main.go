package main

import (
	"log"

	"github.com/nlopes/slack"
	"github.com/vrischmann/envconfig"
)

var conf struct {
	Token string
}

var api *slack.Client
var rtm *slack.RTM

var ids2usernames map[string]string

func main() {
	if err := envconfig.Init(&conf); err != nil {
		log.Fatal(err)
	}

	ids2usernames = make(map[string]string)

	api = slack.New(conf.Token)
	rtm = api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.StarAddedEvent:
				notify(newStar(ev.User, true, ev.Item))
			case *slack.StarRemovedEvent:
				notify(newStar(ev.User, false, ev.Item))
			case *slack.DisconnectedEvent:
				return
			default:
				log.Printf("%+v\n", msg)
				// skip
			}
		}
	}
}

func notify(s Star) {
	_, _, ch, err := api.OpenIMChannel(s.Message.User)
	if err != nil {
		log.Println(err.Error())
		return
	}
	msg, err := s.notification()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	rtm.SendMessage(&slack.OutgoingMessage{
		Channel: ch,
		Text:    msg,
		Type:    "message",
	})
}
