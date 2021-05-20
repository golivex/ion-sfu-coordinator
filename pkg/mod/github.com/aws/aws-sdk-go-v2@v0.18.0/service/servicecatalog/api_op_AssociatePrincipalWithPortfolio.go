// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociatePrincipalWithPortfolioInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The portfolio identifier.
	//
	// PortfolioId is a required field
	PortfolioId *string `min:"1" type:"string" required:"true"`

	// The ARN of the principal (IAM user, role, or group).
	//
	// PrincipalARN is a required field
	PrincipalARN *string `min:"1" type:"string" required:"true"`

	// The principal type. The supported value is IAM.
	//
	// PrincipalType is a required field
	PrincipalType PrincipalType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s AssociatePrincipalWithPortfolioInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociatePrincipalWithPortfolioInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociatePrincipalWithPortfolioInput"}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}

	if s.PrincipalARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("PrincipalARN"))
	}
	if s.PrincipalARN != nil && len(*s.PrincipalARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PrincipalARN", 1))
	}
	if len(s.PrincipalType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PrincipalType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociatePrincipalWithPortfolioOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociatePrincipalWithPortfolioOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociatePrincipalWithPortfolio = "AssociatePrincipalWithPortfolio"

// AssociatePrincipalWithPortfolioRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Associates the specified principal ARN with the specified portfolio.
//
//    // Example sending a request using AssociatePrincipalWithPortfolioRequest.
//    req := client.AssociatePrincipalWithPortfolioRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/AssociatePrincipalWithPortfolio
func (c *Client) AssociatePrincipalWithPortfolioRequest(input *AssociatePrincipalWithPortfolioInput) AssociatePrincipalWithPortfolioRequest {
	op := &aws.Operation{
		Name:       opAssociatePrincipalWithPortfolio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociatePrincipalWithPortfolioInput{}
	}

	req := c.newRequest(op, input, &AssociatePrincipalWithPortfolioOutput{})
	return AssociatePrincipalWithPortfolioRequest{Request: req, Input: input, Copy: c.AssociatePrincipalWithPortfolioRequest}
}

// AssociatePrincipalWithPortfolioRequest is the request type for the
// AssociatePrincipalWithPortfolio API operation.
type AssociatePrincipalWithPortfolioRequest struct {
	*aws.Request
	Input *AssociatePrincipalWithPortfolioInput
	Copy  func(*AssociatePrincipalWithPortfolioInput) AssociatePrincipalWithPortfolioRequest
}

// Send marshals and sends the AssociatePrincipalWithPortfolio API request.
func (r AssociatePrincipalWithPortfolioRequest) Send(ctx context.Context) (*AssociatePrincipalWithPortfolioResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociatePrincipalWithPortfolioResponse{
		AssociatePrincipalWithPortfolioOutput: r.Request.Data.(*AssociatePrincipalWithPortfolioOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociatePrincipalWithPortfolioResponse is the response type for the
// AssociatePrincipalWithPortfolio API operation.
type AssociatePrincipalWithPortfolioResponse struct {
	*AssociatePrincipalWithPortfolioOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociatePrincipalWithPortfolio request.
func (r *AssociatePrincipalWithPortfolioResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
