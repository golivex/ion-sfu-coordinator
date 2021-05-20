// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

type CreateByteMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description of the ByteMatchSet. You can't change Name
	// after you create a ByteMatchSet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateByteMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateByteMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateByteMatchSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateByteMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// A ByteMatchSet that contains no ByteMatchTuple objects.
	ByteMatchSet *waf.ByteMatchSet `type:"structure"`

	// The ChangeToken that you used to submit the CreateByteMatchSet request. You
	// can also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateByteMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateByteMatchSet = "CreateByteMatchSet"

// CreateByteMatchSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Creates a ByteMatchSet. You then use UpdateByteMatchSet to identify the part
// of a web request that you want AWS WAF to inspect, such as the values of
// the User-Agent header or the query string. For example, you can create a
// ByteMatchSet that matches any requests with User-Agent headers that contain
// the string BadBot. You can then configure AWS WAF to reject those requests.
//
// To create and configure a ByteMatchSet, perform the following steps:
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateByteMatchSet request.
//
// Submit a CreateByteMatchSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateByteMatchSet request.
//
// Submit an UpdateByteMatchSet request to specify the part of the request that
// you want AWS WAF to inspect (for example, the header or the URI) and the
// value that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateByteMatchSetRequest.
//    req := client.CreateByteMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateByteMatchSet
func (c *Client) CreateByteMatchSetRequest(input *CreateByteMatchSetInput) CreateByteMatchSetRequest {
	op := &aws.Operation{
		Name:       opCreateByteMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateByteMatchSetInput{}
	}

	req := c.newRequest(op, input, &CreateByteMatchSetOutput{})
	return CreateByteMatchSetRequest{Request: req, Input: input, Copy: c.CreateByteMatchSetRequest}
}

// CreateByteMatchSetRequest is the request type for the
// CreateByteMatchSet API operation.
type CreateByteMatchSetRequest struct {
	*aws.Request
	Input *CreateByteMatchSetInput
	Copy  func(*CreateByteMatchSetInput) CreateByteMatchSetRequest
}

// Send marshals and sends the CreateByteMatchSet API request.
func (r CreateByteMatchSetRequest) Send(ctx context.Context) (*CreateByteMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateByteMatchSetResponse{
		CreateByteMatchSetOutput: r.Request.Data.(*CreateByteMatchSetOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateByteMatchSetResponse is the response type for the
// CreateByteMatchSet API operation.
type CreateByteMatchSetResponse struct {
	*CreateByteMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateByteMatchSet request.
func (r *CreateByteMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
