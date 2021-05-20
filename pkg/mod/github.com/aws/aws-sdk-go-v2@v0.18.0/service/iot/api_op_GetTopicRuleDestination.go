// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetTopicRuleDestinationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the topic rule destination.
	//
	// Arn is a required field
	Arn *string `location:"uri" locationName:"arn" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTopicRuleDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTopicRuleDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTopicRuleDestinationInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTopicRuleDestinationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetTopicRuleDestinationOutput struct {
	_ struct{} `type:"structure"`

	// The topic rule destination.
	TopicRuleDestination *TopicRuleDestination `locationName:"topicRuleDestination" type:"structure"`
}

// String returns the string representation
func (s GetTopicRuleDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTopicRuleDestinationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TopicRuleDestination != nil {
		v := s.TopicRuleDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "topicRuleDestination", v, metadata)
	}
	return nil
}

const opGetTopicRuleDestination = "GetTopicRuleDestination"

// GetTopicRuleDestinationRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about a topic rule destination.
//
//    // Example sending a request using GetTopicRuleDestinationRequest.
//    req := client.GetTopicRuleDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetTopicRuleDestinationRequest(input *GetTopicRuleDestinationInput) GetTopicRuleDestinationRequest {
	op := &aws.Operation{
		Name:       opGetTopicRuleDestination,
		HTTPMethod: "GET",
		HTTPPath:   "/destinations/{arn+}",
	}

	if input == nil {
		input = &GetTopicRuleDestinationInput{}
	}

	req := c.newRequest(op, input, &GetTopicRuleDestinationOutput{})
	return GetTopicRuleDestinationRequest{Request: req, Input: input, Copy: c.GetTopicRuleDestinationRequest}
}

// GetTopicRuleDestinationRequest is the request type for the
// GetTopicRuleDestination API operation.
type GetTopicRuleDestinationRequest struct {
	*aws.Request
	Input *GetTopicRuleDestinationInput
	Copy  func(*GetTopicRuleDestinationInput) GetTopicRuleDestinationRequest
}

// Send marshals and sends the GetTopicRuleDestination API request.
func (r GetTopicRuleDestinationRequest) Send(ctx context.Context) (*GetTopicRuleDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTopicRuleDestinationResponse{
		GetTopicRuleDestinationOutput: r.Request.Data.(*GetTopicRuleDestinationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTopicRuleDestinationResponse is the response type for the
// GetTopicRuleDestination API operation.
type GetTopicRuleDestinationResponse struct {
	*GetTopicRuleDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTopicRuleDestination request.
func (r *GetTopicRuleDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
