package v2_test

import (
	"code.cloudfoundry.org/cli/commands/commandsfakes"
	. "code.cloudfoundry.org/cli/commands/v2"
	"code.cloudfoundry.org/cli/commands/v2/v2fakes"
	"code.cloudfoundry.org/cli/utils/configv3"
	"code.cloudfoundry.org/cli/utils/ui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("Unbind Service Command", func() {
	var (
		cmd        UnbindServiceCommand
		fakeUI     ui.UI
		fakeActor  *v2fakes.FakeUnbindServiceActor
		fakeConfig *commandsfakes.FakeConfig
		executeErr error
	)

	BeforeEach(func() {
		out := NewBuffer()
		fakeUI = ui.NewTestUI(out, out)
		fakeActor = new(v2fakes.FakeUnbindServiceActor)
		fakeConfig = new(commandsfakes.FakeConfig)

		cmd = UnbindServiceCommand{
			UI:     fakeUI,
			Actor:  fakeActor,
			Config: fakeConfig,
		}
	})

	Describe("command prerequisite failures", func() {
		JustBeforeEach(func() {
			executeErr = cmd.Execute([]string{})
		})

		Context("when the api endpoint is not set", func() {
			It("returns an error", func() {
				Expect(executeErr).To(MatchError(NoAPISetError{
					LoginCommand: "cf login",
					APICommand:   "cf api",
				}))
			})
		})

		Context("when the api endpoint is set", func() {
			BeforeEach(func() {
				fakeConfig.TargetReturns("some-url")
			})

			Context("when the user is not logged in", func() {
				It("returns an error", func() {
					Expect(executeErr).To(MatchError(NotLoggedInError{
						LoginCommand: "cf login",
					}))
				})
			})

			Context("when the user is logged in", func() {
				BeforeEach(func() {
					fakeConfig.AccessTokenReturns("some-access-token")
					fakeConfig.RefreshTokenReturns("some-refresh-token")
				})

				Context("when an org is not targeted", func() {
					It("should return an error", func() {
						Expect(executeErr).To(MatchError(NoTargetedOrgError{
							TargetCommand: "cf target",
						}))
					})

					Context("when an org is targeted", func() {
						BeforeEach(func() {
							fakeConfig.TargetedOrganizationReturns(configv3.Organization{
								GUID: "some-org-guid",
							})
						})

						Context("when a space is not targeted", func() {
							It("should return an error", func() {
								Expect(executeErr).To(MatchError(NoTargetedSpaceError{
									TargetCommand: "cf target",
								}))
							})
						})
					})
				})
			})
		})
	})
})
