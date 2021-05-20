// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeFeatureTransformationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the feature transformation to describe.
	//
	// FeatureTransformationArn is a required field
	FeatureTransformationArn *string `locationName:"featureTransformationArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeFeatureTransformationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFeatureTransformationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFeatureTransformationInput"}

	if s.FeatureTransformationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FeatureTransformationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeFeatureTransformationOutput struct {
	_ struct{} `type:"structure"`

	// A listing of the FeatureTransformation properties.
	FeatureTransformation *FeatureTransformation `locationName:"featureTransformation" type:"structure"`
}

// String returns the string representation
func (s DescribeFeatureTransformationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFeatureTransformation = "DescribeFeatureTransformation"

// DescribeFeatureTransformationRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Describes the given feature transformation.
//
//    // Example sending a request using DescribeFeatureTransformationRequest.
//    req := client.DescribeFeatureTransformationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeFeatureTransformation
func (c *Client) DescribeFeatureTransformationRequest(input *DescribeFeatureTransformationInput) DescribeFeatureTransformationRequest {
	op := &aws.Operation{
		Name:       opDescribeFeatureTransformation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFeatureTransformationInput{}
	}

	req := c.newRequest(op, input, &DescribeFeatureTransformationOutput{})
	return DescribeFeatureTransformationRequest{Request: req, Input: input, Copy: c.DescribeFeatureTransformationRequest}
}

// DescribeFeatureTransformationRequest is the request type for the
// DescribeFeatureTransformation API operation.
type DescribeFeatureTransformationRequest struct {
	*aws.Request
	Input *DescribeFeatureTransformationInput
	Copy  func(*DescribeFeatureTransformationInput) DescribeFeatureTransformationRequest
}

// Send marshals and sends the DescribeFeatureTransformation API request.
func (r DescribeFeatureTransformationRequest) Send(ctx context.Context) (*DescribeFeatureTransformationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFeatureTransformationResponse{
		DescribeFeatureTransformationOutput: r.Request.Data.(*DescribeFeatureTransformationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFeatureTransformationResponse is the response type for the
// DescribeFeatureTransformation API operation.
type DescribeFeatureTransformationResponse struct {
	*DescribeFeatureTransformationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFeatureTransformation request.
func (r *DescribeFeatureTransformationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
