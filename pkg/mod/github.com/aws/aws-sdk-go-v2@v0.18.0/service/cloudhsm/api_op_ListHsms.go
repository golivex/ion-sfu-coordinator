// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListHsmsInput struct {
	_ struct{} `type:"structure"`

	// The NextToken value from a previous call to ListHsms. Pass null if this is
	// the first call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHsmsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of the ListHsms operation.
type ListHsmsOutput struct {
	_ struct{} `type:"structure"`

	// The list of ARNs that identify the HSMs.
	HsmList []string `type:"list"`

	// If not null, more results are available. Pass this value to ListHsms to retrieve
	// the next set of items.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHsmsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListHsms = "ListHsms"

// ListHsmsRequest returns a request value for making API operation for
// Amazon CloudHSM.
//
// This is documentation for AWS CloudHSM Classic. For more information, see
// AWS CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/),
// the AWS CloudHSM Classic User Guide (http://docs.aws.amazon.com/cloudhsm/classic/userguide/),
// and the AWS CloudHSM Classic API Reference (http://docs.aws.amazon.com/cloudhsm/classic/APIReference/).
//
// For information about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide (http://docs.aws.amazon.com/cloudhsm/latest/userguide/),
// and the AWS CloudHSM API Reference (http://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
//
// Retrieves the identifiers of all of the HSMs provisioned for the current
// customer.
//
// This operation supports pagination with the use of the NextToken member.
// If more results are available, the NextToken member of the response contains
// a token that you pass in the next call to ListHsms to retrieve the next set
// of items.
//
//    // Example sending a request using ListHsmsRequest.
//    req := client.ListHsmsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/ListHsms
func (c *Client) ListHsmsRequest(input *ListHsmsInput) ListHsmsRequest {
	op := &aws.Operation{
		Name:       opListHsms,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListHsmsInput{}
	}

	req := c.newRequest(op, input, &ListHsmsOutput{})
	return ListHsmsRequest{Request: req, Input: input, Copy: c.ListHsmsRequest}
}

// ListHsmsRequest is the request type for the
// ListHsms API operation.
type ListHsmsRequest struct {
	*aws.Request
	Input *ListHsmsInput
	Copy  func(*ListHsmsInput) ListHsmsRequest
}

// Send marshals and sends the ListHsms API request.
func (r ListHsmsRequest) Send(ctx context.Context) (*ListHsmsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHsmsResponse{
		ListHsmsOutput: r.Request.Data.(*ListHsmsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListHsmsResponse is the response type for the
// ListHsms API operation.
type ListHsmsResponse struct {
	*ListHsmsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHsms request.
func (r *ListHsmsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
