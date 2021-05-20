// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteConnectionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the connection.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConnectionInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about an AWS Direct Connect connection.
type DeleteConnectionOutput struct {
	_ struct{} `type:"structure"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string `locationName:"awsDevice" deprecated:"true" type:"string"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDeviceV2 *string `locationName:"awsDeviceV2" type:"string"`

	// The bandwidth of the connection.
	Bandwidth *string `locationName:"bandwidth" type:"string"`

	// The ID of the connection.
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// The name of the connection.
	ConnectionName *string `locationName:"connectionName" type:"string"`

	// The state of the connection. The following are the possible values:
	//
	//    * ordering: The initial state of a hosted connection provisioned on an
	//    interconnect. The connection stays in the ordering state until the owner
	//    of the hosted connection confirms or declines the connection order.
	//
	//    * requested: The initial state of a standard connection. The connection
	//    stays in the requested state until the Letter of Authorization (LOA) is
	//    sent to the customer.
	//
	//    * pending: The connection has been approved and is being initialized.
	//
	//    * available: The network link is up and the connection is ready for use.
	//
	//    * down: The network link is down.
	//
	//    * deleting: The connection is being deleted.
	//
	//    * deleted: The connection has been deleted.
	//
	//    * rejected: A hosted connection in the ordering state enters the rejected
	//    state if it is deleted by the customer.
	//
	//    * unknown: The state of the connection is not available.
	ConnectionState ConnectionState `locationName:"connectionState" type:"string" enum:"true"`

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy HasLogicalRedundancy `locationName:"hasLogicalRedundancy" type:"string" enum:"true"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `locationName:"jumboFrameCapable" type:"boolean"`

	// The ID of the LAG.
	LagId *string `locationName:"lagId" type:"string"`

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time `locationName:"loaIssueTime" type:"timestamp"`

	// The location of the connection.
	Location *string `locationName:"location" type:"string"`

	// The ID of the AWS account that owns the connection.
	OwnerAccount *string `locationName:"ownerAccount" type:"string"`

	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName *string `locationName:"partnerName" type:"string"`

	// The name of the service provider associated with the connection.
	ProviderName *string `locationName:"providerName" type:"string"`

	// The AWS Region where the connection is located.
	Region *string `locationName:"region" type:"string"`

	// The tags associated with the connection.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`

	// The ID of the VLAN.
	Vlan *int64 `locationName:"vlan" type:"integer"`
}

// String returns the string representation
func (s DeleteConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConnection = "DeleteConnection"

// DeleteConnectionRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Deletes the specified connection.
//
// Deleting a connection only stops the AWS Direct Connect port hour and data
// transfer charges. If you are partnering with any third parties to connect
// with the AWS Direct Connect location, you must cancel your service with them
// separately.
//
//    // Example sending a request using DeleteConnectionRequest.
//    req := client.DeleteConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DeleteConnection
func (c *Client) DeleteConnectionRequest(input *DeleteConnectionInput) DeleteConnectionRequest {
	op := &aws.Operation{
		Name:       opDeleteConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConnectionInput{}
	}

	req := c.newRequest(op, input, &DeleteConnectionOutput{})
	return DeleteConnectionRequest{Request: req, Input: input, Copy: c.DeleteConnectionRequest}
}

// DeleteConnectionRequest is the request type for the
// DeleteConnection API operation.
type DeleteConnectionRequest struct {
	*aws.Request
	Input *DeleteConnectionInput
	Copy  func(*DeleteConnectionInput) DeleteConnectionRequest
}

// Send marshals and sends the DeleteConnection API request.
func (r DeleteConnectionRequest) Send(ctx context.Context) (*DeleteConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConnectionResponse{
		DeleteConnectionOutput: r.Request.Data.(*DeleteConnectionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConnectionResponse is the response type for the
// DeleteConnection API operation.
type DeleteConnectionResponse struct {
	*DeleteConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConnection request.
func (r *DeleteConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
