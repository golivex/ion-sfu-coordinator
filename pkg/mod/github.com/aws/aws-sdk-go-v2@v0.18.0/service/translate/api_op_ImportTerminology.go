// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ImportTerminologyInput struct {
	_ struct{} `type:"structure"`

	// The description of the custom terminology being imported.
	Description *string `type:"string"`

	// The encryption key for the custom terminology being imported.
	EncryptionKey *EncryptionKey `type:"structure"`

	// The merge strategy of the custom terminology being imported. Currently, only
	// the OVERWRITE merge strategy is supported. In this case, the imported terminology
	// will overwrite an existing terminology of the same name.
	//
	// MergeStrategy is a required field
	MergeStrategy MergeStrategy `type:"string" required:"true" enum:"true"`

	// The name of the custom terminology being imported.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The terminology data for the custom terminology being imported.
	//
	// TerminologyData is a required field
	TerminologyData *TerminologyData `type:"structure" required:"true"`
}

// String returns the string representation
func (s ImportTerminologyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportTerminologyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportTerminologyInput"}
	if len(s.MergeStrategy) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("MergeStrategy"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.TerminologyData == nil {
		invalidParams.Add(aws.NewErrParamRequired("TerminologyData"))
	}
	if s.EncryptionKey != nil {
		if err := s.EncryptionKey.Validate(); err != nil {
			invalidParams.AddNested("EncryptionKey", err.(aws.ErrInvalidParams))
		}
	}
	if s.TerminologyData != nil {
		if err := s.TerminologyData.Validate(); err != nil {
			invalidParams.AddNested("TerminologyData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ImportTerminologyOutput struct {
	_ struct{} `type:"structure"`

	// The properties of the custom terminology being imported.
	TerminologyProperties *TerminologyProperties `type:"structure"`
}

// String returns the string representation
func (s ImportTerminologyOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportTerminology = "ImportTerminology"

// ImportTerminologyRequest returns a request value for making API operation for
// Amazon Translate.
//
// Creates or updates a custom terminology, depending on whether or not one
// already exists for the given terminology name. Importing a terminology with
// the same name as an existing one will merge the terminologies based on the
// chosen merge strategy. Currently, the only supported merge strategy is OVERWRITE,
// and so the imported terminology will overwrite an existing terminology of
// the same name.
//
// If you import a terminology that overwrites an existing one, the new terminology
// take up to 10 minutes to fully propagate and be available for use in a translation
// due to cache policies with the DataPlane service that performs the translations.
//
//    // Example sending a request using ImportTerminologyRequest.
//    req := client.ImportTerminologyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/translate-2017-07-01/ImportTerminology
func (c *Client) ImportTerminologyRequest(input *ImportTerminologyInput) ImportTerminologyRequest {
	op := &aws.Operation{
		Name:       opImportTerminology,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportTerminologyInput{}
	}

	req := c.newRequest(op, input, &ImportTerminologyOutput{})
	return ImportTerminologyRequest{Request: req, Input: input, Copy: c.ImportTerminologyRequest}
}

// ImportTerminologyRequest is the request type for the
// ImportTerminology API operation.
type ImportTerminologyRequest struct {
	*aws.Request
	Input *ImportTerminologyInput
	Copy  func(*ImportTerminologyInput) ImportTerminologyRequest
}

// Send marshals and sends the ImportTerminology API request.
func (r ImportTerminologyRequest) Send(ctx context.Context) (*ImportTerminologyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportTerminologyResponse{
		ImportTerminologyOutput: r.Request.Data.(*ImportTerminologyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportTerminologyResponse is the response type for the
// ImportTerminology API operation.
type ImportTerminologyResponse struct {
	*ImportTerminologyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportTerminology request.
func (r *ImportTerminologyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
