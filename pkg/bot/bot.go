package bot

import (
	"context"
	"github.com/ForYaSee/NindoBot/pkg/commands"
	"github.com/andersfylling/disgord"
)

type Bot struct {
	client          *disgord.Client
	commandHandler  *commands.Handler
	commandRegistry *commands.Registry
}

type Config struct {
	Token      string
	ShardIDs   []uint
	ShardCount uint
	Prefix     string
}

func New(conf Config) *Bot {
	client := disgord.New(disgord.Config{
		BotToken:    conf.Token,
		CacheConfig: &disgord.CacheConfig{DisableVoiceStateCaching: true},
		ShardConfig: disgord.ShardConfig{
			ShardCount: conf.ShardCount,
			ShardIDs:   conf.ShardIDs,
		},
	})
	registry := commands.NewRegistry()
	handler := commands.NewHandler(registry, conf.Prefix)
	b := &Bot{
		client:          client,
		commandRegistry: registry,
		commandHandler:  handler,
	}
	b.registerEvents()
	b.registerCommands()
	return b
}

func (b *Bot) registerEvents() {
	b.client.On(disgord.EvtMessageCreate, b.commandHandler.HandleMessageEvent)
}

func (b *Bot) registerCommands() {
	b.commandRegistry.RegisterCommand(commands.Charts())
}

func (b *Bot) Start() error {
	return b.client.StayConnectedUntilInterrupted(context.Background())
}
