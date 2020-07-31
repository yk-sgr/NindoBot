package commands

type Registry struct {
	Commands []*Command
	Invokes  map[string]*Command
}

func NewRegistry() *Registry {
	return &Registry{
		Commands: make([]*Command, 0),
		Invokes:  make(map[string]*Command),
	}
}

func (h *Registry) RegisterCommand(cmd *Command) {
	h.Commands = append(h.Commands, cmd)
	for _, invoke := range cmd.Invocations {
		h.Invokes[invoke] = cmd
	}
}
