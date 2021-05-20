// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RemoveIpRoutesInput struct {
	_ struct{} `type:"structure"`

	// IP address blocks that you want to remove.
	//
	// CidrIps is a required field
	CidrIps []string `type:"list" required:"true"`

	// Identifier (ID) of the directory from which you want to remove the IP addresses.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveIpRoutesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveIpRoutesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveIpRoutesInput"}

	if s.CidrIps == nil {
		invalidParams.Add(aws.NewErrParamRequired("CidrIps"))
	}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveIpRoutesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveIpRoutesOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveIpRoutes = "RemoveIpRoutes"

// RemoveIpRoutesRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Removes IP address blocks from a directory.
//
//    // Example sending a request using RemoveIpRoutesRequest.
//    req := client.RemoveIpRoutesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/RemoveIpRoutes
func (c *Client) RemoveIpRoutesRequest(input *RemoveIpRoutesInput) RemoveIpRoutesRequest {
	op := &aws.Operation{
		Name:       opRemoveIpRoutes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveIpRoutesInput{}
	}

	req := c.newRequest(op, input, &RemoveIpRoutesOutput{})
	return RemoveIpRoutesRequest{Request: req, Input: input, Copy: c.RemoveIpRoutesRequest}
}

// RemoveIpRoutesRequest is the request type for the
// RemoveIpRoutes API operation.
type RemoveIpRoutesRequest struct {
	*aws.Request
	Input *RemoveIpRoutesInput
	Copy  func(*RemoveIpRoutesInput) RemoveIpRoutesRequest
}

// Send marshals and sends the RemoveIpRoutes API request.
func (r RemoveIpRoutesRequest) Send(ctx context.Context) (*RemoveIpRoutesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveIpRoutesResponse{
		RemoveIpRoutesOutput: r.Request.Data.(*RemoveIpRoutesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveIpRoutesResponse is the response type for the
// RemoveIpRoutes API operation.
type RemoveIpRoutesResponse struct {
	*RemoveIpRoutesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveIpRoutes request.
func (r *RemoveIpRoutesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
