// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to return a list of all identities (email addresses
// and domains) that you have attempted to verify under your AWS account, regardless
// of verification status.
type ListIdentitiesInput struct {
	_ struct{} `type:"structure"`

	// The type of the identities to list. Possible values are "EmailAddress" and
	// "Domain". If this parameter is omitted, then all identities will be listed.
	IdentityType IdentityType `type:"string" enum:"true"`

	// The maximum number of identities per page. Possible values are 1-1000 inclusive.
	MaxItems *int64 `type:"integer"`

	// The token to use for pagination.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListIdentitiesInput) String() string {
	return awsutil.Prettify(s)
}

// A list of all identities that you have attempted to verify under your AWS
// account, regardless of verification status.
type ListIdentitiesOutput struct {
	_ struct{} `type:"structure"`

	// A list of identities.
	//
	// Identities is a required field
	Identities []string `type:"list" required:"true"`

	// The token used for pagination.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListIdentitiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListIdentities = "ListIdentities"

// ListIdentitiesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns a list containing all of the identities (email addresses and domains)
// for your AWS account in the current AWS Region, regardless of verification
// status.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using ListIdentitiesRequest.
//    req := client.ListIdentitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListIdentities
func (c *Client) ListIdentitiesRequest(input *ListIdentitiesInput) ListIdentitiesRequest {
	op := &aws.Operation{
		Name:       opListIdentities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListIdentitiesInput{}
	}

	req := c.newRequest(op, input, &ListIdentitiesOutput{})
	return ListIdentitiesRequest{Request: req, Input: input, Copy: c.ListIdentitiesRequest}
}

// ListIdentitiesRequest is the request type for the
// ListIdentities API operation.
type ListIdentitiesRequest struct {
	*aws.Request
	Input *ListIdentitiesInput
	Copy  func(*ListIdentitiesInput) ListIdentitiesRequest
}

// Send marshals and sends the ListIdentities API request.
func (r ListIdentitiesRequest) Send(ctx context.Context) (*ListIdentitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIdentitiesResponse{
		ListIdentitiesOutput: r.Request.Data.(*ListIdentitiesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListIdentitiesRequestPaginator returns a paginator for ListIdentities.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListIdentitiesRequest(input)
//   p := ses.NewListIdentitiesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListIdentitiesPaginator(req ListIdentitiesRequest) ListIdentitiesPaginator {
	return ListIdentitiesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListIdentitiesInput
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

// ListIdentitiesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListIdentitiesPaginator struct {
	aws.Pager
}

func (p *ListIdentitiesPaginator) CurrentPage() *ListIdentitiesOutput {
	return p.Pager.CurrentPage().(*ListIdentitiesOutput)
}

// ListIdentitiesResponse is the response type for the
// ListIdentities API operation.
type ListIdentitiesResponse struct {
	*ListIdentitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIdentities request.
func (r *ListIdentitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
