// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateProductFromPortfolioInput struct {
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

	// The product identifier.
	//
	// ProductId is a required field
	ProductId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateProductFromPortfolioInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateProductFromPortfolioInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateProductFromPortfolioInput"}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}

	if s.ProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductId"))
	}
	if s.ProductId != nil && len(*s.ProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateProductFromPortfolioOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateProductFromPortfolioOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateProductFromPortfolio = "DisassociateProductFromPortfolio"

// DisassociateProductFromPortfolioRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Disassociates the specified product from the specified portfolio.
//
//    // Example sending a request using DisassociateProductFromPortfolioRequest.
//    req := client.DisassociateProductFromPortfolioRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DisassociateProductFromPortfolio
func (c *Client) DisassociateProductFromPortfolioRequest(input *DisassociateProductFromPortfolioInput) DisassociateProductFromPortfolioRequest {
	op := &aws.Operation{
		Name:       opDisassociateProductFromPortfolio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateProductFromPortfolioInput{}
	}

	req := c.newRequest(op, input, &DisassociateProductFromPortfolioOutput{})
	return DisassociateProductFromPortfolioRequest{Request: req, Input: input, Copy: c.DisassociateProductFromPortfolioRequest}
}

// DisassociateProductFromPortfolioRequest is the request type for the
// DisassociateProductFromPortfolio API operation.
type DisassociateProductFromPortfolioRequest struct {
	*aws.Request
	Input *DisassociateProductFromPortfolioInput
	Copy  func(*DisassociateProductFromPortfolioInput) DisassociateProductFromPortfolioRequest
}

// Send marshals and sends the DisassociateProductFromPortfolio API request.
func (r DisassociateProductFromPortfolioRequest) Send(ctx context.Context) (*DisassociateProductFromPortfolioResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateProductFromPortfolioResponse{
		DisassociateProductFromPortfolioOutput: r.Request.Data.(*DisassociateProductFromPortfolioOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateProductFromPortfolioResponse is the response type for the
// DisassociateProductFromPortfolio API operation.
type DisassociateProductFromPortfolioResponse struct {
	*DisassociateProductFromPortfolioOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateProductFromPortfolio request.
func (r *DisassociateProductFromPortfolioResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
