// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type UpdateUserRoutingProfileInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The identifier of the routing profile for the user.
	//
	// RoutingProfileId is a required field
	RoutingProfileId *string `type:"string" required:"true"`

	// The identifier of the user account.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"UserId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUserRoutingProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserRoutingProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserRoutingProfileInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if s.RoutingProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoutingProfileId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserRoutingProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RoutingProfileId != nil {
		v := *s.RoutingProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RoutingProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UserId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateUserRoutingProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateUserRoutingProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserRoutingProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateUserRoutingProfile = "UpdateUserRoutingProfile"

// UpdateUserRoutingProfileRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Assigns the specified routing profile to the specified user.
//
//    // Example sending a request using UpdateUserRoutingProfileRequest.
//    req := client.UpdateUserRoutingProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/UpdateUserRoutingProfile
func (c *Client) UpdateUserRoutingProfileRequest(input *UpdateUserRoutingProfileInput) UpdateUserRoutingProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateUserRoutingProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/users/{InstanceId}/{UserId}/routing-profile",
	}

	if input == nil {
		input = &UpdateUserRoutingProfileInput{}
	}

	req := c.newRequest(op, input, &UpdateUserRoutingProfileOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateUserRoutingProfileRequest{Request: req, Input: input, Copy: c.UpdateUserRoutingProfileRequest}
}

// UpdateUserRoutingProfileRequest is the request type for the
// UpdateUserRoutingProfile API operation.
type UpdateUserRoutingProfileRequest struct {
	*aws.Request
	Input *UpdateUserRoutingProfileInput
	Copy  func(*UpdateUserRoutingProfileInput) UpdateUserRoutingProfileRequest
}

// Send marshals and sends the UpdateUserRoutingProfile API request.
func (r UpdateUserRoutingProfileRequest) Send(ctx context.Context) (*UpdateUserRoutingProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserRoutingProfileResponse{
		UpdateUserRoutingProfileOutput: r.Request.Data.(*UpdateUserRoutingProfileOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserRoutingProfileResponse is the response type for the
// UpdateUserRoutingProfile API operation.
type UpdateUserRoutingProfileResponse struct {
	*UpdateUserRoutingProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserRoutingProfile request.
func (r *UpdateUserRoutingProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
