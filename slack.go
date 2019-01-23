package main

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

func destSlack(optsSlack flagsSlack, activeUsers []string) error {
	slackAPI := slack.New(optsSlack.token)
	var usersIDs []string
	groupID, err := getSlackGroupID(slackAPI, optsSlack.group)
	if err != nil {
		return fmt.Errorf("Error in getSlackGroupID %s", err)
	}

	for _, user := range activeUsers {
		userID, err := getSlackUserID(slackAPI, user)
		if err != nil {
			return err
		}
		usersIDs = append(usersIDs, userID)
	}
	err = updateSlackUserGroup(slackAPI, groupID, strings.Join(usersIDs, ","))
	return err
}

func updateSlackUserGroup(slackAPI *slack.Client, groupID string, usersIDs string) error {
	_, err := slackAPI.UpdateUserGroupMembers(groupID, usersIDs)
	if err != nil {
		return fmt.Errorf("Failed - %s", err)
	}

	return nil
}

func getSlackUserID(slackAPI *slack.Client, userEmail string) (string, error) {
	userDetail, err := slackAPI.GetUserByEmail(userEmail)
	if err != nil {
		return "", fmt.Errorf("Unable to retrieve user details for %s : %s", userEmail, err)
	}
	return userDetail.ID, nil
}

func getSlackGroupID(slackAPI *slack.Client, groupName string) (string, error) {
	groups, err := slackAPI.GetUserGroups()
	if err != nil {
		return "", fmt.Errorf("Slack API error : %s", err)
	}

	for _, group := range groups {
		if group.Name == groupName {
			return group.ID, nil
		}
	}

	return "", fmt.Errorf("Group not found on slack")

}
