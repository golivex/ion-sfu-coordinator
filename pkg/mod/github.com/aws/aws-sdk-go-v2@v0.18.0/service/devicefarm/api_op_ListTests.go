// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the list tests operation.
type ListTestsInput struct {
	_ struct{} `type:"structure"`

	// The test suite's Amazon Resource Name (ARN).
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListTestsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTestsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTestsInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a list tests request.
type ListTestsOutput struct {
	_ struct{} `type:"structure"`

	// If the number of items that are returned is significantly large, this is
	// an identifier that is also returned, which can be used in a subsequent call
	// to this operation to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// Information about the tests.
	Tests []Test `locationName:"tests" type:"list"`
}

// String returns the string representation
func (s ListTestsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTests = "ListTests"

// ListTestsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about tests in a given test suite.
//
//    // Example sending a request using ListTestsRequest.
//    req := client.ListTestsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListTests
func (c *Client) ListTestsRequest(input *ListTestsInput) ListTestsRequest {
	op := &aws.Operation{
		Name:       opListTests,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTestsInput{}
	}

	req := c.newRequest(op, input, &ListTestsOutput{})
	return ListTestsRequest{Request: req, Input: input, Copy: c.ListTestsRequest}
}

// ListTestsRequest is the request type for the
// ListTests API operation.
type ListTestsRequest struct {
	*aws.Request
	Input *ListTestsInput
	Copy  func(*ListTestsInput) ListTestsRequest
}

// Send marshals and sends the ListTests API request.
func (r ListTestsRequest) Send(ctx context.Context) (*ListTestsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTestsResponse{
		ListTestsOutput: r.Request.Data.(*ListTestsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTestsRequestPaginator returns a paginator for ListTests.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTestsRequest(input)
//   p := devicefarm.NewListTestsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTestsPaginator(req ListTestsRequest) ListTestsPaginator {
	return ListTestsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTestsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListTestsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTestsPaginator struct {
	aws.Pager
}

func (p *ListTestsPaginator) CurrentPage() *ListTestsOutput {
	return p.Pager.CurrentPage().(*ListTestsOutput)
}

// ListTestsResponse is the response type for the
// ListTests API operation.
type ListTestsResponse struct {
	*ListTestsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTests request.
func (r *ListTestsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
