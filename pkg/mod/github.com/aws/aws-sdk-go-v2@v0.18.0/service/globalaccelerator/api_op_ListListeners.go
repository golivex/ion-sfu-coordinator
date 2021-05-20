// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListListenersInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the accelerator for which you want to list
	// listener objects.
	//
	// AcceleratorArn is a required field
	AcceleratorArn *string `type:"string" required:"true"`

	// The number of listener objects that you want to return with this call. The
	// default value is 10.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListListenersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListListenersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListListenersInput"}

	if s.AcceleratorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcceleratorArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListListenersOutput struct {
	_ struct{} `type:"structure"`

	// The list of listeners for an accelerator.
	Listeners []Listener `type:"list"`

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListListenersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListListeners = "ListListeners"

// ListListenersRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// List the listeners for an accelerator.
//
//    // Example sending a request using ListListenersRequest.
//    req := client.ListListenersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/ListListeners
func (c *Client) ListListenersRequest(input *ListListenersInput) ListListenersRequest {
	op := &aws.Operation{
		Name:       opListListeners,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListListenersInput{}
	}

	req := c.newRequest(op, input, &ListListenersOutput{})
	return ListListenersRequest{Request: req, Input: input, Copy: c.ListListenersRequest}
}

// ListListenersRequest is the request type for the
// ListListeners API operation.
type ListListenersRequest struct {
	*aws.Request
	Input *ListListenersInput
	Copy  func(*ListListenersInput) ListListenersRequest
}

// Send marshals and sends the ListListeners API request.
func (r ListListenersRequest) Send(ctx context.Context) (*ListListenersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListListenersResponse{
		ListListenersOutput: r.Request.Data.(*ListListenersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListListenersResponse is the response type for the
// ListListeners API operation.
type ListListenersResponse struct {
	*ListListenersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListListeners request.
func (r *ListListenersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
