// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for delete backend environment request.
type DeleteBackendEnvironmentInput struct {
	_ struct{} `type:"structure"`

	// Unique Id of an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name of a backend environment of an Amplify App.
	//
	// EnvironmentName is a required field
	EnvironmentName *string `location:"uri" locationName:"environmentName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBackendEnvironmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBackendEnvironmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBackendEnvironmentInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.EnvironmentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentName"))
	}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackendEnvironmentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnvironmentName != nil {
		v := *s.EnvironmentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "environmentName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure of a delete backend environment result.
type DeleteBackendEnvironmentOutput struct {
	_ struct{} `type:"structure"`

	// Backend environment structure for an Amplify App.
	//
	// BackendEnvironment is a required field
	BackendEnvironment *BackendEnvironment `locationName:"backendEnvironment" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteBackendEnvironmentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackendEnvironmentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackendEnvironment != nil {
		v := s.BackendEnvironment

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "backendEnvironment", v, metadata)
	}
	return nil
}

const opDeleteBackendEnvironment = "DeleteBackendEnvironment"

// DeleteBackendEnvironmentRequest returns a request value for making API operation for
// AWS Amplify.
//
// Delete backend environment for an Amplify App.
//
//    // Example sending a request using DeleteBackendEnvironmentRequest.
//    req := client.DeleteBackendEnvironmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteBackendEnvironment
func (c *Client) DeleteBackendEnvironmentRequest(input *DeleteBackendEnvironmentInput) DeleteBackendEnvironmentRequest {
	op := &aws.Operation{
		Name:       opDeleteBackendEnvironment,
		HTTPMethod: "DELETE",
		HTTPPath:   "/apps/{appId}/backendenvironments/{environmentName}",
	}

	if input == nil {
		input = &DeleteBackendEnvironmentInput{}
	}

	req := c.newRequest(op, input, &DeleteBackendEnvironmentOutput{})
	return DeleteBackendEnvironmentRequest{Request: req, Input: input, Copy: c.DeleteBackendEnvironmentRequest}
}

// DeleteBackendEnvironmentRequest is the request type for the
// DeleteBackendEnvironment API operation.
type DeleteBackendEnvironmentRequest struct {
	*aws.Request
	Input *DeleteBackendEnvironmentInput
	Copy  func(*DeleteBackendEnvironmentInput) DeleteBackendEnvironmentRequest
}

// Send marshals and sends the DeleteBackendEnvironment API request.
func (r DeleteBackendEnvironmentRequest) Send(ctx context.Context) (*DeleteBackendEnvironmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackendEnvironmentResponse{
		DeleteBackendEnvironmentOutput: r.Request.Data.(*DeleteBackendEnvironmentOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackendEnvironmentResponse is the response type for the
// DeleteBackendEnvironment API operation.
type DeleteBackendEnvironmentResponse struct {
	*DeleteBackendEnvironmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackendEnvironment request.
func (r *DeleteBackendEnvironmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
