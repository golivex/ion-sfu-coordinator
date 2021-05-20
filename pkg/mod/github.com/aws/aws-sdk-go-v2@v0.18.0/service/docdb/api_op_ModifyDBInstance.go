// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to ModifyDBInstance.
type ModifyDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether the modifications in this request and any pending modifications
	// are asynchronously applied as soon as possible, regardless of the PreferredMaintenanceWindow
	// setting for the DB instance.
	//
	// If this parameter is set to false, changes to the DB instance are applied
	// during the next maintenance window. Some parameter changes can cause an outage
	// and are applied on the next reboot.
	//
	// Default: false
	ApplyImmediately *bool `type:"boolean"`

	// Indicates that minor version upgrades are applied automatically to the DB
	// instance during the maintenance window. Changing this parameter doesn't result
	// in an outage except in the following case, and the change is asynchronously
	// applied as soon as possible. An outage results if this parameter is set to
	// true during the maintenance window, and a newer minor version is available,
	// and Amazon DocumentDB has enabled automatic patching for that engine version.
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// Indicates the certificate that needs to be associated with the instance.
	CACertificateIdentifier *string `type:"string"`

	// The new compute and memory capacity of the DB instance; for example, db.r5.large.
	// Not all DB instance classes are available in all AWS Regions.
	//
	// If you modify the DB instance class, an outage occurs during the change.
	// The change is applied during the next maintenance window, unless ApplyImmediately
	// is specified as true for this request.
	//
	// Default: Uses existing setting.
	DBInstanceClass *string `type:"string"`

	// The DB instance identifier. This value is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBInstance.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The new DB instance identifier for the DB instance when renaming a DB instance.
	// When you change the DB instance identifier, an instance reboot occurs immediately
	// if you set Apply Immediately to true. It occurs during the next maintenance
	// window if you set Apply Immediately to false. This value is stored as a lowercase
	// string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * The first character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: mydbinstance
	NewDBInstanceIdentifier *string `type:"string"`

	// The weekly time range (in UTC) during which system maintenance can occur,
	// which might result in an outage. Changing this parameter doesn't result in
	// an outage except in the following situation, and the change is asynchronously
	// applied as soon as possible. If there are pending actions that cause a reboot,
	// and the maintenance window is changed to include the current time, changing
	// this parameter causes a reboot of the DB instance. If you are moving this
	// window to the current time, there must be at least 30 minutes between the
	// current time and end of the window to ensure that pending changes are applied.
	//
	// Default: Uses existing setting.
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Must be at least 30 minutes.
	PreferredMaintenanceWindow *string `type:"string"`

	// A value that specifies the order in which an Amazon DocumentDB replica is
	// promoted to the primary instance after a failure of the existing primary
	// instance.
	//
	// Default: 1
	//
	// Valid values: 0-15
	PromotionTier *int64 `type:"integer"`
}

// String returns the string representation
func (s ModifyDBInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBInstanceInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about a DB instance.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s ModifyDBInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDBInstance = "ModifyDBInstance"

// ModifyDBInstanceRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Modifies settings for a DB instance. You can change one or more database
// configuration parameters by specifying these parameters and the new values
// in the request.
//
//    // Example sending a request using ModifyDBInstanceRequest.
//    req := client.ModifyDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/ModifyDBInstance
func (c *Client) ModifyDBInstanceRequest(input *ModifyDBInstanceInput) ModifyDBInstanceRequest {
	op := &aws.Operation{
		Name:       opModifyDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceInput{}
	}

	req := c.newRequest(op, input, &ModifyDBInstanceOutput{})
	return ModifyDBInstanceRequest{Request: req, Input: input, Copy: c.ModifyDBInstanceRequest}
}

// ModifyDBInstanceRequest is the request type for the
// ModifyDBInstance API operation.
type ModifyDBInstanceRequest struct {
	*aws.Request
	Input *ModifyDBInstanceInput
	Copy  func(*ModifyDBInstanceInput) ModifyDBInstanceRequest
}

// Send marshals and sends the ModifyDBInstance API request.
func (r ModifyDBInstanceRequest) Send(ctx context.Context) (*ModifyDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBInstanceResponse{
		ModifyDBInstanceOutput: r.Request.Data.(*ModifyDBInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBInstanceResponse is the response type for the
// ModifyDBInstance API operation.
type ModifyDBInstanceResponse struct {
	*ModifyDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBInstance request.
func (r *ModifyDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
