package utils

import (
	"strings"
)

// list of user-agents from filters collected online
var previewBots = []string{
	"Slackbot",         // slack
	"facebookexternal", // facebook
	"WhatsApp",         // WhatsApp
	"Discordbot",       // discord
	"LinkedInBot",      // LinkedIn
	"TelegramBot",      // telegram
	"SkypeUriPreview",  // skype -> Teams uses this
	"Google-Structured-Data-Testing-Tool",
	"Googlebot", // googlebot
	"Bingbot",   // Microsoft
	"MSOffice",  // Microsoft Office
	"Teams",     // Teams link preview -> haven't actually seen this though but in any case
}

// func to detect bots based on the user agent
func IsBot(userAgent string) bool {
	if strings.TrimSpace(userAgent) == "" { // treat missing user agent as blocked
		return true
	}

	for _, bot := range previewBots {
		if strings.Contains(strings.ToLower(userAgent), strings.ToLower(bot)) {
			return true
		}
	}
	return false
}
