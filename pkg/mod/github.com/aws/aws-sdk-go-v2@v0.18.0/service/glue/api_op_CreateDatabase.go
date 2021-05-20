// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDatabaseInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which to create the database. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The metadata for the database.
	//
	// DatabaseInput is a required field
	DatabaseInput *DatabaseInput `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateDatabaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDatabaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDatabaseInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseInput == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseInput"))
	}
	if s.DatabaseInput != nil {
		if err := s.DatabaseInput.Validate(); err != nil {
			invalidParams.AddNested("DatabaseInput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDatabaseOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateDatabaseOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDatabase = "CreateDatabase"

// CreateDatabaseRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a new database in a Data Catalog.
//
//    // Example sending a request using CreateDatabaseRequest.
//    req := client.CreateDatabaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateDatabase
func (c *Client) CreateDatabaseRequest(input *CreateDatabaseInput) CreateDatabaseRequest {
	op := &aws.Operation{
		Name:       opCreateDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDatabaseInput{}
	}

	req := c.newRequest(op, input, &CreateDatabaseOutput{})
	return CreateDatabaseRequest{Request: req, Input: input, Copy: c.CreateDatabaseRequest}
}

// CreateDatabaseRequest is the request type for the
// CreateDatabase API operation.
type CreateDatabaseRequest struct {
	*aws.Request
	Input *CreateDatabaseInput
	Copy  func(*CreateDatabaseInput) CreateDatabaseRequest
}

// Send marshals and sends the CreateDatabase API request.
func (r CreateDatabaseRequest) Send(ctx context.Context) (*CreateDatabaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDatabaseResponse{
		CreateDatabaseOutput: r.Request.Data.(*CreateDatabaseOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDatabaseResponse is the response type for the
// CreateDatabase API operation.
type CreateDatabaseResponse struct {
	*CreateDatabaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDatabase request.
func (r *CreateDatabaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
