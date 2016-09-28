package v2_test

import (
	"code.cloudfoundry.org/cli/actors"
	"code.cloudfoundry.org/cli/api/cloudcontrollerv2"
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

	JustBeforeEach(func() {
		executeErr = cmd.Execute([]string{})
	})

	Describe("command prerequisite failures", func() {
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

	Describe("no command prerequisite failures", func() {
		BeforeEach(func() {
			fakeConfig.TargetReturns("some-url")
			fakeConfig.AccessTokenReturns("some-access-token")
			fakeConfig.RefreshTokenReturns("some-refresh-token")
			fakeConfig.TargetedOrganizationReturns(configv3.Organization{
				GUID: "some-org-guid",
				Name: "some-org",
			})
			fakeConfig.TargetedSpaceReturns(configv3.Space{
				GUID: "some-space-guid",
				Name: "some-space",
			})
			fakeConfig.CurrentUserReturns(configv3.User{
				Name: "some-user",
			}, nil)

			cmd.RequiredArgs.AppName = "some-app"
			cmd.RequiredArgs.ServiceInstance = "some-service"
		})

		Context("when the app does not exist", func() {
			var appNotFoundErr AppNotFoundError

			BeforeEach(func() {
				fakeActor.GetAppReturns(actors.Application{}, appNotFoundErr)
			})

			It("returns an error", func() {
				Expect(executeErr).To(MatchError(appNotFoundErr))

				Expect(fakeActor.GetAppCallCount()).To(Equal(1))
				Expect(fakeActor.GetAppArgsForCall(0)).To(Equal(
					[]cloudcontrollerv2.Query{
						cloudcontrollerv2.Query{
							Filter:   "name",
							Operator: ":",
							Value:    "some-app",
						},
						cloudcontrollerv2.Query{
							Filter:   "space_guid",
							Operator: ":",
							Value:    "some-space-guid",
						},
					}))
			})
		})

		Context("when the app does exist", func() {
			BeforeEach(func() {
				fakeActor.GetAppReturns(actors.Application{
					GUID: "some-app-guid",
					Name: "some-app",
				}, nil)
			})

			Context("when the service does not exist", func() {
				var serviceInstanceNotFoundErr ServiceInstanceNotFoundError

				BeforeEach(func() {
					fakeActor.GetServiceInstanceReturns(actors.ServiceInstance{}, serviceInstanceNotFoundErr)
				})

				It("returns an error", func() {
					Expect(executeErr).To(MatchError(serviceInstanceNotFoundErr))

					Expect(fakeActor.GetServiceInstanceCallCount()).To(Equal(1))
					Expect(fakeActor.GetServiceInstanceArgsForCall(0)).To(Equal(
						[]cloudcontrollerv2.Query{
							cloudcontrollerv2.Query{
								Filter:   "name",
								Operator: ":",
								Value:    "some-service",
							},
							cloudcontrollerv2.Query{
								Filter:   "space_guid",
								Operator: ":",
								Value:    "some-space-guid",
							},
						}))
				})
			})

			Context("when the service exist", func() {
				BeforeEach(func() {
					fakeActor.GetServiceInstanceReturns(actors.ServiceInstance{
						GUID: "some-service-guid",
						Name: "some-service",
					}, nil)
				})

				Context("when a binding between the app and the service does not exist", func() {
					var serviceBindingNotFoundErr ServiceBindingNotFoundError

					BeforeEach(func() {
						fakeActor.GetServiceBindingReturns(actors.ServiceBinding{}, serviceBindingNotFoundErr)
					})

					It("returns an error", func() {
						Expect(executeErr).To(MatchError(serviceBindingNotFoundErr))

						Expect(fakeActor.GetServiceBindingCallCount()).To(Equal(1))
						Expect(fakeActor.GetServiceBindingArgsForCall(0)).To(Equal(
							[]cloudcontrollerv2.Query{
								cloudcontrollerv2.Query{
									Filter:   "app_guid",
									Operator: ":",
									Value:    "some-app-guid",
								},
								cloudcontrollerv2.Query{
									Filter:   "service_instance_guid",
									Operator: ":",
									Value:    "some-service-guid",
								},
							}))
					})
				})

				Context("when a binding between the app and the service exist", func() {
					BeforeEach(func() {
						fakeActor.GetServiceBindingReturns(actors.ServiceBinding{
							GUID: "some-binding-guid",
						}, nil)
					})

					It("displays the unbinding was successful", func() {
						Expect(executeErr).NotTo(HaveOccurred())

						Expect(fakeUI.Out).To(Say("Unbinding app some-app from service some-service in org some-org / space some-space as some-user..."))
						Eventually(fakeUI.Out).Should(Say("OK"))
					})
				})
			})
		})
	})
})
