// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTypeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the type.
	//
	// Conditional: You must specify TypeName or Arn.
	Arn *string `type:"string"`

	// The kind of type.
	//
	// Currently the only valid value is RESOURCE.
	Type RegistryType `type:"string" enum:"true"`

	// The name of the type.
	//
	// Conditional: You must specify TypeName or Arn.
	TypeName *string `min:"10" type:"string"`

	// The ID of a specific version of the type. The version ID is the value at
	// the end of the Amazon Resource Name (ARN) assigned to the type version when
	// it is registered.
	//
	// If you specify a VersionId, DescribeType returns information about that specific
	// type version. Otherwise, it returns information about the default type version.
	VersionId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTypeInput"}
	if s.TypeName != nil && len(*s.TypeName) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("TypeName", 10))
	}
	if s.VersionId != nil && len(*s.VersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTypeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the type.
	Arn *string `type:"string"`

	// The ID of the default version of the type. The default version is used when
	// the type version is not specified.
	//
	// To set the default version of a type, use SetTypeDefaultVersion .
	DefaultVersionId *string `min:"1" type:"string"`

	// The deprecation status of the type.
	//
	// Valid values include:
	//
	//    * LIVE: The type is registered and can be used in CloudFormation operations,
	//    dependent on its provisioning behavior and visibility scope.
	//
	//    * DEPRECATED: The type has been deregistered and can no longer be used
	//    in CloudFormation operations.
	DeprecatedStatus DeprecatedStatus `type:"string" enum:"true"`

	// The description of the registered type.
	Description *string `min:"1" type:"string"`

	// The URL of a page providing detailed documentation for this type.
	DocumentationUrl *string `type:"string"`

	// The Amazon Resource Name (ARN) of the IAM execution role used to register
	// the type. If your resource type calls AWS APIs in any of its handlers, you
	// must create an IAM execution role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
	// that includes the necessary permissions to call those AWS APIs, and provision
	// that execution role in your account. CloudFormation then assumes that execution
	// role to provide your resource type with the appropriate credentials.
	ExecutionRoleArn *string `min:"1" type:"string"`

	// When the specified type version was registered.
	LastUpdated *time.Time `type:"timestamp"`

	// Contains logging configuration information for a type.
	LoggingConfig *LoggingConfig `type:"structure"`

	// The provisioning behavior of the type. AWS CloudFormation determines the
	// provisioning type during registration, based on the types of handlers in
	// the schema handler package submitted.
	//
	// Valid values include:
	//
	//    * FULLY_MUTABLE: The type includes an update handler to process updates
	//    to the type during stack update operations.
	//
	//    * IMMUTABLE: The type does not include an update handler, so the type
	//    cannot be updated and must instead be replaced during stack update operations.
	//
	//    * NON_PROVISIONABLE: The type does not include all of the following handlers,
	//    and therefore cannot actually be provisioned. create read delete
	ProvisioningType ProvisioningType `type:"string" enum:"true"`

	// The schema that defines the type.
	//
	// For more information on type schemas, see Resource Provider Schema (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html)
	// in the CloudFormation CLI User Guide.
	Schema *string `min:"1" type:"string"`

	// The URL of the source code for the type.
	SourceUrl *string `type:"string"`

	// When the specified type version was registered.
	TimeCreated *time.Time `type:"timestamp"`

	// The kind of type.
	//
	// Currently the only valid value is RESOURCE.
	Type RegistryType `type:"string" enum:"true"`

	// The name of the registered type.
	TypeName *string `min:"10" type:"string"`

	// The scope at which the type is visible and usable in CloudFormation operations.
	//
	// Valid values include:
	//
	//    * PRIVATE: The type is only visible and usable within the account in which
	//    it is registered. Currently, AWS CloudFormation marks any types you register
	//    as PRIVATE.
	//
	//    * PUBLIC: The type is publically visible and usable within any Amazon
	//    account.
	Visibility Visibility `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeType = "DescribeType"

// DescribeTypeRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns detailed information about a type that has been registered.
//
// If you specify a VersionId, DescribeType returns information about that specific
// type version. Otherwise, it returns information about the default type version.
//
//    // Example sending a request using DescribeTypeRequest.
//    req := client.DescribeTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeType
func (c *Client) DescribeTypeRequest(input *DescribeTypeInput) DescribeTypeRequest {
	op := &aws.Operation{
		Name:       opDescribeType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTypeInput{}
	}

	req := c.newRequest(op, input, &DescribeTypeOutput{})
	return DescribeTypeRequest{Request: req, Input: input, Copy: c.DescribeTypeRequest}
}

// DescribeTypeRequest is the request type for the
// DescribeType API operation.
type DescribeTypeRequest struct {
	*aws.Request
	Input *DescribeTypeInput
	Copy  func(*DescribeTypeInput) DescribeTypeRequest
}

// Send marshals and sends the DescribeType API request.
func (r DescribeTypeRequest) Send(ctx context.Context) (*DescribeTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTypeResponse{
		DescribeTypeOutput: r.Request.Data.(*DescribeTypeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTypeResponse is the response type for the
// DescribeType API operation.
type DescribeTypeResponse struct {
	*DescribeTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeType request.
func (r *DescribeTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
