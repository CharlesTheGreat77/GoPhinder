package notifier

import (
	"fmt"
	"net/http"
	"net/url"
)

type TelegramNotifier struct {
	BotToken string
	ChatID   string
}

func NewTelegramNotifier(botToken string, chatID string) *TelegramNotifier {
	return &TelegramNotifier{
		BotToken: botToken,
		ChatID:   chatID,
	}
}

func (t *TelegramNotifier) Notify(title string, fields map[string]string) error {
	msg := "\u2709 *" + title + "*\n"
	for k, v := range fields {
		msg += fmt.Sprintf("*%s:* %s\n", k, v)
	}

	encodedMsg := url.QueryEscape(msg)
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s&parse_mode=Markdown", t.BotToken, t.ChatID, encodedMsg)
	_, err := http.Get(apiURL)
	return err
}
