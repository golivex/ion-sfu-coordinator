// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateWebACLInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource to disassociate from the web
	// ACL.
	//
	// The ARN must be in one of the following formats:
	//
	//    * For a CloudFront distribution: arn:aws:cloudfront::account-id:distribution/distribution-id
	//
	//    * For an Application Load Balancer: arn:aws:elasticloadbalancing: region:account-id:loadbalancer/app/load-balancer-name
	//    /load-balancer-id
	//
	//    * For an Amazon API Gateway stage: arn:aws:apigateway:region ::/restapis/api-id/stages/stage-name
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateWebACLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateWebACLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateWebACLInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateWebACLOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateWebACLOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateWebACL = "DisassociateWebACL"

// DisassociateWebACLRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Disassociates a Web ACL from a regional application resource. A regional
// application can be an Application Load Balancer (ALB) or an API Gateway stage.
//
// For AWS CloudFront, you can disassociate the Web ACL by providing an empty
// WebACLId in the CloudFront API call UpdateDistribution. For information,
// see UpdateDistribution (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html).
//
//    // Example sending a request using DisassociateWebACLRequest.
//    req := client.DisassociateWebACLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/DisassociateWebACL
func (c *Client) DisassociateWebACLRequest(input *DisassociateWebACLInput) DisassociateWebACLRequest {
	op := &aws.Operation{
		Name:       opDisassociateWebACL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateWebACLInput{}
	}

	req := c.newRequest(op, input, &DisassociateWebACLOutput{})
	return DisassociateWebACLRequest{Request: req, Input: input, Copy: c.DisassociateWebACLRequest}
}

// DisassociateWebACLRequest is the request type for the
// DisassociateWebACL API operation.
type DisassociateWebACLRequest struct {
	*aws.Request
	Input *DisassociateWebACLInput
	Copy  func(*DisassociateWebACLInput) DisassociateWebACLRequest
}

// Send marshals and sends the DisassociateWebACL API request.
func (r DisassociateWebACLRequest) Send(ctx context.Context) (*DisassociateWebACLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateWebACLResponse{
		DisassociateWebACLOutput: r.Request.Data.(*DisassociateWebACLOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateWebACLResponse is the response type for the
// DisassociateWebACL API operation.
type DisassociateWebACLResponse struct {
	*DisassociateWebACLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateWebACL request.
func (r *DisassociateWebACLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
