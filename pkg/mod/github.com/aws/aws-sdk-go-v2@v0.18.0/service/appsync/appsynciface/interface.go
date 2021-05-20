// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appsynciface provides an interface to enable mocking the AWS AppSync service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appsynciface

import (
	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

// ClientAPI provides an interface to enable mocking the
// appsync.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWSAppSync.
//    func myFunc(svc appsynciface.ClientAPI) bool {
//        // Make svc.CreateApiCache request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := appsync.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        appsynciface.ClientPI
//    }
//    func (m *mockClientClient) CreateApiCache(input *appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error) {
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
	CreateApiCacheRequest(*appsync.CreateApiCacheInput) appsync.CreateApiCacheRequest

	CreateApiKeyRequest(*appsync.CreateApiKeyInput) appsync.CreateApiKeyRequest

	CreateDataSourceRequest(*appsync.CreateDataSourceInput) appsync.CreateDataSourceRequest

	CreateFunctionRequest(*appsync.CreateFunctionInput) appsync.CreateFunctionRequest

	CreateGraphqlApiRequest(*appsync.CreateGraphqlApiInput) appsync.CreateGraphqlApiRequest

	CreateResolverRequest(*appsync.CreateResolverInput) appsync.CreateResolverRequest

	CreateTypeRequest(*appsync.CreateTypeInput) appsync.CreateTypeRequest

	DeleteApiCacheRequest(*appsync.DeleteApiCacheInput) appsync.DeleteApiCacheRequest

	DeleteApiKeyRequest(*appsync.DeleteApiKeyInput) appsync.DeleteApiKeyRequest

	DeleteDataSourceRequest(*appsync.DeleteDataSourceInput) appsync.DeleteDataSourceRequest

	DeleteFunctionRequest(*appsync.DeleteFunctionInput) appsync.DeleteFunctionRequest

	DeleteGraphqlApiRequest(*appsync.DeleteGraphqlApiInput) appsync.DeleteGraphqlApiRequest

	DeleteResolverRequest(*appsync.DeleteResolverInput) appsync.DeleteResolverRequest

	DeleteTypeRequest(*appsync.DeleteTypeInput) appsync.DeleteTypeRequest

	FlushApiCacheRequest(*appsync.FlushApiCacheInput) appsync.FlushApiCacheRequest

	GetApiCacheRequest(*appsync.GetApiCacheInput) appsync.GetApiCacheRequest

	GetDataSourceRequest(*appsync.GetDataSourceInput) appsync.GetDataSourceRequest

	GetFunctionRequest(*appsync.GetFunctionInput) appsync.GetFunctionRequest

	GetGraphqlApiRequest(*appsync.GetGraphqlApiInput) appsync.GetGraphqlApiRequest

	GetIntrospectionSchemaRequest(*appsync.GetIntrospectionSchemaInput) appsync.GetIntrospectionSchemaRequest

	GetResolverRequest(*appsync.GetResolverInput) appsync.GetResolverRequest

	GetSchemaCreationStatusRequest(*appsync.GetSchemaCreationStatusInput) appsync.GetSchemaCreationStatusRequest

	GetTypeRequest(*appsync.GetTypeInput) appsync.GetTypeRequest

	ListApiKeysRequest(*appsync.ListApiKeysInput) appsync.ListApiKeysRequest

	ListDataSourcesRequest(*appsync.ListDataSourcesInput) appsync.ListDataSourcesRequest

	ListFunctionsRequest(*appsync.ListFunctionsInput) appsync.ListFunctionsRequest

	ListGraphqlApisRequest(*appsync.ListGraphqlApisInput) appsync.ListGraphqlApisRequest

	ListResolversRequest(*appsync.ListResolversInput) appsync.ListResolversRequest

	ListResolversByFunctionRequest(*appsync.ListResolversByFunctionInput) appsync.ListResolversByFunctionRequest

	ListTagsForResourceRequest(*appsync.ListTagsForResourceInput) appsync.ListTagsForResourceRequest

	ListTypesRequest(*appsync.ListTypesInput) appsync.ListTypesRequest

	StartSchemaCreationRequest(*appsync.StartSchemaCreationInput) appsync.StartSchemaCreationRequest

	TagResourceRequest(*appsync.TagResourceInput) appsync.TagResourceRequest

	UntagResourceRequest(*appsync.UntagResourceInput) appsync.UntagResourceRequest

	UpdateApiCacheRequest(*appsync.UpdateApiCacheInput) appsync.UpdateApiCacheRequest

	UpdateApiKeyRequest(*appsync.UpdateApiKeyInput) appsync.UpdateApiKeyRequest

	UpdateDataSourceRequest(*appsync.UpdateDataSourceInput) appsync.UpdateDataSourceRequest

	UpdateFunctionRequest(*appsync.UpdateFunctionInput) appsync.UpdateFunctionRequest

	UpdateGraphqlApiRequest(*appsync.UpdateGraphqlApiInput) appsync.UpdateGraphqlApiRequest

	UpdateResolverRequest(*appsync.UpdateResolverInput) appsync.UpdateResolverRequest

	UpdateTypeRequest(*appsync.UpdateTypeInput) appsync.UpdateTypeRequest
}

var _ ClientAPI = (*appsync.Client)(nil)
