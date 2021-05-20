// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the ConnectDirectory operation.
type ConnectDirectoryInput struct {
	_ struct{} `type:"structure"`

	// A DirectoryConnectSettings object that contains additional information for
	// the operation.
	//
	// ConnectSettings is a required field
	ConnectSettings *DirectoryConnectSettings `type:"structure" required:"true"`

	// A description for the directory.
	Description *string `type:"string"`

	// The fully qualified name of the on-premises directory, such as corp.example.com.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The password for the on-premises user account.
	//
	// Password is a required field
	Password *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// The NetBIOS name of the on-premises directory, such as CORP.
	ShortName *string `type:"string"`

	// The size of the directory.
	//
	// Size is a required field
	Size DirectorySize `type:"string" required:"true" enum:"true"`

	// The tags to be assigned to AD Connector.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s ConnectDirectoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConnectDirectoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConnectDirectoryInput"}

	if s.ConnectSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectSettings"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}
	if s.Password != nil && len(*s.Password) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Password", 1))
	}
	if len(s.Size) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Size"))
	}
	if s.ConnectSettings != nil {
		if err := s.ConnectSettings.Validate(); err != nil {
			invalidParams.AddNested("ConnectSettings", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the results of the ConnectDirectory operation.
type ConnectDirectoryOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the new directory.
	DirectoryId *string `type:"string"`
}

// String returns the string representation
func (s ConnectDirectoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opConnectDirectory = "ConnectDirectory"

// ConnectDirectoryRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Creates an AD Connector to connect to an on-premises directory.
//
// Before you call ConnectDirectory, ensure that all of the required permissions
// have been explicitly granted through a policy. For details about what permissions
// are required to run the ConnectDirectory operation, see AWS Directory Service
// API Permissions: Actions, Resources, and Conditions Reference (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html).
//
//    // Example sending a request using ConnectDirectoryRequest.
//    req := client.ConnectDirectoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/ConnectDirectory
func (c *Client) ConnectDirectoryRequest(input *ConnectDirectoryInput) ConnectDirectoryRequest {
	op := &aws.Operation{
		Name:       opConnectDirectory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ConnectDirectoryInput{}
	}

	req := c.newRequest(op, input, &ConnectDirectoryOutput{})
	return ConnectDirectoryRequest{Request: req, Input: input, Copy: c.ConnectDirectoryRequest}
}

// ConnectDirectoryRequest is the request type for the
// ConnectDirectory API operation.
type ConnectDirectoryRequest struct {
	*aws.Request
	Input *ConnectDirectoryInput
	Copy  func(*ConnectDirectoryInput) ConnectDirectoryRequest
}

// Send marshals and sends the ConnectDirectory API request.
func (r ConnectDirectoryRequest) Send(ctx context.Context) (*ConnectDirectoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ConnectDirectoryResponse{
		ConnectDirectoryOutput: r.Request.Data.(*ConnectDirectoryOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ConnectDirectoryResponse is the response type for the
// ConnectDirectory API operation.
type ConnectDirectoryResponse struct {
	*ConnectDirectoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ConnectDirectory request.
func (r *ConnectDirectoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
