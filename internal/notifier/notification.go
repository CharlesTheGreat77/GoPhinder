package notifier

import (
	"gophinder/utils"
	"log"
)

type NotifierService struct {
	Notifiers []Notifier
}

func NewNotifierService(config utils.Config) (*NotifierService, error) {
	var notifiers []Notifier

	if config.Discord.Enabled {
		discordNotifier := NewDiscordNotifier(config.Discord.WebhookURL)
		notifiers = append(notifiers, discordNotifier)
	}
	if config.Telegram.Enabled {
		telegramNotifier := NewTelegramNotifier(config.Telegram.BotToken, config.Telegram.ChatID)
		notifiers = append(notifiers, telegramNotifier)
	}
	if config.Email.Enabled {
		emailNotifier := NewEmailNotifier(config.Email.SMTPHost, config.Email.SMTPPort, config.Email.Username, config.Email.Password, config.Email.To)
		notifiers = append(notifiers, emailNotifier)
	}

	if len(notifiers) == 0 {
		log.Println("[*] Warning: No notification service configured. Notifications will be disabled.")
	}

	return &NotifierService{
		Notifiers: notifiers,
	}, nil
}

func (ns *NotifierService) Notify(title string, fields map[string]string) error {
	for _, notifier := range ns.Notifiers {
		if err := notifier.Notify(title, fields); err != nil {
			log.Printf("Error notifying with service: %v", err)
			continue
		}
	}
	return nil
}
