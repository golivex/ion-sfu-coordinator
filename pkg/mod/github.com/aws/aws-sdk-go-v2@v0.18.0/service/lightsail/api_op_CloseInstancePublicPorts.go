// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CloseInstancePublicPortsInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance on which you're attempting to close the public ports.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`

	// Information about the public port you are trying to close.
	//
	// PortInfo is a required field
	PortInfo *PortInfo `locationName:"portInfo" type:"structure" required:"true"`
}

// String returns the string representation
func (s CloseInstancePublicPortsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CloseInstancePublicPortsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CloseInstancePublicPortsInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if s.PortInfo == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortInfo"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CloseInstancePublicPortsOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs that contains information about the operation.
	Operation *Operation `locationName:"operation" type:"structure"`
}

// String returns the string representation
func (s CloseInstancePublicPortsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCloseInstancePublicPorts = "CloseInstancePublicPorts"

// CloseInstancePublicPortsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Closes the public ports on a specific Amazon Lightsail instance.
//
// The close instance public ports operation supports tag-based access control
// via resource tags applied to the resource identified by instance name. For
// more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using CloseInstancePublicPortsRequest.
//    req := client.CloseInstancePublicPortsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CloseInstancePublicPorts
func (c *Client) CloseInstancePublicPortsRequest(input *CloseInstancePublicPortsInput) CloseInstancePublicPortsRequest {
	op := &aws.Operation{
		Name:       opCloseInstancePublicPorts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CloseInstancePublicPortsInput{}
	}

	req := c.newRequest(op, input, &CloseInstancePublicPortsOutput{})
	return CloseInstancePublicPortsRequest{Request: req, Input: input, Copy: c.CloseInstancePublicPortsRequest}
}

// CloseInstancePublicPortsRequest is the request type for the
// CloseInstancePublicPorts API operation.
type CloseInstancePublicPortsRequest struct {
	*aws.Request
	Input *CloseInstancePublicPortsInput
	Copy  func(*CloseInstancePublicPortsInput) CloseInstancePublicPortsRequest
}

// Send marshals and sends the CloseInstancePublicPorts API request.
func (r CloseInstancePublicPortsRequest) Send(ctx context.Context) (*CloseInstancePublicPortsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CloseInstancePublicPortsResponse{
		CloseInstancePublicPortsOutput: r.Request.Data.(*CloseInstancePublicPortsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CloseInstancePublicPortsResponse is the response type for the
// CloseInstancePublicPorts API operation.
type CloseInstancePublicPortsResponse struct {
	*CloseInstancePublicPortsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CloseInstancePublicPorts request.
func (r *CloseInstancePublicPortsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
