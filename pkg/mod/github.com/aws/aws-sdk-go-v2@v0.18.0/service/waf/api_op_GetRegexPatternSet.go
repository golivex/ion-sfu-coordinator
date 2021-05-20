// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRegexPatternSetInput struct {
	_ struct{} `type:"structure"`

	// The RegexPatternSetId of the RegexPatternSet that you want to get. RegexPatternSetId
	// is returned by CreateRegexPatternSet and by ListRegexPatternSets.
	//
	// RegexPatternSetId is a required field
	RegexPatternSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRegexPatternSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRegexPatternSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRegexPatternSetInput"}

	if s.RegexPatternSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegexPatternSetId"))
	}
	if s.RegexPatternSetId != nil && len(*s.RegexPatternSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RegexPatternSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRegexPatternSetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the RegexPatternSet that you specified in the GetRegexPatternSet
	// request, including the identifier of the pattern set and the regular expression
	// patterns you want AWS WAF to search for.
	RegexPatternSet *RegexPatternSet `type:"structure"`
}

// String returns the string representation
func (s GetRegexPatternSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRegexPatternSet = "GetRegexPatternSet"

// GetRegexPatternSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Returns the RegexPatternSet specified by RegexPatternSetId.
//
//    // Example sending a request using GetRegexPatternSetRequest.
//    req := client.GetRegexPatternSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetRegexPatternSet
func (c *Client) GetRegexPatternSetRequest(input *GetRegexPatternSetInput) GetRegexPatternSetRequest {
	op := &aws.Operation{
		Name:       opGetRegexPatternSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRegexPatternSetInput{}
	}

	req := c.newRequest(op, input, &GetRegexPatternSetOutput{})
	return GetRegexPatternSetRequest{Request: req, Input: input, Copy: c.GetRegexPatternSetRequest}
}

// GetRegexPatternSetRequest is the request type for the
// GetRegexPatternSet API operation.
type GetRegexPatternSetRequest struct {
	*aws.Request
	Input *GetRegexPatternSetInput
	Copy  func(*GetRegexPatternSetInput) GetRegexPatternSetRequest
}

// Send marshals and sends the GetRegexPatternSet API request.
func (r GetRegexPatternSetRequest) Send(ctx context.Context) (*GetRegexPatternSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRegexPatternSetResponse{
		GetRegexPatternSetOutput: r.Request.Data.(*GetRegexPatternSetOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRegexPatternSetResponse is the response type for the
// GetRegexPatternSet API operation.
type GetRegexPatternSetResponse struct {
	*GetRegexPatternSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRegexPatternSet request.
func (r *GetRegexPatternSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
