// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListVoiceConnectorsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The token to use to retrieve the next page of results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`
}

// String returns the string representation
func (s ListVoiceConnectorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVoiceConnectorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVoiceConnectorsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVoiceConnectorsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListVoiceConnectorsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results.
	NextToken *string `type:"string"`

	// The details of the Amazon Chime Voice Connectors.
	VoiceConnectors []VoiceConnector `type:"list"`
}

// String returns the string representation
func (s ListVoiceConnectorsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVoiceConnectorsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VoiceConnectors != nil {
		v := s.VoiceConnectors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "VoiceConnectors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListVoiceConnectors = "ListVoiceConnectors"

// ListVoiceConnectorsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Lists the Amazon Chime Voice Connectors for the administrator's AWS account.
//
//    // Example sending a request using ListVoiceConnectorsRequest.
//    req := client.ListVoiceConnectorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListVoiceConnectors
func (c *Client) ListVoiceConnectorsRequest(input *ListVoiceConnectorsInput) ListVoiceConnectorsRequest {
	op := &aws.Operation{
		Name:       opListVoiceConnectors,
		HTTPMethod: "GET",
		HTTPPath:   "/voice-connectors",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListVoiceConnectorsInput{}
	}

	req := c.newRequest(op, input, &ListVoiceConnectorsOutput{})
	return ListVoiceConnectorsRequest{Request: req, Input: input, Copy: c.ListVoiceConnectorsRequest}
}

// ListVoiceConnectorsRequest is the request type for the
// ListVoiceConnectors API operation.
type ListVoiceConnectorsRequest struct {
	*aws.Request
	Input *ListVoiceConnectorsInput
	Copy  func(*ListVoiceConnectorsInput) ListVoiceConnectorsRequest
}

// Send marshals and sends the ListVoiceConnectors API request.
func (r ListVoiceConnectorsRequest) Send(ctx context.Context) (*ListVoiceConnectorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVoiceConnectorsResponse{
		ListVoiceConnectorsOutput: r.Request.Data.(*ListVoiceConnectorsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListVoiceConnectorsRequestPaginator returns a paginator for ListVoiceConnectors.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListVoiceConnectorsRequest(input)
//   p := chime.NewListVoiceConnectorsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListVoiceConnectorsPaginator(req ListVoiceConnectorsRequest) ListVoiceConnectorsPaginator {
	return ListVoiceConnectorsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListVoiceConnectorsInput
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

// ListVoiceConnectorsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListVoiceConnectorsPaginator struct {
	aws.Pager
}

func (p *ListVoiceConnectorsPaginator) CurrentPage() *ListVoiceConnectorsOutput {
	return p.Pager.CurrentPage().(*ListVoiceConnectorsOutput)
}

// ListVoiceConnectorsResponse is the response type for the
// ListVoiceConnectors API operation.
type ListVoiceConnectorsResponse struct {
	*ListVoiceConnectorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVoiceConnectors request.
func (r *ListVoiceConnectorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
