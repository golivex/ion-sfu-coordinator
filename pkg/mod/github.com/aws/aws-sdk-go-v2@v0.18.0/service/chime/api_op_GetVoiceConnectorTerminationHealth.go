// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetVoiceConnectorTerminationHealthInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVoiceConnectorTerminationHealthInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVoiceConnectorTerminationHealthInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVoiceConnectorTerminationHealthInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorTerminationHealthInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetVoiceConnectorTerminationHealthOutput struct {
	_ struct{} `type:"structure"`

	// The termination health details.
	TerminationHealth *TerminationHealth `type:"structure"`
}

// String returns the string representation
func (s GetVoiceConnectorTerminationHealthOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorTerminationHealthOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TerminationHealth != nil {
		v := s.TerminationHealth

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TerminationHealth", v, metadata)
	}
	return nil
}

const opGetVoiceConnectorTerminationHealth = "GetVoiceConnectorTerminationHealth"

// GetVoiceConnectorTerminationHealthRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves information about the last time a SIP OPTIONS ping was received
// from your SIP infrastructure for the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using GetVoiceConnectorTerminationHealthRequest.
//    req := client.GetVoiceConnectorTerminationHealthRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetVoiceConnectorTerminationHealth
func (c *Client) GetVoiceConnectorTerminationHealthRequest(input *GetVoiceConnectorTerminationHealthInput) GetVoiceConnectorTerminationHealthRequest {
	op := &aws.Operation{
		Name:       opGetVoiceConnectorTerminationHealth,
		HTTPMethod: "GET",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/termination/health",
	}

	if input == nil {
		input = &GetVoiceConnectorTerminationHealthInput{}
	}

	req := c.newRequest(op, input, &GetVoiceConnectorTerminationHealthOutput{})
	return GetVoiceConnectorTerminationHealthRequest{Request: req, Input: input, Copy: c.GetVoiceConnectorTerminationHealthRequest}
}

// GetVoiceConnectorTerminationHealthRequest is the request type for the
// GetVoiceConnectorTerminationHealth API operation.
type GetVoiceConnectorTerminationHealthRequest struct {
	*aws.Request
	Input *GetVoiceConnectorTerminationHealthInput
	Copy  func(*GetVoiceConnectorTerminationHealthInput) GetVoiceConnectorTerminationHealthRequest
}

// Send marshals and sends the GetVoiceConnectorTerminationHealth API request.
func (r GetVoiceConnectorTerminationHealthRequest) Send(ctx context.Context) (*GetVoiceConnectorTerminationHealthResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVoiceConnectorTerminationHealthResponse{
		GetVoiceConnectorTerminationHealthOutput: r.Request.Data.(*GetVoiceConnectorTerminationHealthOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVoiceConnectorTerminationHealthResponse is the response type for the
// GetVoiceConnectorTerminationHealth API operation.
type GetVoiceConnectorTerminationHealthResponse struct {
	*GetVoiceConnectorTerminationHealthOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVoiceConnectorTerminationHealth request.
func (r *GetVoiceConnectorTerminationHealthResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
