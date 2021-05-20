// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the SetIdentityPoolConfiguration operation.
type SetIdentityPoolConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Options to apply to this identity pool for Amazon Cognito streams.
	CognitoStreams *CognitoStreams `type:"structure"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. This is the ID of the pool to modify.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `location:"uri" locationName:"IdentityPoolId" min:"1" type:"string" required:"true"`

	// Options to apply to this identity pool for push synchronization.
	PushSync *PushSync `type:"structure"`
}

// String returns the string representation
func (s SetIdentityPoolConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetIdentityPoolConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetIdentityPoolConfigurationInput"}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}
	if s.CognitoStreams != nil {
		if err := s.CognitoStreams.Validate(); err != nil {
			invalidParams.AddNested("CognitoStreams", err.(aws.ErrInvalidParams))
		}
	}
	if s.PushSync != nil {
		if err := s.PushSync.Validate(); err != nil {
			invalidParams.AddNested("PushSync", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetIdentityPoolConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CognitoStreams != nil {
		v := s.CognitoStreams

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CognitoStreams", v, metadata)
	}
	if s.PushSync != nil {
		v := s.PushSync

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PushSync", v, metadata)
	}
	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output for the SetIdentityPoolConfiguration operation
type SetIdentityPoolConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Options to apply to this identity pool for Amazon Cognito streams.
	CognitoStreams *CognitoStreams `type:"structure"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito.
	IdentityPoolId *string `min:"1" type:"string"`

	// Options to apply to this identity pool for push synchronization.
	PushSync *PushSync `type:"structure"`
}

// String returns the string representation
func (s SetIdentityPoolConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetIdentityPoolConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CognitoStreams != nil {
		v := s.CognitoStreams

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CognitoStreams", v, metadata)
	}
	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PushSync != nil {
		v := s.PushSync

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PushSync", v, metadata)
	}
	return nil
}

const opSetIdentityPoolConfiguration = "SetIdentityPoolConfiguration"

// SetIdentityPoolConfigurationRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Sets the necessary configuration for push sync.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
//
//    // Example sending a request using SetIdentityPoolConfigurationRequest.
//    req := client.SetIdentityPoolConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/SetIdentityPoolConfiguration
func (c *Client) SetIdentityPoolConfigurationRequest(input *SetIdentityPoolConfigurationInput) SetIdentityPoolConfigurationRequest {
	op := &aws.Operation{
		Name:       opSetIdentityPoolConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/identitypools/{IdentityPoolId}/configuration",
	}

	if input == nil {
		input = &SetIdentityPoolConfigurationInput{}
	}

	req := c.newRequest(op, input, &SetIdentityPoolConfigurationOutput{})
	return SetIdentityPoolConfigurationRequest{Request: req, Input: input, Copy: c.SetIdentityPoolConfigurationRequest}
}

// SetIdentityPoolConfigurationRequest is the request type for the
// SetIdentityPoolConfiguration API operation.
type SetIdentityPoolConfigurationRequest struct {
	*aws.Request
	Input *SetIdentityPoolConfigurationInput
	Copy  func(*SetIdentityPoolConfigurationInput) SetIdentityPoolConfigurationRequest
}

// Send marshals and sends the SetIdentityPoolConfiguration API request.
func (r SetIdentityPoolConfigurationRequest) Send(ctx context.Context) (*SetIdentityPoolConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetIdentityPoolConfigurationResponse{
		SetIdentityPoolConfigurationOutput: r.Request.Data.(*SetIdentityPoolConfigurationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetIdentityPoolConfigurationResponse is the response type for the
// SetIdentityPoolConfiguration API operation.
type SetIdentityPoolConfigurationResponse struct {
	*SetIdentityPoolConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetIdentityPoolConfiguration request.
func (r *SetIdentityPoolConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
