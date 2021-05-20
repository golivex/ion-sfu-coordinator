// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAccountAttributesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeAccountAttributesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeAccountAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The Elastic Beanstalk resource quotas associated with the calling AWS account.
	ResourceQuotas *ResourceQuotas `type:"structure"`
}

// String returns the string representation
func (s DescribeAccountAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAccountAttributes = "DescribeAccountAttributes"

// DescribeAccountAttributesRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Returns attributes related to AWS Elastic Beanstalk that are associated with
// the calling AWS account.
//
// The result currently has one set of attributes—resource quotas.
//
//    // Example sending a request using DescribeAccountAttributesRequest.
//    req := client.DescribeAccountAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeAccountAttributes
func (c *Client) DescribeAccountAttributesRequest(input *DescribeAccountAttributesInput) DescribeAccountAttributesRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAccountAttributesInput{}
	}

	req := c.newRequest(op, input, &DescribeAccountAttributesOutput{})
	return DescribeAccountAttributesRequest{Request: req, Input: input, Copy: c.DescribeAccountAttributesRequest}
}

// DescribeAccountAttributesRequest is the request type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesRequest struct {
	*aws.Request
	Input *DescribeAccountAttributesInput
	Copy  func(*DescribeAccountAttributesInput) DescribeAccountAttributesRequest
}

// Send marshals and sends the DescribeAccountAttributes API request.
func (r DescribeAccountAttributesRequest) Send(ctx context.Context) (*DescribeAccountAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountAttributesResponse{
		DescribeAccountAttributesOutput: r.Request.Data.(*DescribeAccountAttributesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountAttributesResponse is the response type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesResponse struct {
	*DescribeAccountAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountAttributes request.
func (r *DescribeAccountAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
