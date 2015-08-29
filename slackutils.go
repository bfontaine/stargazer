package main

import (
	"os"

	"github.com/nlopes/slack"
)

var slackParams = slack.NewPostMessageParameters()
var slackInfo *slack.Info

func init() {
	slackParams.AsUser = true
	slackParams.LinkNames = 1
	slackParams.UnfurlLinks = true
}

type stringGetter func(string) (string, error)

var getUsername, getChannel stringGetter

func init() {
	if os.Getenv("STARGAZER_TESTING") != "" {
		getUsername = func(id string) (string, error) {
			return "username" + id, nil
		}

		getChannel = func(id string) (string, error) {
			return "channel" + id, nil
		}
		return
	}

	getUsername = memoizeGetter(func(id string) (string, error) {
		u, err := api.GetUserInfo(id)
		if err != nil {
			return "", err
		}
		return u.Name, nil
	})

	getChannel = memoizeGetter(func(id string) (string, error) {
		c, err := api.GetChannelInfo(id)
		if err != nil {
			return "", err
		}
		return c.Name, nil
	})
}

func memoizeGetter(f stringGetter) stringGetter {
	cache := make(map[string]string)

	return func(s string) (string, error) {
		if m, ok := cache[s]; ok {
			return m, nil
		}

		m, err := f(s)
		if err == nil {
			cache[s] = m
		}
		return m, err
	}
}
