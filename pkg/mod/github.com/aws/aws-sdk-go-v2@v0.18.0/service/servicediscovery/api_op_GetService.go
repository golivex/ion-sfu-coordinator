// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetServiceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the service that you want to get settings for.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetServiceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetServiceOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about the service.
	Service *Service `type:"structure"`
}

// String returns the string representation
func (s GetServiceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetService = "GetService"

// GetServiceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Gets the settings for a specified service.
//
//    // Example sending a request using GetServiceRequest.
//    req := client.GetServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/GetService
func (c *Client) GetServiceRequest(input *GetServiceInput) GetServiceRequest {
	op := &aws.Operation{
		Name:       opGetService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetServiceInput{}
	}

	req := c.newRequest(op, input, &GetServiceOutput{})
	return GetServiceRequest{Request: req, Input: input, Copy: c.GetServiceRequest}
}

// GetServiceRequest is the request type for the
// GetService API operation.
type GetServiceRequest struct {
	*aws.Request
	Input *GetServiceInput
	Copy  func(*GetServiceInput) GetServiceRequest
}

// Send marshals and sends the GetService API request.
func (r GetServiceRequest) Send(ctx context.Context) (*GetServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServiceResponse{
		GetServiceOutput: r.Request.Data.(*GetServiceOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetServiceResponse is the response type for the
// GetService API operation.
type GetServiceResponse struct {
	*GetServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetService request.
func (r *GetServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
