// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetVoiceConnectorGroupInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector group ID.
	//
	// VoiceConnectorGroupId is a required field
	VoiceConnectorGroupId *string `location:"uri" locationName:"voiceConnectorGroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVoiceConnectorGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVoiceConnectorGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVoiceConnectorGroupInput"}

	if s.VoiceConnectorGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VoiceConnectorGroupId != nil {
		v := *s.VoiceConnectorGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetVoiceConnectorGroupOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector group details.
	VoiceConnectorGroup *VoiceConnectorGroup `type:"structure"`
}

// String returns the string representation
func (s GetVoiceConnectorGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VoiceConnectorGroup != nil {
		v := s.VoiceConnectorGroup

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VoiceConnectorGroup", v, metadata)
	}
	return nil
}

const opGetVoiceConnectorGroup = "GetVoiceConnectorGroup"

// GetVoiceConnectorGroupRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves details for the specified Amazon Chime Voice Connector group, such
// as timestamps, name, and associated VoiceConnectorItems.
//
//    // Example sending a request using GetVoiceConnectorGroupRequest.
//    req := client.GetVoiceConnectorGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetVoiceConnectorGroup
func (c *Client) GetVoiceConnectorGroupRequest(input *GetVoiceConnectorGroupInput) GetVoiceConnectorGroupRequest {
	op := &aws.Operation{
		Name:       opGetVoiceConnectorGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/voice-connector-groups/{voiceConnectorGroupId}",
	}

	if input == nil {
		input = &GetVoiceConnectorGroupInput{}
	}

	req := c.newRequest(op, input, &GetVoiceConnectorGroupOutput{})
	return GetVoiceConnectorGroupRequest{Request: req, Input: input, Copy: c.GetVoiceConnectorGroupRequest}
}

// GetVoiceConnectorGroupRequest is the request type for the
// GetVoiceConnectorGroup API operation.
type GetVoiceConnectorGroupRequest struct {
	*aws.Request
	Input *GetVoiceConnectorGroupInput
	Copy  func(*GetVoiceConnectorGroupInput) GetVoiceConnectorGroupRequest
}

// Send marshals and sends the GetVoiceConnectorGroup API request.
func (r GetVoiceConnectorGroupRequest) Send(ctx context.Context) (*GetVoiceConnectorGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVoiceConnectorGroupResponse{
		GetVoiceConnectorGroupOutput: r.Request.Data.(*GetVoiceConnectorGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVoiceConnectorGroupResponse is the response type for the
// GetVoiceConnectorGroup API operation.
type GetVoiceConnectorGroupResponse struct {
	*GetVoiceConnectorGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVoiceConnectorGroup request.
func (r *GetVoiceConnectorGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
