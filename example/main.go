package main

import (
	telegramBot "github.com/maverickames/telegramBot"
)

func main() {
	err := telegramBot.SendTelegramBotNotification("<b>bold</b>Hello, Telegram!\nI'm starting up!\n\t - Testing\n", "HTML")
	if err != nil {
		panic(err)
	}
}
