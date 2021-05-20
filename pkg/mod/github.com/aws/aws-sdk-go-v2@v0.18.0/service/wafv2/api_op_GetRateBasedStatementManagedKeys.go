// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRateBasedStatementManagedKeysInput struct {
	_ struct{} `type:"structure"`

	// The name of the rate-based rule to get the keys for.
	//
	// RuleName is a required field
	RuleName *string `min:"1" type:"string" required:"true"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB)
	// or an API Gateway stage.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//    * CLI - Specify the region when you use the CloudFront scope: --scope=CLOUDFRONT
	//    --region=us-east-1.
	//
	//    * API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// Scope is a required field
	Scope Scope `type:"string" required:"true" enum:"true"`

	// The unique identifier for the Web ACL. This ID is returned in the responses
	// to create and list commands. You provide it to operations like update and
	// delete.
	//
	// WebACLId is a required field
	WebACLId *string `min:"1" type:"string" required:"true"`

	// A friendly name of the Web ACL. You cannot change the name of a Web ACL after
	// you create it.
	//
	// WebACLName is a required field
	WebACLName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRateBasedStatementManagedKeysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRateBasedStatementManagedKeysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRateBasedStatementManagedKeysInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}

	if s.WebACLId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLId"))
	}
	if s.WebACLId != nil && len(*s.WebACLId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLId", 1))
	}

	if s.WebACLName == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLName"))
	}
	if s.WebACLName != nil && len(*s.WebACLName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRateBasedStatementManagedKeysOutput struct {
	_ struct{} `type:"structure"`

	// The keys that are of Internet Protocol version 4 (IPv4).
	ManagedKeysIPV4 *RateBasedStatementManagedKeysIPSet `type:"structure"`

	// The keys that are of Internet Protocol version 6 (IPv6).
	ManagedKeysIPV6 *RateBasedStatementManagedKeysIPSet `type:"structure"`
}

// String returns the string representation
func (s GetRateBasedStatementManagedKeysOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRateBasedStatementManagedKeys = "GetRateBasedStatementManagedKeys"

// GetRateBasedStatementManagedKeysRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Retrieves the keys that are currently blocked by a rate-based rule. The maximum
// number of managed keys that can be blocked for a single rate-based rule is
// 10,000. If more than 10,000 addresses exceed the rate limit, those with the
// highest rates are blocked.
//
//    // Example sending a request using GetRateBasedStatementManagedKeysRequest.
//    req := client.GetRateBasedStatementManagedKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/GetRateBasedStatementManagedKeys
func (c *Client) GetRateBasedStatementManagedKeysRequest(input *GetRateBasedStatementManagedKeysInput) GetRateBasedStatementManagedKeysRequest {
	op := &aws.Operation{
		Name:       opGetRateBasedStatementManagedKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRateBasedStatementManagedKeysInput{}
	}

	req := c.newRequest(op, input, &GetRateBasedStatementManagedKeysOutput{})
	return GetRateBasedStatementManagedKeysRequest{Request: req, Input: input, Copy: c.GetRateBasedStatementManagedKeysRequest}
}

// GetRateBasedStatementManagedKeysRequest is the request type for the
// GetRateBasedStatementManagedKeys API operation.
type GetRateBasedStatementManagedKeysRequest struct {
	*aws.Request
	Input *GetRateBasedStatementManagedKeysInput
	Copy  func(*GetRateBasedStatementManagedKeysInput) GetRateBasedStatementManagedKeysRequest
}

// Send marshals and sends the GetRateBasedStatementManagedKeys API request.
func (r GetRateBasedStatementManagedKeysRequest) Send(ctx context.Context) (*GetRateBasedStatementManagedKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRateBasedStatementManagedKeysResponse{
		GetRateBasedStatementManagedKeysOutput: r.Request.Data.(*GetRateBasedStatementManagedKeysOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRateBasedStatementManagedKeysResponse is the response type for the
// GetRateBasedStatementManagedKeys API operation.
type GetRateBasedStatementManagedKeysResponse struct {
	*GetRateBasedStatementManagedKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRateBasedStatementManagedKeys request.
func (r *GetRateBasedStatementManagedKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
