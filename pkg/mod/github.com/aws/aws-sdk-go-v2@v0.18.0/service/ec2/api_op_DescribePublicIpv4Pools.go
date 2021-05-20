// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribePublicIpv4PoolsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The IDs of the address pools.
	PoolIds []string `locationName:"PoolId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribePublicIpv4PoolsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePublicIpv4PoolsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePublicIpv4PoolsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribePublicIpv4PoolsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the address pools.
	PublicIpv4Pools []PublicIpv4Pool `locationName:"publicIpv4PoolSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribePublicIpv4PoolsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePublicIpv4Pools = "DescribePublicIpv4Pools"

// DescribePublicIpv4PoolsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified IPv4 address pools.
//
//    // Example sending a request using DescribePublicIpv4PoolsRequest.
//    req := client.DescribePublicIpv4PoolsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePublicIpv4Pools
func (c *Client) DescribePublicIpv4PoolsRequest(input *DescribePublicIpv4PoolsInput) DescribePublicIpv4PoolsRequest {
	op := &aws.Operation{
		Name:       opDescribePublicIpv4Pools,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribePublicIpv4PoolsInput{}
	}

	req := c.newRequest(op, input, &DescribePublicIpv4PoolsOutput{})
	return DescribePublicIpv4PoolsRequest{Request: req, Input: input, Copy: c.DescribePublicIpv4PoolsRequest}
}

// DescribePublicIpv4PoolsRequest is the request type for the
// DescribePublicIpv4Pools API operation.
type DescribePublicIpv4PoolsRequest struct {
	*aws.Request
	Input *DescribePublicIpv4PoolsInput
	Copy  func(*DescribePublicIpv4PoolsInput) DescribePublicIpv4PoolsRequest
}

// Send marshals and sends the DescribePublicIpv4Pools API request.
func (r DescribePublicIpv4PoolsRequest) Send(ctx context.Context) (*DescribePublicIpv4PoolsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePublicIpv4PoolsResponse{
		DescribePublicIpv4PoolsOutput: r.Request.Data.(*DescribePublicIpv4PoolsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribePublicIpv4PoolsRequestPaginator returns a paginator for DescribePublicIpv4Pools.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribePublicIpv4PoolsRequest(input)
//   p := ec2.NewDescribePublicIpv4PoolsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribePublicIpv4PoolsPaginator(req DescribePublicIpv4PoolsRequest) DescribePublicIpv4PoolsPaginator {
	return DescribePublicIpv4PoolsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribePublicIpv4PoolsInput
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

// DescribePublicIpv4PoolsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribePublicIpv4PoolsPaginator struct {
	aws.Pager
}

func (p *DescribePublicIpv4PoolsPaginator) CurrentPage() *DescribePublicIpv4PoolsOutput {
	return p.Pager.CurrentPage().(*DescribePublicIpv4PoolsOutput)
}

// DescribePublicIpv4PoolsResponse is the response type for the
// DescribePublicIpv4Pools API operation.
type DescribePublicIpv4PoolsResponse struct {
	*DescribePublicIpv4PoolsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePublicIpv4Pools request.
func (r *DescribePublicIpv4PoolsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
