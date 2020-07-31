package commands

import (
	"context"
	"fmt"
	"github.com/ForYaSee/NindoBot/pkg/embed"
	"github.com/ForYaSee/NindoBot/pkg/nindo"
	"github.com/ForYaSee/NindoBot/pkg/util"
	"github.com/andersfylling/disgord"
	"time"
)

func twitterCharts() *Command {
	return &Command{
		Name:        "Twitter",
		Description: "Zeigt die Twitter Charts",
		Invocations: []string{"twitter", "twi"},
		Usage:       "",
		Process: func(ctx *Ctx) error {
			artists, err := nindo.DefaultClient.GetTwitterRankChartsByLikes()
			if err != nil {
				ctx.Message.Reply(context.Background(), ctx.Session, disgord.CreateMessageParams{
					Embed: embed.Error("Wir konnten keine Daten von Nindo bekommen."),
				})
			}
			fields := make([]*disgord.EmbedField, 0)
			for _, artist := range artists {
				fields = append(fields, &disgord.EmbedField{
					Name:   fmt.Sprintf("%d. %s", artist.Rank, artist.ArtistName),
					Value:  fmt.Sprintf("%s Views", util.Comma(int64(artist.Value))),
					Inline: true,
				})
			}
			_, err = ctx.Message.Reply(context.Background(), ctx.Session, disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Description: "Die Top 10 auf Twitter nach Anzahl der durchschnittlichen Likes der letzten 5 Tage.",
					Title:       "Twitter Charts",
					Fields:      fields,
					Color:       0x08a0e9,
					Footer: &disgord.EmbedFooter{
						IconURL: "https://nindo.de/logo.png",
						Text:    "nindo.de",
					},
					Timestamp: disgord.Time{
						Time: time.Now(),
					},
				},
			})
			return err
		},
	}
}
