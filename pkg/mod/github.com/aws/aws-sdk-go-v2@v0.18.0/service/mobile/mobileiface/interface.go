// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mobileiface provides an interface to enable mocking the AWS Mobile service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mobileiface

import (
	"github.com/aws/aws-sdk-go-v2/service/mobile"
)

// ClientAPI provides an interface to enable mocking the
// mobile.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Mobile.
//    func myFunc(svc mobileiface.ClientAPI) bool {
//        // Make svc.CreateProject request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := mobile.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        mobileiface.ClientPI
//    }
//    func (m *mockClientClient) CreateProject(input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error) {
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
	CreateProjectRequest(*mobile.CreateProjectInput) mobile.CreateProjectRequest

	DeleteProjectRequest(*mobile.DeleteProjectInput) mobile.DeleteProjectRequest

	DescribeBundleRequest(*mobile.DescribeBundleInput) mobile.DescribeBundleRequest

	DescribeProjectRequest(*mobile.DescribeProjectInput) mobile.DescribeProjectRequest

	ExportBundleRequest(*mobile.ExportBundleInput) mobile.ExportBundleRequest

	ExportProjectRequest(*mobile.ExportProjectInput) mobile.ExportProjectRequest

	ListBundlesRequest(*mobile.ListBundlesInput) mobile.ListBundlesRequest

	ListProjectsRequest(*mobile.ListProjectsInput) mobile.ListProjectsRequest

	UpdateProjectRequest(*mobile.UpdateProjectInput) mobile.UpdateProjectRequest
}

var _ ClientAPI = (*mobile.Client)(nil)
