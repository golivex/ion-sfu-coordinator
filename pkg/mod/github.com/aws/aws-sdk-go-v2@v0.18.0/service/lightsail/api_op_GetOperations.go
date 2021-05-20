// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetOperationsInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to the next page of results from your get operations
	// request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetOperationsInput) String() string {
	return awsutil.Prettify(s)
}

type GetOperationsOutput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to the next page of results from your get operations
	// request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`

	// An array of key-value pairs containing information about the results of your
	// get operations request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s GetOperationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOperations = "GetOperations"

// GetOperationsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about all operations.
//
// Results are returned from oldest to newest, up to a maximum of 200. Results
// can be paged by making each subsequent call to GetOperations use the maximum
// (last) statusChangedAt value from the previous request.
//
//    // Example sending a request using GetOperationsRequest.
//    req := client.GetOperationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetOperations
func (c *Client) GetOperationsRequest(input *GetOperationsInput) GetOperationsRequest {
	op := &aws.Operation{
		Name:       opGetOperations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOperationsInput{}
	}

	req := c.newRequest(op, input, &GetOperationsOutput{})
	return GetOperationsRequest{Request: req, Input: input, Copy: c.GetOperationsRequest}
}

// GetOperationsRequest is the request type for the
// GetOperations API operation.
type GetOperationsRequest struct {
	*aws.Request
	Input *GetOperationsInput
	Copy  func(*GetOperationsInput) GetOperationsRequest
}

// Send marshals and sends the GetOperations API request.
func (r GetOperationsRequest) Send(ctx context.Context) (*GetOperationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOperationsResponse{
		GetOperationsOutput: r.Request.Data.(*GetOperationsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetOperationsResponse is the response type for the
// GetOperations API operation.
type GetOperationsResponse struct {
	*GetOperationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOperations request.
func (r *GetOperationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
