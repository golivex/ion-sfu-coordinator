// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakera2iruntime

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListHumanLoopsInput struct {
	_ struct{} `type:"structure"`

	// (Optional) The timestamp of the date when you want the human loops to begin.
	// For example, 1551000000.
	CreationTimeAfter *time.Time `location:"querystring" locationName:"CreationTimeAfter" type:"timestamp"`

	// (Optional) The timestamp of the date before which you want the human loops
	// to begin. For example, 1550000000.
	CreationTimeBefore *time.Time `location:"querystring" locationName:"CreationTimeBefore" type:"timestamp"`

	// The total number of items to return. If the total number of available items
	// is more than the value specified in MaxResults, then a NextToken will be
	// provided in the output that you can use to resume pagination.
	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	// A token to resume pagination.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// An optional value that specifies whether you want the results sorted in Ascending
	// or Descending order.
	SortOrder SortOrder `location:"querystring" locationName:"SortOrder" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListHumanLoopsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListHumanLoopsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListHumanLoopsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListHumanLoopsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CreationTimeAfter != nil {
		v := *s.CreationTimeAfter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "CreationTimeAfter",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.CreationTimeBefore != nil {
		v := *s.CreationTimeBefore

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "CreationTimeBefore",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.SortOrder) > 0 {
		v := s.SortOrder

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "SortOrder", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type ListHumanLoopsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects containing information about the human loops.
	//
	// HumanLoopSummaries is a required field
	HumanLoopSummaries []HumanLoopSummary `type:"list" required:"true"`

	// A token to resume pagination.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHumanLoopsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListHumanLoopsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HumanLoopSummaries != nil {
		v := s.HumanLoopSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "HumanLoopSummaries", metadata)
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

const opListHumanLoops = "ListHumanLoops"

// ListHumanLoopsRequest returns a request value for making API operation for
// Amazon Augmented AI Runtime.
//
// Returns information about human loops, given the specified parameters.
//
//    // Example sending a request using ListHumanLoopsRequest.
//    req := client.ListHumanLoopsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-a2i-runtime-2019-11-07/ListHumanLoops
func (c *Client) ListHumanLoopsRequest(input *ListHumanLoopsInput) ListHumanLoopsRequest {
	op := &aws.Operation{
		Name:       opListHumanLoops,
		HTTPMethod: "GET",
		HTTPPath:   "/human-loops",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListHumanLoopsInput{}
	}

	req := c.newRequest(op, input, &ListHumanLoopsOutput{})
	return ListHumanLoopsRequest{Request: req, Input: input, Copy: c.ListHumanLoopsRequest}
}

// ListHumanLoopsRequest is the request type for the
// ListHumanLoops API operation.
type ListHumanLoopsRequest struct {
	*aws.Request
	Input *ListHumanLoopsInput
	Copy  func(*ListHumanLoopsInput) ListHumanLoopsRequest
}

// Send marshals and sends the ListHumanLoops API request.
func (r ListHumanLoopsRequest) Send(ctx context.Context) (*ListHumanLoopsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHumanLoopsResponse{
		ListHumanLoopsOutput: r.Request.Data.(*ListHumanLoopsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListHumanLoopsRequestPaginator returns a paginator for ListHumanLoops.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListHumanLoopsRequest(input)
//   p := sagemakera2iruntime.NewListHumanLoopsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListHumanLoopsPaginator(req ListHumanLoopsRequest) ListHumanLoopsPaginator {
	return ListHumanLoopsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListHumanLoopsInput
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

// ListHumanLoopsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListHumanLoopsPaginator struct {
	aws.Pager
}

func (p *ListHumanLoopsPaginator) CurrentPage() *ListHumanLoopsOutput {
	return p.Pager.CurrentPage().(*ListHumanLoopsOutput)
}

// ListHumanLoopsResponse is the response type for the
// ListHumanLoops API operation.
type ListHumanLoopsResponse struct {
	*ListHumanLoopsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHumanLoops request.
func (r *ListHumanLoopsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
