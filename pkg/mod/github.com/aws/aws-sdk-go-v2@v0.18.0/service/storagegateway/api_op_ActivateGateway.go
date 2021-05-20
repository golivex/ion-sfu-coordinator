// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing one or more of the following fields:
//
//    * ActivateGatewayInput$ActivationKey
//
//    * ActivateGatewayInput$GatewayName
//
//    * ActivateGatewayInput$GatewayRegion
//
//    * ActivateGatewayInput$GatewayTimezone
//
//    * ActivateGatewayInput$GatewayType
//
//    * ActivateGatewayInput$TapeDriveType
//
//    * ActivateGatewayInput$MediumChangerType
type ActivateGatewayInput struct {
	_ struct{} `type:"structure"`

	// Your gateway activation key. You can obtain the activation key by sending
	// an HTTP GET request with redirects enabled to the gateway IP address (port
	// 80). The redirect URL returned in the response provides you the activation
	// key for your gateway in the query string parameter activationKey. It may
	// also include other activation-related parameters, however, these are merely
	// defaults -- the arguments you pass to the ActivateGateway API call determine
	// the actual configuration of your gateway.
	//
	// For more information, see https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html
	// in the Storage Gateway User Guide.
	//
	// ActivationKey is a required field
	ActivationKey *string `min:"1" type:"string" required:"true"`

	// The name you configured for your gateway.
	//
	// GatewayName is a required field
	GatewayName *string `min:"2" type:"string" required:"true"`

	// A value that indicates the AWS Region where you want to store your data.
	// The gateway AWS Region specified must be the same AWS Region as the AWS Region
	// in your Host header in the request. For more information about available
	// AWS Regions and endpoints for AWS Storage Gateway, see Regions and Endpoints
	// (https://docs.aws.amazon.com/general/latest/gr/rande.html#sg_region) in the
	// Amazon Web Services Glossary.
	//
	// Valid Values: See AWS Storage Gateway Regions and Endpoints (https://docs.aws.amazon.com/general/latest/gr/rande.html#sg_region)
	// in the AWS General Reference.
	//
	// GatewayRegion is a required field
	GatewayRegion *string `min:"1" type:"string" required:"true"`

	// A value that indicates the time zone you want to set for the gateway. The
	// time zone is of the format "GMT-hr:mm" or "GMT+hr:mm". For example, GMT-4:00
	// indicates the time is 4 hours behind GMT. GMT+2:00 indicates the time is
	// 2 hours ahead of GMT. The time zone is used, for example, for scheduling
	// snapshots and your gateway's maintenance schedule.
	//
	// GatewayTimezone is a required field
	GatewayTimezone *string `min:"3" type:"string" required:"true"`

	// A value that defines the type of gateway to activate. The type specified
	// is critical to all later functions of the gateway and cannot be changed after
	// activation. The default value is CACHED.
	//
	// Valid Values: "STORED", "CACHED", "VTL", "FILE_S3"
	GatewayType *string `min:"2" type:"string"`

	// The value that indicates the type of medium changer to use for tape gateway.
	// This field is optional.
	//
	// Valid Values: "STK-L700", "AWS-Gateway-VTL"
	MediumChangerType *string `min:"2" type:"string"`

	// A list of up to 50 tags that you can assign to the gateway. Each tag is a
	// key-value pair.
	//
	// Valid characters for key and value are letters, spaces, and numbers that
	// can be represented in UTF-8 format, and the following special characters:
	// + - = . _ : / @. The maximum length of a tag's key is 128 characters, and
	// the maximum length for a tag's value is 256 characters.
	Tags []Tag `type:"list"`

	// The value that indicates the type of tape drive to use for tape gateway.
	// This field is optional.
	//
	// Valid Values: "IBM-ULT3580-TD5"
	TapeDriveType *string `min:"2" type:"string"`
}

// String returns the string representation
func (s ActivateGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ActivateGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ActivateGatewayInput"}

	if s.ActivationKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActivationKey"))
	}
	if s.ActivationKey != nil && len(*s.ActivationKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ActivationKey", 1))
	}

	if s.GatewayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayName"))
	}
	if s.GatewayName != nil && len(*s.GatewayName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayName", 2))
	}

	if s.GatewayRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayRegion"))
	}
	if s.GatewayRegion != nil && len(*s.GatewayRegion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayRegion", 1))
	}

	if s.GatewayTimezone == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayTimezone"))
	}
	if s.GatewayTimezone != nil && len(*s.GatewayTimezone) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayTimezone", 3))
	}
	if s.GatewayType != nil && len(*s.GatewayType) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayType", 2))
	}
	if s.MediumChangerType != nil && len(*s.MediumChangerType) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("MediumChangerType", 2))
	}
	if s.TapeDriveType != nil && len(*s.TapeDriveType) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("TapeDriveType", 2))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// AWS Storage Gateway returns the Amazon Resource Name (ARN) of the activated
// gateway. It is a string made of information such as your account, gateway
// name, and AWS Region. This ARN is used to reference the gateway in other
// API operations as well as resource-based authorization.
//
// For gateways activated prior to September 02, 2015, the gateway ARN contains
// the gateway name rather than the gateway ID. Changing the name of the gateway
// has no effect on the gateway ARN.
type ActivateGatewayOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s ActivateGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opActivateGateway = "ActivateGateway"

// ActivateGatewayRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Activates the gateway you previously deployed on your host. In the activation
// process, you specify information such as the AWS Region that you want to
// use for storing snapshots or tapes, the time zone for scheduled snapshots
// the gateway snapshot schedule window, an activation key, and a name for your
// gateway. The activation process also associates your gateway with your account;
// for more information, see UpdateGatewayInformation.
//
// You must turn on the gateway VM before you can activate your gateway.
//
//    // Example sending a request using ActivateGatewayRequest.
//    req := client.ActivateGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/ActivateGateway
func (c *Client) ActivateGatewayRequest(input *ActivateGatewayInput) ActivateGatewayRequest {
	op := &aws.Operation{
		Name:       opActivateGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ActivateGatewayInput{}
	}

	req := c.newRequest(op, input, &ActivateGatewayOutput{})
	return ActivateGatewayRequest{Request: req, Input: input, Copy: c.ActivateGatewayRequest}
}

// ActivateGatewayRequest is the request type for the
// ActivateGateway API operation.
type ActivateGatewayRequest struct {
	*aws.Request
	Input *ActivateGatewayInput
	Copy  func(*ActivateGatewayInput) ActivateGatewayRequest
}

// Send marshals and sends the ActivateGateway API request.
func (r ActivateGatewayRequest) Send(ctx context.Context) (*ActivateGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ActivateGatewayResponse{
		ActivateGatewayOutput: r.Request.Data.(*ActivateGatewayOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ActivateGatewayResponse is the response type for the
// ActivateGateway API operation.
type ActivateGatewayResponse struct {
	*ActivateGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ActivateGateway request.
func (r *ActivateGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
