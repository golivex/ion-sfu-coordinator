// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSnowballUsageInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetSnowballUsageInput) String() string {
	return awsutil.Prettify(s)
}

type GetSnowballUsageOutput struct {
	_ struct{} `type:"structure"`

	// The service limit for number of Snowballs this account can have at once.
	// The default service limit is 1 (one).
	SnowballLimit *int64 `type:"integer"`

	// The number of Snowballs that this account is currently using.
	SnowballsInUse *int64 `type:"integer"`
}

// String returns the string representation
func (s GetSnowballUsageOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSnowballUsage = "GetSnowballUsage"

// GetSnowballUsageRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// Returns information about the Snowball service limit for your account, and
// also the number of Snowballs your account has in use.
//
// The default service limit for the number of Snowballs that you can have at
// one time is 1. If you want to increase your service limit, contact AWS Support.
//
//    // Example sending a request using GetSnowballUsageRequest.
//    req := client.GetSnowballUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/GetSnowballUsage
func (c *Client) GetSnowballUsageRequest(input *GetSnowballUsageInput) GetSnowballUsageRequest {
	op := &aws.Operation{
		Name:       opGetSnowballUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSnowballUsageInput{}
	}

	req := c.newRequest(op, input, &GetSnowballUsageOutput{})
	return GetSnowballUsageRequest{Request: req, Input: input, Copy: c.GetSnowballUsageRequest}
}

// GetSnowballUsageRequest is the request type for the
// GetSnowballUsage API operation.
type GetSnowballUsageRequest struct {
	*aws.Request
	Input *GetSnowballUsageInput
	Copy  func(*GetSnowballUsageInput) GetSnowballUsageRequest
}

// Send marshals and sends the GetSnowballUsage API request.
func (r GetSnowballUsageRequest) Send(ctx context.Context) (*GetSnowballUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSnowballUsageResponse{
		GetSnowballUsageOutput: r.Request.Data.(*GetSnowballUsageOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSnowballUsageResponse is the response type for the
// GetSnowballUsage API operation.
type GetSnowballUsageResponse struct {
	*GetSnowballUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSnowballUsage request.
func (r *GetSnowballUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
