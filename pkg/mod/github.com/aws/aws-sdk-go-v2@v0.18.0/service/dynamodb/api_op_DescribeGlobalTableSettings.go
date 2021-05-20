// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeGlobalTableSettingsInput struct {
	_ struct{} `type:"structure"`

	// The name of the global table to describe.
	//
	// GlobalTableName is a required field
	GlobalTableName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeGlobalTableSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeGlobalTableSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeGlobalTableSettingsInput"}

	if s.GlobalTableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalTableName"))
	}
	if s.GlobalTableName != nil && len(*s.GlobalTableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("GlobalTableName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeGlobalTableSettingsOutput struct {
	_ struct{} `type:"structure"`

	// The name of the global table.
	GlobalTableName *string `min:"3" type:"string"`

	// The Region-specific settings for the global table.
	ReplicaSettings []ReplicaSettingsDescription `type:"list"`
}

// String returns the string representation
func (s DescribeGlobalTableSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeGlobalTableSettings = "DescribeGlobalTableSettings"

// DescribeGlobalTableSettingsRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Describes Region-specific settings for a global table.
//
// This method only applies to Version 2017.11.29 (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
// of global tables.
//
//    // Example sending a request using DescribeGlobalTableSettingsRequest.
//    req := client.DescribeGlobalTableSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DescribeGlobalTableSettings
func (c *Client) DescribeGlobalTableSettingsRequest(input *DescribeGlobalTableSettingsInput) DescribeGlobalTableSettingsRequest {
	op := &aws.Operation{
		Name:       opDescribeGlobalTableSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeGlobalTableSettingsInput{}
	}

	req := c.newRequest(op, input, &DescribeGlobalTableSettingsOutput{})
	return DescribeGlobalTableSettingsRequest{Request: req, Input: input, Copy: c.DescribeGlobalTableSettingsRequest}
}

// DescribeGlobalTableSettingsRequest is the request type for the
// DescribeGlobalTableSettings API operation.
type DescribeGlobalTableSettingsRequest struct {
	*aws.Request
	Input *DescribeGlobalTableSettingsInput
	Copy  func(*DescribeGlobalTableSettingsInput) DescribeGlobalTableSettingsRequest
}

// Send marshals and sends the DescribeGlobalTableSettings API request.
func (r DescribeGlobalTableSettingsRequest) Send(ctx context.Context) (*DescribeGlobalTableSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGlobalTableSettingsResponse{
		DescribeGlobalTableSettingsOutput: r.Request.Data.(*DescribeGlobalTableSettingsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeGlobalTableSettingsResponse is the response type for the
// DescribeGlobalTableSettings API operation.
type DescribeGlobalTableSettingsResponse struct {
	*DescribeGlobalTableSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGlobalTableSettings request.
func (r *DescribeGlobalTableSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
