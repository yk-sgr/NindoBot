package commands

func Charts() *Command {
	return &Command{
		Name:        "Charts",
		Description: "Zeigt eine Übersicht der SocialMedia Charts",
		Invocations: []string{"charts", "ch"},
		Usage:       "",
		SubCommands: []*Command{youTubeCharts(), twitchCharts(), tiktokCharts(), instagramCharts(), twitterCharts()},
		Process: func(ctx *Ctx) error {
			ctx.SendHelp()
			return nil
		},
	}
}
