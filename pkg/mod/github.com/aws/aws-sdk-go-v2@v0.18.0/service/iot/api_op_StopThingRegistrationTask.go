// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type StopThingRegistrationTaskInput struct {
	_ struct{} `type:"structure"`

	// The bulk thing provisioning task ID.
	//
	// TaskId is a required field
	TaskId *string `location:"uri" locationName:"taskId" type:"string" required:"true"`
}

// String returns the string representation
func (s StopThingRegistrationTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopThingRegistrationTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopThingRegistrationTaskInput"}

	if s.TaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopThingRegistrationTaskInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TaskId != nil {
		v := *s.TaskId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "taskId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StopThingRegistrationTaskOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopThingRegistrationTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopThingRegistrationTaskOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opStopThingRegistrationTask = "StopThingRegistrationTask"

// StopThingRegistrationTaskRequest returns a request value for making API operation for
// AWS IoT.
//
// Cancels a bulk thing provisioning task.
//
//    // Example sending a request using StopThingRegistrationTaskRequest.
//    req := client.StopThingRegistrationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StopThingRegistrationTaskRequest(input *StopThingRegistrationTaskInput) StopThingRegistrationTaskRequest {
	op := &aws.Operation{
		Name:       opStopThingRegistrationTask,
		HTTPMethod: "PUT",
		HTTPPath:   "/thing-registration-tasks/{taskId}/cancel",
	}

	if input == nil {
		input = &StopThingRegistrationTaskInput{}
	}

	req := c.newRequest(op, input, &StopThingRegistrationTaskOutput{})
	return StopThingRegistrationTaskRequest{Request: req, Input: input, Copy: c.StopThingRegistrationTaskRequest}
}

// StopThingRegistrationTaskRequest is the request type for the
// StopThingRegistrationTask API operation.
type StopThingRegistrationTaskRequest struct {
	*aws.Request
	Input *StopThingRegistrationTaskInput
	Copy  func(*StopThingRegistrationTaskInput) StopThingRegistrationTaskRequest
}

// Send marshals and sends the StopThingRegistrationTask API request.
func (r StopThingRegistrationTaskRequest) Send(ctx context.Context) (*StopThingRegistrationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopThingRegistrationTaskResponse{
		StopThingRegistrationTaskOutput: r.Request.Data.(*StopThingRegistrationTaskOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopThingRegistrationTaskResponse is the response type for the
// StopThingRegistrationTask API operation.
type StopThingRegistrationTaskResponse struct {
	*StopThingRegistrationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopThingRegistrationTask request.
func (r *StopThingRegistrationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
