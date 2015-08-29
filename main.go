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

func main() {
	if err := envconfig.Init(&conf); err != nil {
		log.Fatal(err)
	}

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
				log.Println("Disconnection.")
				return
			case *slack.ConnectingEvent:
				log.Println("Connecting....")
			case *slack.ConnectedEvent:
				log.Println("Connected.")
			default:
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

	api.PostMessage(ch, msg, slackParams)
}
