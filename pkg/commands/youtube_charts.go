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

func youTubeCharts() *Command {
	return &Command{
		Name:        "YouTube",
		Description: "Zeigt die YouTube Charts",
		Invocations: []string{"youtube", "yt"},
		Usage:       "",
		Process: func(ctx *Ctx) error {
			artists, err := nindo.DefaultClient.GetYouTubeRankChartsByViews()
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
					Description: "Die Top 10 auf YouTube nach Anzahl der durchschnittlichen Views der letzten 5 Tage.",
					Title:       "YouTube Charts",
					Fields:      fields,
					Color:       0xFF0000,
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
