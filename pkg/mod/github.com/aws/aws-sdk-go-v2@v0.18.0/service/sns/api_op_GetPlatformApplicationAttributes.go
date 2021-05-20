// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for GetPlatformApplicationAttributes action.
type GetPlatformApplicationAttributesInput struct {
	_ struct{} `type:"structure"`

	// PlatformApplicationArn for GetPlatformApplicationAttributesInput.
	//
	// PlatformApplicationArn is a required field
	PlatformApplicationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetPlatformApplicationAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPlatformApplicationAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPlatformApplicationAttributesInput"}

	if s.PlatformApplicationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlatformApplicationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response for GetPlatformApplicationAttributes action.
type GetPlatformApplicationAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Attributes include the following:
	//
	//    * EventEndpointCreated – Topic ARN to which EndpointCreated event notifications
	//    should be sent.
	//
	//    * EventEndpointDeleted – Topic ARN to which EndpointDeleted event notifications
	//    should be sent.
	//
	//    * EventEndpointUpdated – Topic ARN to which EndpointUpdate event notifications
	//    should be sent.
	//
	//    * EventDeliveryFailure – Topic ARN to which DeliveryFailure event notifications
	//    should be sent upon Direct Publish delivery failure (permanent) to one
	//    of the application's endpoints.
	Attributes map[string]string `type:"map"`
}

// String returns the string representation
func (s GetPlatformApplicationAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPlatformApplicationAttributes = "GetPlatformApplicationAttributes"

// GetPlatformApplicationAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Retrieves the attributes of the platform application object for the supported
// push notification services, such as APNS and FCM. For more information, see
// Using Amazon SNS Mobile Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
//
//    // Example sending a request using GetPlatformApplicationAttributesRequest.
//    req := client.GetPlatformApplicationAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/GetPlatformApplicationAttributes
func (c *Client) GetPlatformApplicationAttributesRequest(input *GetPlatformApplicationAttributesInput) GetPlatformApplicationAttributesRequest {
	op := &aws.Operation{
		Name:       opGetPlatformApplicationAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPlatformApplicationAttributesInput{}
	}

	req := c.newRequest(op, input, &GetPlatformApplicationAttributesOutput{})
	return GetPlatformApplicationAttributesRequest{Request: req, Input: input, Copy: c.GetPlatformApplicationAttributesRequest}
}

// GetPlatformApplicationAttributesRequest is the request type for the
// GetPlatformApplicationAttributes API operation.
type GetPlatformApplicationAttributesRequest struct {
	*aws.Request
	Input *GetPlatformApplicationAttributesInput
	Copy  func(*GetPlatformApplicationAttributesInput) GetPlatformApplicationAttributesRequest
}

// Send marshals and sends the GetPlatformApplicationAttributes API request.
func (r GetPlatformApplicationAttributesRequest) Send(ctx context.Context) (*GetPlatformApplicationAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPlatformApplicationAttributesResponse{
		GetPlatformApplicationAttributesOutput: r.Request.Data.(*GetPlatformApplicationAttributesOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPlatformApplicationAttributesResponse is the response type for the
// GetPlatformApplicationAttributes API operation.
type GetPlatformApplicationAttributesResponse struct {
	*GetPlatformApplicationAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPlatformApplicationAttributes request.
func (r *GetPlatformApplicationAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
