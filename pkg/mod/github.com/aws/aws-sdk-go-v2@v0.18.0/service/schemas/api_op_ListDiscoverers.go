// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDiscoverersInput struct {
	_ struct{} `type:"structure"`

	DiscovererIdPrefix *string `location:"querystring" locationName:"discovererIdPrefix" type:"string"`

	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	SourceArnPrefix *string `location:"querystring" locationName:"sourceArnPrefix" type:"string"`
}

// String returns the string representation
func (s ListDiscoverersInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDiscoverersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DiscovererIdPrefix != nil {
		v := *s.DiscovererIdPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "discovererIdPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArnPrefix != nil {
		v := *s.SourceArnPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "sourceArnPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListDiscoverersOutput struct {
	_ struct{} `type:"structure"`

	Discoverers []DiscovererSummary `type:"list"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDiscoverersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDiscoverersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Discoverers != nil {
		v := s.Discoverers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Discoverers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDiscoverers = "ListDiscoverers"

// ListDiscoverersRequest returns a request value for making API operation for
// Schemas.
//
// List the discoverers.
//
//    // Example sending a request using ListDiscoverersRequest.
//    req := client.ListDiscoverersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/ListDiscoverers
func (c *Client) ListDiscoverersRequest(input *ListDiscoverersInput) ListDiscoverersRequest {
	op := &aws.Operation{
		Name:       opListDiscoverers,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/discoverers",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDiscoverersInput{}
	}

	req := c.newRequest(op, input, &ListDiscoverersOutput{})
	return ListDiscoverersRequest{Request: req, Input: input, Copy: c.ListDiscoverersRequest}
}

// ListDiscoverersRequest is the request type for the
// ListDiscoverers API operation.
type ListDiscoverersRequest struct {
	*aws.Request
	Input *ListDiscoverersInput
	Copy  func(*ListDiscoverersInput) ListDiscoverersRequest
}

// Send marshals and sends the ListDiscoverers API request.
func (r ListDiscoverersRequest) Send(ctx context.Context) (*ListDiscoverersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDiscoverersResponse{
		ListDiscoverersOutput: r.Request.Data.(*ListDiscoverersOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDiscoverersRequestPaginator returns a paginator for ListDiscoverers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDiscoverersRequest(input)
//   p := schemas.NewListDiscoverersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDiscoverersPaginator(req ListDiscoverersRequest) ListDiscoverersPaginator {
	return ListDiscoverersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDiscoverersInput
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

// ListDiscoverersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDiscoverersPaginator struct {
	aws.Pager
}

func (p *ListDiscoverersPaginator) CurrentPage() *ListDiscoverersOutput {
	return p.Pager.CurrentPage().(*ListDiscoverersOutput)
}

// ListDiscoverersResponse is the response type for the
// ListDiscoverers API operation.
type ListDiscoverersResponse struct {
	*ListDiscoverersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDiscoverers request.
func (r *ListDiscoverersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
