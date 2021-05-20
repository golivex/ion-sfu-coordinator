// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetCloudFormationStackRecordsInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to a specific page of results for your get cloud
	// formation stack records request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetCloudFormationStackRecordsInput) String() string {
	return awsutil.Prettify(s)
}

type GetCloudFormationStackRecordsOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects describing the CloudFormation stack records.
	CloudFormationStackRecords []CloudFormationStackRecord `locationName:"cloudFormationStackRecords" type:"list"`

	// A token used for advancing to the next page of results of your get relational
	// database bundles request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetCloudFormationStackRecordsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCloudFormationStackRecords = "GetCloudFormationStackRecords"

// GetCloudFormationStackRecordsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the CloudFormation stack record created as a result of the create
// cloud formation stack operation.
//
// An AWS CloudFormation stack is used to create a new Amazon EC2 instance from
// an exported Lightsail snapshot.
//
//    // Example sending a request using GetCloudFormationStackRecordsRequest.
//    req := client.GetCloudFormationStackRecordsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetCloudFormationStackRecords
func (c *Client) GetCloudFormationStackRecordsRequest(input *GetCloudFormationStackRecordsInput) GetCloudFormationStackRecordsRequest {
	op := &aws.Operation{
		Name:       opGetCloudFormationStackRecords,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCloudFormationStackRecordsInput{}
	}

	req := c.newRequest(op, input, &GetCloudFormationStackRecordsOutput{})
	return GetCloudFormationStackRecordsRequest{Request: req, Input: input, Copy: c.GetCloudFormationStackRecordsRequest}
}

// GetCloudFormationStackRecordsRequest is the request type for the
// GetCloudFormationStackRecords API operation.
type GetCloudFormationStackRecordsRequest struct {
	*aws.Request
	Input *GetCloudFormationStackRecordsInput
	Copy  func(*GetCloudFormationStackRecordsInput) GetCloudFormationStackRecordsRequest
}

// Send marshals and sends the GetCloudFormationStackRecords API request.
func (r GetCloudFormationStackRecordsRequest) Send(ctx context.Context) (*GetCloudFormationStackRecordsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCloudFormationStackRecordsResponse{
		GetCloudFormationStackRecordsOutput: r.Request.Data.(*GetCloudFormationStackRecordsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCloudFormationStackRecordsResponse is the response type for the
// GetCloudFormationStackRecords API operation.
type GetCloudFormationStackRecordsResponse struct {
	*GetCloudFormationStackRecordsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCloudFormationStackRecords request.
func (r *GetCloudFormationStackRecordsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
