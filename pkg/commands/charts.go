package commands

import (
	"context"
	"fmt"
	"github.com/ForYaSee/NindoBot/pkg/embed"
	"github.com/ForYaSee/NindoBot/pkg/nindo"
	"github.com/ForYaSee/NindoBot/pkg/util"
	"github.com/andersfylling/disgord"
	"github.com/getsentry/sentry-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	chartsCalls = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "bot_command_charts_calls",
	}, []string{"platform"})
)

func Charts() *Command {
	return &Command{
		Name:        "Charts",
		Description: "Zeigt eine Übersicht der SocialMedia Charts",
		Invocations: []string{"charts", "ch"},
		Usage:       "",
		SubCommands: []*Command{youTubeCharts(), twitchCharts(), tiktokCharts(), instagramCharts(), twitterCharts()},
		Process: func(ctx *Ctx) error {
			chartsCalls.WithLabelValues("/").Inc()
			ctx.SendHelp()
			return nil
		},
	}
}

func twitterCharts() *Command {
	return &Command{
		Name:        "Twitter",
		Description: "Zeigt die Twitter Charts",
		Invocations: []string{"twitter", "twi"},
		Usage:       "",
		Process: func(ctx *Ctx) error {
			chartsCalls.WithLabelValues("twitter").Inc()
			artists, err := nindo.DefaultClient.GetTwitterRankChartsByLikes()
			if err != nil {
				sentry.CaptureException(err)
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
					Footer:      NindoCreditsFooter,
					Timestamp:   DefaultTimestamp,
				},
			})
			return err
		},
	}
}

func youTubeCharts() *Command {
	return &Command{
		Name:        "YouTube",
		Description: "Zeigt die YouTube Charts",
		Invocations: []string{"youtube", "yt"},
		Usage:       "",
		Process: func(ctx *Ctx) error {
			chartsCalls.WithLabelValues("youtube").Inc()
			artists, err := nindo.DefaultClient.GetYouTubeRankChartsByViews()
			if err != nil {
				sentry.CaptureException(err)
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
					Footer:      NindoCreditsFooter,
					Timestamp:   DefaultTimestamp,
				},
			})
			return err
		},
	}
}

func tiktokCharts() *Command {
	return &Command{
		Name:        "TikTok",
		Description: "Zeigt die TikTok Charts",
		Invocations: []string{"tiktok", "tt"},
		Usage:       "",
		Process: func(ctx *Ctx) error {
			chartsCalls.WithLabelValues("tiktok").Inc()
			artists, err := nindo.DefaultClient.GetTikTokRankChartsByLikes()
			if err != nil {
				sentry.CaptureException(err)
				ctx.Message.Reply(context.Background(), ctx.Session, disgord.CreateMessageParams{
					Embed: embed.Error("Wir konnten keine Daten von Nindo bekommen."),
				})
			}
			fields := make([]*disgord.EmbedField, 0)
			for _, artist := range artists {
				fields = append(fields, &disgord.EmbedField{
					Name:   fmt.Sprintf("%d. %s", artist.Rank, artist.ArtistName),
					Value:  fmt.Sprintf("%s Likes", util.Comma(int64(artist.Value))),
					Inline: true,
				})
			}
			_, err = ctx.Message.Reply(context.Background(), ctx.Session, disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Description: "Die Top 10 auf TikTok nach Anzahl der durchschnittlichen Likes der letzten 5 Tage.",
					Title:       "TikTok Charts",
					Fields:      fields,
					Color:       0x69C9D0,
					Footer:      NindoCreditsFooter,
					Timestamp:   DefaultTimestamp,
				},
			})
			return err
		},
	}
}

func twitchCharts() *Command {
	return &Command{
		Name:        "Twitch",
		Description: "Zeigt die Twitch Charts",
		Invocations: []string{"twitch", "tw"},
		Usage:       "",
		Process: func(ctx *Ctx) error {
			chartsCalls.WithLabelValues("twitch").Inc()
			artists, err := nindo.DefaultClient.GetTwitchRankChartsByViewers()
			if err != nil {
				sentry.CaptureException(err)
				ctx.Message.Reply(context.Background(), ctx.Session, disgord.CreateMessageParams{
					Embed: embed.Error("Wir konnten keine Daten von Nindo bekommen."),
				})
			}
			fields := make([]*disgord.EmbedField, 0)
			for _, artist := range artists {
				fields = append(fields, &disgord.EmbedField{
					Name:   fmt.Sprintf("%d. %s", artist.Rank, artist.ArtistName),
					Value:  fmt.Sprintf("%s Likes", util.Comma(int64(artist.Value))),
					Inline: true,
				})
			}
			_, err = ctx.Message.Reply(context.Background(), ctx.Session, disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Description: "Die Top 10 auf Twitch nach Anzahl der durchschnittlichen Viewer.",
					Title:       "Twitch Charts",
					Fields:      fields,
					Color:       0x6441A4,
					Footer:      NindoCreditsFooter,
					Timestamp:   DefaultTimestamp,
				},
			})
			return err
		},
	}
}

func instagramCharts() *Command {
	return &Command{
		Name:        "Instagram",
		Description: "Zeigt die Instagram Charts",
		Invocations: []string{"instagram", "insta"},
		Usage:       "",
		Process: func(ctx *Ctx) error {
			chartsCalls.WithLabelValues("instagram").Inc()
			artists, err := nindo.DefaultClient.GetInstagramRankChartsByLikes()
			if err != nil {
				sentry.CaptureException(err)
				ctx.Message.Reply(context.Background(), ctx.Session, disgord.CreateMessageParams{
					Embed: embed.Error("Wir konnten keine Daten von Nindo bekommen."),
				})
			}
			fields := make([]*disgord.EmbedField, 0)
			for _, artist := range artists {
				fields = append(fields, &disgord.EmbedField{
					Name:   fmt.Sprintf("%d. %s", artist.Rank, artist.ArtistName),
					Value:  fmt.Sprintf("%s Likes", util.Comma(int64(artist.Value))),
					Inline: true,
				})
			}
			_, err = ctx.Message.Reply(context.Background(), ctx.Session, disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Description: "Die Top 10 auf Instagram nach Anzahl der durchschnittlichen Likes der letzten 5 Tage.",
					Title:       "Instagram Charts",
					Fields:      fields,
					Color:       0xaa7a50,
					Footer:      NindoCreditsFooter,
					Timestamp:   DefaultTimestamp,
				},
			})
			return err
		},
	}
}
