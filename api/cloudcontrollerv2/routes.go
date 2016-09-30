package cloudcontrollerv2

import "github.com/tedsuo/rata"

const (
	InfoRequest = "Info"
	AppsRequest = "Apps"
)

var Routes = rata.Routes{
	{Path: "/v2/info", Method: "GET", Name: InfoRequest},
	{Path: "/v2/apps", Method: "GET", Name: AppsRequest},
}
