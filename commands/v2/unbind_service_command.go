package v2

import (
	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/commands/flags"
)

//go:generate counterfeiter . UnbindServiceActor

type UnbindServiceActor interface {
}

type UnbindServiceCommand struct {
	RequiredArgs    flags.BindServiceArgs `positional-args:"yes"`
	usage           interface{}           `usage:"CF_NAME unbind-service APP_NAME SERVICE_INSTANCE"`
	relatedCommands interface{}           `related_commands:"apps, delete-service, services"`

	UI     commands.UI
	Actor  UnbindServiceActor
	Config commands.Config
}

func (cmd UnbindServiceCommand) Setup(config commands.Config, ui commands.UI) error {
	cmd.UI = ui
	cmd.Config = config
	return nil
}

func (cmd UnbindServiceCommand) Execute(args []string) error {
	if cmd.Config.Target() == "" {
		return NoAPISetError{
			LoginCommand: "cf login",
			APICommand:   "cf api",
		}
	}

	if cmd.Config.AccessToken() == "" && cmd.Config.RefreshToken() == "" {
		return NotLoggedInError{
			LoginCommand: "cf login",
		}
	}

	if cmd.Config.TargetedOrganization().GUID == "" {
		return NoTargetedOrgError{
			TargetCommand: "cf target",
		}
	}

	if cmd.Config.TargetedSpace().GUID == "" {
		return NoTargetedSpaceError{
			TargetCommand: "cf target",
		}
	}

	return nil
}
