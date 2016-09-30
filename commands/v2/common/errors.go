package common

import "fmt"

type APIRequestError struct {
	Err error
}

func (e APIRequestError) Error() string {
	return "Request error: {{.Error}}\nTIP: If you are behind a firewall and require an HTTP proxy, verify the https_proxy environment variable is correctly set. Else, check your network connection."
}

func (e APIRequestError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"Error": e.Err,
	})
}

type InvalidSSLCertError struct {
	API string
}

func (e InvalidSSLCertError) Error() string {
	return "Invalid SSL Cert for {{.API}}\nTIP: Use 'cf api --skip-ssl-validation' to continue with an insecure API endpoint"
}

func (e InvalidSSLCertError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"API": e.API,
	})
}

type NoAPISetError struct {
	BinaryName string
}

func (e NoAPISetError) Error() string {
	return "No API endpoint set. Use '{{.LoginCommand}}' or '{{.ApiCommand}}' to target an endpoint."
}

func (e NoAPISetError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"LoginCommand": fmt.Sprintf("%s login", e.BinaryName),
		"ApiCommand":   fmt.Sprintf("%s api", e.BinaryName),
	})
}

type NotLoggedInError struct {
	BinaryName string
}

func (e NotLoggedInError) Error() string {
	return "Not logged in. Use '{{.LoginCommand}}' to log in."
}

func (e NotLoggedInError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"LoginCommand": fmt.Sprintf("%s login", e.BinaryName),
	})
}

type NoTargetedOrgError struct {
	BinaryName string
}

func (e NoTargetedOrgError) Error() string {
	return "No org targeted. Use '{{.TargetCommand}}' to target an org."
}

func (e NoTargetedOrgError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"TargetCommand": fmt.Sprintf("%s login", e.BinaryName),
	})
}

type NoTargetedSpaceError struct {
	BinaryName string
}

func (e NoTargetedSpaceError) Error() string {
	return "No space targeted. Use '{{.TargetCommand}}' to target a space."
}

func (e NoTargetedSpaceError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"TargetCommand": fmt.Sprintf("%s login", e.BinaryName),
	})
}

type AppNotFoundError struct {
	AppName string
}

func (e AppNotFoundError) Error() string {
	return "App {{.AppName}} not found"
}

func (e AppNotFoundError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"AppName": e.AppName,
	})
}

type ServiceInstanceNotFoundError struct {
	ServiceInstance string
}

func (e ServiceInstanceNotFoundError) Error() string {
	return "Service instance {{.ServiceInstance}} not found"
}

func (e ServiceInstanceNotFoundError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"ServiceInstance": e.ServiceInstance,
	})
}

type ServiceBindingNotFoundError struct {
	AppName         string
	ServiceInstance string
}

func (e ServiceBindingNotFoundError) Error() string {
	return "Binding between {{.ServiceInstance}} and {{.AppName}} did not exist"
}

func (e ServiceBindingNotFoundError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"AppName":         e.AppName,
		"ServiceInstance": e.ServiceInstance,
	})
}
