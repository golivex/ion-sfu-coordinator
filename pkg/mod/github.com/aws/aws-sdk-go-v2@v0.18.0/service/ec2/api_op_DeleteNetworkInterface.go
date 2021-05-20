// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for DeleteNetworkInterface.
type DeleteNetworkInterfaceInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNetworkInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNetworkInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNetworkInterfaceInput"}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteNetworkInterfaceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNetworkInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteNetworkInterface = "DeleteNetworkInterface"

// DeleteNetworkInterfaceRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified network interface. You must detach the network interface
// before you can delete it.
//
//    // Example sending a request using DeleteNetworkInterfaceRequest.
//    req := client.DeleteNetworkInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteNetworkInterface
func (c *Client) DeleteNetworkInterfaceRequest(input *DeleteNetworkInterfaceInput) DeleteNetworkInterfaceRequest {
	op := &aws.Operation{
		Name:       opDeleteNetworkInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNetworkInterfaceInput{}
	}

	req := c.newRequest(op, input, &DeleteNetworkInterfaceOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteNetworkInterfaceRequest{Request: req, Input: input, Copy: c.DeleteNetworkInterfaceRequest}
}

// DeleteNetworkInterfaceRequest is the request type for the
// DeleteNetworkInterface API operation.
type DeleteNetworkInterfaceRequest struct {
	*aws.Request
	Input *DeleteNetworkInterfaceInput
	Copy  func(*DeleteNetworkInterfaceInput) DeleteNetworkInterfaceRequest
}

// Send marshals and sends the DeleteNetworkInterface API request.
func (r DeleteNetworkInterfaceRequest) Send(ctx context.Context) (*DeleteNetworkInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNetworkInterfaceResponse{
		DeleteNetworkInterfaceOutput: r.Request.Data.(*DeleteNetworkInterfaceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNetworkInterfaceResponse is the response type for the
// DeleteNetworkInterface API operation.
type DeleteNetworkInterfaceResponse struct {
	*DeleteNetworkInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNetworkInterface request.
func (r *DeleteNetworkInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
