// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteVoiceConnectorStreamingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVoiceConnectorStreamingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVoiceConnectorStreamingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVoiceConnectorStreamingConfigurationInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorStreamingConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVoiceConnectorStreamingConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVoiceConnectorStreamingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorStreamingConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteVoiceConnectorStreamingConfiguration = "DeleteVoiceConnectorStreamingConfiguration"

// DeleteVoiceConnectorStreamingConfigurationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the streaming configuration for the specified Amazon Chime Voice
// Connector.
//
//    // Example sending a request using DeleteVoiceConnectorStreamingConfigurationRequest.
//    req := client.DeleteVoiceConnectorStreamingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteVoiceConnectorStreamingConfiguration
func (c *Client) DeleteVoiceConnectorStreamingConfigurationRequest(input *DeleteVoiceConnectorStreamingConfigurationInput) DeleteVoiceConnectorStreamingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteVoiceConnectorStreamingConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/streaming-configuration",
	}

	if input == nil {
		input = &DeleteVoiceConnectorStreamingConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteVoiceConnectorStreamingConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVoiceConnectorStreamingConfigurationRequest{Request: req, Input: input, Copy: c.DeleteVoiceConnectorStreamingConfigurationRequest}
}

// DeleteVoiceConnectorStreamingConfigurationRequest is the request type for the
// DeleteVoiceConnectorStreamingConfiguration API operation.
type DeleteVoiceConnectorStreamingConfigurationRequest struct {
	*aws.Request
	Input *DeleteVoiceConnectorStreamingConfigurationInput
	Copy  func(*DeleteVoiceConnectorStreamingConfigurationInput) DeleteVoiceConnectorStreamingConfigurationRequest
}

// Send marshals and sends the DeleteVoiceConnectorStreamingConfiguration API request.
func (r DeleteVoiceConnectorStreamingConfigurationRequest) Send(ctx context.Context) (*DeleteVoiceConnectorStreamingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVoiceConnectorStreamingConfigurationResponse{
		DeleteVoiceConnectorStreamingConfigurationOutput: r.Request.Data.(*DeleteVoiceConnectorStreamingConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVoiceConnectorStreamingConfigurationResponse is the response type for the
// DeleteVoiceConnectorStreamingConfiguration API operation.
type DeleteVoiceConnectorStreamingConfigurationResponse struct {
	*DeleteVoiceConnectorStreamingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVoiceConnectorStreamingConfiguration request.
func (r *DeleteVoiceConnectorStreamingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
