// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartAvailabilityMonitorTestInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s StartAvailabilityMonitorTestInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartAvailabilityMonitorTestInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartAvailabilityMonitorTestInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartAvailabilityMonitorTestOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s StartAvailabilityMonitorTestOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartAvailabilityMonitorTest = "StartAvailabilityMonitorTest"

// StartAvailabilityMonitorTestRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Start a test that verifies that the specified gateway is configured for High
// Availability monitoring in your host environment. This request only initiates
// the test and that a successful response only indicates that the test was
// started. It doesn't indicate that the test passed. For the status of the
// test, invoke the DescribeAvailabilityMonitorTest API.
//
// Starting this test will cause your gateway to go offline for a brief period.
//
//    // Example sending a request using StartAvailabilityMonitorTestRequest.
//    req := client.StartAvailabilityMonitorTestRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/StartAvailabilityMonitorTest
func (c *Client) StartAvailabilityMonitorTestRequest(input *StartAvailabilityMonitorTestInput) StartAvailabilityMonitorTestRequest {
	op := &aws.Operation{
		Name:       opStartAvailabilityMonitorTest,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartAvailabilityMonitorTestInput{}
	}

	req := c.newRequest(op, input, &StartAvailabilityMonitorTestOutput{})
	return StartAvailabilityMonitorTestRequest{Request: req, Input: input, Copy: c.StartAvailabilityMonitorTestRequest}
}

// StartAvailabilityMonitorTestRequest is the request type for the
// StartAvailabilityMonitorTest API operation.
type StartAvailabilityMonitorTestRequest struct {
	*aws.Request
	Input *StartAvailabilityMonitorTestInput
	Copy  func(*StartAvailabilityMonitorTestInput) StartAvailabilityMonitorTestRequest
}

// Send marshals and sends the StartAvailabilityMonitorTest API request.
func (r StartAvailabilityMonitorTestRequest) Send(ctx context.Context) (*StartAvailabilityMonitorTestResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartAvailabilityMonitorTestResponse{
		StartAvailabilityMonitorTestOutput: r.Request.Data.(*StartAvailabilityMonitorTestOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartAvailabilityMonitorTestResponse is the response type for the
// StartAvailabilityMonitorTest API operation.
type StartAvailabilityMonitorTestResponse struct {
	*StartAvailabilityMonitorTestOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartAvailabilityMonitorTest request.
func (r *StartAvailabilityMonitorTestResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
