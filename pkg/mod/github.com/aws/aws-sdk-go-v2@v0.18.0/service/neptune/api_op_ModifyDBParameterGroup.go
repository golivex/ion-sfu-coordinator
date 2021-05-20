// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyDBParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB parameter group.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBParameterGroup.
	//
	// DBParameterGroupName is a required field
	DBParameterGroupName *string `type:"string" required:"true"`

	// An array of parameter names, values, and the apply method for the parameter
	// update. At least one parameter name, value, and apply method must be supplied;
	// subsequent arguments are optional. A maximum of 20 parameters can be modified
	// in a single request.
	//
	// Valid Values (for the application method): immediate | pending-reboot
	//
	// You can use the immediate value with dynamic parameters only. You can use
	// the pending-reboot value for both dynamic and static parameters, and changes
	// are applied when you reboot the DB instance without failover.
	//
	// Parameters is a required field
	Parameters []Parameter `locationNameList:"Parameter" type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyDBParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBParameterGroupInput"}

	if s.DBParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupName"))
	}

	if s.Parameters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Parameters"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyDBParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// Provides the name of the DB parameter group.
	DBParameterGroupName *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDBParameterGroup = "ModifyDBParameterGroup"

// ModifyDBParameterGroupRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Modifies the parameters of a DB parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName, ParameterValue,
// and ApplyMethod. A maximum of 20 parameters can be modified in a single request.
//
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot without failover to the DB instance associated
// with the parameter group before the change can take effect.
//
// After you modify a DB parameter group, you should wait at least 5 minutes
// before creating your first DB instance that uses that DB parameter group
// as the default parameter group. This allows Amazon Neptune to fully complete
// the modify action before the parameter group is used as the default for a
// new DB instance. This is especially important for parameters that are critical
// when creating the default database for a DB instance, such as the character
// set for the default database defined by the character_set_database parameter.
// You can use the Parameter Groups option of the Amazon Neptune console or
// the DescribeDBParameters command to verify that your DB parameter group has
// been created or modified.
//
//    // Example sending a request using ModifyDBParameterGroupRequest.
//    req := client.ModifyDBParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/ModifyDBParameterGroup
func (c *Client) ModifyDBParameterGroupRequest(input *ModifyDBParameterGroupInput) ModifyDBParameterGroupRequest {
	op := &aws.Operation{
		Name:       opModifyDBParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBParameterGroupInput{}
	}

	req := c.newRequest(op, input, &ModifyDBParameterGroupOutput{})
	return ModifyDBParameterGroupRequest{Request: req, Input: input, Copy: c.ModifyDBParameterGroupRequest}
}

// ModifyDBParameterGroupRequest is the request type for the
// ModifyDBParameterGroup API operation.
type ModifyDBParameterGroupRequest struct {
	*aws.Request
	Input *ModifyDBParameterGroupInput
	Copy  func(*ModifyDBParameterGroupInput) ModifyDBParameterGroupRequest
}

// Send marshals and sends the ModifyDBParameterGroup API request.
func (r ModifyDBParameterGroupRequest) Send(ctx context.Context) (*ModifyDBParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBParameterGroupResponse{
		ModifyDBParameterGroupOutput: r.Request.Data.(*ModifyDBParameterGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBParameterGroupResponse is the response type for the
// ModifyDBParameterGroup API operation.
type ModifyDBParameterGroupResponse struct {
	*ModifyDBParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBParameterGroup request.
func (r *ModifyDBParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
