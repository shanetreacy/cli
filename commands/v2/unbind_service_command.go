package v2

import (
	"code.cloudfoundry.org/cli/actors"
	"code.cloudfoundry.org/cli/api/cloudcontrollerv2"
	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/commands/flags"
)

//go:generate counterfeiter . UnbindServiceActor

type UnbindServiceActor interface {
	GetApp(query []cloudcontrollerv2.Query) (actors.Application, error)
	GetServiceInstance(query []cloudcontrollerv2.Query) (actors.ServiceInstance, error)
	GetServiceBinding(query []cloudcontrollerv2.Query) (actors.ServiceBinding, error)
	UnbindService(serviceBindingGUID string) error
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

	org := cmd.Config.TargetedOrganization()
	space := cmd.Config.TargetedSpace()

	app, err := cmd.Actor.GetApp([]cloudcontrollerv2.Query{
		cloudcontrollerv2.Query{
			Filter:   "name",
			Operator: ":",
			Value:    cmd.RequiredArgs.AppName,
		},
		cloudcontrollerv2.Query{
			Filter:   "space_guid",
			Operator: ":",
			Value:    space.GUID,
		},
	})
	if err != nil {
		return err
	}

	serviceInstance, err := cmd.Actor.GetServiceInstance([]cloudcontrollerv2.Query{
		cloudcontrollerv2.Query{
			Filter:   "name",
			Operator: ":",
			Value:    cmd.RequiredArgs.ServiceInstance,
		},
		cloudcontrollerv2.Query{
			Filter:   "space_guid",
			Operator: ":",
			Value:    space.GUID,
		},
	})
	if err != nil {
		return err
	}

	serviceBinding, err := cmd.Actor.GetServiceBinding([]cloudcontrollerv2.Query{
		cloudcontrollerv2.Query{
			Filter:   "app_guid",
			Operator: ":",
			Value:    app.GUID,
		},
		cloudcontrollerv2.Query{
			Filter:   "service_instance_guid",
			Operator: ":",
			Value:    serviceInstance.GUID,
		},
	})
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err //return translatable err?
	}

	cmd.UI.DisplayHeaderFlavorText("Unbinding app {{.AppName}} from service {{.ServiceInstance}} in org {{.Org}} / space {{.Space}} as {{.User}}...", map[string]interface{}{
		"AppName":         app.Name,
		"ServiceInstance": serviceInstance.Name,
		"Org":             org.Name,
		"Space":           space.Name,
		"User":            user.Name,
	})
	cmd.UI.DisplayOK()

	err = cmd.Actor.UnbindService(serviceBinding.GUID)
	if err != nil {
		return err
	}

	return nil
}
