package commands

import (
	"github.com/andersfylling/disgord"
	"time"
)

var (
	DefaultTimestamp = disgord.Time{
		Time: time.Now(),
	}
	NindoCreditsFooter = &disgord.EmbedFooter{
		IconURL: "https://nindo.de/logo.png",
		Text:    "nindo.de",
	}
)
