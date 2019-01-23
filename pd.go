package main

import (
	"fmt"
	"time"

	pagerduty "github.com/PagerDuty/go-pagerduty"
)

func srcPD(pd flagsPD) ([]string, error) {
	usersLIst, err := getPDUsers(pd)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch active users : %s", err)
	}
	return usersLIst, nil
}

func getPDUsers(pd flagsPD) ([]string, error) {
	var opts pagerduty.ListOnCallUsersOptions
	var usersList []string
	opts.Since = time.Now().UTC().String()
	opts.Until = time.Now().UTC().String()
	client := pagerduty.NewClient(pd.key)
	for _, schedule := range pd.schedules {
		users, err := client.ListOnCallUsers(schedule, opts)
		if err != nil {
			return nil, fmt.Errorf("Error while retrieving oncall users : %s", err)
		}
		for _, user := range users {
			usersList = append(usersList, user.Email)
		}
	}
	return usersList, nil
}
