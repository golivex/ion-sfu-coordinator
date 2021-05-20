// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSubscriptionInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// The AWS Shield Advanced subscription details for an account.
	Subscription *Subscription `type:"structure"`
}

// String returns the string representation
func (s DescribeSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSubscription = "DescribeSubscription"

// DescribeSubscriptionRequest returns a request value for making API operation for
// AWS Shield.
//
// Provides details about the AWS Shield Advanced subscription for an account.
//
//    // Example sending a request using DescribeSubscriptionRequest.
//    req := client.DescribeSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DescribeSubscription
func (c *Client) DescribeSubscriptionRequest(input *DescribeSubscriptionInput) DescribeSubscriptionRequest {
	op := &aws.Operation{
		Name:       opDescribeSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSubscriptionInput{}
	}

	req := c.newRequest(op, input, &DescribeSubscriptionOutput{})
	return DescribeSubscriptionRequest{Request: req, Input: input, Copy: c.DescribeSubscriptionRequest}
}

// DescribeSubscriptionRequest is the request type for the
// DescribeSubscription API operation.
type DescribeSubscriptionRequest struct {
	*aws.Request
	Input *DescribeSubscriptionInput
	Copy  func(*DescribeSubscriptionInput) DescribeSubscriptionRequest
}

// Send marshals and sends the DescribeSubscription API request.
func (r DescribeSubscriptionRequest) Send(ctx context.Context) (*DescribeSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSubscriptionResponse{
		DescribeSubscriptionOutput: r.Request.Data.(*DescribeSubscriptionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSubscriptionResponse is the response type for the
// DescribeSubscription API operation.
type DescribeSubscriptionResponse struct {
	*DescribeSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSubscription request.
func (r *DescribeSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
