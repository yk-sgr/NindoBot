package embed

import "github.com/andersfylling/disgord"

func Error(message string) *disgord.Embed {
	return &disgord.Embed{
		Title:       "Fehler!",
		Description: message,
		Color:       0xeb4034,
	}
}
