// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutMailboxPermissionsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the user, group, or resource for which to update mailbox
	// permissions.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The identifier of the user, group, or resource to which to grant the permissions.
	//
	// GranteeId is a required field
	GranteeId *string `min:"12" type:"string" required:"true"`

	// The identifier of the organization under which the user, group, or resource
	// exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The permissions granted to the grantee. SEND_AS allows the grantee to send
	// email as the owner of the mailbox (the grantee is not mentioned on these
	// emails). SEND_ON_BEHALF allows the grantee to send email on behalf of the
	// owner of the mailbox (the grantee is not mentioned as the physical sender
	// of these emails). FULL_ACCESS allows the grantee full access to the mailbox,
	// irrespective of other folder-level permissions set on the mailbox.
	//
	// PermissionValues is a required field
	PermissionValues []PermissionType `type:"list" required:"true"`
}

// String returns the string representation
func (s PutMailboxPermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutMailboxPermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutMailboxPermissionsInput"}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.GranteeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GranteeId"))
	}
	if s.GranteeId != nil && len(*s.GranteeId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("GranteeId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.PermissionValues == nil {
		invalidParams.Add(aws.NewErrParamRequired("PermissionValues"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutMailboxPermissionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutMailboxPermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutMailboxPermissions = "PutMailboxPermissions"

// PutMailboxPermissionsRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Sets permissions for a user, group, or resource. This replaces any pre-existing
// permissions.
//
//    // Example sending a request using PutMailboxPermissionsRequest.
//    req := client.PutMailboxPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/PutMailboxPermissions
func (c *Client) PutMailboxPermissionsRequest(input *PutMailboxPermissionsInput) PutMailboxPermissionsRequest {
	op := &aws.Operation{
		Name:       opPutMailboxPermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutMailboxPermissionsInput{}
	}

	req := c.newRequest(op, input, &PutMailboxPermissionsOutput{})
	return PutMailboxPermissionsRequest{Request: req, Input: input, Copy: c.PutMailboxPermissionsRequest}
}

// PutMailboxPermissionsRequest is the request type for the
// PutMailboxPermissions API operation.
type PutMailboxPermissionsRequest struct {
	*aws.Request
	Input *PutMailboxPermissionsInput
	Copy  func(*PutMailboxPermissionsInput) PutMailboxPermissionsRequest
}

// Send marshals and sends the PutMailboxPermissions API request.
func (r PutMailboxPermissionsRequest) Send(ctx context.Context) (*PutMailboxPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutMailboxPermissionsResponse{
		PutMailboxPermissionsOutput: r.Request.Data.(*PutMailboxPermissionsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutMailboxPermissionsResponse is the response type for the
// PutMailboxPermissions API operation.
type PutMailboxPermissionsResponse struct {
	*PutMailboxPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutMailboxPermissions request.
func (r *PutMailboxPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
