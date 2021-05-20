// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartAccessLoggingInput struct {
	_ struct{} `type:"structure"`

	// The name of the container that you want to start access logging on.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartAccessLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartAccessLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartAccessLoggingInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartAccessLoggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartAccessLoggingOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartAccessLogging = "StartAccessLogging"

// StartAccessLoggingRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Starts access logging on the specified container. When you enable access
// logging on a container, MediaStore delivers access logs for objects stored
// in that container to Amazon CloudWatch Logs.
//
//    // Example sending a request using StartAccessLoggingRequest.
//    req := client.StartAccessLoggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/StartAccessLogging
func (c *Client) StartAccessLoggingRequest(input *StartAccessLoggingInput) StartAccessLoggingRequest {
	op := &aws.Operation{
		Name:       opStartAccessLogging,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartAccessLoggingInput{}
	}

	req := c.newRequest(op, input, &StartAccessLoggingOutput{})
	return StartAccessLoggingRequest{Request: req, Input: input, Copy: c.StartAccessLoggingRequest}
}

// StartAccessLoggingRequest is the request type for the
// StartAccessLogging API operation.
type StartAccessLoggingRequest struct {
	*aws.Request
	Input *StartAccessLoggingInput
	Copy  func(*StartAccessLoggingInput) StartAccessLoggingRequest
}

// Send marshals and sends the StartAccessLogging API request.
func (r StartAccessLoggingRequest) Send(ctx context.Context) (*StartAccessLoggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartAccessLoggingResponse{
		StartAccessLoggingOutput: r.Request.Data.(*StartAccessLoggingOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartAccessLoggingResponse is the response type for the
// StartAccessLogging API operation.
type StartAccessLoggingResponse struct {
	*StartAccessLoggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartAccessLogging request.
func (r *StartAccessLoggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
