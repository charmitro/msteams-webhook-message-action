package main

import (
	"fmt"
	"os"

	goteamsnotify "github.com/dasrick/go-teams-notify/v2"
)

func main() {
	err := sendTheMessage()
	if err != nil {
		panic(err.Error())
	}
}

func sendTheMessage() error {
	fmt.Println(os.Getenv("TITLE"))
	mstClient := goteamsnotify.NewClient()

	webhookURL :=
		"https://outlook.office.com/webhook/" + os.Getenv("WEBHOOK_URL")

	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = os.Getenv("TITLE")
	msgCard.Text = os.Getenv("MESSAGE")
	msgCard.ThemeColor = "#DF813D"

	return mstClient.Send(string(webhookURL), msgCard)
}
