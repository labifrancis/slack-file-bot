package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_API_TOKEN", "KEY")
	os.Setenv("CHANNEL_ID", "ID")
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{""}

}
