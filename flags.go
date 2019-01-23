package main

import (
	"flag"
	"fmt"
	"strings"
)

type flagsPD struct {
	key       string
	schedules []string
}

type flagsSlack struct {
	token string
	group string
}

type flags struct {
	pd    flagsPD
	slack flagsSlack
}

func parseFlags() (flags, error) {
	var parsedFlags flags
	flagPdKey := flag.String("pd-key", "", "Pagerduty API Key")
	flagSlackToken := flag.String("slack-token", "", "Slack Token")
	flagPdSchedules := flag.String("pd-schedules", "", "Comma separated list of Pagerduty schedules")
	flagSlackGroup := flag.String("slack-group", "", "Slack user group to update")
	flag.Parse()
	parsedFlags.pd.key = *flagPdKey
	listPdSchedules := *flagPdSchedules
	parsedFlags.slack.token = *flagSlackToken
	parsedFlags.slack.group = *flagSlackGroup

	if parsedFlags.pd.key == "" {
		return parsedFlags, fmt.Errorf("-pd-key should be defined")
	}

	if listPdSchedules == "" {
		return parsedFlags, fmt.Errorf("-pd-schedules should be defined")
	}
	parsedFlags.pd.schedules = strings.Split(listPdSchedules, ",")

	if parsedFlags.slack.token == "" {
		return parsedFlags, fmt.Errorf("-slack-token should be defined")
	}

	if parsedFlags.slack.group == "" {
		return parsedFlags, fmt.Errorf("-slack-group should be defined")
	}
	return parsedFlags, nil

}
