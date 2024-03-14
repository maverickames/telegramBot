package telegrambot

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func SendTelegramBotNotification(message string, parseMode string) error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN env not set")
	}

	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	if chatID == "" {
		return fmt.Errorf("TELEGRAM_CHAT_ID env not set")
	}

	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	response, err := http.PostForm(endpoint,
		url.Values{
			"chat_id":    {chatID},
			"text":       {message},
			"parse_mode": {parseMode},
		})
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("telegram response status Code: %d", response.StatusCode)
	}

	// defer response.Body.Close()
	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(body))
	return nil
}
