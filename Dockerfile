FROM golang:1.14.6
WORKDIR /go/src/github.com/nawadanp/Slack-PD-Sync
COPY . .
RUN make build
CMD ["./Slack-PD-Sync"]
