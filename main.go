package main

import (
	"log"

	"github.com/bfontaine/stargazer/starsbot"
	"github.com/vrischmann/envconfig"
)

func main() {
	if err := envconfig.Init(&conf); err != nil {
		log.Fatal(err)
	}

	bot := starsbot.New(conf.Slack.Token)

	go getStars(bot)

	bot.Start()
}

func getStars(bot *starsbot.Bot) {
	for {
		select {
		case s := <-bot.Stars:
			log.Printf("star event: %+v\n", s)
			// TODO
		}
	}
}
