package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
	"github.com/vrischmann/envconfig"
)

var conf struct {
	Token string
}

var api *slack.Client
var rtm *slack.RTM

var whitelist = make(map[string]struct{})

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
			case *slack.MessageEvent:
				if strings.HasPrefix(ev.Channel, "D") {
					resp := gotDM(ev.Channel, ev.User, ev.Text)
					if resp != "" {
						api.PostMessage(ev.Channel, resp, slackParams)
					}
				}
			default:
				// skip
			}
		}
	}
}

func notify(s Star) {
	if _, ok := whitelist[s.Message.User]; !ok {
		log.Printf("%s is not whitelisted, skipping.", s.Message.User)
		return
	}

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

var reEnableMsg = regexp.MustCompile(`\b(enable|yes|oui|want|please)\b`)
var reDisableMsg = regexp.MustCompile(`\b(disable|no|non|stop|fuck you|go away)\b`)

func gotDM(channel, user, text string) string {
	if reEnableMsg.MatchString(strings.ToLower(strings.TrimSpace(text))) {
		if _, ok := whitelist[user]; ok {
			return "I already know that ;)"
		}

		whitelist[user] = struct{}{}
		return "Got it! Happy stalking!"

	} else if reDisableMsg.MatchString(text) {
		if _, ok := whitelist[user]; !ok {
			return "I didn’t want to disturb anyway."
		}

		delete(whitelist, user)
		return "OK got it. I won't disturb you anymore."
	}

	return ""
}
