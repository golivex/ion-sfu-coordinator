// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartSmartHomeApplianceDiscoveryInput struct {
	_ struct{} `type:"structure"`

	// The room where smart home appliance discovery was initiated.
	//
	// RoomArn is a required field
	RoomArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartSmartHomeApplianceDiscoveryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSmartHomeApplianceDiscoveryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSmartHomeApplianceDiscoveryInput"}

	if s.RoomArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartSmartHomeApplianceDiscoveryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartSmartHomeApplianceDiscoveryOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartSmartHomeApplianceDiscovery = "StartSmartHomeApplianceDiscovery"

// StartSmartHomeApplianceDiscoveryRequest returns a request value for making API operation for
// Alexa For Business.
//
// Initiates the discovery of any smart home appliances associated with the
// room.
//
//    // Example sending a request using StartSmartHomeApplianceDiscoveryRequest.
//    req := client.StartSmartHomeApplianceDiscoveryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/StartSmartHomeApplianceDiscovery
func (c *Client) StartSmartHomeApplianceDiscoveryRequest(input *StartSmartHomeApplianceDiscoveryInput) StartSmartHomeApplianceDiscoveryRequest {
	op := &aws.Operation{
		Name:       opStartSmartHomeApplianceDiscovery,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartSmartHomeApplianceDiscoveryInput{}
	}

	req := c.newRequest(op, input, &StartSmartHomeApplianceDiscoveryOutput{})
	return StartSmartHomeApplianceDiscoveryRequest{Request: req, Input: input, Copy: c.StartSmartHomeApplianceDiscoveryRequest}
}

// StartSmartHomeApplianceDiscoveryRequest is the request type for the
// StartSmartHomeApplianceDiscovery API operation.
type StartSmartHomeApplianceDiscoveryRequest struct {
	*aws.Request
	Input *StartSmartHomeApplianceDiscoveryInput
	Copy  func(*StartSmartHomeApplianceDiscoveryInput) StartSmartHomeApplianceDiscoveryRequest
}

// Send marshals and sends the StartSmartHomeApplianceDiscovery API request.
func (r StartSmartHomeApplianceDiscoveryRequest) Send(ctx context.Context) (*StartSmartHomeApplianceDiscoveryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartSmartHomeApplianceDiscoveryResponse{
		StartSmartHomeApplianceDiscoveryOutput: r.Request.Data.(*StartSmartHomeApplianceDiscoveryOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartSmartHomeApplianceDiscoveryResponse is the response type for the
// StartSmartHomeApplianceDiscovery API operation.
type StartSmartHomeApplianceDiscoveryResponse struct {
	*StartSmartHomeApplianceDiscoveryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartSmartHomeApplianceDiscovery request.
func (r *StartSmartHomeApplianceDiscoveryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
