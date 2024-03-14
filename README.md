# Telegram Bot

Simple telegram bot notifier. 

Set the env variable. 
TELEGRAM_CHAT_ID=<chatid>
TELEGRAM_BOT_TOKEN=<bot_token>

To run the example
env TELEGRAM_BOT_TOKEN="<bot_token>" TELEGRAM_CHAT_ID=<chatID> go run main.go

You can find your chat ID by hitting this url. It returns json data and will show you what channels it has access to and will have the id as a key/value. 
https://api.telegram.org/<bot_token>/getUpdates

Refer to the documentation for how to format the message 
https://core.telegram.org/bots/api#formatting-options