// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeRuleInput struct {
	_ struct{} `type:"structure"`

	// The event bus associated with the rule. If you omit this, the default event
	// bus is used.
	EventBusName *string `min:"1" type:"string"`

	// The name of the rule.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRuleInput"}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
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

type DescribeRuleOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the rule.
	Arn *string `min:"1" type:"string"`

	// The description of the rule.
	Description *string `type:"string"`

	// The event bus associated with the rule.
	EventBusName *string `min:"1" type:"string"`

	// The event pattern. For more information, see Event Patterns (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)
	// in the Amazon EventBridge User Guide.
	EventPattern *string `type:"string"`

	// If this is a managed rule, created by an AWS service on your behalf, this
	// field displays the principal name of the AWS service that created the rule.
	ManagedBy *string `min:"1" type:"string"`

	// The name of the rule.
	Name *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the IAM role associated with the rule.
	RoleArn *string `min:"1" type:"string"`

	// The scheduling expression: for example, "cron(0 20 * * ? *)" or "rate(5 minutes)".
	ScheduleExpression *string `type:"string"`

	// Specifies whether the rule is enabled or disabled.
	State RuleState `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRule = "DescribeRule"

// DescribeRuleRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Describes the specified rule.
//
// DescribeRule doesn't list the targets of a rule. To see the targets associated
// with a rule, use ListTargetsByRule.
//
//    // Example sending a request using DescribeRuleRequest.
//    req := client.DescribeRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/DescribeRule
func (c *Client) DescribeRuleRequest(input *DescribeRuleInput) DescribeRuleRequest {
	op := &aws.Operation{
		Name:       opDescribeRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRuleInput{}
	}

	req := c.newRequest(op, input, &DescribeRuleOutput{})
	return DescribeRuleRequest{Request: req, Input: input, Copy: c.DescribeRuleRequest}
}

// DescribeRuleRequest is the request type for the
// DescribeRule API operation.
type DescribeRuleRequest struct {
	*aws.Request
	Input *DescribeRuleInput
	Copy  func(*DescribeRuleInput) DescribeRuleRequest
}

// Send marshals and sends the DescribeRule API request.
func (r DescribeRuleRequest) Send(ctx context.Context) (*DescribeRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRuleResponse{
		DescribeRuleOutput: r.Request.Data.(*DescribeRuleOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRuleResponse is the response type for the
// DescribeRule API operation.
type DescribeRuleResponse struct {
	*DescribeRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRule request.
func (r *DescribeRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
