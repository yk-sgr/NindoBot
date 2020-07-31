package commands

import (
	"context"
	"fmt"
	"github.com/andersfylling/disgord"
	"strings"
	"time"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Invocations []string
	SubCommands []*Command
	Process     func(ctx *Ctx) error
}

type Ctx struct {
	Invoke     string
	FullInvoke string
	Args       []string
	Message    *disgord.Message
	Author     *disgord.User
	ChannelID  disgord.Snowflake
	Session    disgord.Session
	Handler    *Handler
	Command    *Command
}

func (c *Ctx) SendHelp() {
	e := &disgord.Embed{
		Title:       fmt.Sprintf("%s - Hilfe", c.Command.Name),
		Description: c.Command.Description,
		Color:       0x7749a0,
		Timestamp: disgord.Time{
			Time: time.Now(),
		},
	}

	fields := make([]*disgord.EmbedField, 0)
	fields = append(fields, &disgord.EmbedField{
		Name:   "Aliase",
		Value:  strings.Join(c.Command.Invocations, ", "),
		Inline: true,
	})
	for _, sub := range c.Command.SubCommands {
		fields = append(fields, &disgord.EmbedField{
			Name:  fmt.Sprintf("%s - %s", sub.Name, sub.Description),
			Value: fmt.Sprintf("%s %s", c.FullInvoke, sub.Invocations[0]),
		})
	}
	e.Fields = fields

	c.Message.Reply(context.Background(), c.Session, disgord.CreateMessageParams{
		Embed: e,
	})
}
