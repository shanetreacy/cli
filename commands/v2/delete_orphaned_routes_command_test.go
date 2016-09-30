package v2_test

import (
	"code.cloudfoundry.org/cli/commands/v2"
	"code.cloudfoundry.org/cli/utils/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("DeletedOrphanedRoutes Command", func() {
	var (
		cmd    v2.DeleteOrphanedRoutesCommand
		fakeUI ui.UI
		// fakeActor  *v2fakes.FakeAPIConfigActor
		// fakeConfig *commandsfakes.FakeConfig
	)

	BeforeEach(func() {
		out := NewBuffer()
		fakeUI = ui.NewTestUI(out, out)
		// fakeActor = new(v2fakes.FakeDeleteOrphanedRoutesActor)
		// fakeConfig = new(commandsfakes.FakeConfig)

		cmd = v2.DeleteOrphanedRoutesCommand{
			UI: fakeUI,
			// Actor:  fakeActor,
			// Config: fakeConfig,
		}
	})

	Context("when orphaned routes exist", func() {
		It("deletes the orphaned routes", func() {
		})
	})

	Context("when orphaned routes exceed more than a page", func() {
		It("deletes all the orphaned routes", func() {
		})
	})

	FDescribe("Force flag", func() {
		Context("when -f is provided", func() {
			var err error
			JustBeforeEach(func() {
				cmd.Force = true
				err = cmd.Execute([]string{})
			})

			It("does not give a prompt", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(fakeUI.Out).ToNot(Say("Really delete orphaned routes?"))
			})
		})

		Context("when -f is not provided", func() {
			var err error
			JustBeforeEach(func() {
				err = cmd.Execute([]string{})
			})

			It("does give a prompt", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(fakeUI.Out).To(Say("Really delete orphaned routes?"))
			})
		})
	})
})
