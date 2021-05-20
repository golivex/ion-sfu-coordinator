// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package inspectoriface provides an interface to enable mocking the Amazon Inspector service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package inspectoriface

import (
	"github.com/aws/aws-sdk-go-v2/service/inspector"
)

// ClientAPI provides an interface to enable mocking the
// inspector.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Inspector.
//    func myFunc(svc inspectoriface.ClientAPI) bool {
//        // Make svc.AddAttributesToFindings request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := inspector.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        inspectoriface.ClientPI
//    }
//    func (m *mockClientClient) AddAttributesToFindings(input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AddAttributesToFindingsRequest(*inspector.AddAttributesToFindingsInput) inspector.AddAttributesToFindingsRequest

	CreateAssessmentTargetRequest(*inspector.CreateAssessmentTargetInput) inspector.CreateAssessmentTargetRequest

	CreateAssessmentTemplateRequest(*inspector.CreateAssessmentTemplateInput) inspector.CreateAssessmentTemplateRequest

	CreateExclusionsPreviewRequest(*inspector.CreateExclusionsPreviewInput) inspector.CreateExclusionsPreviewRequest

	CreateResourceGroupRequest(*inspector.CreateResourceGroupInput) inspector.CreateResourceGroupRequest

	DeleteAssessmentRunRequest(*inspector.DeleteAssessmentRunInput) inspector.DeleteAssessmentRunRequest

	DeleteAssessmentTargetRequest(*inspector.DeleteAssessmentTargetInput) inspector.DeleteAssessmentTargetRequest

	DeleteAssessmentTemplateRequest(*inspector.DeleteAssessmentTemplateInput) inspector.DeleteAssessmentTemplateRequest

	DescribeAssessmentRunsRequest(*inspector.DescribeAssessmentRunsInput) inspector.DescribeAssessmentRunsRequest

	DescribeAssessmentTargetsRequest(*inspector.DescribeAssessmentTargetsInput) inspector.DescribeAssessmentTargetsRequest

	DescribeAssessmentTemplatesRequest(*inspector.DescribeAssessmentTemplatesInput) inspector.DescribeAssessmentTemplatesRequest

	DescribeCrossAccountAccessRoleRequest(*inspector.DescribeCrossAccountAccessRoleInput) inspector.DescribeCrossAccountAccessRoleRequest

	DescribeExclusionsRequest(*inspector.DescribeExclusionsInput) inspector.DescribeExclusionsRequest

	DescribeFindingsRequest(*inspector.DescribeFindingsInput) inspector.DescribeFindingsRequest

	DescribeResourceGroupsRequest(*inspector.DescribeResourceGroupsInput) inspector.DescribeResourceGroupsRequest

	DescribeRulesPackagesRequest(*inspector.DescribeRulesPackagesInput) inspector.DescribeRulesPackagesRequest

	GetAssessmentReportRequest(*inspector.GetAssessmentReportInput) inspector.GetAssessmentReportRequest

	GetExclusionsPreviewRequest(*inspector.GetExclusionsPreviewInput) inspector.GetExclusionsPreviewRequest

	GetTelemetryMetadataRequest(*inspector.GetTelemetryMetadataInput) inspector.GetTelemetryMetadataRequest

	ListAssessmentRunAgentsRequest(*inspector.ListAssessmentRunAgentsInput) inspector.ListAssessmentRunAgentsRequest

	ListAssessmentRunsRequest(*inspector.ListAssessmentRunsInput) inspector.ListAssessmentRunsRequest

	ListAssessmentTargetsRequest(*inspector.ListAssessmentTargetsInput) inspector.ListAssessmentTargetsRequest

	ListAssessmentTemplatesRequest(*inspector.ListAssessmentTemplatesInput) inspector.ListAssessmentTemplatesRequest

	ListEventSubscriptionsRequest(*inspector.ListEventSubscriptionsInput) inspector.ListEventSubscriptionsRequest

	ListExclusionsRequest(*inspector.ListExclusionsInput) inspector.ListExclusionsRequest

	ListFindingsRequest(*inspector.ListFindingsInput) inspector.ListFindingsRequest

	ListRulesPackagesRequest(*inspector.ListRulesPackagesInput) inspector.ListRulesPackagesRequest

	ListTagsForResourceRequest(*inspector.ListTagsForResourceInput) inspector.ListTagsForResourceRequest

	PreviewAgentsRequest(*inspector.PreviewAgentsInput) inspector.PreviewAgentsRequest

	RegisterCrossAccountAccessRoleRequest(*inspector.RegisterCrossAccountAccessRoleInput) inspector.RegisterCrossAccountAccessRoleRequest

	RemoveAttributesFromFindingsRequest(*inspector.RemoveAttributesFromFindingsInput) inspector.RemoveAttributesFromFindingsRequest

	SetTagsForResourceRequest(*inspector.SetTagsForResourceInput) inspector.SetTagsForResourceRequest

	StartAssessmentRunRequest(*inspector.StartAssessmentRunInput) inspector.StartAssessmentRunRequest

	StopAssessmentRunRequest(*inspector.StopAssessmentRunInput) inspector.StopAssessmentRunRequest

	SubscribeToEventRequest(*inspector.SubscribeToEventInput) inspector.SubscribeToEventRequest

	UnsubscribeFromEventRequest(*inspector.UnsubscribeFromEventInput) inspector.UnsubscribeFromEventRequest

	UpdateAssessmentTargetRequest(*inspector.UpdateAssessmentTargetInput) inspector.UpdateAssessmentTargetRequest
}

var _ ClientAPI = (*inspector.Client)(nil)
