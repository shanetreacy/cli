package v2

import (
	"code.cloudfoundry.org/cli/commands"
)

type DeleteOrphanedRoutesCommand struct {
	Force           bool        `short:"f" description:"Force deletion without confirmation"`
	usage           interface{} `usage:"CF_NAME delete-orphaned-routes [-f]"`
	relatedCommands interface{} `related_commands:"delete-route, routes"`
	UI              commands.UI
}

func (cmd DeleteOrphanedRoutesCommand) Setup(config commands.Config, ui commands.UI) error {
	cmd.UI = ui
	return nil
}

func (cmd DeleteOrphanedRoutesCommand) Execute(args []string) error {
	// cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	if cmd.Force {
		//run the actor command directly
		return nil
	}

	ui.DisplayPrompt("Really delete orphaned routes?")
	//run the actor command directly
	return nil
}
