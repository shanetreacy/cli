package common_test

import (
	. "code.cloudfoundry.org/cli/commands/v2/common"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Translatable Errors", func() {
	translateFunc := func(s string, vars ...interface{}) string {
		return "translated " + s
	}

	Describe("APIRequestError", func() {
		Describe("Error", func() {
			It("returns the error template", func() {
				e := APIRequestError{}
				Expect(e).To(MatchError("Request error: {{.Error}}\nTIP: If you are behind a firewall and require an HTTP proxy, verify the https_proxy environment variable is correctly set. Else, check your network connection."))
			})
		})

		Describe("Translate", func() {
			It("returns the translated error", func() {
				e := APIRequestError{}
				Expect(e.Translate(translateFunc)).To(Equal("translated Request error: {{.Error}}\nTIP: If you are behind a firewall and require an HTTP proxy, verify the https_proxy environment variable is correctly set. Else, check your network connection."))
			})
		})
	})

	Describe("InvalidSSLCertError", func() {
		Describe("Error", func() {
			It("returns the error template", func() {
				e := InvalidSSLCertError{}
				Expect(e).To(MatchError("Invalid SSL Cert for {{.API}}\nTIP: Use 'cf api --skip-ssl-validation' to continue with an insecure API endpoint"))
			})
		})

		Describe("Translate", func() {
			It("returns the translated error", func() {
				e := InvalidSSLCertError{}
				Expect(e.Translate(translateFunc)).To(Equal("translated Invalid SSL Cert for {{.API}}\nTIP: Use 'cf api --skip-ssl-validation' to continue with an insecure API endpoint"))
			})
		})
	})

	Describe("NoAPISetError", func() {
		Describe("Error", func() {
			It("returns the error template", func() {
				e := NoAPISetError{}
				Expect(e).To(MatchError("No API endpoint set. Use '{{.LoginCommand}}' or '{{.ApiCommand}}' to target an endpoint."))
			})
		})

		Describe("Translate", func() {
			It("returns the translated error", func() {
				e := NoAPISetError{}
				Expect(e.Translate(translateFunc)).To(Equal("translated No API endpoint set. Use '{{.LoginCommand}}' or '{{.ApiCommand}}' to target an endpoint."))
			})
		})
	})

	Describe("NotLoggedInError", func() {
		Describe("Error", func() {
			It("returns the error template", func() {
				e := NotLoggedInError{}
				Expect(e).To(MatchError("Not logged in. Use '{{.LoginCommand}}' to log in."))
			})
		})

		Describe("Translate", func() {
			It("returns the translated error", func() {
				e := NotLoggedInError{}
				Expect(e.Translate(translateFunc)).To(Equal("translated Not logged in. Use '{{.LoginCommand}}' to log in."))
			})
		})
	})

	Describe("NoTargetedOrgError", func() {
		Describe("Error", func() {
			It("returns the error template", func() {
				e := NoTargetedOrgError{}
				Expect(e).To(MatchError("No org targeted. Use '{{.TargetCommand}}' to target an org."))
			})
		})

		Describe("Translate", func() {
			It("returns the translated error", func() {
				e := NoTargetedOrgError{}
				Expect(e.Translate(translateFunc)).To(Equal("translated No org targeted. Use '{{.TargetCommand}}' to target an org."))
			})
		})
	})

	Describe("NoTargetedSpaceError", func() {
		Describe("Error", func() {
			It("returns the error template", func() {
				e := NoTargetedSpaceError{}
				Expect(e).To(MatchError("No space targeted. Use '{{.TargetCommand}}' to target a space."))
			})
		})

		Describe("Translate", func() {
			It("returns the translated error", func() {
				e := NoTargetedSpaceError{}
				Expect(e.Translate(translateFunc)).To(Equal("translated No space targeted. Use '{{.TargetCommand}}' to target a space."))
			})
		})
	})

	Describe("AppNotFoundError", func() {
		Describe("Error", func() {
			It("returns the error template", func() {
				e := AppNotFoundError{}
				Expect(e).To(MatchError("App {{.AppName}} not found"))
			})
		})

		Describe("Translate", func() {
			It("returns the translated error", func() {
				e := AppNotFoundError{}
				Expect(e.Translate(translateFunc)).To(Equal("translated App {{.AppName}} not found"))
			})
		})
	})

	Describe("ServiceInstanceNotFoundError", func() {
		Describe("Error", func() {
			It("returns the error template", func() {
				e := ServiceInstanceNotFoundError{}
				Expect(e).To(MatchError("Service instance {{.ServiceInstance}} not found"))
			})
		})

		Describe("Translate", func() {
			It("returns the translated error", func() {
				e := ServiceInstanceNotFoundError{}
				Expect(e.Translate(translateFunc)).To(Equal("translated Service instance {{.ServiceInstance}} not found"))
			})
		})
	})

	Describe("ServiceBindingNotFoundError", func() {
		Describe("Error", func() {
			It("returns the error template", func() {
				e := ServiceBindingNotFoundError{}
				Expect(e).To(MatchError("Binding between {{.ServiceInstance}} and {{.AppName}} did not exist"))
			})
		})

		Describe("Translate", func() {
			It("returns the translated error", func() {
				e := ServiceBindingNotFoundError{}
				Expect(e.Translate(translateFunc)).To(Equal("translated Binding between {{.ServiceInstance}} and {{.AppName}} did not exist"))
			})
		})
	})
})
