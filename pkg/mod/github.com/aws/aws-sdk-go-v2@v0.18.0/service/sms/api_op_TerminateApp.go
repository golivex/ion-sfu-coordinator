// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TerminateAppInput struct {
	_ struct{} `type:"structure"`

	// ID of the application to terminate.
	AppId *string `locationName:"appId" type:"string"`
}

// String returns the string representation
func (s TerminateAppInput) String() string {
	return awsutil.Prettify(s)
}

type TerminateAppOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TerminateAppOutput) String() string {
	return awsutil.Prettify(s)
}

const opTerminateApp = "TerminateApp"

// TerminateAppRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Terminates the stack for an application.
//
//    // Example sending a request using TerminateAppRequest.
//    req := client.TerminateAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/TerminateApp
func (c *Client) TerminateAppRequest(input *TerminateAppInput) TerminateAppRequest {
	op := &aws.Operation{
		Name:       opTerminateApp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TerminateAppInput{}
	}

	req := c.newRequest(op, input, &TerminateAppOutput{})
	return TerminateAppRequest{Request: req, Input: input, Copy: c.TerminateAppRequest}
}

// TerminateAppRequest is the request type for the
// TerminateApp API operation.
type TerminateAppRequest struct {
	*aws.Request
	Input *TerminateAppInput
	Copy  func(*TerminateAppInput) TerminateAppRequest
}

// Send marshals and sends the TerminateApp API request.
func (r TerminateAppRequest) Send(ctx context.Context) (*TerminateAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TerminateAppResponse{
		TerminateAppOutput: r.Request.Data.(*TerminateAppOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TerminateAppResponse is the response type for the
// TerminateApp API operation.
type TerminateAppResponse struct {
	*TerminateAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TerminateApp request.
func (r *TerminateAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
