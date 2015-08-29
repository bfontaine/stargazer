package main

import "github.com/nlopes/slack"

var slackParams = slack.NewPostMessageParameters()

func init() {
	slackParams.AsUser = true
	slackParams.LinkNames = 1
	slackParams.UnfurlLinks = true
}

var getUsername = memoizeGetter(func(id string) (string, error) {
	u, err := api.GetUserInfo(id)
	if err != nil {
		return "", err
	}
	return u.Name, nil
})

var getChannel = memoizeGetter(func(id string) (string, error) {
	c, err := api.GetChannelInfo(id)
	if err != nil {
		return "", err
	}
	return c.Name, nil
})

func memoizeGetter(f func(string) (string, error)) func(string) (string, error) {
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
