// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteUserInput struct {
	_ struct{} `type:"structure"`

	// The authentication type for the user. You must specify USERPOOL.
	//
	// AuthenticationType is a required field
	AuthenticationType AuthenticationType `type:"string" required:"true" enum:"true"`

	// The email address of the user.
	//
	// Users' email addresses are case-sensitive.
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s DeleteUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserInput"}
	if len(s.AuthenticationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AuthenticationType"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUser = "DeleteUser"

// DeleteUserRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Deletes a user from the user pool.
//
//    // Example sending a request using DeleteUserRequest.
//    req := client.DeleteUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteUser
func (c *Client) DeleteUserRequest(input *DeleteUserInput) DeleteUserRequest {
	op := &aws.Operation{
		Name:       opDeleteUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserInput{}
	}

	req := c.newRequest(op, input, &DeleteUserOutput{})
	return DeleteUserRequest{Request: req, Input: input, Copy: c.DeleteUserRequest}
}

// DeleteUserRequest is the request type for the
// DeleteUser API operation.
type DeleteUserRequest struct {
	*aws.Request
	Input *DeleteUserInput
	Copy  func(*DeleteUserInput) DeleteUserRequest
}

// Send marshals and sends the DeleteUser API request.
func (r DeleteUserRequest) Send(ctx context.Context) (*DeleteUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserResponse{
		DeleteUserOutput: r.Request.Data.(*DeleteUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserResponse is the response type for the
// DeleteUser API operation.
type DeleteUserResponse struct {
	*DeleteUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUser request.
func (r *DeleteUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
