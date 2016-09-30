package v2

import (
	"code.cloudfoundry.org/cli/actors/v2actions"
	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/commands/flags"
	"code.cloudfoundry.org/cli/commands/v2/common"
)

//go:generate counterfeiter . UnbindServiceActor

type UnbindServiceActor interface {
	UnbindServiceBySpace(appName string, serviceInstanceName string, spaceGUID string) (v2actions.Warnings, error)
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

	client, err := common.NewCloudControllerClient(config)
	if err != nil {
		return err
	}
	cmd.Actor = v2actions.NewActor(client)

	return nil
}

func (cmd UnbindServiceCommand) Execute(args []string) error {
	err := common.CheckTarget(cmd.Config, true, true)
	if err != nil {
		return err
	}

	space := cmd.Config.TargetedSpace()
	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err //return translatable err?
	}

	cmd.UI.DisplayHeaderFlavorText("Unbinding app {{.AppName}} from service {{.ServiceInstance}} in org {{.Org}} / space {{.Space}} as {{.User}}...", map[string]interface{}{
		"AppName":         cmd.RequiredArgs.AppName,
		"ServiceInstance": cmd.RequiredArgs.ServiceInstanceName,
		"Org":             cmd.Config.TargetedOrganization().Name,
		"Space":           space.Name,
		"User":            user.Name,
	})

	_, err = cmd.Actor.UnbindServiceBySpace(cmd.RequiredArgs.AppName, cmd.RequiredArgs.ServiceInstanceName, space.GUID)
	if err != nil {
		return err
	}

	cmd.UI.DisplayOK()

	// if binding does not exist display binding d.n.e. message

	return nil
}
