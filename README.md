# Slack-PD-Sync
Synchronize Slack group with PD schedules.

## Build sources 

```
go get github.com/nawadanp/Slack-PD-Sync
cd $GOPATH/src/github.com/nawadanp/Slack-PD-Sync
make build
./Slack-PD-Sync
```

## Usage

```
./Slack-PD-Sync --help
Usage of ./Slack-PD-Sync:
  -pd-key string
    	Pagerduty API Key
  -pd-schedules string
    	Comma separated list of Pagerduty schedules
  -slack-group string
    	Slack user group to update
  -slack-token string
    	Slack Token
```

You can use either flags or env variable to run the tool.
If both are set, flags values will be used.

| Flag | Env variable |
| ---- | ------------ |
| pd-key | PD_KEY |
| pd-schedules | PD_SCHEDULES |
| slack-group | SLACK_GROUP |
| slack-token | SLACK_TOKEN |


## Example

```
./Slack-PD-Sync -pd-key xxx -pd-schedules XXXXXX  -slack-group 'Support' -slack-token xoxp-xxx
Run PD Function
Active users : [anthony@example.com]
Run Slack Function
Slack user group updated
```
