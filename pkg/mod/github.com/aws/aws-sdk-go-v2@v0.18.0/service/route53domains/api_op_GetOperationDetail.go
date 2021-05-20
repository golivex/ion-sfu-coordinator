// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The GetOperationDetail request includes the following element.
type GetOperationDetailInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the operation for which you want to get the status. Amazon
	// Route 53 returned the identifier in the response to the original request.
	//
	// OperationId is a required field
	OperationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetOperationDetailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOperationDetailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOperationDetailInput"}

	if s.OperationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OperationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The GetOperationDetail response includes the following elements.
type GetOperationDetailOutput struct {
	_ struct{} `type:"structure"`

	// The name of a domain.
	DomainName *string `type:"string"`

	// Detailed information on the status including possible errors.
	Message *string `type:"string"`

	// The identifier for the operation.
	OperationId *string `type:"string"`

	// The current status of the requested operation in the system.
	Status OperationStatus `type:"string" enum:"true"`

	// The date when the request was submitted.
	SubmittedDate *time.Time `type:"timestamp"`

	// The type of operation that was requested.
	Type OperationType `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetOperationDetailOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOperationDetail = "GetOperationDetail"

// GetOperationDetailRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation returns the current status of an operation that is not completed.
//
//    // Example sending a request using GetOperationDetailRequest.
//    req := client.GetOperationDetailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/GetOperationDetail
func (c *Client) GetOperationDetailRequest(input *GetOperationDetailInput) GetOperationDetailRequest {
	op := &aws.Operation{
		Name:       opGetOperationDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOperationDetailInput{}
	}

	req := c.newRequest(op, input, &GetOperationDetailOutput{})
	return GetOperationDetailRequest{Request: req, Input: input, Copy: c.GetOperationDetailRequest}
}

// GetOperationDetailRequest is the request type for the
// GetOperationDetail API operation.
type GetOperationDetailRequest struct {
	*aws.Request
	Input *GetOperationDetailInput
	Copy  func(*GetOperationDetailInput) GetOperationDetailRequest
}

// Send marshals and sends the GetOperationDetail API request.
func (r GetOperationDetailRequest) Send(ctx context.Context) (*GetOperationDetailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOperationDetailResponse{
		GetOperationDetailOutput: r.Request.Data.(*GetOperationDetailOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetOperationDetailResponse is the response type for the
// GetOperationDetail API operation.
type GetOperationDetailResponse struct {
	*GetOperationDetailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOperationDetail request.
func (r *GetOperationDetailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
