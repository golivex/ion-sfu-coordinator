// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopSentimentDetectionJobInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the sentiment detection job to stop.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopSentimentDetectionJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopSentimentDetectionJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopSentimentDetectionJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopSentimentDetectionJobOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the sentiment detection job to stop.
	JobId *string `min:"1" type:"string"`

	// Either STOP_REQUESTED if the job is currently running, or STOPPED if the
	// job was previously stopped with the StopSentimentDetectionJob operation.
	JobStatus JobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s StopSentimentDetectionJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopSentimentDetectionJob = "StopSentimentDetectionJob"

// StopSentimentDetectionJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Stops a sentiment detection job in progress.
//
// If the job state is IN_PROGRESS the job is marked for termination and put
// into the STOP_REQUESTED state. If the job completes before it can be stopped,
// it is put into the COMPLETED state; otherwise the job is be stopped and put
// into the STOPPED state.
//
// If the job is in the COMPLETED or FAILED state when you call the StopDominantLanguageDetectionJob
// operation, the operation returns a 400 Internal Request Exception.
//
// When a job is stopped, any documents already processed are written to the
// output location.
//
//    // Example sending a request using StopSentimentDetectionJobRequest.
//    req := client.StopSentimentDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/StopSentimentDetectionJob
func (c *Client) StopSentimentDetectionJobRequest(input *StopSentimentDetectionJobInput) StopSentimentDetectionJobRequest {
	op := &aws.Operation{
		Name:       opStopSentimentDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopSentimentDetectionJobInput{}
	}

	req := c.newRequest(op, input, &StopSentimentDetectionJobOutput{})
	return StopSentimentDetectionJobRequest{Request: req, Input: input, Copy: c.StopSentimentDetectionJobRequest}
}

// StopSentimentDetectionJobRequest is the request type for the
// StopSentimentDetectionJob API operation.
type StopSentimentDetectionJobRequest struct {
	*aws.Request
	Input *StopSentimentDetectionJobInput
	Copy  func(*StopSentimentDetectionJobInput) StopSentimentDetectionJobRequest
}

// Send marshals and sends the StopSentimentDetectionJob API request.
func (r StopSentimentDetectionJobRequest) Send(ctx context.Context) (*StopSentimentDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopSentimentDetectionJobResponse{
		StopSentimentDetectionJobOutput: r.Request.Data.(*StopSentimentDetectionJobOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopSentimentDetectionJobResponse is the response type for the
// StopSentimentDetectionJob API operation.
type StopSentimentDetectionJobResponse struct {
	*StopSentimentDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopSentimentDetectionJob request.
func (r *StopSentimentDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
