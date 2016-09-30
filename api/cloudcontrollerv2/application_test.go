package cloudcontrollerv2_test

import (
	"net/http"

	. "code.cloudfoundry.org/cli/api/cloudcontrollerv2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Application", func() {
	var (
		client *CloudControllerClient
	)

	BeforeEach(func() {
		client = NewTestClient()
	})

	Describe("GetApplications", func() {
		BeforeEach(func() {
			response1 := `{
				"next_url": "/v2/apps?q=space_guid:some-space-guid&page=2",
				"resources": [
					{
						"metadata": {
							"guid": "some-app-guid-1"
						},
						"entity": {
							"name": "some-app-name-1"
						}
					},
					{
						"metadata": {
							"guid": "some-app-guid-2"
						},
						"entity": {
							"name": "some-app-name-2"
						}
					}
				]
			}`
			response2 := `{
				"next_url": null,
				"resources": [
					{
						"metadata": {
							"guid": "some-app-guid-3"
						},
						"entity": {
							"name": "some-app-name-3"
						}
					},
					{
						"metadata": {
							"guid": "some-app-guid-4"
						},
						"entity": {
							"name": "some-app-name-4"
						}
					}
				]
			}`
			server.AppendHandlers(
				CombineHandlers(
					VerifyRequest("GET", "/v2/apps", "q=space_guid:some-space-guid"),
					RespondWith(http.StatusOK, response1, http.Header{"X-Cf-Warnings": {"this is a warning"}}),
				),
			)
			server.AppendHandlers(
				CombineHandlers(
					VerifyRequest("GET", "/v2/apps", "q=space_guid:some-space-guid&page=2"),
					RespondWith(http.StatusOK, response2, http.Header{"X-Cf-Warnings": {"this is another warning"}}),
				),
			)
		})

		Context("when apps exist", func() {
			It("returns all the queried apps", func() {
				apps, warnings, err := client.GetApplications([]Query{{
					Filter:   SpaceGUIDFilter,
					Operator: EqualOperator,
					Value:    "some-space-guid",
				}})
				Expect(err).NotTo(HaveOccurred())
				Expect(apps).To(ConsistOf([]Application{
					{Name: "some-app-name-1", GUID: "some-app-guid-1"},
					{Name: "some-app-name-2", GUID: "some-app-guid-2"},
					{Name: "some-app-name-3", GUID: "some-app-guid-3"},
					{Name: "some-app-name-4", GUID: "some-app-guid-4"},
				}))
				Expect(warnings).To(ConsistOf(Warnings{"this is a warning", "this is another warning"}))
			})
		})
	})
})
