// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAcceleratorAttributesInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the accelerator with the attributes that
	// you want to describe.
	//
	// AcceleratorArn is a required field
	AcceleratorArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAcceleratorAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAcceleratorAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAcceleratorAttributesInput"}

	if s.AcceleratorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcceleratorArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeAcceleratorAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The attributes of the accelerator.
	AcceleratorAttributes *AcceleratorAttributes `type:"structure"`
}

// String returns the string representation
func (s DescribeAcceleratorAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAcceleratorAttributes = "DescribeAcceleratorAttributes"

// DescribeAcceleratorAttributesRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Describe the attributes of an accelerator.
//
//    // Example sending a request using DescribeAcceleratorAttributesRequest.
//    req := client.DescribeAcceleratorAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/DescribeAcceleratorAttributes
func (c *Client) DescribeAcceleratorAttributesRequest(input *DescribeAcceleratorAttributesInput) DescribeAcceleratorAttributesRequest {
	op := &aws.Operation{
		Name:       opDescribeAcceleratorAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAcceleratorAttributesInput{}
	}

	req := c.newRequest(op, input, &DescribeAcceleratorAttributesOutput{})
	return DescribeAcceleratorAttributesRequest{Request: req, Input: input, Copy: c.DescribeAcceleratorAttributesRequest}
}

// DescribeAcceleratorAttributesRequest is the request type for the
// DescribeAcceleratorAttributes API operation.
type DescribeAcceleratorAttributesRequest struct {
	*aws.Request
	Input *DescribeAcceleratorAttributesInput
	Copy  func(*DescribeAcceleratorAttributesInput) DescribeAcceleratorAttributesRequest
}

// Send marshals and sends the DescribeAcceleratorAttributes API request.
func (r DescribeAcceleratorAttributesRequest) Send(ctx context.Context) (*DescribeAcceleratorAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAcceleratorAttributesResponse{
		DescribeAcceleratorAttributesOutput: r.Request.Data.(*DescribeAcceleratorAttributesOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAcceleratorAttributesResponse is the response type for the
// DescribeAcceleratorAttributes API operation.
type DescribeAcceleratorAttributesResponse struct {
	*DescribeAcceleratorAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAcceleratorAttributes request.
func (r *DescribeAcceleratorAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
