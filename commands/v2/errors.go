package v2

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
	LoginCommand string
	APICommand   string
}

func (e NoAPISetError) Error() string {
	return "No API endpoint set. Use '{{.LoginCommand}}' or '{{.ApiCommand}}' to target an endpoint."
}

func (e NoAPISetError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"LoginCommand": e.LoginCommand,
		"ApiCommand":   e.APICommand,
	})
}

type NotLoggedInError struct {
	LoginCommand string
}

func (e NotLoggedInError) Error() string {
	return "Not logged in. Use '{{.LoginCommand}}' to log in."
}

func (e NotLoggedInError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"LoginCommand": e.LoginCommand,
	})
}

type NoTargetedOrgError struct {
	TargetCommand string
}

func (e NoTargetedOrgError) Error() string {
	return "No org targeted. Use '{{.TargetCommand}}' to target an org."
}

func (e NoTargetedOrgError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"TargetCommand": e.TargetCommand,
	})
}

type NoTargetedSpaceError struct {
	TargetCommand string
}

func (e NoTargetedSpaceError) Error() string {
	return "No space targeted. Use '{{.TargetCommand}}' to target a space."
}

func (e NoTargetedSpaceError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"TargetCommand": e.TargetCommand,
	})
}
