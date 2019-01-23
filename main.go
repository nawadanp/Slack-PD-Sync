package main

import "fmt"

func main() {
	opts, err := parseFlags()
	if err != nil {
		fmt.Printf("Unable to parse the flags : %s\n", err)
		return
	}

	fmt.Println("Run PD Function")
	activeUsers, err := srcPD(opts.pd)
	if err != nil {
		fmt.Printf("Unable to call PagerDuty : %s\n", err)
		return
	}
	fmt.Printf("Active users : %s\n", activeUsers)
	fmt.Println("Run Slack Function")
	err = destSlack(opts.slack, activeUsers)
	if err != nil {
		fmt.Printf("Unable to update user group on Slack : %s\n", err)
		return
	}
	fmt.Println("Slack user group updated")
	return
}
