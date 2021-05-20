// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StartMonitoringScheduleInput struct {
	_ struct{} `type:"structure"`

	// The name of the schedule to start.
	//
	// MonitoringScheduleName is a required field
	MonitoringScheduleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartMonitoringScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartMonitoringScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartMonitoringScheduleInput"}

	if s.MonitoringScheduleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MonitoringScheduleName"))
	}
	if s.MonitoringScheduleName != nil && len(*s.MonitoringScheduleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MonitoringScheduleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartMonitoringScheduleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartMonitoringScheduleOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartMonitoringSchedule = "StartMonitoringSchedule"

// StartMonitoringScheduleRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Starts a previously stopped monitoring schedule.
//
// New monitoring schedules are immediately started after creation.
//
//    // Example sending a request using StartMonitoringScheduleRequest.
//    req := client.StartMonitoringScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StartMonitoringSchedule
func (c *Client) StartMonitoringScheduleRequest(input *StartMonitoringScheduleInput) StartMonitoringScheduleRequest {
	op := &aws.Operation{
		Name:       opStartMonitoringSchedule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartMonitoringScheduleInput{}
	}

	req := c.newRequest(op, input, &StartMonitoringScheduleOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StartMonitoringScheduleRequest{Request: req, Input: input, Copy: c.StartMonitoringScheduleRequest}
}

// StartMonitoringScheduleRequest is the request type for the
// StartMonitoringSchedule API operation.
type StartMonitoringScheduleRequest struct {
	*aws.Request
	Input *StartMonitoringScheduleInput
	Copy  func(*StartMonitoringScheduleInput) StartMonitoringScheduleRequest
}

// Send marshals and sends the StartMonitoringSchedule API request.
func (r StartMonitoringScheduleRequest) Send(ctx context.Context) (*StartMonitoringScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartMonitoringScheduleResponse{
		StartMonitoringScheduleOutput: r.Request.Data.(*StartMonitoringScheduleOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartMonitoringScheduleResponse is the response type for the
// StartMonitoringSchedule API operation.
type StartMonitoringScheduleResponse struct {
	*StartMonitoringScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartMonitoringSchedule request.
func (r *StartMonitoringScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
