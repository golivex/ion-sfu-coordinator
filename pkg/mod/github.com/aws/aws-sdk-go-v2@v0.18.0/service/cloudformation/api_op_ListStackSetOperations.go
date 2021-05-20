// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListStackSetOperationsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListStackSetOperations again and assign that
	// token to the request object's NextToken parameter. If there are no remaining
	// results, the previous response object's NextToken parameter is set to null.
	NextToken *string `min:"1" type:"string"`

	// The name or unique ID of the stack set that you want to get operation summaries
	// for.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListStackSetOperationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStackSetOperationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStackSetOperationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListStackSetOperationsOutput struct {
	_ struct{} `type:"structure"`

	// If the request doesn't return all results, NextToken is set to a token. To
	// retrieve the next set of results, call ListOperationResults again and assign
	// that token to the request object's NextToken parameter. If there are no remaining
	// results, NextToken is set to null.
	NextToken *string `min:"1" type:"string"`

	// A list of StackSetOperationSummary structures that contain summary information
	// about operations for the specified stack set.
	Summaries []StackSetOperationSummary `type:"list"`
}

// String returns the string representation
func (s ListStackSetOperationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListStackSetOperations = "ListStackSetOperations"

// ListStackSetOperationsRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns summary information about operations performed on a stack set.
//
//    // Example sending a request using ListStackSetOperationsRequest.
//    req := client.ListStackSetOperationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackSetOperations
func (c *Client) ListStackSetOperationsRequest(input *ListStackSetOperationsInput) ListStackSetOperationsRequest {
	op := &aws.Operation{
		Name:       opListStackSetOperations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListStackSetOperationsInput{}
	}

	req := c.newRequest(op, input, &ListStackSetOperationsOutput{})
	return ListStackSetOperationsRequest{Request: req, Input: input, Copy: c.ListStackSetOperationsRequest}
}

// ListStackSetOperationsRequest is the request type for the
// ListStackSetOperations API operation.
type ListStackSetOperationsRequest struct {
	*aws.Request
	Input *ListStackSetOperationsInput
	Copy  func(*ListStackSetOperationsInput) ListStackSetOperationsRequest
}

// Send marshals and sends the ListStackSetOperations API request.
func (r ListStackSetOperationsRequest) Send(ctx context.Context) (*ListStackSetOperationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStackSetOperationsResponse{
		ListStackSetOperationsOutput: r.Request.Data.(*ListStackSetOperationsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListStackSetOperationsResponse is the response type for the
// ListStackSetOperations API operation.
type ListStackSetOperationsResponse struct {
	*ListStackSetOperationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStackSetOperations request.
func (r *ListStackSetOperationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
